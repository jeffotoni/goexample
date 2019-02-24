// Go Api server
// @jeffotoni
// 2019-02-22

package main

import (
    "log"

    "gopkg.in/gographics/imagick.v2/imagick"
)

func main() {

    pdfName := "ref.pdf"
    imageName := "test.jpg"

    if err := ConvertPdfToJpg(pdfName, imageName); err != nil {
        log.Fatal(err)
    }
}

// ConvertPdfToJpg will take a filename of a pdf file and convert the file into an
// image which will be saved back to the same location. It will save the image as a
// high resolution jpg file with minimal compression.
func ConvertPdfToJpg(pdfName string, imageName string) error {

    // Setup
    imagick.Initialize()
    defer imagick.Terminate()

    mw := imagick.NewMagickWand()
    defer mw.Destroy()

    // Must be *before* ReadImageFile
    // Make sure our image is high quality
    if err := mw.SetResolution(300, 300); err != nil {
        return err
    }

    // Load the image file into imagick
    if err := mw.ReadImage(pdfName); err != nil {
        return err
    }

    // Must be *after* ReadImageFile
    // Flatten image and remove alpha channel, to prevent alpha turning black in jpg
    if err := mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_FLATTEN); err != nil {
        return err
    }

    // Set any compression (100 = max quality)
    if err := mw.SetCompressionQuality(95); err != nil {
        return err
    }

    // Select only first page of pdf
    mw.SetIteratorIndex(0)

    // Convert into JPG
    if err := mw.SetFormat("jpg"); err != nil {
        return err
    }

    // Save File
    return mw.WriteImage(imageName)
}
