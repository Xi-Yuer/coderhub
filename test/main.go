package main

import (
	"coderhub/model"
	"coderhub/shared/storage"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	cfg := &elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
		Username:  "elastic",
		Password:  "2214380963Wx!!",
	}
	c, err := storage.NewElasticSearchClient(cfg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 构造搜索查询
	ids, err := c.SearchByFields("academic_navigators", map[string]interface{}{
		"user_id":   "",
		"content":   "",
		"education": "",
		"major":     "",
		"school":    "",
	})
	
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("id type: %T\n", ids)
	fmt.Println(ids)

	// 数据库查询
	db := storage.NewGorm()

	var List []*model.AcademicNavigator // 声明一个指针切片
	if err := db.Model(&model.AcademicNavigator{}).Where("id IN (?)", ids).Find(&List).Error; err != nil {
		fmt.Println("Error querying database:", err)
		return
	}

	// 输出查询结果
	if len(List) > 0 {
		fmt.Println(List[0].Content)
	} else {
		fmt.Println("No results found")
	}
}