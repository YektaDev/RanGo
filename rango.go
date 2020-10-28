package rango

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	"log"
	mathRand "math/rand"
	"time"
)

//Characters of random string generator
var charsLowercase = "abcdefghijklmnopqrstuvwxyz"
var charsUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var charsNumber = "0123456789"
var charsSpecial = "[](){}^;:#?.,&|!~`_%$@\"'=+-*/\\ <>"

//Generate a random int
func RnInt(startIncluded int, endNotIncluded int) (randomInt int) {
	rnInit()
	randomInt = startIncluded + mathRand.Intn(endNotIncluded-startIncluded)

	return randomInt
}

//Generate a random int
func RandomInt(startIncluded int, endNotIncluded int) (randomInt int, isSeedTimeDependent bool) {
	isSeedTimeDependent = rnInit()
	randomInt = startIncluded + mathRand.Intn(endNotIncluded-startIncluded)

	return randomInt, isSeedTimeDependent
}

//Generate a random string
func RnString(length int, containsLowercase bool, containsUppercase bool, containsNumber bool, containsSpecial bool) (randomString string) {
	if !containsLowercase && !containsUppercase && !containsNumber && !containsSpecial {
		return ""
	}
	if length < 1 {
		return ""
	}

	var chars []rune

	if containsLowercase {
		chars = append(chars, []rune(charsLowercase)...)
	}
	if containsUppercase {
		chars = append(chars, []rune(charsUppercase)...)
	}
	if containsNumber {
		chars = append(chars, []rune(charsNumber)...)
	}
	if containsSpecial {
		chars = append(chars, []rune(charsSpecial)...)
	}

	rnInit()

	b := make([]rune, length)
	for i := range b {
		b[i] = chars[mathRand.Intn(len(chars))]
	}
	return string(b)
}

//Generate a random string
func RandomString(length int, containsLowercase bool, containsUppercase bool, containsNumber bool, containsSpecial bool) (randomString string, isSeedTimeDependent bool) {
	if !containsLowercase && !containsUppercase && !containsNumber && !containsSpecial {
		return "", false
	}
	if length < 1 {
		return "", false
	}

	var chars []rune

	if containsLowercase {
		chars = append(chars, []rune(charsLowercase)...)
	}
	if containsUppercase {
		chars = append(chars, []rune(charsUppercase)...)
	}
	if containsNumber {
		chars = append(chars, []rune(charsNumber)...)
	}
	if containsSpecial {
		chars = append(chars, []rune(charsSpecial)...)
	}

	isSeedTimeDependent = rnInit()

	b := make([]rune, length)
	for i := range b {
		b[i] = chars[mathRand.Intn(len(chars))]
	}
	return string(b), isSeedTimeDependent
}

//Generate a random string with the given characters
func RnStringFrom(length int, chars string) (randomString string) {
	if length < 1 {
		return ""
	}

	rnInit()

	b := make([]rune, length)
	for i := range b {
		b[i] = []rune(chars)[mathRand.Intn(len([]rune(chars)))]
	}
	return string(b)
}

//Generate a random string with the given characters
func RandomStringFrom(length int, chars string) (randomString string, isSeedTimeDependent bool) {
	if length < 1 {
		return "", false
	}

	isSeedTimeDependent = rnInit()

	b := make([]rune, length)
	for i := range b {
		b[i] = []rune(chars)[mathRand.Intn(len([]rune(chars)))]
	}
	return string(b), isSeedTimeDependent
}

//Generate seed
func rnInit() (isSeedTimeDependent bool) {
	var b [8]byte
	_, err := cryptoRand.Read(b[:])
	if err != nil {
		log.Println("Unable to generate seed with cryptographically secure random number generator. Time-dependent seed has been used instead.")
		mathRand.Seed(time.Now().UnixNano())
		return true
	}
	mathRand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
	return false
}
