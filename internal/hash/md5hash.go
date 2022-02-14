package hash

import (
	"crypto/md5"
	"fmt"
)

type Md5Hasher struct {
}

func NewMd5Hasher() Hasher {
	return &Md5Hasher{}
}

func (m *Md5Hasher) Hash(message []byte) string {
	return fmt.Sprintf("%x", md5.Sum(message))
}
