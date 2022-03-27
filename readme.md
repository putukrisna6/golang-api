# Simple REST API Go App

## Features:
- Gin for the web framework
- Gorm for ORM
- jwt-go for authentication
- Redis for caching
- MariaDB for main db
- phpMyAdmin for db management
- Portainer for container management

## How to use:
- Clone the repo
- Setup your .env file
- Run `docker-compose up -d`

## Misc info:
### phpMyAdmin
- access via `http://localhost:8085/` or through the "Open in Browser" in Docker Desktop
- by default: server is `db`, username is `root`, and password is `test`. configurable within `docker-compose.yml`
### Portainer
- access via `http://localhost:9000/`
- feel free to set the password and username when running it for the first time
### Testing
- open your browser and go to `http://localhost:8080/`, if all is fine you would be greeted by `Hello, World`
- use Postman or similar tools and set the url to `http://localhost:8080/`
- do not forget to set the `Authorization` key within `Headers` on your requests on some routes protected by the jwt middleware
- available routes: 
  - POST   /api/auth/login
  - POST   /api/auth/register
  - GET    /api/users/profile
  - PUT    /api/users
  - GET    /api/books
  - POST   /api/books
  - GET    /api/books/:id
  - PUT    /api/books
  - DELETE /api/books/:id
  - GET    /api/receipts/all
  - POST   /api/receipts
  - GET    /api/receipts/:id
  - PUT    /api/receipts
  - DELETE /api/receipts/:id
  - GET    /       

- some jquery request I used to test the APIs from another domain
  ```js
  $.ajax({
        url: "http://localhost:8080/api/receipts/all"
  })
  ```

  ```js
  $.ajax({
        url: "http://localhost:8080/api/receipts/1"
  })
  ```
  ```js
  $.ajax({
        type: 'POST',
        url: "http://localhost:8080/api/receipts",
        data: {"amount":3,"total":67},
        dataType: "json",
        success: function(resultData) { alert("Save Complete") }
  });
  ```
  ```js
  $.ajax({
        type: 'PUT',
        url: "http://localhost:8080/api/receipts",
        data: {"id":4,"amount":323,"total":6723.2},
        dataType: "json",
        success: function(resultData) { alert("Update Complete") }
  });
  ```
  ```js
  $.ajax({
      type: 'DELETE',
      url: "http://localhost:8080/api/receipts/4",
      dataType: "json",
      success: function(resultData) { alert("Delete Complete") }
  });
  ```
