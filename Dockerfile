FROM golang:latest
WORKDIR /build
RUN git clone https://github.com/naumanahmedkhan/continent-api.git
WORKDIR /build/continent-api/main
RUN go build
EXPOSE 5050
CMD ./main

# docker run --network host -p 5050:5050 gonomi