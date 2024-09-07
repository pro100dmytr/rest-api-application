# rest-api-practice

# user-server

# REST API

# endpoints
GET /users -- get list of users -- 200, 404, 500
GET /users/:id --get user by id -- 200, 404, 500
POST /users/:id -- create user -- 204, 4xx, Header Location: url
PUT /users/:id -- fully update user -- 204/200, 404, 400, 500
PATCH /users/:id -- partially update user -- 204/200, 404, 400, 500
DELETE /users/:id -- delete user by id -- 200, 404, 400

# unit test
go test ./internal/user -v
