# syntax=docker/dockerfile:1.4
FROM node:18.1.0 AS client
WORKDIR /workspace

RUN mkdir client
COPY client/package.json client/yarn.lock ./client
RUN <<EOF
  cd client
  yarn install --frozen-lockfile
EOF

COPY client ./client
RUN <<EOF
  cd client
  yarn build
EOF


FROM golang:1.18.1 AS server
WORKDIR /workspace

RUN mkdir server
COPY server/go.mod server/go.sum ./server
RUN <<EOF
  cd server
  go mod download
EOF

COPY server server
COPY --from=client /workspace/client/dist /workspace/server/assets

RUN <<EOF
  cd server
  CGO_ENABLED=0 go build -trimpath -ldflags "-s -w -buildid="
EOF


FROM gcr.io/distroless/static
WORKDIR /workspace
COPY --from=server /workspace/server/server ./
CMD ["/workspace/server"]
