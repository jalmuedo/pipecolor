// Utilities to color some words in a text string.
//
// José Almuedo [https://github.com/jalmuedo]
// This code is in the public domain.
package main

import (
	"regexp"

	"github.com/gookit/color"
)

func replaceCaseInsensitive(text string, chunk string, colorChunk string) string {
	re := regexp.MustCompile(`(?i)` + chunk)
	return re.ReplaceAllString(text, colorChunk)
}

func colorizeChunks(text string, chunks []string) string {

	for _, chunk := range chunks {
		chunkWithColor := color.BgGreen.Sprintf("%s", chunk)
		text = replaceCaseInsensitive(text, chunk, chunkWithColor)
	}
	return text
}
