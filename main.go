package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"os"

	"github.com/disintegration/imaging"
)

func main() {
	// Command-line arguments
	inputPath := flag.String("input", "", "Path to the input image")
	outputPath := flag.String("output", "output.jpg", "Path to save the cropped image")
	x0 := flag.Int("x0", 0, "X coordinate for the top-left corner of the crop area")
	y0 := flag.Int("y0", 0, "Y coordinate for the top-left corner of the crop area")
	width := flag.Int("width", 100, "Width of the crop area")
	height := flag.Int("height", 100, "Height of the crop area")

	// Parse the arguments
	flag.Parse()

	// Check if the input path is provided
	if *inputPath == "" {
		fmt.Println("Error: Please specify the input image path using the -input argument")
		flag.Usage()
		os.Exit(1)
	}

	// Open the input image
	src, err := imaging.Open(*inputPath)
	if err != nil {
		log.Fatalf("Error opening the image: %v", err)
	}

	// Define the crop area
	rect := image.Rect(*x0, *y0, *x0+*width, *y0+*height)

	// Crop the image
	cropped := imaging.Crop(src, rect)

	// Save the cropped image
	err = imaging.Save(cropped, *outputPath)
	if err != nil {
		log.Fatalf("Error saving the image: %v", err)
	}

	fmt.Printf("Image successfully cropped and saved as %s\n", *outputPath)
}
