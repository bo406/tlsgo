## A simple https web server for go

### Ca:
$openssl genrsa -out ca.key 2048   
$openssl req -x509 -new -nodes -key ca.key -subj "/CN=botls.com" -days 5000 -out ca.crt

**私钥文件 ca.key**  
**数字证书 ca.crt**

### Server:
$openssl genrsa -out server.key 2048   
$openssl req -new -key server.key -subj "/CN=localhost" -out server.csr  
$openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000

**私钥文件 server.key**  
**数字证书 server.crt**

### Client:
openssl genrsa -out client.key 2048   
openssl req -new -key client.key -subj "/CN=botls_cn" -out client.csr  
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 5000

**私钥文件 client.key**  
**数字证书 client.crt**

### Client:
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -extfile client.ext -out client.crt -days 5000

### Run

$go run server.go

$go run client.go
