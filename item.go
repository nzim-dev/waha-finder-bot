package info

import "fmt"

//Item that we get after scraping
type Item struct {
	Title string
	Link  string
}

func FormatOutput(items []Item) string {
	var output string

	for _, item := range items {
		output += fmt.Sprintf("%s\n%s\n", item.Title, item.Link)
	}

	return output
}
