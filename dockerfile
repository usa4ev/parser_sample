FROM golang:1.19.6-alpine3.17

RUN apk update \
# Install git
    && apk add --no-cache --upgrade git \
# Install make
    && apk add --no-cache make 

EXPOSE 8080
EXPOSE 9090

#Get server src
WORKDIR /parser_sample 
RUN git clone https://github.com/usa4ev/parser_sample ./\
    # Build server
    && make build-srv-linux && chmod +x ./bin/ParserServer-linux\
    # Cleanup
    && apk del git && apk del make
 
CMD ["./bin/ParserServer-linux"] 



