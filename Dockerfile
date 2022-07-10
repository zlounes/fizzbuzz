FROM redhat/ubi8-minimal


RUN microdnf install shadow-utils procps less nc vi  &&\
    rm -rf /var/cache/yum

USER root
RUN groupadd --gid 1000 fizzbuzz && \
    useradd --uid 1000  fizzbuzz -g fizzbuzz

RUN mkdir /go && microdnf install golang && microdnf clean all

WORKDIR /go/src/app
COPY . .

RUN go build -o fizzbuzz  main.go
RUN chown -R fizzbuzz:fizzbuzz .

#RUN mkdir fizzbuzz
#ENV GOBIN=/go/src/app/bin/
#RUN go get -d -v ./...
#RUN go install -v ./...

USER fizzbuzz

WORKDIR /go/src/app
EXPOSE 8080
CMD ["./fizzbuzz"]


