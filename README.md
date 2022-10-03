# goctl-rest-discover

* 生成基础框架
```shell
goctl api go -api exa.api -dir . --style go_zero
```
* 为api生成调用客户端
```shell
goctl api plugin -p goctl-rest-discover="rest-discover" -api exa.api -dir .
```

调用api像rpc一样简单

