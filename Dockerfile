FROM golang:alpine as builder

RUN apk update \
	&& apk upgrade \
	&& apk add --no-cache \
	ca-certificates \
	git \
	curl \
	coreutils \
	&& curl -sSL -o /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub \
	&& curl -sSL -O https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.28-r0/glibc-2.28-r0.apk \
	&& apk add --no-cache \
	glibc-2.28-r0.apk \
	&& rm glibc-2.28-r0.apk \
	&& update-ca-certificates 2>/dev/null || true

RUN addgroup -S gogeta \
	&& adduser -S gogeta -G gogeta

WORKDIR /app

COPY . .

RUN go get ./...

RUN env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o gogeta .

FROM scratch

WORKDIR /

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app/gogeta .

USER gogeta

CMD [ "/gogeta" ]
