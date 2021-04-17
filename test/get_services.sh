grpcurl -plaintext -d '{"type": "chat"}' \
    127.0.0.1:9100 cloudcade.grpc.center.CenterService/GetServices