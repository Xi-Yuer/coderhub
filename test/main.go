package main

import (
	"coderhub/shared/storage"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	cfg := &elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
		Username:  "c",
		Password:  "2214380963Wx!!",
	}
	c, err := storage.NewElasticSearchClient(cfg)
	if err != nil {
		return
	}
	ids, err := c.SearchByFields("users", map[string]interface{}{"user_name": "123456"})
	if err != nil {
		return
	}
	fmt.Printf("id type: %T\n", ids[0])
	fmt.Println(ids)
}
