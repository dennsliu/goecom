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
"password":"123456",
"type":0
}'

$ curl -i -X POST \
 http://127.0.0.1:6001/v1/lib/merchant/search \
 -H 'Content-Type: application/json' \
 -d '{
"page":1,
"lastid":0,
"keyword":"te",
"pagesize":3,
"status":1,
"ordertype":"desc"
}'
$ curl -i -X POST \
 http://127.0.0.1:6001/v1/lib/merchant/add \
 -H 'Content-Type: application/json' \
 -d '{
"name": "testing1",
"status":1
}'
$ curl -i -X POST \
 http://127.0.0.1:6001/v1/lib/merchant/user/add \
 -H 'Content-Type: application/json' \
 -d '{
"username":"testing1",
"nickname":"testing1",
"email":"testing1@testing1.com",
"password":"123456",
"telephone":"13511111111",
"mobliephone":"13511111111",
"merchantid":1,
"status":1
}'

$ curl -i -X POST \
 http://127.0.0.1:6001/v1/lib/merchant/user/get \
 -H 'Content-Type: application/json' \
 -d '{
"id":1
}'
$ curl -i -X POST \
 http://127.0.0.1:6001/v1/lib/merchant/user/search \
 -H 'Content-Type: application/json' \
 -d '{
"page":1,
"lastid":0,
"keyword":"",
"pagesize":3,
"status":1,
"ordertype":"desc"
}'
