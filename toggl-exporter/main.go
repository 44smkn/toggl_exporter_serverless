package main

import (
	"bytes"
	"os"

	"github.com/44smkn/toggl_exporter/pkg/config"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"
	"github.com/prometheus/common/promlog"
	"github.com/prometheus/common/version"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)
	exporter := config.InitExporter(promlogConfig, logger)
	prometheus.MustRegister(exporter)
	prometheus.MustRegister(version.NewCollector("toggl_exporter"))

	mfs, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		level.Error(logger).Log("msg", err.Error())
		os.Exit(1)
	}

	buffer := &bytes.Buffer{}
	enc := expfmt.NewEncoder(buffer, expfmt.FmtText)
	for _, mf := range mfs {
		enc.Encode(mf)
	}

	return events.APIGatewayProxyResponse{
		Body:       buffer.String(),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
