package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Hold all the data gets written into the "read" function in the io.Reader interface
// Custom type created, which to be a Writer interface, must implement write method
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil { // 'nil' === no value
		fmt.Println("Error:", err)
		os.Exit(1) // Indicates some error happened and exist the program
	}

	// Printing the response directly, response's body doesn't appear at all
	// Reason. Response's body in go is a io.ReadCloser
	fmt.Println(resp)

	// Get info from the read function
	// 1) Initialize an element []byte previously which writes in that element
	//   1.1) Normal
	//bs := []byte{}
	//   1.2) With make
	// 9999 -- Number of the elements initialized with whitespace
	bsTwo := make([]byte, 99999)

	resp.Body.Read(bsTwo)
	//fmt.Println("bsTwo: ", bsTwo)
	fmt.Println("bsTwo: ", string(bsTwo)) // Conver to string, to be more readable

	// 2) Via io.Copy copying to output source
	// io.copy(ElementWhichImplementsWriterInterface, ElementWhichImplementsReaderInterface)
	//   2.1) Using a type from GO which already implements Writer interface
	io.Copy(os.Stdout, resp.Body)
	//   2.2) Custom type which implements Writer interface === It's a Write type
	lw := logWriter{}
	io.Copy(lw, resp.Body)
	// Copy resp.Body to lw till EOF or an error happens
}

// Method implemented to indicate that logWriter is a writer type
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	// Since you are implementing a Writer interface, you could return any int, error that you want
	// but it would be better to return what they normally mean in a Writer interface --> len(bs)
	return len(bs), nil
}
