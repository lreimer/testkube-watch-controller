FROM gcr.io/distroless/static-debian11

ENV TKW_HOME=/

COPY dist/testkube-watch_linux_amd64_v1/testkube-watch /

ENTRYPOINT ["/testkube-watch"]
