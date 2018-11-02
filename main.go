package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"regexp"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
)

func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	match := re.FindAllSubmatch(contents, -1)
	for _, m := range match {
		fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
	}
	fmt.Printf("Matches Found: %d\n", len(match))
}

func main() {

	printCityList(all)
}
