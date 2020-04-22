package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func getSrcPath(path string) string {
	return fmt.Sprintf("./files/%s", path)
}
func getDestPath(path string) string {
	return fmt.Sprintf("./docs/chapters/%s", path)
}
func trimAllSpaces(text string) string {
	replacer := strings.NewReplacer(" ", "", "\n", "")
	return replacer.Replace(text)
}

func getFileName(title string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	title = reg.ReplaceAllString(title, "")
	name := trimAllSpaces(strings.ToLower(title))
	return fmt.Sprintf("%s.html", name)

}

func getFileAsString(path string) (string, error) {
	chapterHTMLFile, err := ioutil.ReadFile(getSrcPath(path))
	if err != nil {
		return "", err
	}
	return string(chapterHTMLFile), nil
}

func writeHTMLFile(path string, content string) error {
	toWrite := []byte(content)
	err := ioutil.WriteFile(getDestPath(path), toWrite, 0644)
	if err != nil {
		return err
	}
	return nil
}

// GetChaptersFromFile gets a chapterPage from the filename
func (pages *Page) GetChaptersFromFile(path string) error {
	chaptersFile, err := ioutil.ReadFile(getSrcPath(path))
	if err != nil {
		return err
	}
	var chapters []Chapter
	json.Unmarshal(chaptersFile, &chapters)
	for _, chapter := range chapters {
		pages.Content = append(pages.Content, chapter)
	}
	return nil
}
