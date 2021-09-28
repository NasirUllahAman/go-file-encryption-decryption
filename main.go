package main

import (
	//"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/NasirUllahAman/go-file-encryption-decryption/pkg"
)

func main() {
	key := "passphrasewhichneedstobe32bytes!"
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}

	choice := os.Args
	if choice[1] == "y" {
		encryptedText, err := pkg.Encrypt([]byte(key), string(data))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Encrypted text", encryptedText)

		ioutil.WriteFile("test.txt", []byte(encryptedText), 0644)
	} else if choice[1] == "n" {

		decryptedText, _ := pkg.Decrypt([]byte(key), string(data))

		ioutil.WriteFile("test.txt", []byte(decryptedText), 0644)

		fmt.Println("Decrypted text", decryptedText)
	}
}
