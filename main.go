package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
)

func readOverlay(overlayFilepath string, height uint) image.Image {

	//Read in the overlay file
	inOverlay, err := os.Open(overlayFilepath)
	if err != nil {
		log.Fatalln(err)
	}
	defer inOverlay.Close()

	//Determine file type
	extension := filepath.Ext(overlayFilepath)

	//Decode file based on type
	var decodedOverlay image.Image

	if extension == ".jpg" {
		decodedOverlay, err = jpeg.Decode(inOverlay)
	} else if extension == ".png" {
		decodedOverlay, err = png.Decode(inOverlay)
	} else {
		return nil
	}

	//Log out any decoding errors
	if err != nil {
		log.Fatalln(err)
	}

	//Resize overlay image
	resizedOverlay := resize.Resize(0, height, decodedOverlay, resize.Lanczos3)
	return resizedOverlay
}

func buildNewParrot(decodedGif *gif.GIF, overlayImage image.Image, numFrames int) *gif.GIF {
	//For each frame
	for x := 0; x < numFrames; x++ {
		t := float64(x) / float64(numFrames)
		ellipseOffset := math.Pi / 4

		//Get decoded gif frame
		frame := decodedGif.Image[x]

		//Create new frame
		newFrame := image.NewPaletted(frame.Bounds(), frame.Palette)

		//Calculate the position for the offset image
		offset := image.Pt(int(40+26*math.Cos(t*2*math.Pi+ellipseOffset)), int(35+-5*math.Sin(t*2*math.Pi+ellipseOffset)))

		//Write frame from decoded gif
		draw.Draw(newFrame, frame.Bounds(), frame, image.ZP, draw.Src)

		//Add overlay image to the frame
		draw.Draw(newFrame, frame.Bounds().Add(offset), overlayImage, image.ZP, draw.Over)

		//Overwrite decoded gif with the newly created frame
		decodedGif.Image[x] = newFrame
	}

	return decodedGif
}

func main() {
	//Get filepath argument
	var overlayFilepath string
	if len(os.Args) >= 2 {
		overlayFilepath = os.Args[1]
	} else {
		log.Fatalln("ERROR: No overlay filepath provided!")
	}

	//Open the base parrot gif
	baseGif, err := os.Open("./parrot.gif")
	if err != nil {
		log.Fatalln(err)
	}
	defer baseGif.Close()

	//Decode base parrot gif into frames
	decodedGif, err := gif.DecodeAll(baseGif)
	if err != nil {
		log.Fatalln(err)
	}

	//Decode the overlay image
	overlayImage := readOverlay(overlayFilepath, 70)

	//Get number of frames in the decoded gif
	numFrames := len(decodedGif.Image)

	//Build new parrot gif
	parrotGif := buildNewParrot(decodedGif, overlayImage, numFrames)

	//Create output file
	outGif, err := os.Create("./parrot_out.gif")
	if err != nil {
		log.Fatalln(err)
	}
	defer outGif.Close()

	//Encode frames into output gif and write it out
	err = gif.EncodeAll(outGif, parrotGif)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("A new party parrot has been born!")
	}
}
