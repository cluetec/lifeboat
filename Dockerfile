FROM scratch
WORKDIR /app
ENTRYPOINT ["/app/lb"]
COPY LICENSE README.md ./
COPY out/lb lb
