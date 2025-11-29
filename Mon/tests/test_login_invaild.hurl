# ---------------------------------------
# 1. Invalid login
# ---------------------------------------
POST http://localhost:8080/vote
Content-Type: application/json

{
  "username": "admin",
  "password": "wrong",
  "animal": "cat"
}

HTTP 401
[Asserts]
header "Content-Type" contains "text/plain"
body contains "Unauthorized"

