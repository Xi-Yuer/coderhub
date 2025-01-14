package main

import (
	"coderhub/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func main() {
	var articles []*model.ArticleAndAuthInfo
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:2214380963Wx!!@tcp(localhost:3306)/coderhub?charset=utf8&parseTime=True&loc=Local", // Data source name
		DefaultStringSize:         256,                                                                                      // Default string size
		DisableDatetimePrecision:  true,                                                                                     // Disable datetime precision for MySQL < 5.6
		DontSupportRenameIndex:    true,                                                                                     // Drop & create index for renaming in MySQL < 5.7
		DontSupportRenameColumn:   true,                                                                                     // Use 'change' for renaming columns
		SkipInitializeWithVersion: false,                                                                                    // Automatically configure based on MySQL version
	}), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level set to Info for development
				IgnoreRecordNotFoundError: false,       // Log ErrRecordNotFound errors
				ParameterizedQueries:      true,        // Include params in SQL log for easier debugging
				Colorful:                  true,        // Enable colored output
			},
		),
	})
	err = db.Table("articles AS a").
		Select(`
        a.*,
		u.*,
        u.id AS author_id,
        GROUP_CONCAT(img.url) AS images
    	`).
		Joins("JOIN users u ON a.author_id = u.id").
		Joins("LEFT JOIN image_relations ir ON a.id = ir.entity_id AND ir.entity_type = ?", "article_content").
		Joins("LEFT JOIN images img ON ir.image_id = img.id").
		Where("a.type = ? AND a.status = ?", "micro_post", "published").
		Group("a.id").
		Order("a.id DESC").
		Limit(10).
		Scan(&articles).Error
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, article := range articles {
		fmt.Printf("%#v\n", article)
	}
}
