# `cryptembed`

... is a Go library designed for encrypting and decrypting embedded files in Go applications. It enables easy management of sensitive data by encrypting files that are then embedded into your Go binaries.

## Warning

**This library is maintained primarily for my personal use and may not be actively supported or updated.** It is provided as-is and is not guaranteed to be reliable for production use. Developers are encouraged to fork, adapt, and extend the code for their own projects and ensure it meets their needs through their own testing and validation.

## Getting Started

### Installation

Add `cryptembed` to your Go project using:

```bash
go get github.com/toxyl/cryptembed
```

### Using The CLI Tool And Library
1. **Build The CLI Tool**

   To build the CLI tool, run:

   ```bash
   sudo go build -o /usr/local/bin/cembed app/cli/main.go
   ```

2. **Encrypt Project Files**

   To encrypt files before building your application, run:

   ```bash
   cembed -dir <directory> -passphrase <passphrase> -encrypt
   ```

   - `-dir <directory>`: Directory to scan for Go files with `// @encrypt` directives.
   - `-passphrase <passphrase>`: Passphrase for encryption.
   - `-encrypt`: Flag to indicate encryption mode.

3. **Decrypt Project Files**

   To decrypt files, use:

   ```bash
   cembed -dir <directory> -passphrase <passphrase> -decrypt
   ```

   - `-decrypt`: Flag to indicate decryption mode.

4. **Decrypt Embedded Files**
   
   Within your application you have to use the `cryptembed.DecryptData` method to decrypt the encrypted embed. See this [example implementation](app/example/main.go).

## Contributing

If you find this library useful, you are welcome to fork it or copy the code for your own projects. Please adapt it as necessary and implement your own test suites to ensure it meets your requirements. 


## License

This project is released into the public domain under the UNLICENSE.
