package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"net/url"
	"os"
)

func main() {
	startPage := 1
	endPage := 15134
	outputFile := "grab_domain.txt"

	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Failed to create file: %s\n", err.Error())
		return
	}
	defer file.Close()
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64)"),
	)
	c.OnHTML("table[style='padding: 5px 10px;width:100%'] tr", func(e *colly.HTMLElement) {
		rawURL := e.ChildAttr("td:nth-child(2) a", "href")
		if rawURL != "" {
			parsedURL, err := url.Parse(rawURL)
			if err == nil {
				domain := parsedURL.Hostname()
				if domain != "" {
					_, err := file.WriteString(domain + "\n")
					if err != nil {
						fmt.Printf("Failed to write to file: %s\n", err.Error())
					}
				}
			}
		}
	})
	for page := startPage; page <= endPage; page++ {
		url := fmt.Sprintf("https://www.websitescrawl.com/domain-list-%d", page)
		fmt.Printf("Scraping page %d: %s\n", page, url)
		err := c.Visit(url)
		if err != nil {
			fmt.Printf("Failed to scrape page %d: %s\n", page, err.Error())
			continue
		}
	}

	fmt.Printf("Domains saved to %s\n", outputFile)
}
