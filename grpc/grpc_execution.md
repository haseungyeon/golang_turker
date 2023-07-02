### 프로젝트 폴더구조

```pre
grpc-go
    └─greet
        ├─client
        ├─proto
        └─server
```

### protocol buffers 컴파일

```powershell
protoc -Igreet/proto --go_out=. --go_opt=module=github.com/Clement-Jean/grpc-go-course --go-grpc_out=. --go-grpc_opt=module=github.com/Clement-Jean/grpc-go-course .\greet\proto\dummy.proto
```


주어진 명령어는 protoc 컴파일러를 사용하여 gRPC와 Protocol Buffers를 사용하는 파일을 컴파일하는 명령어입니다. 각 옵션과 인자의 의미는 다음과 같습니다:

- protoc: Protocol Buffers 컴파일러인 protoc를 실행하는 명령어입니다.
- Igreet/proto: greet/proto 디렉토리를 포함한 추가적인 import 경로를 지정합니다. 이 경로는 컴파일러가 필요한 다른 프로토콜 버퍼 파일을 찾는 데 사용됩니다.
- --go_out=.: Go 언어용으로 Protocol Buffers 파일을 컴파일하여 .pb.go 파일을 현재 디렉토리에 출력합니다. 이 파일은 Protocol Buffers 메시지를 Go 코드로 사용할 수 있게 해줍니다.
- --go_opt=module=github.com/Clement-Jean/grpc-go-course/greet/proto: Go 모듈 경로를 지정합니다. 이 옵션은 Go 모듈을 사용하는 경우에만 필요하며, 컴파일된 Go 코드의 패키지 경로를 설정합니다.
- --go-grpc_out=.: gRPC용으로 Protocol Buffers 파일을 컴파일하여 .pb.gw.go 파일을 현재 디렉토리에 출력합니다. 이 파일은 gRPC 서비스를 Go 코드로 사용할 수 있게 해줍니다.
- --go-grpc_opt=module=github.com/Clement-Jean/grpc-go-course/greet/proto: Go 모듈 경로를 지정합니다. 이 옵션은 Go 모듈을 사용하는 경우에만 필요하며, 컴파일된 gRPC 코드의 패키지 경로를 설정합니다.
- .\greet\proto\dummy.proto: 컴파일할 Protocol Buffers 파일의 경로와 파일 이름입니다. 이 경우, greet/proto 디렉토리에 위치한 dummy.proto 파일을 컴파일합니다.

따라서, 주어진 명령어는 greet/proto 디렉토리에 위치한 dummy.proto 파일을 컴파일하여 Go 언어와 gRPC용 코드를 현재 디렉토리에 출력합니다. 컴파일된 Go 파일은 Protocol Buffers 메시지와 gRPC 서비스를 사용할 수 있게 해줍니다. --go_opt와 --go-grpc_opt 옵션을 사용하여 Go 모듈 경로를 지정할 수 있습니다. 이렇게 하면 컴파일된 Go 코드의 패키지 경로가 설정되며, 모듈을 사용하는 경우에 유용합니다.