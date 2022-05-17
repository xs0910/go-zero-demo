1、启动etcd 

2、启动 user rpc

```shell
cd mall/user/rpc
go run user.go
```

3、启动 order api

```shell
cd mall/order/api
go run order.go
```

4、访问order api

```shell
curl -i -X GET http://localhost:8888/api/order/get/1
```

