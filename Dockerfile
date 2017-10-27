FROM scratch

MAINTAINER skuenzli@qualimente.com

COPY bin/linux/amd64/heartbeat /heartbeat

ENTRYPOINT ["/heartbeat"]
