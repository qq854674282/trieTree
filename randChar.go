package randChar

import (
	"math/rand"
	"time"
)

func RandChar(index int) string {
	var r *rand.Rand
	var string_len int

	r = rand.New(rand.NewSource(time.Now().UnixNano() + int64(index)))

	string_len = r.Intn(99) + 1 //生成字符串的长度在1-100之间
	var slice = make([]byte, string_len, string_len)
	var res string

	for i := 0; i < string_len; i++ {
		slice[i] = byte(r.Intn(74) + 48)

		if slice[i] > 57 && slice[i] < 65 {
			slice[i] += 8
		} else if slice[i] > 90 && slice[i] < 97 {
			slice[i] += 7
		}
	}

	res = string(slice)

	return res
	//return slice
}
