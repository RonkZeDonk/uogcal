GOOS=linux
GOARCH=amd64

ENTRYPOINT=cmd/main.go

OUTPUT=main-${GOOS}-${GOARCH}
PACKAGE=uogcal.tar.gz

.PHONY: build styles package fmt clean

build: styles
ifeq ($(GOARM),)
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${OUTPUT} ${ENTRYPOINT}
else
	GOOS=${GOOS} GOARCH=${GOARCH} GOARM=${GOARM} go build -o ${OUTPUT} ${ENTRYPOINT}
endif

styles:
	NODE_ENV=production pnpm buildStyles

package: ${OUTPUT}
	tar -czvf ${PACKAGE} ${OUTPUT} views public robots.txt

fmt:
	go fmt -x ./...

clean:
	rm -f ${OUTPUT} ${PACKAGE}
