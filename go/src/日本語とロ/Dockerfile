FROM golang:latest
RUN export GOROOT=/usr/local/go
RUN export GOPATH=$HOME/Projects/Proj1
RUN export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
COPY kanji.go /home/
WORKDIR /home/
