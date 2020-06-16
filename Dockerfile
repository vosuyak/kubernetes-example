# Define the base image
FROM golang:1.14.1
# Create an app directory to hold container sources
RUN mkdir /app
# Copy EVERYTHING from the root, to the container (local -> container)
ADD . /app
# Officially define the work app location 
WORKDIR /app
# Build the app into binary
RUN go build -o main .
# Run the binary executable ( workdir / main.go !.go)
CMD ["/app/main"]