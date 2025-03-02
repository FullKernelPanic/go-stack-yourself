FROM golang:1.24-alpine AS base

WORKDIR /opt/app

RUN apk add --update git nodejs npm
RUN npm install -g esbuild

ENV GOBIN=/go/bin
ENV PATH=$GOBIN:$PATH

COPY go.mod go.sum ./

RUN go install github.com/a-h/templ/cmd/templ@v0.3.833 && \
    go mod download

FROM base AS dev

RUN go install github.com/air-verse/air@latest

COPY . .

CMD ["tail", "-f", "/dev/null"]

FROM base AS test

COPY . .

CMD ["tail", "-f", "/dev/null"]

FROM base AS prod

COPY . .

CMD ["tail", "-f", "/dev/null"]
