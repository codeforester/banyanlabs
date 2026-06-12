# URL Shortener cookie authentication smoke tests.

POST {{base_url}}/auth/signup
Content-Type: application/json
{
  "username": "alice",
  "email": "alice@example.com",
  "password": "correct horse battery staple"
}
HTTP 201
[Captures]
signup_session: cookie "banyan_url_shortener_session"
[Asserts]
header "Content-Type" contains "application/json"
header "Set-Cookie" contains "banyan_url_shortener_session="
header "Set-Cookie" contains "HttpOnly"
header "Set-Cookie" contains "SameSite=Lax"
jsonpath "$.user_id" isInteger
jsonpath "$.username" == "alice"

POST {{base_url}}/auth/login
Content-Type: application/json
{
  "username": "alice",
  "password": "wrong password"
}
HTTP 401
[Asserts]
jsonpath "$.error" == "invalid credentials"

POST {{base_url}}/auth/login
Content-Type: application/json
{
  "username": "alice",
  "password": "correct horse battery staple"
}
HTTP 200
[Captures]
login_session: cookie "banyan_url_shortener_session"
[Asserts]
header "Content-Type" contains "application/json"
header "Set-Cookie" contains "banyan_url_shortener_session="
header "Set-Cookie" contains "HttpOnly"
jsonpath "$.user_id" isInteger
jsonpath "$.username" == "alice"

POST {{base_url}}/auth/logout
Cookie: banyan_url_shortener_session={{login_session}}
Content-Type: application/json
{}
HTTP 204
[Asserts]
header "Set-Cookie" contains "banyan_url_shortener_session="
header "Set-Cookie" contains "Max-Age=0"
