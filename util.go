package main

import (
	"crypto/rand"
	"fmt"
	"github.com/disintegration/imaging"
	"log"
	"mime/multipart"
)

//resizeAndSave function resize and saves the form input file to filepath(./web/img/temp) with the given filename.
func resizeAndSave(file *multipart.File, filepath, filename string) error {
	img, err := imaging.Decode(*file)
	if err != nil {
		log.Println("image couldn't be decoded:")
		return err
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
	return imaging.Save(dstImage8060, filepath+filename)
}

//randomString function generates and returns n bytes sized string
func randomString(n int) string {
	randSuffix := make([]byte, n)
	if _, err := rand.Read(randSuffix); err != nil {
		log.Println("random string generation failed:", err)
	}
	return fmt.Sprintf("%X", randSuffix)
}
