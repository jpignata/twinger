FROM golang:alpine AS build
RUN apk add --update make git
WORKDIR /src
ADD main.go Makefile /src/
RUN make

FROM alpine
ADD https://curl.haxx.se/ca/cacert.pem /etc/ssl/certs/
EXPOSE 79
WORKDIR /app
COPY --from=build /src/twinger /app/
CMD ["/app/twinger"]
