package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	green := color.RGBA{0, 255, 0, 255}
	rectImg := image.NewRGBA(image.Rect(0, 0, 200, 200))


	red:=color.RGBA{200,30,30,255}


	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)
	for y :=0; y<200;y=y+20 {
	for x := 0; x < 200; x++ {
		rectImg.Set(x, y, red)
	}
	}
	
	for x:=0; x<200;x=x+20 {
		for y:=0; y<200;y++ {
			rectImg.Set(x, y, red)

		}


	}
	
	
	
	//y :=100
    //rectImg.Set(x, y, red)
	//for y :=0; y<200;y=y+10

	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
       defer file.Close()
	png.Encode(file, rectImg)
}