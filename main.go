package main

import (
	//"bufio"
	"fmt"
	"io/ioutil"
	"os"

	c1 "github.com/NasirUllahAman/go-file-encryption-decryption/pkg"
)

func main() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}

	c := c1.NewCaesar(1)
	choice := os.Args
	if choice[1] == "y" {
		encryptedText := c.Encryption(string(data))
		fmt.Println("Encrypted text", encryptedText)

		ioutil.WriteFile("test.txt", []byte(encryptedText), 0644)
	} else if choice[1] == "n" {

		decryptedText := c.Decryption(string(data))

		ioutil.WriteFile("test.txt", []byte(decryptedText), 0644)

		fmt.Println("Decrypted text", decryptedText)
	}
}
