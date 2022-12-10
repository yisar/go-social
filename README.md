# 3n5b


## 扩展安装
```shell
go get -u github.com/gin-gonic/gin
go get go.mongodb.org/mongo-driver/mongo
go get github.com/dgrijalva/jwt-go
```

## Docker 安装 mongoDB
```shell
docker run -d --name mongo -e MONGO_INITDB_ROOT_USERNAME=ht -e MONGO_INITDB_ROOT_PASSWORD=123456 -p 27017:27017 mongo
```