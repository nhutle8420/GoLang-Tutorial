package main

func main() {
	listbook := createBooks()
	listbook2, listbook3 := inventory(listbook, 13)
	listbook2.in()
	listbook3.in()
}
