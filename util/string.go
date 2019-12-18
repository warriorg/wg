package util

import (
	"errors"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// RandString 生成随机字符串
func RandString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// PasswordStrong 检查密码强度
func PasswordStrong(password string) (int, error) {
	indNum := [4]int{0, 0, 0, 0}
	spCode := []byte{'!', '@', '#', '$', '%', '^', '&', '*', '_', '-'}
	if len(password) < 8 {
		return 0, nil
	}

	passwdByte := []byte(password)

	for _, i := range passwdByte {
		if i >= 'A' && i <= 'Z' {
			indNum[0] = 1
			continue
		}
		if i >= 'a' && i <= 'z' {
			indNum[1] = 1
			continue
		}

		if i >= '0' && i <= '9' {
			indNum[2] = 1
			continue
		}

		notEnd := 0
		for _, s := range spCode {
			if i == s {
				indNum[3] = 1
				notEnd = 1
				break
			}
		}

		if notEnd != 1 {
			return 0, errors.New("Unsupport code")
		}
	}

	codeCount := 0

	for _, i := range indNum {
		codeCount += i
	}

	return codeCount, nil
}
