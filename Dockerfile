FROM golang

RUN mkdir -p /clearmove/studapi
ADD . /clearmove/studapi/
WORKDIR /clearmove/studapi
RUN go build -o studapi
EXPOSE 8080
CMD ['/clearmove/studapi']
