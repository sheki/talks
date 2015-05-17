FROM golang:onbuild 
RUN mkdir -p :/app
COPY *.slide /app
COPY images /app
RUN go get golang.org/x/tools/cmd/present
ENTRYPOINT cd /app && /go/bin/present -http=":8080"
EXPOSE 8080
