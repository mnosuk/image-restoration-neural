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

const currentFileCount = 17

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
func saveImageFrom1DArray(img []float64) {
	gray := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{100, 100}})
	indexCount := 0
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			grayColor := (color.Gray{uint8(math.Ceil(img[indexCount] * 255))})
			gray.Set(x, y, grayColor)
			indexCount++
		}
	}

	outfilename := "./answer/test.png"
	outfile, err := os.Create(outfilename)
	if err != nil {
		// replace this with real error handling
		panic(err.Error())
	}
	defer outfile.Close()
	png.Encode(outfile, gray)
}
func main() {
	/* create some random noise from ./input directory to ./output directory */
	// buildTrainingSet()

	inputImage0 := get1DArrayImage("./output/0.png")
	outputImage0 := get1DArrayImage("./input/0.png")

	inputImage1 := get1DArrayImage("./output/1.png")
	outputImage1 := get1DArrayImage("./input/1.png")

	inputImage2 := get1DArrayImage("./output/2.png")
	outputImage2 := get1DArrayImage("./input/2.png")

	inputImage3 := get1DArrayImage("./output/3.png")
	outputImage3 := get1DArrayImage("./input/3.png")

	inputImage4 := get1DArrayImage("./output/4.png")
	outputImage4 := get1DArrayImage("./input/4.png")

	inputImage5 := get1DArrayImage("./output/5.png")
	outputImage5 := get1DArrayImage("./input/5.png")

	inputImage6 := get1DArrayImage("./output/6.png")
	outputImage6 := get1DArrayImage("./input/6.png")

	inputImage7 := get1DArrayImage("./output/7.png")
	outputImage7 := get1DArrayImage("./input/7.png")

	inputImage8 := get1DArrayImage("./output/8.png")
	outputImage8 := get1DArrayImage("./input/8.png")

	inputImage9 := get1DArrayImage("./output/9.png")
	outputImage9 := get1DArrayImage("./input/9.png")

	inputImage10 := get1DArrayImage("./output/10.png")
	outputImage10 := get1DArrayImage("./input/10.png")

	inputImage11 := get1DArrayImage("./output/11.png")
	outputImage11 := get1DArrayImage("./input/11.png")

	inputImage12 := get1DArrayImage("./output/12.png")
	outputImage12 := get1DArrayImage("./input/12.png")

	inputImage13 := get1DArrayImage("./output/13.png")
	outputImage13 := get1DArrayImage("./input/13.png")

	inputImage14 := get1DArrayImage("./output/14.png")
	outputImage14 := get1DArrayImage("./input/14.png")

	inputImage15 := get1DArrayImage("./output/15.png")
	outputImage15 := get1DArrayImage("./input/15.png")

	inputImage16 := get1DArrayImage("./output/16.png")
	outputImage16 := get1DArrayImage("./input/16.png")
	fmt.Println("load training set done")
	patterns := [][][]float64{
		{inputImage0, outputImage0},
		{inputImage1, outputImage1},
		{inputImage2, outputImage2},
		{inputImage3, outputImage3},
		{inputImage4, outputImage4},
		{inputImage5, outputImage5},
		{inputImage6, outputImage6},
		{inputImage7, outputImage7},
		{inputImage8, outputImage8},
		{inputImage9, outputImage9},
		{inputImage10, outputImage10},
		{inputImage11, outputImage11},
		{inputImage12, outputImage12},
		{inputImage13, outputImage13},
		{inputImage14, outputImage14},
		{inputImage15, outputImage15},
		{inputImage16, outputImage16},
	}

	// // Training Phase
	start := time.Now()
	ff := &gobrain.FeedForward{}
	ff.Init(10000, 40, 10000)
	ff.Train(patterns, 1000, 0.4, 0.2, false)
	elapsed := time.Since(start)
	fmt.Printf("training time : %+v\n", elapsed)

	// Testing Phase
	inputs := inputImage0
	answer := ff.Update(inputs)

	saveImageFrom1DArray(answer)
	fmt.Println("image saved")

}
