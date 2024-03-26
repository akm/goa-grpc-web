# See https://github.com/grpc/grpc-web/blob/master/net/grpc/gateway/docker/envoy/Dockerfile

FROM envoyproxy/envoy:v1.22.0

COPY envoy/envoy.yaml /etc/envoy/envoy.yaml

ENTRYPOINT [ "/usr/local/bin/envoy" ]
CMD [ "-c /etc/envoy/envoy.yaml", "-l trace", "--log-path /tmp/envoy_info.log" ]
