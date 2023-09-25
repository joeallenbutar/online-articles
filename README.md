# online-articles

## How To Run

1. Run docker compose with command `docker compose up`
2. Run application with command `go run main.go`

## Documentation API
- Postman : import "Online Articles.postman_collection.json" to Postman app.
- Swagger : http://localhost:8000/online-articles/index.html

## Bucket MinIO
- http://localhost:9001/login
    - username: admin
    - password: Strong#Pass#23

Optinal for set bucket to public:
- sh-4.4# mc alias set myminio http://localhost:9000 admin Strong#Pass#23
- sh-4.4# mc anonymous set download myminio/articles
![Alt text](/img/1.png)

## Deskripsi Aplikasi
# Aplikasi Artikel Online
Dalam aplikasi ini, terdapat 7 HTTP REST API Endpoint, yakni:
1. Registrasi User : [POST] http://localhost:8000/v1/api/register
   - Melakukan registrasi akun user.
2. Login User : [POST] http://localhost:8000/v1/api/login
   - Melakukan login dan melakukan generate JWT token yang akan digunakan pada module article untuk authorization.
3. Create Article : [POST] http://localhost:8000/v1/api/article
   - Melakukan proses create article. 
4. Update Article : [PUT] http://localhost:8000/v1/api/article/[article_id]
   - Melakukan proses update article. 
5. Delete Article : [DELETE] http://localhost:8000/v1/api/article/[article_id]
   - Melakukan proses delete article. 
6. Get Article By Id : [GET] http://localhost:8000/v1/api/article/[article_id]
   - Melakukan proses pengambilan data artikel berdasarkan id. 
7. Get All Article : [GET] http://localhost:8000/v1/api/article
   - Melakukan proses pengambilan semua data artikel. 