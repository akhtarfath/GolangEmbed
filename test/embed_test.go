package test

import (
	"GoEmbed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

func TestString(t *testing.T) {
	fmt.Println("Text from File:", GoEmbed.FileText)
}

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("haerin.png", GoEmbed.FileImage, fs.ModePerm)
	if err != nil {
		panic(err.Error())
	}
}

func TestMultipleFiles(t *testing.T) {
	// ReadFile return []byte, error. Must full path of file
	a, _ := GoEmbed.FileMultiple.ReadFile("files/a.txt")
	b, _ := GoEmbed.FileMultiple.ReadFile("files/b.txt")
	c, _ := GoEmbed.FileMultiple.ReadFile("files/c.txt")

	fmt.Println("a.txt:", string(a))
	fmt.Println("b.txt:", string(b))
	fmt.Println("c.txt:", string(c))
}

func TestPathMatcher(t *testing.T) {
	dirEntry, _ := GoEmbed.FilePath.ReadDir("files") // ReadDir return []fs.DirEntry, error
	for _, entry := range dirEntry {                 // fs.DirEntry is a type of os.FileInfo
		if !entry.IsDir() { // check if entry is a file
			fmt.Println("=========================================")
			fmt.Println("File:", entry.Name())                                 // entry.Name() return file name
			content, err := GoEmbed.FilePath.ReadFile("files/" + entry.Name()) // ReadFile return []byte, error
			if err != nil {
				panic(err.Error())
			}
			fmt.Println("Content:", string(content))
		}
	}
}
