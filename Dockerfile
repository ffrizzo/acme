FROM golang:alpine as builder

RUN apk update \
    && apk add git

COPY ./ /go/src/github.com/ffrizzo/acme

WORKDIR /go/src/github.com/ffrizzo/acme

RUN go get -u github.com/golang/dep/cmd/dep \
    && dep ensure    

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/acme cmd/main.go

FROM scratch
EXPOSE 7070
COPY --from=builder /go/src/github.com/ffrizzo/acme/bin/acme .
CMD ["./acme"]