package httpadapter

import "strings"

func checkImage(fileStr string) bool {
	imageExtensions := []string{".jpg", ".jpeg", ".jfif", ".pjpeg", ".pjp", ".avif", ".gif", ".png"}
	for _, imgExtension := range imageExtensions {
		if strings.HasSuffix(fileStr, imgExtension) {
			return true
		}
	}
	return false
}
