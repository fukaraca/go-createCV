package main

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"log"
	"mime/multipart"
)

//resizer function resize and saves the form input file to filepath(./web/img/temp) with the given filename.
func resizer(file *multipart.File, filepath, filename string) (error, *image.NRGBA) {
	img, err := imaging.Decode(*file)
	if err != nil {
		log.Println("image couldn't be decoded:")
		return err, nil
	}

	//aspect ratio
	srcX := float32(img.Bounds().Size().X)
	srcY := float32(img.Bounds().Size().Y)
	if srcX > srcY {
		srcY = srcY * (res / srcX)
		srcX = res
	} else {
		srcX = srcX * (res / srcY)
		srcY = res
	}

	// Resize srcImage to size = highestRes aspect ratio using the Lanczos filter.
	dstImage8060 := imaging.Resize(img, int(srcX), int(srcY), imaging.Lanczos)
	fmt.Println(dstImage8060.Pix)
	return nil, dstImage8060
}
