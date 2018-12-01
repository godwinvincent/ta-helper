FROM alpine
RUN apk add --no-cache ca-certificates
COPY email /email
EXPOSE 80
ENTRYPOINT ["/email"]