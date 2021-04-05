FROM golang:1.16-alpine3.13 AS build

WORKDIR /build

COPY server/go.mod server/go.sum ./
RUN go mod download

COPY server ./
RUN CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo -trimpath -ldflags '-s -w -extldflags "-static"'


FROM alpine:3.13

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=build /build/server /app/server

CMD [ "/app/server" ]
