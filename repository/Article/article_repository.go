package repository

import (
	"coderhub/model"
	"coderhub/shared/CacheDB"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	CreateArticle(article *model.Articles) error
	GetArticleByID(id int64) (*model.Articles, error)
	UpdateArticle(article *model.Articles) error
	DeleteArticle(id int64) error
}
type ArticleRepositoryImpl struct {
	DB    *gorm.DB
	Redis CacheDB.RedisDB
}

func NewArticleRepositoryImpl(db *gorm.DB, rdb CacheDB.RedisDB) *ArticleRepositoryImpl {
	return &ArticleRepositoryImpl{
		DB:    db,
		Redis: rdb,
	}
}

func (r *ArticleRepositoryImpl) CreateArticle(article *model.Articles) error {
	return r.DB.Create(article).Error
}

func (r *ArticleRepositoryImpl) GetArticleByID(id int64) (*model.Articles, error) {
	var article model.Articles
	key := article.CacheKeyByID(id)

	// 尝试从缓存获取
	if cached, err := r.getCache(key); err == nil {
		return cached, nil
	} else {
		// 从数据库获取
		fmt.Printf("获取的文章ID: %v\n", id)
		if err := r.DB.Model(&model.Articles{}).Where("id = ?", id).First(&article).Error; err != nil {
			return nil, err
		}
		// 缓存结果
		err := r.setCache(key, &article)
		if err != nil {
			return nil, err
		}
		return &article, nil
	}
}

func (r *ArticleRepositoryImpl) UpdateArticle(article *model.Articles) error {
	oldArticle, err := r.GetArticleByID(article.ID)
	if err != nil {
		return err
	}
	if oldArticle.ID != article.ID {
		return errors.New("非法操作")
	}
	err = r.delCache(oldArticle.CacheKeyByID(oldArticle.ID))
	if err != nil {
		return err
	}
	if err := r.DB.Model(&model.Articles{}).Where("id = ?", article.ID).Updates(article).Error; err != nil {
		return err
	}

	err = r.setCache(article.CacheKeyByID(article.ID), article)
	if err != nil {
		return err
	}
	return nil
}

func (r *ArticleRepositoryImpl) DeleteArticle(id int64) error {
	article, err := r.GetArticleByID(id)
	if err != nil {
		return err
	}
	err = r.delCache(article.CacheKeyByID(id))
	if err != nil {
		return err
	}
	return r.DB.Delete(&model.Articles{}, id).Error
}

func (r *ArticleRepositoryImpl) getCache(key string) (*model.Articles, error) {
	var article model.Articles
	data, err := r.Redis.Get(key)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(data), &article); err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *ArticleRepositoryImpl) setCache(key string, article *model.Articles) error {
	data, err := json.Marshal(article)
	if err != nil {
		return err
	}
	return r.Redis.Set(key, string(data))
}

func (r *ArticleRepositoryImpl) delCache(key string) error {
	return r.Redis.Del(key)
}
