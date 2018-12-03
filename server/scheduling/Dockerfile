FROM alpine
RUN apk add --no-cache ca-certificates
COPY scheduling /scheduling
EXPOSE 80
ENTRYPOINT ["/scheduling"]