FROM golang

WORKDIR /app

RUN export GO111MODULE=on

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]