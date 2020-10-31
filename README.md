# golang_grpc_client_server_stream


### Golang:
   First published in 2009, Go is an open-source programming language developed by a team at Google and the combined effort of other contributors. It is meant to simplify the process of software development, particularly for complex architecture and processes.

#### Advantages:
1. Golang Is Fast
2. Golang Is Easy To Learn
3. Golang Is Well-Scaled
4. Comprehensive Programming Tools
5. Growing Pool Of Talent

### gRPC:
   gRPC is a high performance, open source universal RPC Framework.In simple words, it enables the server and client applications to communicate transparently and build connected systems. GRPC is developed and open sourced by Google.

#### RPC:
   The remote Procedure call is quite an interesting concept. Remote Procedure Call is a high-level model for client-server communication.Assume there are two computers, computer A(on local) and computer B(on some network). Computer B provides some API’s, let’s say it has some procedures which can be run and these procedures can be run on computer B itself.

#### GRPC Works:
   A client application is able to call methods directly on a server-side application present on other machines.
   Service is defined, methods are specified which can be further remotely called with their parameters and return types.
   On the other hand, the server runs a GRPC server to handle client calls.
   It uses protocol buffers as the Interface Definition Language to enable communication between two different systems used for describing the service interface and the structure of payload messages.
   HTTP/2 – GRPC is basically a protocol built on top of HTTP/2. HTTP/2 is used as a transport.
   Protobuf serialization – Messages that we serialize both for the request and response are encoded with protocol buffers.    
   Clients open one long-lived connection to GRPC server.
   A new HTTP/2 stream for each RPC call.
   Allows Client-Side and Server-Side Streaming.
   Bidirectional Streaming.

#### Benefits of Adopting GRPC:
   
1. HTTP/REST Features
2. Easy to understand.
3. Web infrastructure already built on top of HTTP.
4. Great tooling for testing, inspection, and modification.
5. Loose coupling between clients/server makes changes easy.
6. High-quality HTTP implementations in every language.

#### Type of gRPC stream:

1. Unary Stream
2. Server Streaming
3. Client Streaming
4. Bidirectional Streaming

#### Why GRPC:
1. Ability to break free from the call-and-response architecture. GRPC is built on HTTP/2, which supports traditional Request/Response model and bidirectional streams.
2. Switch from JSON to protocol buffers.
3. Multiplexing.
4. Duplex Streaming.
5. Because of binary data format, it gets much lighter.
6. Polyglot.


#### Server Streaming:
   In this type of streaming, Server delivers the content or information to the client either live or on demand.

#### Client Streaming:
   In Client streaming, let’s say an example, a client writes a sequence of messages and after that send them to the server as a stream. And then the client will wait for the server to read them and return the response.


### Implmented gRPC methods:

#### Server Streaming:

    GetPlants(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (PlantService_GetPlantsClient, error)

#### Unary stream:

    GetPlant(ctx context.Context, in *PlantID, opts ...grpc.CallOption) (*Plant, error)
	
#### Client Streaming:

    CreatePlant(ctx context.Context, opts ...grpc.CallOption) (PlantService_CreatePlantClient, error)
    UpdatePlant(ctx context.Context, opts ...grpc.CallOption) (PlantService_UpdatePlantClient, error) 
    DeletePlant(ctx context.Context, opts ...grpc.CallOption) (PlantService_DeletePlantClient, error)


### Used Packages:

    go get -u google.golang.org/grpc
    go get -u github.com/golang/protobuf/protoc-gen-go


### Commands:

#### make clean
   Remove go files in the pb directory

#### make gen
   Generate go protobuf file based on proto files in the proto directory

#### make sever
   Start gRPC server usig port :9090

#### make client
   Dial gRPC server methods
