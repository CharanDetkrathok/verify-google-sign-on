FROM golang:1.18.3-alpine3.16
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o main .
EXPOSE 8885
CMD [ "/app/main" ]

