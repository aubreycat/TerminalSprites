package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"time"

	"github.com/fatih/color"
	"github.com/nfnt/resize"
)

func main() {
	folderPath := "sprites"

	speed := flag.Int("speed", 500, "speed :3")
	loop := flag.Bool("loop", false, "loop! ?!?!")
	flag.Parse()

	for {
		err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				img, err := loadImage(path)
				if err != nil {
					fmt.Printf("D: it no worky :( %s| %v\n", path, err)
					return nil
				}

				resizedImg := resize.Resize(40, 20, img, resize.Lanczos3)

				clear()
				displayImage(resizedImg)
				time.Sleep(time.Duration(*speed) * time.Millisecond)
			}

			return nil
		})

		if err != nil {
			fmt.Printf("nuh uh directory no no work :( | %v\n", err)
			break
		}

		if !*loop {
			break
		}
	}
}

func loadImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func clear) {
	fmt.Print("\033[H\033[2J")
}

func displayImage(img image.Image) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			greyscale := (r + g + b) / 3

			switch {
			case greyscale < 0x5555:
				color.New(color.BgBlack).Print(" ")
			case greyscale < 0xAAAA:
				color.New(color.BgWhite).Print(" ")
			case greyscale < 0xFFFF:
				color.New(color.BgHiBlack).Print(" ")
			default:
				color.New(color.BgHiWhite).Print(" ")
			}
		}
		fmt.Println()
	}
}