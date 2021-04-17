grpcurl -plaintext -d '{"name": "tom"}' \
    127.0.0.1:9100 CenterService/GetUserOnlineTime
