package main

func main() {
	listbook := createBooks()

	listbook.SaveFile("Book.txt")
	listbook = append(listbook, "Filter")
	listbook.in()

	listbook = ReadFile("Book.txt")
	listbook.in()
}
