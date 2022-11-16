package main

import (
	"Bookstore/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bxcodec/faker/v4"
	"log"
	"net/http"
	"sync"
)

const BaseUri = "http://localhost:1323/"
const CreateBookUri = BaseUri + "books"
const GetBooksUri = BaseUri + "books"

var bookIdChannel = make(chan int)
var wg = new(sync.WaitGroup)
var mutex = &sync.Mutex{}

func main() {
	wg.Add(4)

	go createBook(bookIdChannel)
	go printBook(bookIdChannel)
	go createBook(bookIdChannel)
	go printBook(bookIdChannel)

	wg.Wait()
}

func createBook(bookIdChannel chan int) *models.Book {
	defer wg.Done()

	fakeBook := new(models.Book)

	mutex.Lock()
	err := faker.FakeData(&fakeBook)
	if err != nil {
		log.Fatal(err)
	}
	mutex.Unlock()

	postBody, _ := json.Marshal(fakeBook)

	resp, err := http.Post(CreateBookUri, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		log.Fatalln(err)
	}

	book := new(models.Book)

	err = json.NewDecoder(resp.Body).Decode(&book)
	if err != nil {
		log.Fatal(err)
	}

	bookIdChannel <- book.Id

	return book
}

func printBook(bookIdChannel chan int) {
	defer wg.Done()
	bookId := <-bookIdChannel

	resp, err := http.Get(fmt.Sprintf("%s/%d", GetBooksUri, bookId))

	if err != nil {
		log.Fatalln(err)
	}

	book := new(models.Book)

	err = json.NewDecoder(resp.Body).Decode(&book)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", book)
}
