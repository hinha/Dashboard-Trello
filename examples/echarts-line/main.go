package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"io"
	"math/rand"
	"os"
)

var (
	fruits      = []string{"Apple", "Banana", "Peach ", "Lemon", "Pear", "Cherry"}
)

func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(fruits); i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func lineBase() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic line example", Subtitle: "This is the subtitle."}),
	)

	line.SetXAxis(fruits).
		AddSeries("Category A", generateLineItems())
	return line
}

func lineShowLabel() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "title and label options",
			Subtitle: "go-echarts is an awesome chart library written in Golang",
			Link:     "https://github.com/go-echarts/go-echarts",
		}),
	)

	line.SetXAxis(fruits).
		AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
		)
	return line
}

func lineSmoothArea() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "smooth area"}),
	)

	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
			charts.WithAreaStyleOpts(opts.AreaStyle{
				Opacity: 0.2,
			}),
			charts.WithLineChartOpts(opts.LineChart{
				Smooth: true,
			}),
		)
	return line
}

func lineSmooth() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "smooth style",
		}),
	)

	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{
				Smooth: true,
			}),
		)
	return line
}

func lineMarkPoint() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "markpoint options",
		}),
	)

	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(
			charts.WithMarkPointNameTypeItemOpts(
				opts.MarkPointNameTypeItem{Name: "Maximum", Type: "max"},
				opts.MarkPointNameTypeItem{Name: "Average", Type: "average"},
				opts.MarkPointNameTypeItem{Name: "Minimum", Type: "min"},
			),
			charts.WithMarkPointStyleOpts(
				opts.MarkPointStyle{Label: &opts.Label{Show: true}}),
		)
	return line
}

func lineMulti() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithLegendOpts(opts.Legend{
			Data: []string{"TO DO", "DONE"},
		}),
		charts.WithTitleOpts(opts.Title{
			Title: "multi lines",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Trigger: "axis",
			Show: true,
		}),
		charts.WithDataZoomOpts(opts.DataZoom{Type: "inside", Start: 0, End: 100}, opts.DataZoom{Start: 0, End: 10}),
		charts.WithYAxisOpts(opts.YAxis{Type: "value"}),
		charts.WithXAxisOpts(opts.XAxis{Type: "category", Data: fruits}),
	)

	line.SetXAxis(fruits).
		AddSeries("TO DO", generateLineItems()).
		AddSeries("DONE", generateLineItems())
	return line
}





func main() {
	page := components.NewPage()
	page.AddCharts(
		lineShowLabel(),
		lineSmoothArea(),
		lineSmooth(),
		lineMarkPoint(),
		lineMulti(),
	)
	f, err := os.Create("examples/line.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
