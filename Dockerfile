FROM golang:1.19.2-alpine as build-env
RUN apk add build-base
RUN go install -v github.com/diego95root/nuclei/v2/cmd/nuclei@latest

FROM alpine:3.16.2
RUN apk add --no-cache bind-tools ca-certificates chromium
COPY --from=build-env /go/bin/nuclei /usr/local/bin/nuclei
ENTRYPOINT ["nuclei"]
