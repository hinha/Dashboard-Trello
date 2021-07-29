package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	b64 "encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"

	"github.com/mergermarket/go-pkcs7"
)

var (
	// ErrEncryptBlockSize indicates data blockSize
	ErrEncryptBlockSize = errors.New("some data is not a multiple of the block size")
	// ErrEncryptShort indicates data is to short
	ErrEncryptShort = errors.New("encrypt too short")
)

type BearerCipher struct {
	secretIV     string
	secretHeader string
	secretCookie string
}

func NewCipher(secretIV, secretHeader string, secretCookie string) *BearerCipher {
	return &BearerCipher{secretIV, secretHeader, secretCookie}
}

func (b *BearerCipher) PKCS5Padding(ciphertext []byte, blockSize int, _ int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func (b *BearerCipher) EncryptAES256(plaintext string, key string, iv string, blockSize int) []byte {
	bKey := []byte(key)
	bIV := []byte(iv)
	bPlaintext := b.PKCS5Padding([]byte(plaintext), blockSize, len(plaintext))
	block, _ := aes.NewCipher(bKey)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)

	return ciphertext
}

func (b *BearerCipher) DecryptAES256(crypt []byte, key string, iv string) ([]byte, error) {
	bKey := []byte(key)
	bIV := []byte(iv)
	block, err := aes.NewCipher(bKey)
	if err != nil {
		return nil, err
	}
	if len(crypt) == 0 {
		return nil, err
	}
	ecb := cipher.NewCBCDecrypter(block, bIV)
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)
	return b.PKCS5Trimming(decrypted), nil
}

func (b *BearerCipher) PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func (b *BearerCipher) Binary(s string) string {
	res := ""
	for _, c := range s {
		res = fmt.Sprintf("%s%.8b", res, c)
	}
	return res
}

// HashTo32Bytes will compute a cryptographically useful hash of the input string.
func (b *BearerCipher) HashTo32Bytes(input string) []byte {
	data := sha256.Sum256([]byte(input))
	return data[0:]
}

func (b *BearerCipher) PadIV(role, char, secretKey string) (string, string) {
	vector := b.Binary(char)
	initialIV := base64.URLEncoding.EncodeToString(b.EncryptAES256(role, secretKey, vector, aes.BlockSize))

	pad := len(initialIV) / 2
	start := initialIV[:pad-4] // salah ini ambil kebelakang
	end := initialIV[pad-4:]
	return start, end
}

// encryptCBC encrypts plain text string into cipher text string
func encryptCBC(key, message []byte) (string, error) {
	plainText, err := pkcs7.Pad(message, aes.BlockSize)
	if err != nil {
		return "", fmt.Errorf(`plainText: "%s" has error`, plainText)
	}
	if len(plainText)%aes.BlockSize != 0 {
		err := fmt.Errorf(`plainText: "%s" has the wrong block size`, plainText)
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)
	return b64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%x", cipherText))), nil
}

func (b *BearerCipher) EncryptStringCBC(plainText []byte) (string, error) {

	key := b.HashTo32Bytes(b.GetSecretCookie())
	encrypted, err := encryptCBC(key, plainText)
	if err != nil {
		return "", err
	}

	return encrypted, nil
}

// decryptCBC decrypts cipher text string into plain text string
func decryptCBC(key, message []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(message) < aes.BlockSize {
		return "", ErrEncryptShort
	}
	iv := message[:aes.BlockSize]
	message = message[aes.BlockSize:]
	if len(message)%aes.BlockSize != 0 {
		return "", ErrEncryptBlockSize
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(message, message)

	message, _ = pkcs7.Unpad(message, aes.BlockSize)
	return fmt.Sprintf("%s", message), nil
}

// The function will output the resulting plain text string with an error variable.
func (b *BearerCipher) DecryptStringCBC(cryptoText string) (plainTextString string, err error) {
	sDec, err := b64.URLEncoding.DecodeString(cryptoText)
	if err != nil {
		return "", err
	}

	cipherText, _ := hex.DecodeString(string(sDec))
	decrypted, err := decryptCBC(b.HashTo32Bytes(b.GetSecretCookie()), cipherText)
	if err != nil {
		return "", err
	}

	return decrypted, nil
}

func (b *BearerCipher) GetSecretIV() []byte {
	return b.HashTo32Bytes(b.secretIV)
}
func (b *BearerCipher) GetSecretHeader() []byte {
	return b.HashTo32Bytes(b.secretHeader)
}

func (b *BearerCipher) GetSecretCookie() string {
	return b.secretCookie
}
