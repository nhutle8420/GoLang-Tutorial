package main

import "fmt"

type Books []string

func (n Books) in() {
	fmt.Println(n)
}
func createBooks() Books {
	kq := Books{}
	chapter := []string{"chapter 1", "Chapter 2", "chapter 3", "chapter 4", "chapter 5"}
	book := []string{"Conna", "Pokeomon", "Onepice", "Doaremon"}
	for _, booki := range book {
		for _, chapi := range chapter {

			newbook := booki + "-" + chapi
			kq = append(kq, newbook)
		}
	}
	return kq
}
func inventory(n Books, Sl int) (Books, Books) {
	return n[:Sl], n[Sl:]
}
