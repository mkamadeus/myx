# build
FROM golang:alpine as builder

RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy 
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myx .

# run
FROM python:3.9-alpine
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/myx .

ENTRYPOINT ["./myx"]