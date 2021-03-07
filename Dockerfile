FROM golang:1.16-alpine as build-env

RUN apk add --update --no-cache ca-certificates git

WORKDIR /go/src/github.com/army4d/housing-break-even-calculator

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /go/bin/calculator

FROM scratch

COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=build-env /go/bin/calculator /go/bin/calculator

COPY --from=build-env /go/src/github.com/army4d/housing-break-even-calculator/config ${HOME}/.config/calculator

EXPOSE 50051

ENTRYPOINT ["/go/bin/calculator"]
