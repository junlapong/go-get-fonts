package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type fontFace struct {
	Subset     string
	FontFamily string
	FontStyle  string
	FontWeight string
	URL        string
}

var fontsPath string = "./fonts"
var cssFileName string = "fonts.css"

func init() {
	initDir(fontsPath)
}

func main() {
	Do()

	fontname := fontsPath + "/Sarabun-normal-400.woff2"
	fonturl := "https://fonts.gstatic.com/s/sarabun/v7/DtVjJx26TKEr37c9aBVJmw.eot"
	println(fontname)

	DownloadFile(fontname, fonturl)
}

// Do download
func Do() {

	url := "https://fonts.googleapis.com/css?family=Sarabun:200,400"

	userAgents := map[string]string{
		//"eot": "Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 6.1; Trident/4.0)",
		//"woff": "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:27.0) Gecko/20100101 Firefox/27.0",
		"woff2": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36",
		//"svg": "Mozilla/4.0 (iPad; CPU OS 4_0_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/4.1 Mobile/9A405 Safari/7534.48.3",
		//"ttf": "Mozilla/5.0 (Unknown; Linux x86_64) AppleWebKit/538.1 (KHTML, like Gecko) Safari/538.1 Daum/4.1",
	}

	for t, ua := range userAgents {
		fmt.Printf("%s => %s\n", t, ua)
		DownloadCSS(url, ua)
	}
}

// DownloadCSS for download fonts
func DownloadCSS(url, userAgent string) {

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatalln(err)
	}

	request.Header.Set("User-Agent", userAgent)
	response, err := client.Do(request)

	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()

	fmt.Println(response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)
	cssBody := string(body)
	println(cssBody)
	writeCSS(cssBody)

	// parese css to list download font
	paresFontFace(cssBody)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeCSS(css string) error {

	cssFile := fontsPath + "/" + cssFileName
	f, err := os.OpenFile(cssFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)

	defer f.Close()

	if _, err = f.WriteString(css); err != nil {
		panic(err)
	}

	return nil
}

func initDir(dirname string) error {

	_, err := os.Stat(dirname)

	if os.IsExist(err) {
		err = os.Remove(dirname)
		fmt.Printf("%s is exists\n", dirname)
	} else {
		err = os.Mkdir(dirname, 0755)

		// create dir
		if err == nil || os.IsExist(err) {
			abs, err := filepath.Abs(dirname)
			if err == nil {
				fmt.Println("Absolute:", abs)
			}
			return nil
		}
	}

	return nil
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
