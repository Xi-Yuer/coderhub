package repository

import (
	"coderhub/model"
	"coderhub/shared/CacheDB"
	"errors"
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
		lastSyncTime:  time.Now(),
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
				// 在goroutine中处理锁的续期
				stopChan := make(chan struct{})
				go func() {
					ticker := time.NewTicker(10 * time.Second)
					defer ticker.Stop()
					for {
						select {
						case <-ticker.C:
							// 每10秒续期一次，将过期时间重置为30秒
							_ = r.Redis.Expire(lockKey, 30*time.Second)
						case <-stopChan:
							return
						}
					}
				}()

				// 执行同步
				err := r.SyncIncrementalPVToDB()
				// 同步完成后，关闭续期goroutine并释放锁
				close(stopChan)
				_ = r.Redis.Del(lockKey)

				if err == nil {
					_ = r.Redis.Set("article:pv:last_sync_time", time.Now().Format(time.RFC3339))
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
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
		err = r.DB.Model(&model.ArticlePV{}).Where("article_id = ?", articleID).Update("count", count).Error
		if err != nil {
			continue
		}

		// 同步成功后从待同步集合中删除
		_ = r.Redis.SRem("article:pv:need_sync", articleIDStr)
	}

	return nil
}
