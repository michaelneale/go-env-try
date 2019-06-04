FROM scratch
EXPOSE 8080
ENTRYPOINT ["/go-env-try"]
COPY ./bin/ /