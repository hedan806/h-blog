FROM docker.antfact.com/platform/alpine:3.7

RUN apk upgrade && apk add --no-cache ca-certificates
WORKDIR /app
COPY . .

EXPOSE 5000

ENTRYPOINT ["./h-blog-unix"]