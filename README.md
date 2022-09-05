创建 mall 服务
mkdir goecom
cd goecom
go mod init goecom

lib model:
goctl model mysql ddl -src merchant.sql -dir . -c -style goZero
goctl model mysql ddl -src merchant_user.sql -dir . -c -style goZero
goctl model mysql ddl -src store.sql -dir . -c -style goZero
goctl model mysql ddl -src languages.sql -dir . -c -style goZero

product model:
goctl model mysql ddl -src product_brand.sql -dir . -c -style goZero

api:
goctl api go -api product.api -dir . -style goZero
goctl api go -api lib.api -dir . -style goZero

#install swagger
brew tap go-swagger/go-swagger
brew install go-swagger
#intsll goctl-swagger
https://github.com/zeromicro/goctl-swagger
goctl api plugin -plugin goctl-swagger="swagger -filename lib.json" -api lib.api -dir . -style goZero #生成 Swagger json 文件
goctl api plugin -plugin goctl-swagger="swagger -filename lib.json -host 127.0.0.1:6001" -api lib.api -dir . #生成 Swagger json 文件 指定 host
生成 rpc 代码：
goctl rpc protoc lib.proto --go_out=./types --go-grpc_out=./types --zrpc_out=. -style goZero

各位好，问一下:
现在我们搭的项目， 写好 API 后用下面命令生成 API Json 文件；
goctl api plugin -plugin goctl-swagger="swagger -filename lib.json" -api lib.api -dir . -style goZero
后台管理用的是 Ant pro design, 把上面生成的 Json 文件放到 config 文件夹用 npm run openapi 生成 API 调用的 JS 代码；
现在遇到的问题是， 在没有 JWT Token 时候是可以的， 加入 JWT 后

各位好，问一下:
goctl api plugin -plugin goctl-swagger="swagger -filename lib.json" -api lib.api -dir . -style goZero
这个生命令有没有方法把 JWT token 的 header 参数生成进去
