package main

import (
	"github.com/ethereum/go-ethereum/metrics"
	cli "gopkg.in/urfave/cli.v1"
)

var MetricsPrometheusEndpointFlag = cli.StringFlag{
	Name:  "metrics.prometheus.endpoint",
	Usage: "Prometheus API endpoint to report metrics to",
	Value: ":19090",
}

var (
	reg = metrics.NewRegistry()

	txCountSentMeter = metrics.NewRegisteredCounter("tx_count_sent", reg)
	txCountGotMeter  = metrics.NewRegisteredCounter("tx_count_got", reg)
	txTpsMeter       = metrics.NewRegisteredHistogram("tx_tps", reg, metrics.NewUniformSample(500))
)

func SetupPrometheus(ctx *cli.Context) {
	return
}
