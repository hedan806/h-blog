FROM docker.antfact.com/platform/alpine:3.7

RUN apk upgrade && apk add --no-cache ca-certificates
WORKDIR /app
COPY h-blog-unix /app/

EXPOSE 5000

ENTRYPOINT ["./h-blog-unix"]