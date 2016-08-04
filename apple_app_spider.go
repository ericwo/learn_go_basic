package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type App struct {
	gorm.Model

	Name   string
	Link   string
	Letter string
}

func Spider() {
	db, err := gorm.Open("mysql", "root:@/applist?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	url := "https://itunes.apple.com/cn/genre/ios-tu-shu/id6018?mt=8&letter=A&page="

	for i := 1; i <= 29; i++ {
		doc, err := goquery.NewDocument(url + strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}

		doc.Find("#selectedcontent .column ul li").Each(func(i int, s *goquery.Selection) {
			target := s.Find("a")
			band := target.Text()
			link, _ := target.Attr("href")

			app := App{Name: band, Link: link, Letter: "A"}
			db.Create(&app)

			fmt.Printf("%d: %s-%s\n", i, band, link)
		})
	}
}

func main() {
	Spider()
}
