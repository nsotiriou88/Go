package main

import (
	"image/color"
	"math/rand"

	"github.com/gonum/stat/distuv"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	data := make([]float64, 100)
	// Draw some random values from the standard
	// normal distribution.
	// rand.Seed(int64(0))
	for i := range data {
		data[i] = rand.NormFloat64()
	}
	// v := make(plotter.Values, 10)
	v := plotter.Values(data)

	// Make a plot and set its title.
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Histogram"

	// Create a histogram of our values drawn
	// from the standard normal.
	h, err := plotter.NewHist(v, 16)
	if err != nil {
		panic(err)
	}
	// Normalize the area under the histogram to
	// sum to one.
	h.Normalize(1)
	p.Add(h)

	// The normal distribution function
	norm := plotter.NewFunction(distuv.UnitNormal.Prob)
	norm.Color = color.RGBA{R: 255, A: 255}
	norm.Width = vg.Points(2)
	p.Add(norm)
	// fmt.Println(norm)
	// fmt.Printf("%T\n", norm)

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "hist.png"); err != nil {
		panic(err)
	}
}
