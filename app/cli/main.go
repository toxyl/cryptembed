package main

import (
	"flag"
	"log"

	"github.com/toxyl/cryptembed/cryptembed"
)

func main() {
	dir := flag.String("dir", ".", "Directory to scan")
	passphrase := flag.String("passphrase", "", "Encryption/decryption passphrase")
	encrypt := flag.Bool("encrypt", false, "Encrypt files")
	decrypt := flag.Bool("decrypt", false, "Decrypt files")
	flag.Parse()

	if *passphrase == "" {
		log.Fatal("Passphrase is required")
	}
	if *encrypt && *decrypt {
		log.Fatal("Cannot specify both encrypt and decrypt")
	}
	if !*encrypt && !*decrypt {
		log.Fatal("Either encrypt or decrypt must be specified")
	}

	cryptembed.ProcessDirectory(*dir, *passphrase, *encrypt)
}
