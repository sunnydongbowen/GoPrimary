package main

import "fmt"
import "testing"
type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type ReaderWriter interface {
	Reader
	Writer
}

type Book struct {
}

func (b *Book) ReadBook() {
	fmt.Println("read book")
}

func (b *Book) WriteBook() {
	fmt.Println("write book")
}

func TestInterface7(t *testing.T) {
	var r Reader
	r = &Book{}
	r.ReadBook()

	var rw ReaderWriter
	rw = &Book{}
	rw.WriteBook()
	rw.ReadBook()	

}