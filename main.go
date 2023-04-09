package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed files/text1.txt
var teks1 string

//go:embed gambar.jpg
var pic []byte

func DupPic() {
	err := ioutil.WriteFile("gambar-new.jpg", pic, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/*.txt
var path embed.FS

func AllFiles() {
	driEntries, _ := path.ReadDir("files")
	for _, entry := range driEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}

func main() {
	fmt.Println(teks1)
	DupPic()
	AllFiles()

}
