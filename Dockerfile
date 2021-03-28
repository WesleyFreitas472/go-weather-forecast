FROM ubuntu:18.04
ADD ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
COPY build/weather-forecast /app
COPY app.yaml /app
ENTRYPOINT ["/app/weather-forecast"]