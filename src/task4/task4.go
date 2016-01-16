package task4

import (
	"fmt"
	"crypto/md5"
	"strconv"
)

func Solve() {
	md5 := md5.New()
	key := "bgvyzdsv"
	i := 1

	for {
		data := []byte(key + strconv.Itoa(i))

		md5.Write(data)
		hash := fmt.Sprintf("%x\n", md5.Sum(nil))
		i++
		md5.Reset()

		if (hash[0:5] == "00000") {
			fmt.Println("Five zeroes", string(data), hash)
		}

		if (hash[0:6] == "000000") {
			fmt.Println("Six zeroes", string(data), hash)
			return
		}
	}
}