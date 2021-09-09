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
	itemCntPie = 4
	seasons    = []string{"Spring", "Summer", "Autumn ", "Winter"}
)

func generatePieItems() []opts.PieData {
	items := make([]opts.PieData, 0)
	for i := 0; i < itemCntPie; i++ {
		items = append(items, opts.PieData{Name: seasons[i], Value: rand.Intn(100)})
	}
	return items
}

func pieShowLabel() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithTooltipOpts(opts.Tooltip{Trigger: "item", Show: true}),
	)

	pie.AddSeries("pie", generatePieItems()).
		SetSeriesOptions(
			charts.WithSunburstOpts(opts.SunburstChart{Animation: true, SelectedMode: true}),
			charts.WithLabelOpts(opts.Label{Show: true, Formatter: "{b}: {c}"}),
			charts.WithEmphasisOpts(opts.Emphasis{ItemStyle: &opts.ItemStyle{}}),
		)
	return pie
}


func main() {

	page := components.NewPage()
	page.AddCharts(
		pieShowLabel(),
	)
	f, err := os.Create("examples/pie.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
