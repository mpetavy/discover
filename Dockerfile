FROM alpine:latest
EXPOSE 15000/udp
WORKDIR /
COPY ./discover /
ENTRYPOINT ["/discover"]
CMD ["-c",":15000","-service","simulate","-info","Hello World :-)"]

