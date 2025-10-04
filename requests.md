curl.exe -i http://localhost:8080/health
curl.exe -i http://localhost:8080/tasks
curl.exe -i -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d '{"title":"Купить молоко"}'
curl.exe -i http://localhost:8080/tasks/1
curl.exe -i -X PATCH http://localhost:8080/tasks/1 -H "Content-Type: application/json" -d '{"done":true}'
curl.exe -i -X DELETE http://localhost:8080/tasks/1
curl.exe -i -X OPTIONS http://localhost:8080/tasks
