FROM golang
RUN mkdir -p /app
COPY . /app
RUN go get golang.org/x/tools/cmd/present
ENTRYPOINT cd /app && /go/bin/present -http=":8080"
EXPOSE 8080
