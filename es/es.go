package es

import (
	"crypto/tls"
	"esAlert/config"
	"net/http"

	"github.com/elastic/go-elasticsearch/v7"
)

var ES *elasticsearch.Client

func init() {
	ES = newEsClient()
}

func newEsClient() *elasticsearch.Client {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Username:  config.Configs.ES.User,
		Password:  config.Configs.ES.Password,
		Addresses: config.Configs.ES.Hosts,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	})
	if err != nil {
		panic(err)
	}
	return es
}
