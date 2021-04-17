grpcurl -plaintext -d '{"type": "chat", "address": "127.0.0.1:9101"}' \
    127.0.0.1:9100 cloudcade.grpc.center.CenterService/RegisterService

grpcurl -plaintext -d '{"type": "chat", "address": "127.0.0.1:9102"}' \
    127.0.0.1:9100 cloudcade.grpc.center.CenterService/RegisterService

grpcurl -plaintext -d '{"type": "chat", "address": "127.0.0.1:9103"}' \
    127.0.0.1:9100 cloudcade.grpc.center.CenterService/RegisterService