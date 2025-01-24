FROM golang:1.22

WORKDIR /app/

COPY --chown=app:app ./app .
COPY --chown=app:app ./cli .

RUN go mod download && go mod verify

RUN go build -v -o / ./...

ENV PATH="/app:${PATH}"

CMD ["app"]