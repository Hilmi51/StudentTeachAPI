FROM golang:1.20

WORKDIR /go/src/student-teach-api

COPY . .

RUN go get -v

RUN go build -o main .

CMD ["./main"]