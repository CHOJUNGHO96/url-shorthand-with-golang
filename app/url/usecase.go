package url

import (
	"crypto/sha1"
	"fmt"
	"github.com/catinello/base62"
	"log"
)

func GetHashUrl(url string) string {
	h := sha1.New()
	h.Write([]byte(url))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func GetDecodeBase62(hashData string) int {
	decodedData, err := base62.Decode(hashData)
	if err != nil {
		log.Printf("Error decoding hashData: %v", err)
		return -1
	}
	return decodedData
}

func GetEncodeBase62(Base62EncodeData int) string {
	return base62.Encode(Base62EncodeData)
}
