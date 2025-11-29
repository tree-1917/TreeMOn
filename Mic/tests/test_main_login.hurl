# ---------------------
# Test Vaild Login 
# ---------------------
POST http://localhost:8080/main_login
Content-Type: application/json

{
    "username": "admin",
    "password": "1234"
}

HTTP 200
[Asserts]
header "Content-type" contains "application/json"
body contains "ok"


