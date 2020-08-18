package main

import (
	"fmt"
	"gmssl"
)

func main() {

	versions := gmssl.GetVersions()
	fmt.Println("GmSSL Versions:")
	for _, version := range versions {
		fmt.Println(" " + version)
	}
	fmt.Println("")

	fmt.Print("Digest Algorithms:")
	digests := gmssl.GetDigestNames()
	for _, digest := range digests {
		fmt.Print(" " + digest)
	}
	fmt.Println("\n")

	fmt.Print("Ciphers:")
	ciphers := gmssl.GetCipherNames()
	for _, cipher := range ciphers {
		fmt.Print(" " + cipher)
	}
	fmt.Println("\n")

	/* Generate random key and IV */
	keylen, _ := gmssl.GetCipherKeyLength("SMS4")
	key, _ := gmssl.GenerateRandom(keylen)
	ivlen, _ := gmssl.GetCipherIVLength("SMS4")
	iv, _ := gmssl.GenerateRandom(ivlen)
	/* SMS4-CBC Encrypt/Decrypt */
	encryptor, _ := gmssl.NewCipherContext("SMS4", key, iv, true)
	ciphertext1, _ := encryptor.Update([]byte("hello"))
	ciphertext2, _ := encryptor.Final()
	ciphertext := make([]byte, 0, len(ciphertext1)+len(ciphertext2))
	ciphertext = append(ciphertext, ciphertext1...)
	ciphertext = append(ciphertext, ciphertext2...)

	decryptor, _ := gmssl.NewCipherContext("SMS4", key, iv, false)
	plaintext1, _ := decryptor.Update(ciphertext)
	plaintext2, _ := decryptor.Final()
	plaintext := make([]byte, 0, len(plaintext1)+len(plaintext2))
	plaintext = append(plaintext, plaintext1...)
	plaintext = append(plaintext, plaintext2...)

	fmt.Printf("sms4(\"%s\") = %x\n", plaintext, ciphertext)
	fmt.Println()

}
