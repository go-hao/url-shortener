package base62

import (
	"log"
	"math"
	"strings"
)

// default
// 0-9: 0-9
// a-z: 10-35
// A-Z: 36-61

type Base62Conf struct {
	Seed62 string `json:",default=0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"` // must be a rearrangement of "0-9A-Za-z"
}

type Base62 struct {
	base62String string
}

func MustNew(c Base62Conf) *Base62 {
	if len(c.Seed62) != base {
		log.Fatalf("error: MustNew: %s", ErrInvalidSeed62.Error())
	}

	b := &Base62{}

	bc := defaultBase62String
	for _, v := range c.Seed62 {
		if string(v) == replacer {
			log.Fatalf("error: MustNew: %s", ErrInvalidSeed62.Error())
		}
		index := strings.Index(bc, string(v))
		if index < 0 {
			log.Fatalf("error: MustNew: %s", ErrInvalidSeed62.Error())
		}
		bc = strings.Replace(bc, string(v), replacer, 1)
		b.base62String += string(v)
	}

	return b
}

func (b *Base62) Encode(num uint64) string {
	if num == 0 {
		return string(b.base62String[0])
	}
	ret := []byte{}
	for num > 0 {
		mod := num % 62
		div := num / 62
		ret = append(ret, b.base62String[mod])
		num = div
	}
	reverse(ret)

	return string(ret)
}

func (b *Base62) Decode(s string) (uint64, error) {
	sSlice := []byte(s)
	sSlice = reverse(sSlice)
	var ret uint64 = 0
	for i, sMember := range sSlice {
		index := strings.Index(b.base62String, string(sMember))
		if index < 0 {
			return 0, ErrInvalidBase62Value
		}
		base := math.Pow(62, float64(i))

		ret += uint64(index) * uint64(base)
	}

	return ret, nil
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
