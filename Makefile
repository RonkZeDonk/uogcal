GOOS=linux
GOARCH=amd64

ENTRYPOINT=cmd/main.go

OUTPUT=main-${GOOS}-${GOARCH}
PACKAGE=uogcal.tar.gz

.PHONY: build backend frontend package fmt clean

build: backend frontend

${OUTPUT}: backend
dist: frontend

backend:
ifeq ($(GOARM),)
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${OUTPUT} ${ENTRYPOINT}
else
	GOOS=${GOOS} GOARCH=${GOARCH} GOARM=${GOARM} go build -o ${OUTPUT} ${ENTRYPOINT}
endif

frontend:
	NODE_ENV=production pnpm build

package: ${OUTPUT} dist
	tar -czvf ${PACKAGE} ${OUTPUT} dist

fmt:
	go fmt -x ./...

clean:
	rm -rf ${OUTPUT} ${PACKAGE} dist
