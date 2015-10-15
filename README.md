### HTTPS IN GO

$openssl genrsa -out server.key 2048

$openssl req -new -x509 -key server.key -out server.crt -days 365


Country Name (2 letter code) [AU]:<br/>
State or Province Name (full name) [Some-State]:<br/>
Locality Name (eg, city) []:<br/>
Organization Name (eg, company) [Internet Widgits Pty Ltd]:<br/>
Organizational Unit Name (eg, section) []:<br/>
Common Name (e.g. server FQDN or YOUR name) []:localhost<br/>
Email Address []:<br/>
