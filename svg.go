package main

import "github.com/xyproto/tinysvg"

// Generate a new SVG image
func svgImage() []byte {
	document, svg := tinysvg.NewTinySVG(128, 64)
	svg.Describe("Hello SVG")

	// x, y, radius, color
	svg.Circle(30, 10, 5, "red")
	svg.Circle(110, 30, 2, "green")
	svg.Circle(80, 40, 7, "blue")

	// x, y, font size, font family, text and color
	svg.Text(3, 60, 6, "Courier", "There will be cake", "#394851")

	return document.Bytes()
}
