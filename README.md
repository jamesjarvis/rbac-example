# rbac-example

Example Golang application with Role Based Access Control.

It is a HTTP Server, mimicking a string:string dictionary, with the following endpoints:

- `GET /v1/map/{key}`
  - Returns a string `value` with status 200 if found
  - Returns `NOT_FOUND` with status 404 if not found
- `POST /v1/map/{key}`
  - The POST body is the `value`
  - Returns the same string `value` with status 200 if set
  - Returns `UNAUTHORISED` with status 403 if not allowed
