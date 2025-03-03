FROM golang:1.24

WORKDIR /app
COPY . .

WORKDIR /app/backend
RUN go mod download && go mod verify

RUN go build -o app ./src
CMD ["./app"]
