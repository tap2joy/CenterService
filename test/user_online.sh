grpcurl -plaintext -d '{"name": "tom", "gate": "127.0.0.1:9101"}' \
    127.0.0.1:9100 cloudcade.grpc.center.CenterService/UserOnline

grpcurl -plaintext -d '{"name": "lucy", "gate": "127.0.0.1:9101"}' \
    127.0.0.1:9100 cloudcade.grpc.center.CenterService/UserOnline

grpcurl -plaintext -d '{"name": "jack", "gate": "127.0.0.1:9102"}' \
    127.0.0.1:9100 cloudcade.grpc.center.CenterService/UserOnline