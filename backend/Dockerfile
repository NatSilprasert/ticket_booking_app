FROM golang:1.25-alpine

WORKDIR /src/app

RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod tidy

CMD ["sh", "-c", "until pg_isready -h db -U $DB_USER; do echo waiting for db; sleep 2; done; air -c .air.toml"]