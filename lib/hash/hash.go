package hash

import "crypto/sha1"

func Sha1(str string) string {
	sha := sha1.New()
	sha.Write([]byte(str))
	return string(sha.Sum(nil))
}
