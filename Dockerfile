FROM golang:1.16-alpine

EXPOSE 80

WORKDIR /app

COPY go.* .
RUN go mod download

COPY *.go .

RUN go build -o ./whoami

CMD [ "./whoami" ]