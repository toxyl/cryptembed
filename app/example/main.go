package main

import (
	_ "embed"
	"flag"
	"fmt"

	"github.com/toxyl/cryptembed/cryptembed"
)

//go:embed a.txt
var unencrypted string // cryptembed will leave this file alone (embed directive is not preceded by "// @encrypt")

// @encrypt
//
//go:embed b.txt
var encrypted string // cryptembed will encrypt this file (signaled by "// @encrypt")

func main() {
	passphrase := flag.String("passphrase", "", "Encryption/decryption passphrase")
	flag.Parse()
	decrypted, err := cryptembed.DecryptData([]byte(encrypted), *passphrase)
	if err != nil {
		fmt.Printf("Could not decrypt embedded data: %s\n", err)
		return
	}
	fmt.Printf("Unencrypted data:\n%s\n\nEncrypted data before:\n%s\n\nEncrypted data after:\n%s\n\n", unencrypted, encrypted, decrypted)
}
