package exporter

import (
	"esAlert/config"
	"esAlert/handler"
	"sort"

	"github.com/prometheus/client_golang/prometheus"
)

type Exporter struct {
	gaugeDesc *prometheus.Desc
}

type Collector interface {
	Describe(chan<- *prometheus.Desc)
	Collect(chan<- prometheus.Metric)
}

var labelNames []string

func init() {
	set := make(map[string]bool)
	for _, m := range config.Configs.Metrics {
		for k := range m.Labels {
			set[k] = true
		}
	}
	labelNames = []string{"index"}
	for k := range set {
		labelNames = append(labelNames, k)
	}
	sort.Strings(labelNames[1:])
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.gaugeDesc
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	for _, m := range config.Configs.Metrics {
		vals := make([]string, len(labelNames))
		vals[0] = m.Index
		for i, n := range labelNames[1:] {
			vals[i+1] = m.Labels[n]
		}
		ch <- prometheus.MustNewConstMetric(
			e.gaugeDesc,
			prometheus.GaugeValue,
			handler.Count(m.Keywords, m.Index),
			vals...,
		)
	}
}

func NewExporter() *Exporter {
	return &Exporter{
		gaugeDesc: prometheus.NewDesc(
			"xxljob",
			"xxljob",
			labelNames,
			prometheus.Labels{},
		),
	}
}
