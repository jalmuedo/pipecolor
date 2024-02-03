// Utilities to color some words in a text string.
//
// José Almuedo [https://github.com/jalmuedo]
// This code is in the public domain.
package main

import (
	"regexp"

	"github.com/gookit/color"
)

func replaceTxt(text string, chunk string, colorChunk color.Color) string {
	re := regexp.MustCompile(`(?i)` + chunk)
	return re.ReplaceAllString(text, colorChunk.Render(chunk))
}

func colorize(text string, chunks []string) string {
	colors := []color.Color{
		color.BgRed,
		color.BgGreen,
		color.BgCyan,
		color.BgHiMagenta,
	}

	for i, chunk := range chunks {
		chunkWithColor := colors[i%len(colors)]
		text = replaceTxt(text, chunk, chunkWithColor)
	}
	return text
}
