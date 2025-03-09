FROM golang:1.24-alpine AS base

RUN mkdir -m777 -p /opt/go
RUN mkdir -m777 -p /.npm

WORKDIR /opt/app

RUN apk add --update git nodejs npm
RUN npm install -g esbuild

ENV GOBIN=/go/bin
ENV PATH=$GOBIN:$PATH

COPY go.mod go.sum ./
RUN go install github.com/a-h/templ/cmd/templ@v0.3.833 && go mod download

COPY web/package.json web/package-lock.json ./web/
RUN npm install --prefix /opt/app/web

FROM base AS development

RUN go install github.com/air-verse/air@latest

COPY . .

CMD ["tail", "-f", "/dev/null"]

FROM base AS testing

COPY . .

CMD ["tail", "-f", "/dev/null"]

FROM base AS production

COPY . .

CMD ["tail", "-f", "/dev/null"]
