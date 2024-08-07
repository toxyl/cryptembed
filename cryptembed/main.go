package cryptembed

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/toxyl/cipherutils/aesgcm"
	"github.com/toxyl/errors"
	"github.com/toxyl/flo"
)

func EncryptData(plainText []byte, passphrase string) ([]byte, error) {
	b, err := aesgcm.EncryptBytes(plainText, passphrase)
	if err != nil {
		return nil, err
	}
	return append([]byte{0xDE, 0xAD}, b...), nil
}

func DecryptData(cipherText []byte, passphrase string) ([]byte, error) {
	if len(cipherText) < 2 || cipherText[0] != 0xDE || cipherText[1] != 0xAD {
		return nil, errors.Newf("invalid magic bytes")
	}
	return aesgcm.DecryptBytes(cipherText[2:], passphrase)
}

func ProcessDirectory(dir, passphrase string, encrypt bool) {
	flo.Dir(dir).Each(func(f *flo.FileObj) {
		if strings.HasSuffix(f.Name(), ".go") {
			if err := processFile(f.Path(), passphrase, encrypt); err != nil {
				errors.Panic(err, "Processing directory failed")
			}
		}
	}, nil)
}

func processFile(goFilePath, passphrase string, encrypt bool) error {
	file, err := os.Open(goFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	baseDir := filepath.Dir(goFilePath)

	var embedPaths []string
	scanner := bufio.NewScanner(file)

	foundEncryptDirective := false

	for scanner.Scan() {
		line := scanner.Text()
		if foundEncryptDirective {
			if strings.HasPrefix(line, "//go:embed") {
				embedPaths = append(embedPaths, filepath.Join(baseDir, strings.TrimSpace(strings.TrimPrefix(line, "//go:embed"))))
				foundEncryptDirective = false
			} else if strings.TrimSpace(line) != "//" {
				foundEncryptDirective = false
			}
		} else if strings.HasPrefix(line, "// @encrypt") {
			foundEncryptDirective = true
		}
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	if len(embedPaths) == 0 {
		return errors.Newf("no embedded file paths found")
	}

	for _, embedPath := range embedPaths {
		if err := processEmbeddedFile(embedPath, passphrase, encrypt); err != nil {
			return err
		}
	}

	return nil
}

func processEmbeddedFile(filePath, passphrase string, encrypt bool) error {
	f := flo.File(filePath)
	var processedData []byte
	var err error
	if encrypt {
		processedData, err = EncryptData(f.AsBytes(), passphrase)
	} else {
		processedData, err = DecryptData(f.AsBytes(), passphrase)
	}

	if err != nil {
		return err
	}

	return f.StoreBytes(processedData)
}
