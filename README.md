# k-console
Kafka manage console

## QuickStart
1. Start Kafka
```
cd k-console/ && docker-compose up -d
```
2. Run
```
cd k-console/ && KAFKA_STRING=localhost:9094 go run .
```
3. Test
```
curl 127.0.0.1:7777:/ping
# output: {"message":"pong"}
curl 127.0.0.1:7777/api/v1/topics
# output: {}
curl 127.0.0.1:7777/api/v1/groups
# output: {}
```

4. Swagger

Open `http://127.0.0.1:7777/swagger/index.html` in browser

