创建 mall 服务
mkdir goecom
cd goecom
go mod init goecom

lib model:
goctl model mysql ddl -src merchant.sql -dir . -c -style goZero
goctl model mysql ddl -src merchant_user.sql -dir . -c -style goZero

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

goctl api plugin -plugin goctl-swagger="swagger -filename lib.json" -api lib.api -dir . -style goZero
