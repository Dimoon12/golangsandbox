package main

import (
	"fmt"
	"io"
	"os/exec"
	"net/http"
	"os"
	"strings"
	"strconv"
)

func decompress(filename string){
	cmd := exec.Command("/usr/bin/tar", "-xf " + "data/"+filename+ " -C data")
	fmt.Printf("Пытаюсь распаковать\n")
	err := cmd.Run()
	fmt.Printf("Завершено с ошибкой: %v", err)
}


func resetdata(){
	fmt.Print("Стираю временные данные\n")
	os.RemoveAll("data")

}

func initial() {
	fmt.Printf("Инициализация\n")
	os.Mkdir("data", 0750)
	fmt.Printf("Инициализация завершена\n")
}

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

	for i := 0; i < lenght; i++ {
        fmt.Println("Запуск загрузки: ",newlist[i], "Индекс: ", i)
		index := strconv.Itoa(i)
		download(newlist[i], index)
}
}

func main(){
initial()
readfile()
decompress("0")
//resetdata()
}



func download(url string, name string) {
	prefix := "https://"
	err := DownloadFile("data/"+name, prefix+url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Загружено по url: " + url + " Имя: "+name)
}

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
