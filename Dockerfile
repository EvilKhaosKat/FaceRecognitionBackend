FROM umputun/baseimage:buildgo-latest as builder
#FROM golang:1.12 as builder
WORKDIR /go/FaceRecognitionBackend
COPY . /go/FaceRecognitionBackend
RUN go build ./cmd/web

FROM umputun/baseimage:app-latest
#FROM amd64/alpine:latest
WORKDIR /srv
COPY --from=builder /go/FaceRecognitionBackend/web /srv/web
RUN \
    chown -R app:app /srv && \
    chmod +x /srv/web
CMD ["/srv/web"]

EXPOSE 10080