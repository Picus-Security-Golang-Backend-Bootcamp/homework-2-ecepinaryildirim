package helper

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Author struct {
	name, surname string
}

type Book struct {
	ID           int
	page_number  int
	price        int
	stock_number int
	book_name    string
	ISBN         string
	stock_code   string
	author       Author
	is_deleted   bool
}

type Deletable interface {
	Delete()
}

var err_book_already_deleted = errors.New("book has already deleted")
var err_not_enough_book_stock = errors.New("there is not enough book stock")
var err_invalid_id = errors.New("book ID is invalid")
var err_book_not_available = errors.New("searched book is not available")
var err_unexpected_error = errors.New("unexpected behaviour")

func InitBookStruct(ID, page_number, price, stock_number int, book_name, ISBN, stock_code, author_info string) Book {
	b := Book{
		ID:           ID,
		page_number:  page_number,
		price:        price,
		stock_number: stock_number,
		book_name:    book_name,
		ISBN:         ISBN,
		stock_code:   stock_code,
		author:       Author{name: strings.Fields(author_info)[0], surname: strings.Fields(author_info)[1]},
	}

	return b
}

func SearchBooks(books_struct_slice []Book, input_arguments []string) error {
	//for keeping all the search results inside books struct slice
	var big_list []string

	for i := 0; i < (len(books_struct_slice)); i++ {
		big_list = append(big_list, books_struct_slice[i].author.name, " ", books_struct_slice[i].author.surname)
		big_list = append(big_list, books_struct_slice[i].book_name)
		big_list = append(big_list, books_struct_slice[i].stock_code)
		big_list = append(big_list, books_struct_slice[i].ISBN)
	}
	//getting a string that contains match cases
	books_list := ASubSlice(big_list, input_arguments)

	if len(books_list) != 0 {
		for i := 0; i < len(books_list); i++ {
			fmt.Println(books_list[i])
		}
	} else {
		return err_book_not_available
	}

	return nil
}

func ListBooks(books []string) {
	for i := 0; i < len(books); i++ {
		fmt.Println(books[i])
	}
}

func ASubSlice(slice_one, slice_two []string) []string {
	var return_slice []string

	//making a string from a string array for strings.contains method
	str := strings.Join(slice_two, " ")

	//eliminating upper-lower case problem
	str = strings.ToUpper(str)

	if len(slice_one) < len(slice_two) {
		return return_slice
	}

	for i := 0; i < len(slice_one); i++ {
		if strings.Contains(strings.ToUpper(slice_one[i]), str) {
			return_slice = append(return_slice, slice_one[i])
		}
	}
	return return_slice
}

func Buy(id_str, wanted_amount_str string, book_struct []Book) error {
	ID, er := strconv.Atoi(id_str)
	if er != nil {
		return err_unexpected_error
	}

	//if ID is bigger than the length, that means it is invalid because ID starts from 1 and increases by 1 for each book
	if ID > len(book_struct) {
		return err_invalid_id
	}
	wanted_amount, err := strconv.Atoi(wanted_amount_str)
	if err != nil {
		return err_unexpected_error
	}
	if wanted_amount > book_struct[ID].stock_number {
		return err_not_enough_book_stock
	} else {
		//updating the struct by passing reference, we want this change to be actually done on the object
		updateBookStruct(&book_struct[ID], wanted_amount)
	}
	return nil
}

func updateBookStruct(b *Book, wanted_amount int) {
	fmt.Printf("Stock updated! Old amount: %d ", b.stock_number)
	b.stock_number -= wanted_amount
	fmt.Printf("New amount: %d ", b.stock_number)
}

func DeleteBook(id_str string, book_struct []Book) error {
	ID, er := strconv.Atoi(id_str)
	if er != nil {
		return err_unexpected_error
	}
	//if ID is bigger than the length, that means it is invalid because ID starts from 1 and increases by 1 for each book
	if ID > len(book_struct) {
		return err_invalid_id
	}
	b := book_struct[ID]
	if !b.is_deleted {
		b.Delete()
		fmt.Println("Book deleted!")
	} else {
		return err_book_already_deleted
	}
	return nil
}

func (b *Book) Delete() {
	b.is_deleted = true
}
