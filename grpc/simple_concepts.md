## GRPC

### RPC란?

- Remote procedure call로 원격지에 존재하는 코드의 함수들을 호출하는 기술이다.

### protocol buffers role

- rest api를 전송할 때 지켜야 하는 통신 규격이 있듯이 데이터 전송 규격을 정의함.

```proto
syntax = "proto3";
message Greeting {
    string first_name = 1;
}

message GreetRequest {
    Greeting greeting = 1;
}

service GreetService {
    rpc Greet(GreetRequest) returns (GreetResponse) {}
}
```

### GRPC에서 proto buffers를 사용하는 이유

- grpc는 데이터를 주고받기 위해 프로토콜 버퍼를 이용하는데 이 프로토콜 버퍼는 더 작은 사이즈를 사용한다.

```JSON
{
    "age": 26,
    "first_name": "Clement",
    "last_name": "JEAN"
}
```
=> 이 JSON 데이터는 52byte의 크기를 차지한다.

```proto
syntax = "proto3";

message Person {
    uint32 age = 1;
    string first_name = 2;
    string last_name = 3;
}
```
=> 반면 proto buffer는 17byte 크기밖에 차지하지 않는다.

- JSON의 경우 인간이 읽기 형태이기 때문에 파싱하는데 cpu가 많이 소모되는 반면, protocol buffers는 바이너리 형태이고 실제로 메모리에서 표시되는 방식과 매우 가깝기 때문에 cpu 집약도가 낮기 때문에 훨씬 빠르다.

### GRPC api types

#### Unary

- resp api와 유사. 클라이언트는 하나의 request를 서버로 보내고 하나의 response를 받는다. 전통적인 response requirement가 요구될 때 사용. (HTTP1)

#### Server streaming

- 클라이언트는 하나의 request를 서버로 보내고 서버는 하나 이상의 response를 보낸다. real time, 문자 전송에 사용된다.

#### Client streaming

- 클라이언트는 여러개의 request를 서버로 보내고 서버는 하나의 response를 보낸다. uploading이나 정보 updating에 주로 이용된다.

#### Bi directional streaming

- 클라이언트는 다수의 request를 보내고 서버로부터 다수의 response를 받는다.

```proto
// GRPC api examples
service GreetService {
    // Unary
    rpc Greet(GreetRequest) returns (GreetResponse) {};

    //Server streaming
    rpc GreetManyTimes(GreetRequest) returns (Stream GreetResponse) {};

    // Client streaming
    rpc LongGreet(stream GreetRequest) returns (GreetResponse) {};

    // Bi directional streaming
    rpc GreetEveryone(stream GreetRequest) returns (stream GreetResponse) {};
}
```

### GRPC vs REST apis


|                |GRPC            |REST                  |
|----------------|----------------|----------------------|
|data structure  |protocol buffers|JSON                  |
|HTTP version    |HTTP2           |HTTP1                 |
|transport method|streaming       |Unary                 |
|communications  |Bi directional  |Cline->Server         |
|schema          |Free design     |GET/POST/UPDATE/DELETE|
=> 또한 HTTP2는 HTTP1에 비해 보안 측면에서 더 안정적이며 전송되는 데이터 규격이 자유롭지만 데이터 타입과 크기를 미리 지정하므로 더 strict 하다.