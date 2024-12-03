package repository

import (
	"coderhub/model"
	"coderhub/shared/CacheDB"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type ArticlePVRepository interface {
	CreateArticlePV(articlePV *model.ArticlePV) error
	GetArticlePVByArticleID(articleID int64) (*model.ArticlePV, error)
	SyncIncrementalPVToDB() error
}

type ArticlePVRepositoryImpl struct {
	DB            *gorm.DB
	Redis         CacheDB.RedisDB
	lastSyncTime  time.Time
	syncThreshold int64
}

func NewArticlePVRepositoryImpl(db *gorm.DB, rdb CacheDB.RedisDB) *ArticlePVRepositoryImpl {
	return &ArticlePVRepositoryImpl{
		DB:            db,
		Redis:         rdb,
		syncThreshold: 100,
	}
}

func (r *ArticlePVRepositoryImpl) CreateArticlePV(articlePV *model.ArticlePV) error {
	articleIDStr := fmt.Sprintf("%d", articlePV.ArticleID)

	// 增加PV计数
	newCount, err := r.Redis.HIncrBy("article:pv", articleIDStr, 1)
	if err != nil {
		return err
	}

	// 将文章ID添加到待同步集合
	err = r.Redis.SAdd("article:pv:need_sync", articleIDStr)
	if err != nil {
		return err
	}

	// 检查是否需要触发同步
	shouldSync := false

	// 尝试获取分布式锁
	lockKey := "article:pv:sync:lock"
	// 设置锁的过期时间为30秒
	locked, err := r.Redis.SetNX(lockKey, "1", 30*time.Second)
	if err != nil {
		return err
	}

	if locked {
		defer r.Redis.Del(lockKey)

		// 数量阈值触发
		if newCount%r.syncThreshold == 0 {
			shouldSync = true
		}

		// 获取上次同步时间
		lastSyncStr, err := r.Redis.Get("article:pv:last_sync_time")
		if err == nil && lastSyncStr != "" {
			lastSyncTime, err := time.Parse(time.RFC3339, lastSyncStr)
			if err == nil && time.Since(lastSyncTime) > 5*time.Minute {
				shouldSync = true
			}
		}

		if shouldSync {
			go func() {
				err := r.SyncIncrementalPVToDB()
				if err == nil {
					// 更新最后同步时间
					r.Redis.Set("article:pv:last_sync_time", time.Now().Format(time.RFC3339))
				}
			}()
		}
	}

	return nil
}

func (r *ArticlePVRepositoryImpl) GetArticlePVByArticleID(articleID int64) (*model.ArticlePV, error) {
	count, err := r.Redis.HGet("article:pv", fmt.Sprintf("%d", articleID))
	if err == nil && count != "" {
		countInt, err := strconv.ParseInt(count, 10, 64)
		if err == nil {
			return &model.ArticlePV{ArticleID: articleID, Count: countInt}, nil
		}
	}

	var articlePV model.ArticlePV
	err = r.DB.Where("article_id = ?", articleID).First(&articlePV).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &model.ArticlePV{ArticleID: articleID, Count: 0}, nil
		}
		return nil, err
	}
	return &articlePV, nil
}

func (r *ArticlePVRepositoryImpl) SyncIncrementalPVToDB() error {
	// 获取需要同步的文章ID集合
	needSyncSet, err := r.Redis.SMembers("article:pv:need_sync")
	if err != nil {
		return err
	}

	if len(needSyncSet) == 0 {
		return nil
	}

	// 批量获取这些文章的PV数据
	for _, articleIDStr := range needSyncSet {
		articleID, err := strconv.ParseInt(articleIDStr, 10, 64)
		if err != nil {
			continue
		}

		// 获取PV数据
		countStr, err := r.Redis.HGet("article:pv", articleIDStr)
		if err != nil || countStr == "" {
			continue
		}

		count, err := strconv.ParseInt(countStr, 10, 64)
		if err != nil {
			continue
		}

		// 更新数据库
		err = r.DB.Exec(`
			INSERT INTO article_pvs (article_id, count, last_sync_at, updated_at, created_at)
			VALUES (?, ?, NOW(), NOW(), NOW())
			ON DUPLICATE KEY UPDATE 
			count = ?,
			last_sync_at = NOW(),
			updated_at = NOW()
		`, articleID, count, count).Error

		if err != nil {
			continue
		}

		// 同步成功后从待同步集合中删除
		r.Redis.SRem("article:pv:need_sync", articleIDStr)
	}

	return nil
}
