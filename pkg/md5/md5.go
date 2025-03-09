package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// Sum convert data to md5
func Sum(data string) string {
	h := md5.New()
	h.Write([]byte(data))

	return hex.EncodeToString(h.Sum(nil))
}
