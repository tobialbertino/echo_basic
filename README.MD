## Architecture

per-modules(src):  
Delivery -> Service -> Repository(if Exist)

## Postman Collection
Import postman collection on folder ```./script/Simple.postman_collection.json```

docker build -t go-be -f Dockerfile .
docker run -p 8000:8000 --name be go-be