FROM golang:alpine

RUN apk update && apk add git

ARG DEFAULT_PORT

ENV PORT=${DEFAULT_PORT}
ENV INSTANCE_ID="dari dockerfile"

RUN echo "isi dari argument DEFAULT_PORT = ${DEFAULT_PORT}"
RUN echo "isi dari env var PORT = ${PORT}"
RUN echo "isi dari env var INSTANCE_ID= ${INSTANCE_ID}"

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o binary

ENTRYPOINT [ "./binary" ]
