# Confluence REST service on Golang

https://developer.atlassian.com/server/confluence/confluence-rest-api-examples/

Groovy version - https://github.com/AndriiMaliuta/confluence-rest-service-groovy

get services
```
grpcurl -plaintext localhost:9093 list
```

getPage
```curl
./grpcurl -d '{"id": "123123"}' -plaintext localhost:50051 andmal.PageService.GetPage
```