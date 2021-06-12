FROM golang

WORKDIR /app

COPY . .
RUN go mod download
RUN go install

ENTRYPOINT /go/bin/runestats

EXPOSE 8080
