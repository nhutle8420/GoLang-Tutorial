package main

import "os"

func main() {
	listbook := createBooks()
	listbook.SaveFile("Books.txt")
	listbook.in()
	listbook = append(listbook, "Filter")
	listbook.changSttBook()
	listbook.in()
	listbook, err := ReadFile("Books.txt")
	if err != nil {
		os.Exit(1)
	}
	listbook.in()
}
