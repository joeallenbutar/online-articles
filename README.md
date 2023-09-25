# online-articles

## How To Run

1. Run docker compose with command `docker compose up`
2. Run application with command `go run main.go`

## Documentation API
- Postman : import "Online Articles.postman_collection.json" to Postman app.
- Swagger : http://localhost:8000/online-articles/index.html

## Bucket MinIO
- http://localhost:9001/login
    username: admin
    password: Strong#Pass#23

Optinal for set bucket to public:
- sh-4.4# mc alias set myminio http://localhost:9000 admin Strong#Pass#23
- sh-4.4# mc anonymous set download myminio/articles
![Alt text](/img/1.png)