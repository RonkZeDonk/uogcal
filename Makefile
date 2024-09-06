GOOS=linux
GOARCH=amd64

OUTPUT=main-${GOOS}-${GOARCH}

.PHONY: build backend frontend package fmt distclean

build: frontend backend

${OUTPUT}: backend
dist: frontend

backend:
ifeq ($(GOARM),)
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${OUTPUT}
else
	GOOS=${GOOS} GOARCH=${GOARCH} GOARM=${GOARM} go build -o ${OUTPUT}
endif

frontend:
	NODE_ENV=production pnpm build

fmt:
	go fmt -x ./...

clean:
	rm main-*-*

distclean:
	rm -rf main-*-* dist
