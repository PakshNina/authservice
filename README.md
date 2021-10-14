## Authserver

#### How to run server
- install go, set GOROOT, GOPATH, install [migrate](https://github.com/golang-migrate/migrate), protoc-gen-go, protoc-gen-go-grpc
- Create .env file from .env.example
- Create db:
```shell script
make first
```
- Apply migrations:
```shell script
make m_up
```
- Run server:
```shell script
make server
```

#### Test client
Run client in second terminal with test credentials 'user' & 'P@ssword':
```shell script
make testclient username=user password=P@ssword
```

#### Run tests
```shell script
make test
```