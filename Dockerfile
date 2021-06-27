RUN mkdir -p /go/src/scalable-api
WORKDIR /go/src/scalable-api
COPY . /go/src/scalable-api
RUN go install scalable-api
CMD /go/bin/scalable-api
EXPOSE 8080