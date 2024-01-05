FROM scratch
WORKDIR /app
ENTRYPOINT ["/app/lb"]
COPY out/lb lb
