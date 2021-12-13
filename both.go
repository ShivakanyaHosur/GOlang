package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	contents, err := ioutil.ReadFile("C:\\Users\\Shivakanya\\go\\read1.txt")
	if err != nil {
		fmt.Println("Error: ", err)

	}
	fmt.Println(string(contents))
	var i string
	var j int
	for i, j = range wordrepetation(string(contents)) {
		fmt.Println(i, "=", j)
	}
	var k rune
	var l int
	for k, l = range chararepetation(string(contents)) {
		fmt.Println(string(k), "=", l)
	}

}

func wordrepetation(st string) map[string]int {
	contents := strings.Fields(st)
	wordcount := make(map[string]int)

	for _, word := range contents {

		_, matched := wordcount[word]

		//fmt.Println(i, matched)
		if matched == true {
			wordcount[word] = wordcount[word] + 1
		} else {
			wordcount[word] = 1
		}
	}
	return wordcount
}

func chararepetation(st string) map[rune]int {
	str := []string{st}

	content := strings.Join(str, "")
	charcount := make(map[rune]int)

	for _, char := range content {
		_, matched := charcount[char]
		//fmt.Println(i, matched)
		if matched == true {
			charcount[char] = charcount[char] + 1
		} else {
			charcount[char] = 1
		}
	}
	return charcount

}
