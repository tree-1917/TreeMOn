# ---------------------------------------
# Test vote with valid credentials
# ---------------------------------------
POST http://localhost:8080/main_vote
Content-Type: application/json

{
  "username": "admin",
  "password": "1234",
  "animal": "cat"
}

HTTP 200

[Asserts]
header "Content-Type" contains "application/json"
body contains "cat"


