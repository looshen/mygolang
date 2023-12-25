test:
1、local test
go build -o httpserver *.go

docker build -t myhttpsvc .

docker run -itd --name myhttpsvc2 -p 8080:8080 myhttpsvc:latest


http://localhost:8080/healthz

http://localhost:8080/metrics


2、k8s apply yaml test registry.cn-hangzhou.aliyuncs.com/rosentest01/docker.mylearn:v1.0.8:

kubectl apply -f deploy-probe.yaml

kubectl get svc # get nodeport ip

10.96.142.206/healthz
10.96.142.206/metrics
