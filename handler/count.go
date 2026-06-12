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

func Count(keywords []string, index string) float64 {
	keyword := strings.Join(keywords, " ")
	dsl := fmt.Sprintf(Count_DSL, keyword)
	respESBody := requestESCount(index, dsl)
	return parseESCountRespBody(respESBody)
}

var Count_DSL = `
{
  "track_total_hits": true,
  "query": {
    "bool": {
      "must": [
	    {
	      "query_string": {
		    "query": "%s",
			"default_field": "message",
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
}`

type CountResp struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
	} `json:"hits"`
}

func requestESCount(index, dsl string) []byte {
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

func parseESCountRespBody(body []byte) float64 {
	var r CountResp
	if err := json.Unmarshal(body, &r); err != nil {
		return 0
	}
	return float64(r.Hits.Total.Value)
}
