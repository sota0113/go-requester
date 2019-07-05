FROM golang:latest AS build-env
WORKDIR /go/src/app
COPY ./main.go ./
RUN go get github.com/kelseyhightower/envconfig
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:latest
WORKDIR /go/src/app
COPY --from=build-env /go/src/app/app ./
ENV RQ_HOST="localhost"
ENV RQ_PROTOCOL="http"
ENV RQ_PATH="/list"
ENV RQ_PORT="8080"
ENV RP_PORT="8080"
CMD ["./app"]
