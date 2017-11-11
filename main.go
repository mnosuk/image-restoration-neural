package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/goml/gobrain"
)

const currentFileCount = 4

func buildTrainingSet() {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < currentFileCount; i++ {
		iStr := strconv.Itoa(i)
		filename := "./input/" + iStr + ".png"
		infile, err := os.Open(filename)
		if err != nil {
			panic(err.Error())
		}
		defer infile.Close()

		src, _, err := image.Decode(infile)
		if err != nil {
			panic(err.Error())
		}

		// Create a new grayscale image
		bounds := src.Bounds()
		w, h := bounds.Max.X, bounds.Max.Y
		gray := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{w, h}})
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				oldColor := src.At(x, y)
				r, g, b, _ := oldColor.RGBA()
				avg := 0.2125*float64(r) + 0.7154*float64(g) + 0.0721*float64(b)
				grayColor := (color.Gray{uint8(255)})
				if perfectRand := r1.Intn(10); perfectRand != 1 {
					grayColor = (color.Gray{uint8(math.Ceil(avg))})
				} else {
					grayColor = (color.Gray{uint8(255)})
				}
				gray.Set(x, y, grayColor)
			}
		}

		// Encode the grayscale image to the output file
		outfilename := "./output/" + iStr + ".png"
		outfile, err := os.Create(outfilename)
		if err != nil {
			// replace this with real error handling
			panic(err.Error())
		}
		defer outfile.Close()
		png.Encode(outfile, gray)
	}
}
func get1DArrayImage(imagePath string) []float64 {
	output := make([]float64, 10000)

	infile, err := os.Open(imagePath)
	if err != nil {
		panic(err.Error())
	}
	defer infile.Close()
	src, _, err := image.Decode(infile)
	if err != nil {
		panic(err.Error())
	}
	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	indexCount := 0
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			oldColor := src.At(x, y)
			r, g, b, _ := oldColor.RGBA()
			// fmt.Println(r)

			avg := 0.2125*float64(r) + 0.7154*float64(g) + 0.0721*float64(b)
			// fmt.Println(avg)
			output[indexCount] = float64(uint8(math.Ceil(avg))) / 255.0
			indexCount++
		}
	}
	return output
}
func main() {
	/* create some random noise from ./input directory to ./output directory */
	buildTrainingSet()
	inputImage0 := get1DArrayImage("./output/0.png")
	fmt.Println("load 0.png done")
	inputImage1 := get1DArrayImage("./output/1.png")
	fmt.Println("load 1.png done")
	inputImage2 := get1DArrayImage("./output/2.png")
	fmt.Println("load 2.png done")
	inputImage3 := get1DArrayImage("./output/3.png")
	fmt.Println("load 3.png done")
	fmt.Printf("%+v\n", inputImage0)
	patterns := [][][]float64{
		{inputImage0, {0, 0, 0, 0, 0, 0, 0}},
		{inputImage1, {0, 1, 1, 0, 0, 0, 0}},
		{inputImage2, {1, 1, 0, 1, 1, 0, 1}},
		{inputImage3, {1, 1, 1, 1, 0, 0, 1}},
	}

	// // Training Phase
	ff := &gobrain.FeedForward{}
	ff.Init(10000, 10000, 10000)
	ff.Train(patterns, 1000, 0.6, 0.4, false)

	// Testing Phase
	// inputs := []float64{1}
	// answer := ff.Update(inputs)
}
