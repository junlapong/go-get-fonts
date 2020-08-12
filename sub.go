package main

import (
	"fmt"
	"regexp"
)

// https://regex101.com/codegen?language=golang

func paresFontFace(css string) error {

	var re = regexp.MustCompile(`(?i)\s*(?:\/\*\s*(.*?)\s*\*\/)?[^@]*?@font-face\s*{(?:[^}]*?)}\s*`)

	for i, match := range re.FindAllString(css, -1) {
		fmt.Println(match, "found at index", i)
	}

	return nil
}

func parseFontFamily(str string) error {
	var re = regexp.MustCompile(`(?i)font-family\s*:\s*(?:\'|")?([^;]*?)(?:\'|")?\s*;`)

	if len(re.FindStringIndex(str)) > 0 {
		fmt.Println(re.FindString(str), "found at index", re.FindStringIndex(str)[0])
		fmt.Printf("%#v\n", re.FindStringSubmatch(str))
		fmt.Printf("%s\n", re.FindStringSubmatch(str)[1])
	}

	return nil
}

func parseFontStyle(str string) error {
	var re = regexp.MustCompile(`(?i)font-style\s*\:\s*[\'"]?([^\;]+?)[\'"]?\;`)

	if len(re.FindStringIndex(str)) > 0 {
		fmt.Println(re.FindString(str), "found at index", re.FindStringIndex(str)[0])
		fmt.Printf("%#v\n", re.FindStringSubmatch(str))
		fmt.Printf("%s\n", re.FindStringSubmatch(str)[1])
	}

	return nil
}

func parseFontWeight(str string) error {
	var re = regexp.MustCompile(`(?i)font-weight\s*:\s*([^;]*?)\s*;`)

	if len(re.FindStringIndex(str)) > 0 {
		fmt.Println(re.FindString(str), "found at index", re.FindStringIndex(str)[0])
		fmt.Printf("%#v\n", re.FindStringSubmatch(str))
		fmt.Printf("%s\n", re.FindStringSubmatch(str)[1])
	}

	return nil
}

func parseFontURL(str string) error {

	var re = regexp.MustCompile(`(?i)url\s*\(\s*(?:\'|")?\s*([^]*?)\s*(?:\'|")?\s*\)\s*?`)

	for i, match := range re.FindAllString(str, -1) {
		fmt.Println(match, "found at index", i)
	}

	// if len(re.FindStringIndex(str)) > 0 {
	// 	fmt.Println(re.FindString(str), "found at index", re.FindStringIndex(str)[0])
	// 	fmt.Printf("%#v\n", re.FindStringSubmatch(str))
	// 	fmt.Printf("%s\n", re.FindStringSubmatch(str)[1])
	// }

	return nil
}
