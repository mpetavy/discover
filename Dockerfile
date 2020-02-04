FROM golang:1.13.7
RUN mkdir /src
WORKDIR /src
RUN git clone https://github.com/mpetavy/discover
WORKDIR /src/discover
RUN go get -u ./...
RUN go build -o discover
CMD ["-c",":15000","-info","HelloWorld","-service","simulate"]
ENTRYPOINT ["discover"]

