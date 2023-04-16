FROM golang:1.19.6-alpine3.17

RUN apk update \
# Install git
    && apk add --no-cache --upgrade git \
# Install bash
    && apk add --no-cache --upgrade bash \
# Install make
    && apk add --no-cache make 

#Get ghostorange src
WORKDIR /ghostorange 
RUN git clone https://github.com/usa4ev/parser_sample ./\
    # Build ghostorange
    && make build-srv-linux && chmod +x ./bin/ParserServer-linux\
    # Cleanup
    && apk del git && apk del bash && apk del make 

EXPOSE 8080

WORKDIR /parser_sample/bin
 
CMD ["./ParserServer-linux"] 



