// Go Api server
// @jeffotoni
// 2019-01-04

package main

import (
    "image"
    "image/jpeg"
    "os"

    graphics "code.google.com/p/graphics-go/graphics"
)

func main() {
    imagePath, _ := os.Open("./golang1.png")
    defer imagePath.Close()
    srcImage, _, _ := image.Decode(imagePath)

    // Dimension of new thumbnail 80 X 80
    dstImage := image.NewRGBA(image.Rect(0, 0, 100, 100))
    // Thumbnail function of Graphics
    graphics.Thumbnail(dstImage, srcImage)

    newImage, _ := os.Create("thumbnail.png")
    defer newImage.Close()
    jpeg.Encode(newImage, dstImage, &jpeg.Options{jpeg.DefaultQuality})
}
