# k8s

kubectl port-forward service/goserver-service 80:80
 
kubectl proxy --port=80# k8s

kubectl run -it fortio --rm --image=fortio/fortio -- load -qps 800 -t 120s -c 70 "http://goserver-service/healthz"