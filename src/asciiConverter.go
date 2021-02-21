package main

import (
	_ "image/jpeg"
	_ "image/png"

	"github.com/qeesung/image2ascii/convert"
)

//ConvertImageToASCII is a function that converts a given image to a ascii representation
//of that image
func ConvertImageToASCII(imageName string) string {
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 100
	convertOptions.FixedHeight = 40

	converter := convert.NewImageConverter()
	asciiImage := converter.ImageFile2ASCIIString(imageName, &convertOptions)
	return asciiImage
}
