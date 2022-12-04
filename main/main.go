package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/dongzeli95/tcb-openapi-go-sdk"
	"github.com/dongzeli95/tcb-openapi-go-sdk/component/database/query"
	"github.com/dongzeli95/tcb-openapi-go-sdk/config"
	"github.com/dongzeli95/tcb-openapi-go-sdk/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CLOUDBASE_URL = "https://tcb-api.tencentcloudapi.com"

	ERR_EMPTY_DOUBLE_TYPE = "The double type is empty."
	ERR_EMPTY_INT_TYPE    = "The int type is empty."
)

type dataResult struct {
	DataList []string `json:"list"`
}

type findResult struct {
	Data *dataResult `json:"data"`
}

func main() {
	log.Println("Holy shit")

	client := tcb.NewTcb(&config.Config{
		// EnvId:     DEV_ENV_ID,
		Timeout:   time.Duration(100) * time.Second,
		LogPrefix: "tcb",
		Debug:     true,
		// SecretId:  SECRET_ID,
		// SecretKey: SECRET_KEY,
	})

	data, err := client.GetDatabase().FindBySortOrder("videos", query.NewQuery(), 10, 0, primitive.M{}, util.OrderedMap{
		Order: []string{"weight", "created_at"},
		Map: map[string]interface{}{
			"weight":     -1,
			"created_at": -1,
		},
	})
	if err != nil {
		log.Println("err: " + err.Error())
	} else {
		var res *findResult
		err = json.Unmarshal([]byte(data), &res)
		if err != nil {
			log.Println("err: " + err.Error())
		} else {
			for _, str := range res.Data.DataList {
				log.Println(str)
			}
		}
	}
}
