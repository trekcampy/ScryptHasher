/*
 *  ScryptHasher.go
 *
 * This is a simple utility that outputs a 64 byte hash/key using Scrypt API
 *   Usage : ScryptHasher --passphrase <input string used in generating hash>
 *
 */

package main

import (
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"log"

	"golang.org/x/crypto/scrypt"
)

var passPhrase = flag.String("passphrase", "passphrase", "pass phrase used in generating hash")
var saltIn = flag.String("salt", "", "8 byte salt in hex format")

const (
	N  = 1 << 20
	r  = 8
	p  = 1
	sz = 64
)

func main() {
	flag.Parse()

	salt, err := getSalt()
	if err != nil {
		log.Fatalf("Failed to get salt : error=(%v)", err)
	}

	dk, err := scrypt.Key([]byte("some password"), salt, N, r, p, sz)
	if err != nil {
		log.Fatalf("Failed to generate hash : error=(%v)", err)
	}

	fmt.Printf("Passphrase %s, Salt %s => Hash %s\n",
		*passPhrase,
		hex.EncodeToString(salt),
		hex.EncodeToString(dk))

}

func getSalt() ([]byte, error) {

	var salt []byte
	var err error

	if *saltIn != "" {

		salt, err = hex.DecodeString(*saltIn)
		if err != nil {
			return salt, err
		}
	} else {
		salt = gen64BitRand()
	}

	return salt, nil
}
func gen64BitRand() []byte {

	buf := make([]byte, 8)
	rand.Read(buf)

	return buf
}
