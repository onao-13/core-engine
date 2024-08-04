package utils

import (
	"core-engine/internal/common"
	"strings"
)

func IsMedia(filename string, format string) bool {
	return strings.Contains(filename, format)
}

func IsImage(filename string) bool {
	return IsMedia(filename, common.MediaFormatPNG) ||
		IsMedia(filename, common.MediaFormatJPEG) ||
		IsMedia(filename, common.MediaFormatJPG)
}
