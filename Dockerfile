FROM index.caicloud.io/debian:jessie
MAINTAINER zhoushaolei <shaolei@caicloud.io>

# Set the timezone to Shanghai
RUN echo "Asia/Shanghai" > /etc/timezone && \
    dpkg-reconfigure -f noninteractive tzdata && \
    sed -i "s/httpredir.debian.org/mirrors.163.com/g" /etc/apt/sources.list && \
    apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates

WORKDIR /
ADD hierarchy_exporter /

ENTRYPOINT ["/hierarchy_exporter"]
CMD ["-h"]
