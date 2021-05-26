FROM debian:bullseye

ARG GO_GILE=go1.16.4.linux-amd64.tar.gz
ARG SRC_MODULE=github.com/mkaiho/go-deploy-sample

ENV GOPATH /go
ENV PATH $PATH:/usr/local/go/bin:$GOPATH/bin

RUN groupadd -r devusers && useradd -r -g devusers devuser

RUN apt-get update && apt-get install -y \
make curl git

RUN curl -OL https://golang.org/dl/${GO_GILE} \
&& tar -C /usr/local -xzf ${GO_GILE} \
&& rm ${GO_GILE}

RUN mkdir -p $GOPATH/src/$SRC_MODULE && chmod -R 777 $GOPATH

WORKDIR $GOPATH/src/$SRC_MODULE
