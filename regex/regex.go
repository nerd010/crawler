package main

import (
	"fmt"
	"regexp"
)

const text = `My email is mark@gmail.com
 mayun@abc.com
    xieyu@163.com`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)

	for _, mail := range match {
		fmt.Println(i, mail)
	}
	fmt.Println(match)
}
