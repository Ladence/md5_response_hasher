package hash

import (
	"crypto/md5"
	"fmt"
)

type Md5 struct {
}

func (m *Md5) Hash(message []byte) string {
	return fmt.Sprintf("%x", md5.Sum(message))
}
