FROM redhat/ubi8-minimal as builder

RUN mkdir /go && microdnf install golang && microdnf clean all

WORKDIR /go/src/app
COPY . .

RUN go build -o fizzbuzz  main.go

FROM redhat/ubi8-minimal 

RUN microdnf install shadow-utils procps less nc vi  &&\
    rm -rf /var/cache/yum

USER root
RUN groupadd --gid 1000 fizzbuzz && \
    useradd --uid 1000  fizzbuzz -g fizzbuzz

COPY --chown=mycomosi:mycomosi --from=builder /go/src/app/fizzbuzz /usr/bin/fizzbuzz

WORKDIR /go/src/app
COPY . .

USER fizzbuzz

WORKDIR /usr/bin/
RUN echo "server port : $SERVER_PORT" 
CMD ./fizzbuzz -port $SERVER_PORT


