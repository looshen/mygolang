go build -o httpserver *.go
docker build -t myhttpsvc .
docker run -itd --name myhttpsvc2 -p 8080:8080 myhttpsvc:latest

http://localhost:8080/healthz
http://localhost:8080/metrics