package main

import (
	"fmt"
	"strings"
)

// Paragraph is the content inside chapter
type Paragraph struct {
	Image string `json:"image"`
	Text  string `json:"text"`
	Title string `json:"title"`
}

// Chapter is a Episode in a Learning Module
type Chapter struct {
	Book          string      `json:"book"`
	ChapterNumber int         `json:"chapterNumber"`
	Title         string      `json:"title"`
	SubTitle      string      `json:"subTitle"`
	Description   string      `json:"description"`
	FeaturedImage string      `json:"featuredImage"`
	Content       []Paragraph `json:"content"`
}

func (chapter Chapter) getImageURL() string {
	return fmt.Sprintf("../../doc/assets/images/%s", chapter.FeaturedImage)
}
func (paragraph Paragraph) getImageURL() string {
	return fmt.Sprintf("../../doc/assets/images/%s", paragraph.Image)
}

func (chapter Chapter) getParagraphs() string {
	paraHTML := `<div class="paragraphs"><ul>`
	for _, para := range chapter.Content {
		paraHTML = paraHTML + fmt.Sprintf(`
		<li>
		<h3>%s</h3>
		<img src="%s" alt="%s" class="para_img">
		<p>%s</p>
		</li>
		<hr>
		`, para.Title, para.getImageURL(),
			chapter.Title, para.Text)
	}
	paraHTML = paraHTML + `</ul> </div>`
	return paraHTML
}

// CreateHTML will create HTML for chapters
func (chapter Chapter) CreateHTML() error {
	chapterHTM, err := getFileAsString("chapter.html")
	if err != nil {
		return err
	}

	replacer := strings.NewReplacer("[FEATURED_IMAGE]", chapter.getImageURL(),
		"[TITLE]", chapter.Title, "[SUBTITLE]", chapter.SubTitle, "[DESCRIPTION]",
		chapter.Description, "[CONTENT]", chapter.getParagraphs())
	result := replacer.Replace(chapterHTM)

	err = writeHTMLFile(getFileName(chapter.Title), result)
	if err != nil {
		return err
	}
	return nil
}
