FROM alpine:latest  
WORKDIR /root/
COPY whoami .
CMD ["./whoami"]  
