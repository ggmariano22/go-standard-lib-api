# First step build (runtime)
FROM golang:alpine AS build

RUN apk update && apk add --no-cache git bash curl

WORKDIR /src

COPY . .

RUN go install github.com/cosmtrek/air@latest

RUN go mod download
RUN go mod verify

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o api ./cmd/main.go

CMD ["air", ".air.toml"]

# Second step build (binary)
#FROM scratch

#COPY --from=build src/api .

#ENTRYPOINT ["air", ".air.toml"]