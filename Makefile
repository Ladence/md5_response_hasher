GOOS?=darwin
GOARCH?=amd64

APP?=myhttp

clean:
	rm -f ${APP}

build: clean
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${APP}

test: build
	go test -v -race ./...