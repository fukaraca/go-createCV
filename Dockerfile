FROM golang:1.18-alpine as build_base

WORKDIR /src

RUN apk add --no-cache --update \
  git  \
    ca-certificates

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base AS server_builder

COPY . /src



RUN go install ./

FROM surnet/alpine-wkhtmltopdf:3.15.0-0.12.6-small as wkhtmltopdf

FROM alpine AS projectservice

RUN apk add --no-cache libstdc++ libx11 libxrender libxext libssl1.1 fontconfig freetype ttf-dejavu ttf-droid ttf-freefont ttf-liberation && apk add --no-cache --virtual .build-deps msttcorefonts-installer && update-ms-fonts && fc-cache -f && rm -rf /tmp/* && apk del .build-deps

RUN apk add ca-certificates

COPY --from=server_builder /go/bin/go-createCV /bin/go-createCV
COPY --from=server_builder /src/ /src
COPY --from=wkhtmltopdf /bin/wkhtmltopdf /bin/wkhtmltopdf
WORKDIR /src

EXPOSE 8080
ENTRYPOINT /bin/go-createCV


