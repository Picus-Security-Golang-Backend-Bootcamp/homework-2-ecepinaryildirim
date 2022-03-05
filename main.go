package main

import (
	"fmt"
	helper "homework-2-ecepinaryildirim/pkg"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var books_struct_slice []helper.Book
var books_slice []string
var authors_slice []string
var random_stock_code_slice []string
var ID int

func init() {
	//books slice
	books_slice = []string{"Harry Potter and the Prisoner of Azkaban", "Harry Potter and the Goblet of Fire",
		"It", "Replay", "Everything's Eventual", "The Shining"}
	//authors slice, in order
	authors_slice = []string{"J.K. Rowling", "J.K. Rowling",
		"Stephen King", "Ken Grimwood", "Stephen King", "Stephen King"}
	//random stock code slice, in order
	random_stock_code_slice = []string{"HarryRowP", "HRowlingG",
		"ItKing", "PlayKG", "EventualKing", "StephenShin"}

	//without seed, random numbers are always the same
	rand.Seed(time.Now().UnixNano())
	for i, book_name := range books_slice {
		ID++
		ISBN := strconv.Itoa(rand.Intn(9999999999))
		newBookStruct := helper.InitBookStruct(ID, rand.Intn(100), rand.Intn(100), rand.Intn(100), book_name, ISBN, random_stock_code_slice[i], authors_slice[i])
		books_struct_slice = append(books_struct_slice, newBookStruct)
	}
}

func main() {
	args := os.Args
	var input_arguments []string

	if len(args) == 1 {
		fmt.Printf("Commands that are available :\n search => for searching the book in the list \n list => for listing the books \n delete => for deleting the book \n buy => for buying the book")
		return
	}

	if args[1] == "search" {
		//after the search word, all words are going to be searched inside books
		for i := 2; i < len(args); i++ {
			input_arguments = append(input_arguments, args[i])
		}
		helper.SearchBooks(books_struct_slice, input_arguments)

	} else if args[1] == "list" {
		helper.ListBooks(books_slice)

	} else if args[1] == "delete" {
		helper.DeleteBook(args[2], books_struct_slice)
	} else if args[1] == "buy" {
		helper.Buy(args[2], args[3], books_struct_slice)
	} else {
		//user has entered invalid input
		fmt.Println("Invalid input!")
		return
	}
}
