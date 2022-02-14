package functional

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ladence/md5_response_hasher/internal/hash"
	"github.com/Ladence/md5_response_hasher/internal/transport"
)

// test case with a static data response (expected result received by md5sum)
func TestClientDummyResponseMd5(t *testing.T) {
	const static = "dummy"
	const expectedHash = "275876e34cf609db118f3d84b799a790"

	hasher := hash.NewMd5Hasher()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(static))
	}))
	defer srv.Close()

	c := transport.NewClient()
	respBytes, err := c.SendRequest(srv.URL)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	res := hasher.Hash(respBytes)
	if res != expectedHash {
		t.Errorf("wrong result! expected: %v, got: %v", expectedHash, res)
	}
}
