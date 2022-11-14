package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

type Books []string

func (n Books) in() {
	fmt.Println(n)
}
func createBooks() Books {
	kq := Books{}
	chapter := []string{"chapter 1", "Chapter 2"}
	book := []string{"Conna", "Doaremon", "Lulu"}
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
func (b Books) ChangeToString() string {
	return strings.Join(b, ",")
}
func (b Books) SaveFile(NameFile string) error {
	data := []byte(b.ChangeToString())
	return ioutil.WriteFile(NameFile, data, 0666)
}
func ReadFile(NameFile string) (Books, error) {
	data, err := ioutil.ReadFile(NameFile)
	if err != nil {
		log.Printf("Erro: %v", err)
		return Books{}, err
	}
	book := string(data)
	arry := strings.Split(book, ",")
	listbook := Books(arry)
	return listbook, nil
}

func (b Books) changSttBook() {
	rand.Seed(time.Now().UnixNano())
	for index := range b {
		r := rand.Intn(len(b))
		b[index], b[r] = b[r], b[index]

	}
}
