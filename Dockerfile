FROM golang:1.24.2

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download && go mod verify
COPY . .
RUN go mod tidy
RUN go build -v -o main .

CMD ["/app/main"]
