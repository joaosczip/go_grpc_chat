# Go gRPC chat

A simple chat made with golang and gRPC

## How it works

[gRPC](https://grpc.io/) is a rpc framework made by Google.

Your goal is to delivery a simple and high performance communication tool.
With gRPC we can build distributed systems with any programming language, and then, generate "stubs" that give to us a contracts to communicate between these services (a great advantage over REST)

---
**Bidirectional Streaming**

As gRPC works over [HTTP/2](https://developers.google.com/web/fundamentals/performance/http2?hl=pt-br), it's possible to do **multiplexing** communication between the client and the server, in this way, making possible **bidirectional** communications (when the client sends a stream to the server and receives another).

From the gPRC docs: 
> Bidirectional streaming RPCs where both sides send a sequence of messages using a read-write stream

---

The image below illustrates how this process works.

![gRPC Bidirectional Streaming](https://miro.medium.com/max/2400/1*Ug3CAac6nPclg87bxmRBoA.png)
<p align = "center">
https://levelup.gitconnected.com/grpc-how-to-make-bi-directional-streaming-calls-70b4a0569b5b
</p>
