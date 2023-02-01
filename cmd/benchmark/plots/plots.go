package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	grpc "github.com/anilsenay/go-http-vs-grpc/cmd/benchmark/grpc/benchmark"
	http "github.com/anilsenay/go-http-vs-grpc/cmd/benchmark/http/benchmark"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func main() {
	var requestNumber int
	var concurrency int = 1

	argsWithoutProg := os.Args[1:]

	r, err := strconv.Atoi(argsWithoutProg[0])
	if err != nil {
		panic(err)
	}
	requestNumber = r

	if len(argsWithoutProg) >= 2 {
		c, err := strconv.Atoi(argsWithoutProg[1])
		if err != nil {
			panic(err)
		}
		concurrency = c
	}

	http_results := http.RunBenchmark(requestNumber, concurrency)
	fmt.Println("http requests are done. Waiting for 5 seconds before starting grpc requests")
	time.Sleep(5 * time.Second)
	grpc_results := grpc.RunBenchmark(requestNumber, concurrency)

	grpc_xaxis, grpc_series := generateSeries(grpc_results)
	http_xaxis, http_series := generateSeries(http_results)

	drawPlot("grpc_plot", grpc_xaxis, grpc_series)
	drawPlot("http_plot", http_xaxis, http_series)

	drawPlots("grpc_vs_http", grpc_xaxis,
		Series{
			Name: "gRPC",
			Data: grpc_series,
		},
		Series{
			Name: "HTTP",
			Data: http_series,
		},
	)
}

func generateSeries(timeSeries []time.Duration) ([]int, []opts.LineData) {
	xaxis := make([]int, 0)
	items := make([]opts.LineData, 0)
	for i, v := range timeSeries {
		xaxis = append(xaxis, (i))
		items = append(items, opts.LineData{Name: "", Value: v})
	}
	return xaxis, items
}

func drawPlot(filename string, xaxis []int, items []opts.LineData) {
	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeInfographic,
		}),
	)

	line.SetXAxis(xaxis).
		AddSeries("Requests", items).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	f, _ := os.Create(filename + ".html")
	_ = line.Render(f)
}

type Series struct {
	Name string
	Data []opts.LineData
}

func drawPlots(filename string, xaxis []int, series ...Series) {
	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeInfographic,
		}),
		charts.WithLegendOpts(opts.Legend{
			Show: true,
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type: "inside",
		}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show:  true,
			Right: "20%",
			Feature: &opts.ToolBoxFeature{
				SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
					Show:  true,
					Type:  "png",
					Title: "Download as PNG",
				},
				DataView: &opts.ToolBoxFeatureDataView{
					Show:  true,
					Title: "DataView",
					Lang:  []string{"data view", "turn off", "refresh"},
				},
			}},
		),
		charts.WithXAxisOpts(opts.XAxis{AxisLabel: &opts.AxisLabel{Show: true, Formatter: "{value} th request"}}),
		charts.WithYAxisOpts(opts.YAxis{
			AxisLabel: &opts.AxisLabel{Show: true, Formatter: "{value} ns"},
		}),
	)

	line = line.SetXAxis(xaxis)
	for _, item := range series {
		line = line.AddSeries(item.Name, item.Data)
	}

	line.
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{Smooth: true}),
		)
	f, _ := os.Create(filename + ".html")
	_ = line.Render(f)
}
