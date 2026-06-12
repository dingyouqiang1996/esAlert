package handler

import (
	"context"
	"encoding/json"
	"esAlert/es"
	"fmt"
	"io"
	"log"
	"strings"
)

func Search(keywords []string, index string) float64 {
	keyword := strings.Join(keywords, " ")
	dsl := fmt.Sprintf(Search_DSL, keyword)
	respESBody := requestEsSearch(index, dsl) // 请求es
	return parseESSearachRespBody(respESBody) // 处理响应
}

var Search_DSL = `
{
  "track_total_hits": true,
  "query": {
    "bool": {
      "must": [
        {
          "query_string": {
            "query": "%s",
            "default_operator": "AND"
          }
        }
      ],
      "filter": [
        {
          "range": {
            "@timestamp": {
              "gte": "now-1m"
            }
          }
        }
      ]
    }
  },
  "size": 0
}
`

type SearchResp struct {
	Hits struct {
		Total struct {
			Value int64 `json:"value"`
		} `json:"total"`
	} `json:"hits"`
}

func requestEsSearch(index, dsl string) []byte {
	res, _ := es.ES.Search(
		es.ES.Search.WithContext(context.Background()),
		es.ES.Search.WithIndex(index),
		es.ES.Search.WithBody(strings.NewReader(dsl)),
	)
	defer res.Body.Close()
	log.Println(res)
	body, _ := io.ReadAll(res.Body)
	return body
}

func parseESSearachRespBody(body []byte) float64 {
	var resp struct {
		Hits struct {
			Total struct {
				Value int64 `json:"value"`
			} `json:"total"`
		} `json:"hits"`
	}
	json.Unmarshal(body, &resp)
	if resp.Hits.Total.Value > 0 {
		return 1
	}
	return 0
}
