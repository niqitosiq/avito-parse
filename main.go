package main

import (
	"fmt"
	"flag"
	"github.com/gocolly/colly"

)


func main(){

	c := colly.NewCollector()
	
	c.OnHTML(".item-description-title-link", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnHTML(".item-view-content", func(e *colly.HTMLElement){
		fmt.Println("Name:", e.ChildText(".title-info-title-text"))
		fmt.Println("Description:", e.ChildText(".item-description-text"), "\n\n")
	})

	var link_main string

	flag.StringVar(&link_main, "link", "", "Please input a link;")
	flag.Parse()

	fmt.Println("Let's look:", link_main)
	c.Visit(link_main)
}