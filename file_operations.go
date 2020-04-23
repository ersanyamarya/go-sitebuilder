package main

import (
	"encoding/json"
	"io/ioutil"

	utility "./utility"
)

func getFileAsString(path string) (string, error) {
	chapterHTMLFile, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(chapterHTMLFile), nil
}

func writeHTMLFile(path string, content string) error {
	toWrite := []byte(content)
	err := ioutil.WriteFile(utility.GetDestPath(path), toWrite, 0644)
	if err != nil {
		return err
	}
	return nil
}

// GetChaptersFromFile gets a chapterPage from the filename
func (pages *Page) GetChaptersFromFile(path string) error {
	chaptersFile, err := ioutil.ReadFile(utility.GetJSONSrcPath(path))
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
