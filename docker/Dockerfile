FROM golang:stretch as builder
# VictoriaLogs Builder
RUN apt update && apt install -y git wget
RUN cd /tmp && git clone https://github.com/faceair/VictoriaLogs
WORKDIR /tmp/VictoriaLogs
RUN go mod download && make all
RUN echo '#!/bin/bash \n\
/usr/bin/vmstorage && \n\
/usr/bin/vmselect -storageNode 127.0.0.1:8401 && \n\
/usr/bin/vminsert -storageNode 127.0.0.1:8400 -importerListenAddr 127.0.0.1:2003 \n'\
> /tmp/entrypoint.sh

FROM debian:stretch
# VictoriaLogs Stripped
WORKDIR /root/
COPY --from=builder /tmp/VictoriaLogs/bin/vm* /usr/bin/
COPY --from=builder /tmp/entrypoint.sh /
RUN chmod +x /entrypoint.sh
EXPOSE 8401
ENTRYPOINT ["/entrypoint.sh"]
