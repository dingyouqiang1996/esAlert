package main

import (
	_ "esAlert/config"
	"esAlert/es"
	"esAlert/exporter"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

/*
  1. 要可配置 ✅️
  2. 告警发送到e家 ❌️
  3. 返回0或1 ❌️

  todo:
  1. xxl AND fail ✅️
  2. labels
  3. counter 类型
  4. histogram 类型
  5. collector: https://cloud.tencent.com/developer/article/2391335
*/

func main() {
	YouKnowForSearch()
	reg := prometheus.NewRegistry()
	export := exporter.NewExporter()
	reg.MustRegister(export)
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	log.Println("esAlert start...")
	http.ListenAndServe(":2112", nil)
}

func YouKnowForSearch() {
	res, err := es.ES.Info()
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println(res.String())
}
