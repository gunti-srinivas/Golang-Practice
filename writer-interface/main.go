package main

import (
	"bytes"
	"fmt"
)

func main() {
	// f, err := os.Create("output.txt")
	// if err != nil {
	// 	log.Fatalf("error %v", err)
	// }
	// defer f.Close()

	// s := []byte("hello gophers")
	// _, err = f.Write(s)
	// if err != nil {
	// 	log.Fatalf("error %v ", err)
	// }

	b := bytes.NewBufferString("hello this is srinivas")
	fmt.Println(b.String())
	b.WriteString("hii my second line")
	fmt.Println(b.String())
	b.Write([]byte("third argument"))
	fmt.Println(b.String())
}
