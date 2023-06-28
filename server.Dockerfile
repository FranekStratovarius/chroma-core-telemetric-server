FROM golang AS builder
WORKDIR /builder
COPY ./server /builder
RUN go get
RUN go build
RUN ls -al

FROM ubuntu:latest
WORKDIR /server
COPY --from=builder /builder/chroma-core-telemetric-server ./chroma-core-telemetric-server
RUN ls -al
RUN pwd
RUN chmod +x ./chroma-core-telemetric-server
ENTRYPOINT ["./chroma-core-telemetric-server"]
EXPOSE 1234