kubectl delete all --all

minikube image build -t chaincue-real-estate-go .
minikube image ls
minikube addons enable ingress
minikube ip -> sudo nano /etc/hosts -> "minikube_ip" local.chaincue.com

```
192.168.00.0 local.chaincue.com
```

kubectl -f postgres.yml apply
kubectl exec -it postgres-backend-0 -- psql -U admin -d postgres -c "CREATE DATABASE \"chaincue-real-estate-postgres\";"

kubectl -f redis.yml apply

kubectl -f backend.yml apply
