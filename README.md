# nats-streaming-server-embedded

This code demonstrates that creating a nats standalone server and subscription in the same projects, causes a protobuff error:

```
2019/06/17 12:03:40 proto: duplicate proto type registered: pb.PubMsg
2019/06/17 12:03:40 proto: duplicate proto type registered: pb.PubAck
2019/06/17 12:03:40 proto: duplicate proto type registered: pb.MsgProto
2019/06/17 12:03:40 proto: duplicate proto type registered: pb.Ack
2019/06/17 12:03:40 proto: duplicate proto type registered: pb.ConnectRequest
2019/06/17 12:03:40 proto: duplicate proto type registered: pb.ConnectResponse
2019/06/17 12:03:40 proto: duplicate proto type registered: pb.Ping
2019/06/17 12:03:40 proto: duplicate proto type registered: pb.PingResponse
2019/06/17 12:03:40 proto: duplicate proto type registered: pb.SubscriptionRequest
2019/06/17 12:03:40 proto: duplicate proto type registered: pb.SubscriptionResponse
2019/06/17 12:03:40 proto: duplicate proto type registered: pb.UnsubscribeRequest
2019/06/17 12:03:40 proto: duplicate proto type registered: pb.CloseRequest
2019/06/17 12:03:40 proto: duplicate proto type registered: pb.CloseResponse
panic: proto: duplicate enum registered: pb.StartPosition

goroutine 1 [running]:
github.com/mvcatsifma/nats-streaming-server-embedded/vendor/github.com/gogo/protobuf/proto.RegisterEnum(...)
	/Users/tdelnoij/go/src/github.com/mvcatsifma/nats-streaming-server-embedded/vendor/github.com/gogo/protobuf/proto/properties.go:509
github.com/mvcatsifma/nats-streaming-server-embedded/vendor/github.com/nats-io/stan.go/pb.init.0()
	/Users/tdelnoij/go/src/github.com/mvcatsifma/nats-streaming-server-embedded/vendor/github.com/nats-io/stan.go/pb/protocol.pb.go:263 +0x4e7
```
