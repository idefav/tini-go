FROM golang:alpine as builder
WORKDIR /app
COPY main.go .
COPY ./go.mod .
RUN ls -l
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM ubuntu:18.04 as run
COPY ./tini-go /tini-go
COPY --from=0 ./app .
ENTRYPOINT ["/tini-go","--"]
CMD ["./app"]