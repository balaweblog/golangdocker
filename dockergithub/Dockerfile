FROM golang:1.7

ADD . /home
        
WORKDIR /home

RUN \
       go get -u github.com/balaweblog/gosqlapi

CMD ["go","run","main.go"]