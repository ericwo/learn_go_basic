package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type App struct {
	gorm.Model

	Name string
	Link string
}

func Spider() {
	doc, err := goquery.NewDocument("https://itunes.apple.com/cn/genre/ios-tu-shu/id6018?mt=8&letter=A&page=5")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#selectedcontent .column ul li").Each(func(i int, s *goquery.Selection) {
		band := s.Find("a").Text()
		fmt.Printf("%d: %s\n", i, band)
	})
}

func main() {
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}

	Spider()
}
