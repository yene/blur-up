package main

import (
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"

	"encoding/base64"

	"github.com/nfnt/resize"
)

func main() {
	//flag.Parse()
	//path := flag.Arg(0)

	// open "test.jpg"
	file, err := os.Open("original.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(30, 0, img, resize.Lanczos3)

	out, err := os.Create("original_resized.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)

	dat, _ := ioutil.ReadFile("original_resized.jpg")
	str := base64.StdEncoding.EncodeToString(dat)
	log.Println("url('data:image/jpeg;base64," + str + "')")

	os.Remove("original_resized.jpg")
}
