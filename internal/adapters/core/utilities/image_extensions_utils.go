package utilities

import "strings"

func CheckImage(fileStr string) bool {
	imageExtensions := []string{".jpg", ".jpeg", ".jfif", ".pjpeg", ".pjp", ".avif", ".gif", ".png"}
	for _, imgExtension := range imageExtensions {
		if strings.HasSuffix(fileStr, imgExtension) {
			return true
		}
	}
	return false
}

func ExtractExtension(fileStr string) string {
	index := strings.LastIndex(fileStr, ".")
	if index == -1 {
		return ""
	}
	return fileStr[index:]
}
