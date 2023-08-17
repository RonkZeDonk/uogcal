GOOS=linux
GOARCH=amd64

ENTRYPOINT=cmd/main.go

OUTPUT=main-${GOOS}-${GOARCH}
PACKAGE=uogcal.tar.gz

.PHONY: build package fmt clean

build:
	NODE_ENV=production pnpm buildStyles
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${OUTPUT} ${ENTRYPOINT}

package: ${OUTPUT}
	tar -czvf ${PACKAGE} ${OUTPUT} views public robots.txt

fmt:
	go fmt -x ./...

clean:
	rm -f ${OUTPUT} ${PACKAGE}