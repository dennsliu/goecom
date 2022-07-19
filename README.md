创建 mall 服务
mkdir goecom
cd goecom
go mod init goecom

model:
goctl model mysql ddl -src product_brand.sql -dir . -c
goctl model mysql ddl -src merchant.sql -dir . -c
goctl model mysql ddl -src merchant_user.sql -dir . -c

api:
goctl api go -api product.api -dir .
goctl api go -api lib.api -dir .

$ curl -i -X POST \
 http://127.0.0.1:6001/v1/lib/merchant/user/login \
 -H 'Content-Type: application/json' \
 -d '{
"username":"testing1",
"password":"123456"
}'
