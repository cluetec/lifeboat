FROM scratch
ENTRYPOINT ["/app/lb", "backup"]
COPY LICENSE README.md /app/
COPY ./out/lb /app/
