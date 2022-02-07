FROM golang:latest
WORKDIR /build
RUN git clone https://github.com/naumanahmedkhan/goRestAPI.git
WORKDIR /build/goRestAPI/main
RUN go build
EXPOSE 5050
CMD ./main

# docker run --network host -p 5050:5050 gonomi