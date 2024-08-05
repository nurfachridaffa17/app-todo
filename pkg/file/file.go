package file

import (
	"fmt"
	"os"
	"regexp"
)

func GetRootDirectory() string {
	projectDir := os.Getenv("PROJECT_DIR")
	projectName := regexp.MustCompile(`^(.*` + projectDir + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	return string(rootPath)
}

func GenerateUniqueFileName(prefix string, filename string) string {
	uniqueFilename := fmt.Sprintf("%s_%s.png", prefix, filename)
	return uniqueFilename
}
