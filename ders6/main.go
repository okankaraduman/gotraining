package main

import (
	"os"
)

var (
	newFile *os.File
	err     error
)

func main() {
	////Creating a file
	//newFile, err = os.Create("demo.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	/*
		Bir dosya 100 byte'a kesilmelidir.
		Eğer dosya 100 byte'tan az ise içerik aklır, geri kalan kısım boş byte ile dolacaktır.
		Eğer dosya 100 byte'ın üzeirnde ise 100 byte'tan sonraki her şey kaybolur.

	*/
	//err := os.Truncate("demo.txt", 100)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fi, err := os.Stat("demo.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("File Name:", fi.Name())
	//fmt.Println("Size in Bytes", fi.Size())
	//fmt.Println("Permissions", fi.Mode())
	//fmt.Println("Last Modified", fi.ModTime())
	//fmt.Println("Is Directory", fi.IsDir())
	//fmt.Println("System Info \n:", fi.Sys())

	//Rename and Move a File

	//originalPath := "demo.txt"
	//NewPath := "./moved/test.txt"
	//err := os.Rename(originalPath, NewPath)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//Delete a file
	//err := os.Remove("demo.txt")
	//
	//if err != nil {
	//	log.Fatal(err)
	//}

}
