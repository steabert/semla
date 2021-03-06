package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func main() {
	err := mandelbrot(complex(-2.0, -1.0), complex(3.0, 2.0), 500.0)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	} else {
		fmt.Println("Done!")
	}
}

func escapeTime(c complex128, limit uint8) uint8 {
	var z complex128 = complex(0.0, 0.0)
	var i uint8
	for i = 0; i < limit; i++ {
		z = z*z + c
		if real(z)*real(z)+imag(z)*imag(z) > 4.0 {
			return i
		}
	}
	return limit
}

func mandelbrot(
	z_offset complex128,
	z_range complex128,
	res float64,
) error {
	var width uint64 = uint64(res * real(z_range))
	var height uint64 = uint64(res * imag(z_range))
	var matrix []uint8 = make([]uint8, width*height)
	var row uint64
	var col uint64
	for row = 0; row < height; row++ {
		for col = 0; col < width; col++ {
			c := z_offset + complex(float64(col)/res, float64(row)/res)
			matrix[col+width*row] = uint8(escapeTime(c, 100) * 2)
		}
	}

	img := image.Gray{
		Pix:    matrix,
		Stride: int(width),
		Rect: image.Rectangle{
			Min: image.Point{X: 0, Y: 0},
			Max: image.Point{X: int(width), Y: int(height)},
		},
	}

	f, err := os.Create("mandelbrot.png")
	if err != nil {
		fmt.Println("Error: could not create image file: ", err.Error())
		return err
	}
	defer f.Close()

	err = png.Encode(f, &img)
	if err != nil {
		fmt.Println("Error: could not create image file: ", err.Error())
		return err
	}
	return nil
}
