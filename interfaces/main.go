package main

import (
	"bytes"
	"fmt"
)

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	fmt.Println("Writing")
	return 0, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	fmt.Println("Closing")
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello, world!"))

	var wc WriterCloser = NewBufferedWriterCloser()
	defer wc.Close()
	wc.Write([]byte("Hello, world!"))

	r, ok := wc.(*BufferedWriterCloser)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Can't convert")
	}
}
