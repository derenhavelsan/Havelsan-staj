FROM golang:latest


# Set destination for COPY
WORKDIR /app

# Download Go modules -> bunları tanımlamadığımız zaman build edemiyor RUN go build -o çalışmıyor..
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy go codes
COPY main/ .

# Build -> paketlerin dependency'lerine göre compile edilmesini sağlıyor
RUN go build -o /go-get-request

# This is for documentation purposes only.
# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 8080

# (Optional) environment variable that our dockerised
# application can make use of. The value of environment
# variables can also be set via parameters supplied
# to the docker command on the command line.
#ENV HTTP_PORT=8081

# Run
CMD [ "/go-get-request" ]


