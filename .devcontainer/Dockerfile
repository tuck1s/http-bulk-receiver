# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /workspace

# Install common development tools
RUN apt-get update && apt-get -y install git curl

RUN go build