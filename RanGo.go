package RanGo

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	"log"
	mathRand "math/rand"
	"time"
)

func RnInt(startIncluded int, endNotIncluded int) int {
	rnInit()
	return mathRand.Intn(endNotIncluded-startIncluded) + startIncluded
}
func rnInit() {
	var b [8]byte
	_, err := cryptoRand.Read(b[:])
	if err != nil {
		log.Println("Unable to generate seed with cryptographically secure random number generator. Time-dependent seed has been used instead.")
		mathRand.Seed(time.Now().UnixNano())
	} else {
		mathRand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
	}
}
