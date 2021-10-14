## Authserver

#### How to run server
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