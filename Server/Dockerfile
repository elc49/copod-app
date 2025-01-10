FROM golang:1.23

RUN mkdir -p go/src/copod
WORKDIR /go/src/copod
COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux go build -v -o copod

EXPOSE 4545
EXPOSE 4434

CMD ["./copod"]
