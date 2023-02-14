// Go Api server
// @jeffotoni
// 2019-01-04

package main

import (
	"fmt"
	"log"

	"github.com/disintegration/imaging"
)

func main() {
	// Open a test image.
	srcImage, err := imaging.Open("./golang1.png")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
		return
	}
	// Resize srcImage to size = 128x128px using the Lanczos filter.
	dstImage128 := imaging.Resize(srcImage, 100, 100, imaging.Lanczos)
	// dstImageFill := imaging.Fill(srcImage, 100, 100, imaging.Center, imaging.Lanczos)

	err = imaging.Save(dstImage128, "golang100x100.png")

	if err != nil {
		log.Fatalf("failed to load image: %v", err)
		return
	}

	fmt.Println("success imagem")
	//fmt.Println(dstImage128)
}
