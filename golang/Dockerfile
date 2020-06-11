FROM golang:alpine AS builder

WORKDIR /src/webapp

COPY ./src/webapp .

RUN CGO_ENABLED=0 GOOS=linux go build -o webapp -a -installsuffix cgo -ldflags '-w -extldflags "-static"'

RUN ls

WORKDIR /bin
RUN cp /src/webapp/webapp ./webapp

FROM scratch

COPY --from=builder /bin .

ENTRYPOINT ["./webapp"]

