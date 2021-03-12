FROM golang:1.15.10 as builder

RUN mkdir /app
COPY ./source/xmen /app
WORKDIR /app


# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
######## Start a new stage from scratch #######
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .


EXPOSE 3009

#CMD go run main.go
# Command to run the executable
CMD ["./main"] 