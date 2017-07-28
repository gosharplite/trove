FROM google/debian:wheezy
ADD trove trove
VOLUME /trove_files
ENTRYPOINT ["/trove"]
