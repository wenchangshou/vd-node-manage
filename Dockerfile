FROM golang:1.17.5-alpine as builder
RUN apk add --no-cache ca-certificates git
RUN apk add build-base
WORKDIR /src
COPY go.mod  go.sum ./
RUN go mod download
COPY . .
RUN ls
ARG SKAFFOLD_GO_GCFLAGS
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o /go/bin/server ./module/server

FROM alpine as release
RUN apk add --no-cache ca-certificates \
    busybox-extras net-tools bind-tools
WORKDIR /src
COPY --from=builder /go/bin/server /src/server
COPY ./module/server/cfg-docker.json /src/cfg.json
EXPOSE 6031
ENTRYPOINT ["/src/server"]