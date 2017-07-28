# trove
A simple file server.

[Dockerfile]
FROM google/debian:wheezy
ADD trove trove
VOLUME /trove_files
ENTRYPOINT ["/trove"]

CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .

sudo docker build -t gosharplite/trove:v4 .

sudo docker run --publish 8080:8080 -v /trove_files:/trove_files gosharplite/trove:v4
