package main

import (
	"flag"
	"fmt"
	"github.com/jackdanger/collectlinks"
	"net/http"
	"os"
)

func main() {
	flag.Parse()

	args := flag.Args()
	fmt.Println(args)
	if len(args) < 1 {
		fmt.Println("Please specify start page")
		os.Exit(1)
	}

	resp, err := http.Get(args[0])
	if err != nil {
		return
	}
	defer resp.Body.Close()

	links := collectlinks.All(resp.Body)

	for _, link := range links {
		fmt.Println(link)
	}
}
