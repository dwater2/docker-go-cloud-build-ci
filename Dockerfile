FROM golang:1.12-alpine as builder
ADD src/sum /src/sum
WORKDIR /src/sum
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w"

FROM scratch
ADD src/sum /src/sum
WORKDIR /src/sum
COPY --from=builder /src/sum .
CMD ["./sum"]