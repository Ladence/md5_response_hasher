package hash

type Hasher interface {
	Hash([]byte) string
}
