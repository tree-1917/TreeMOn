# ---------------------------------------
# Test Invalid Login (wrong password)
# ---------------------------------------
POST http://localhost:8080/vote
Content-Type: application/json

{
  "username": "admin",
  "password": "1234",
  "animal": "cat"
}

HTTP 200
[Asserts]
header "Content-Type" contains "text/plain; charset=utf-8"
jsonpath "$.cat" exists
jsonpath "$.dog" exists


