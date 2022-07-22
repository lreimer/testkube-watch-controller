FROM gcr.io/distroless/static-debian11
COPY testkube-watch /

ENTRYPOINT ["/testkube-watch"]
CMD ["version"]
