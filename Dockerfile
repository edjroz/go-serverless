FROM golang:1.12 as builder

# Copy local code to the container image.
WORKDIR /go/src/github.com/edjroz/go-serverless
COPY . .

ENV GO111MODULE=on
# Build the command inside the container.
RUN go get -v -d
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN CGO_ENABLED=0 GOOS=linux go build -v -o sign 

FROM alpine
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/src/github.com/edjroz/go-serverless/sign /sign

# Run the web service on container startup.
CMD ["/sign"]