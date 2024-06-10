# Use the official Alpine image as the base
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the built Go binary from the local machine to the container
COPY /bin/de .

# Command to run the executable
CMD ["./de"]