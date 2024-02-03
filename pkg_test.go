// Package tests.
//
// José Almuedo [https://github.com/jalmuedo]
// This code is in the public domain.
package main

import (
	"testing"

	"github.com/gookit/color"
)

func TestColorize(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		chunks   []string
		expected string
	}{
		{
			name:     "Colorize single chunk",
			text:     "The quick brown fox",
			chunks:   []string{"quick"},
			expected: "The \x1b[41mquick\x1b[0m brown fox",
		},
		{
			name:     "Colorize multiple chunks",
			text:     "quick quick quick",
			chunks:   []string{"quick"},
			expected: "\x1b[41mquick\x1b[0m \x1b[41mquick\x1b[0m \x1b[41mquick\x1b[0m",
		},
		{
			name:     "No colorization",
			text:     "The lazy dog",
			chunks:   []string{"quick"},
			expected: "The lazy dog",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := colorize(tt.text, tt.chunks)
			if result != tt.expected {
				t.Errorf("Expected: %s, Got: %s", tt.expected, result)
			}
		})
	}
}

func TestReplaceTxt(t *testing.T) {
	tests := []struct {
		name       string
		text       string
		chunk      string
		colorChunk color.Color
		expected   string
	}{
		{
			name:       "Replace single occurrence",
			text:       "The quick brown fox",
			chunk:      "quick",
			colorChunk: color.FgGreen,
			expected:   "The \x1b[32mquick\x1b[0m brown fox",
		},
		{
			name:       "Replace multiple occurrences",
			text:       "quick quick quick",
			chunk:      "quick",
			colorChunk: color.FgRed,
			expected:   "\x1b[31mquick\x1b[0m \x1b[31mquick\x1b[0m \x1b[31mquick\x1b[0m",
		},
		{
			name:       "No replacement",
			text:       "The lazy dog",
			chunk:      "quick",
			colorChunk: color.FgBlue,
			expected:   "The lazy dog",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := replaceTxt(tt.text, tt.chunk, tt.colorChunk)
			if result != tt.expected {
				t.Errorf("Expected: %s, Got: %s", tt.expected, result)
			}
		})
	}
}
