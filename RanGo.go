package RanGo

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	"log"
	mathRand "math/rand"
	"time"
)

func RnInt(startIncluded int, endNotIncluded int) (randomInt int) {
	rnInit()
	randomInt = startIncluded + mathRand.Intn(endNotIncluded-startIncluded)

	return randomInt
}
func RandomInt(startIncluded int, endNotIncluded int) (randomInt int, isSeedTimeDependent bool) {
	isSeedTimeDependent = rnInit()
	randomInt = startIncluded + mathRand.Intn(endNotIncluded-startIncluded)

	return randomInt, isSeedTimeDependent
}
func rnInit() (isSeedTimeDependent bool) {
	var b [8]byte
	_, err := cryptoRand.Read(b[:])
	if err != nil {
		log.Println("Unable to generate seed with cryptographically secure random number generator. Time-dependent seed has been used instead.")
		mathRand.Seed(time.Now().UnixNano())

		return true
	} else {
		mathRand.Seed(int64(binary.LittleEndian.Uint64(b[:])))

		return false
	}
}
