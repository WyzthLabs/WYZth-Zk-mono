FROM golang:1.19.3 as builder

ARG PACKAGE=eventindexer

RUN apt install git curl

RUN mkdir /wyzth_zkevm-mono

WORKDIR /wyzth_zkevm-mono

COPY . .

RUN go mod download

WORKDIR /wyzth_zkevm-mono/packages/$PACKAGE

RUN CGO_ENABLED=0 GOOS=linux go build -o /wyzth_zkevm-mono/packages/$PACKAGE/bin/${PACKAGE} /wyzth_zkevm-mono/packages/$PACKAGE/cmd/main.go

FROM alpine:latest

ARG PACKAGE

RUN apk add --no-cache ca-certificates

COPY --from=builder /wyzth_zkevm-mono/packages/$PACKAGE/bin/$PACKAGE /usr/local/bin/

ENTRYPOINT ["$PACKAGE"]