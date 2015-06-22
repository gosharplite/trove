FROM google/debian:wheezy
ADD trove trove
VOLUME /trove_files
EXPOSE 8080
ENTRYPOINT ["/trove"]