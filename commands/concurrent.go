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

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(4)

	bookIdChannel := make(chan int)

	go createBook(wg, bookIdChannel)
	go printBook(wg, bookIdChannel)
	go createBook(wg, bookIdChannel)
	go printBook(wg, bookIdChannel)

	wg.Wait()
}

func createBook(wg *sync.WaitGroup, bookIdChannel chan int) *models.Book {
	defer wg.Done()

	fakeBook := new(models.Book)

	err := faker.FakeData(&fakeBook)
	if err != nil {
		log.Fatal(err)
	}

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

func printBook(wg *sync.WaitGroup, bookIdChannel chan int) {
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
	fmt.Printf("%v", book)
	fmt.Println()
}
