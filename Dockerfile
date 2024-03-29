FROM centos:centos7

ENV ServiceName cd_cmdb

WORKDIR /opt

ADD $ServiceName /opt
ADD entrypoint.sh /opt
RUN set -ex \
&& chmod u+x ./$ServiceName \
&& chmod u+x ./entrypoint.sh \
&& mkdir -p /var/log/go_log

ENTRYPOINT ["./entrypoint.sh"]
