# Authserver

### How to run server
- Create db
```shell script
make db
```
- Compile proto files
```shell script
make compile
```
- Run server
```shell script
make server
```

### How to test with client
Run in second terminal
```shell script
make testclient username=user password=P@ssword
```
