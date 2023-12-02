package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/gbeltramo/go-23d/internal/load23d"
)

const width, height = 640, 480

func main() {

	// init image to full white
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	// NOTE iterating of iY index in reverse order so that the xy coords
	// NOTE in the 2D image correspond to the xy coordinates in my head
	for iX := 0; iX < width; iX++ {
		for iY := 1; iY < height+1; iY++ {
			img.Set(iX, height-iY, color.NRGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 255,
			})
		}
	}

	// Load points
	triang, err := load23d.LoadSTL("./data/bunny.stl")
	if err != nil {
		log.Fatal(err)
	}

	// Plot points on image
	for _, p := range triang.Vertices {
		// z perspective divide with z translation to put bunny in view
		// NOTE The correct equations are
		// pProjX := - p.X / (p.Z + zTranslation)
		// pProjY := - p.Y / (p.Z + zTranslation)

		// NOTE But we use the following because we need to rotate the
		// NOTE bunny in order to see it
		yTranslation := float32(20.0) // NOTE this is a trick
		zTranslation := float32(40.0) // NOTE this is a trick
		pProjX := p.X / (p.Y + zTranslation)
		pProjY := (p.Z - yTranslation) / (p.Y + zTranslation)

		// mapping to pixel space
		xPix := int(((pProjX + 1) / 2) * width)
		yPix := int(((pProjY + 1) / 2) * height)

		if xPix == width {
			xPix = width - 1
		}

		if yPix == 0 {
			yPix = 1
		}
		
		img.Set(xPix, height-yPix, color.NRGBA{
			R: 50 + uint8(yPix / 5), // uint8((iX + iY) & 255),
			G: 50, // uint8((iX + iY) << 1 & 255),
			B: 50 + uint8(xPix / 10), // uint8((iX + iY) << 2 & 255),
			A: 255,
		})
	}
	
	// save image
	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
