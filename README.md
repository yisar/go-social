# 3n5b


## 扩展安装
```shell
go get -u github.com/gin-gonic/gin
go get go.mongodb.org/mongo-driver/mongo
go get github.com/dgrijalva/jwt-go
```

## Docker 安装 mongoDB
```shell
docker run -d --network some-network --name some-mongo -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=admin -p 27017:27017 mixixibi
```
