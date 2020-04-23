package utility

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// GetSrcPath is the directory path for template files
func GetSrcPath(fileName string) string {
	return fmt.Sprintf("%s%s", srcPathString, fileName)
}

// GetDestPath is the directory path for output files
func GetDestPath(fileName string) string {
	return fmt.Sprintf("%s%s", destPathString, fileName)
}

// TrimAllSpaces trims the space around any text
func TrimAllSpaces(text string) string {
	replacer := strings.NewReplacer(" ", "", "\n", "")
	return replacer.Replace(text)
}

// GetFileName takes the title as the input and converts it into a valid file name
func GetFileName(title string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	title = reg.ReplaceAllString(title, "")
	name := TrimAllSpaces(strings.ToLower(title))
	return fmt.Sprintf("%s.html", name)
}
