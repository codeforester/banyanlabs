# URL Shortener health endpoint smoke tests.

GET {{base_url}}/healthz
HTTP 200
[Asserts]
header "Content-Type" contains "application/json"
header "X-Request-ID" exists
jsonpath "$.service" == "url-shortener"
jsonpath "$.status" == "ok"
jsonpath "$.uptime" isString

HEAD {{base_url}}/healthz
HTTP 200
[Asserts]
header "X-Request-ID" exists

POST {{base_url}}/healthz
HTTP 405
[Asserts]
header "Allow" contains "GET"
header "Allow" contains "HEAD"
body contains "method not allowed"
