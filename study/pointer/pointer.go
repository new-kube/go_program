package main

import "fmt"

type BookResource struct {
}

var resourceTable = BookResource{}

func (br *BookResource) TableName() string {
	return "xes_book_resource"
}

func main() {
	fmt.Println(resourceTable.TableName())
}
