FROM alpine
RUN apk add --no-cache ca-certificates
COPY gateway /gateway
EXPOSE 80
ENTRYPOINT ["/gateway"]