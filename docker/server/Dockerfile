FROM golang:1.16 AS build
WORKDIR /build

COPY server/go.mod server/go.sum ./
RUN go mod download

COPY server ./
RUN CGO_ENABLED=0 go build -trimpath -ldflags "-s -w -buildid="


FROM gcr.io/distroless/static:latest
WORKDIR /app

COPY --from=build /build/server /app/server

CMD [ "/app/server" ]
