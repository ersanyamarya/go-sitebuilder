package main

// PageItem is types that can have a HTML output
type PageItem interface {
	CreateHTML() error
}

// Page is types that can have a HTML output
type Page struct {
	Content []PageItem
}
