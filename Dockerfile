FROM golang

ADD . /redir

RUN cd /redir && go build redir.go

ENTRYPOINT /redir/redir

EXPOSE 8888

