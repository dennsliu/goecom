$ curl -i -X POST \
 http://127.0.0.1:6001/v1/lib/gettoken \
 -H 'Content-Type: application/json' \
 -d '{
"id":1
}'

$ curl -i -X POST \
 http://127.0.0.1:6001/v1/lib/merchant/user/login \
 -H 'Content-Type: application/json' \
 -d '{
"username":"testing1",
"password":"123456"
}'

$ curl -i -X POST \
 http://127.0.0.1:6001/v1/lib/merchant/search \
 -H 'Content-Type: application/json' \
 -d '{
"page":1,
"lastval":1,
"keyword":"test",
"pagesize":1,
"orderby":"id",
"status":1
}'
$ curl -i -X POST \
 http://127.0.0.1:6001/v1/lib/merchant/user/get \
 -H 'Content-Type: application/json' \
 -d '{
"id":1
}'
