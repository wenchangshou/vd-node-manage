FROM golang:latest

WORKDIR /build
COPY  . .
RUN go mod init server
RUN go mod tidy
RUN go build -o app .
WORKDIR /app
RUN cp /build/app  .

EXPOSE 6031
CMD ["/server/app"]
