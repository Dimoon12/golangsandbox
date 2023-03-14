package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readfile(){
    dat, err := os.ReadFile("listoflinks.txt")
	check(err)
	newlist := strings.Fields(string(dat))
	lenght:=len(newlist)
	fmt.Print(newlist[0])
	fmt.Print(lenght)
}

func main(){
readfile()
//download("img.freepik.com/premium-photo/british-shorthair-kitten-3-5-months-old-sitting-looking-up_191971-4591.jpg", "test.png")
}



func download(url string, name string) {
	prefix := "https://"
	err := DownloadFile(name, prefix+url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Загружено по url: " + url + "Имя: "+name)
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
