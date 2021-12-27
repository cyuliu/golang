### deploy httpserver
```
kubectl create ns tracing
kubectl label ns tracing istio-injection=enabled
kubectl -n tracing apply -f service0.yaml
kubectl -n tracing apply -f service1.yaml
kubectl -n tracing apply -f service2.yaml 
kubectl apply -f istio-specs.yaml -n tracing
```
```
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=cyuliu Inc./CN=*.cyuliu.io' -keyout cyuliu.io.key -out cyuliu.io.crt
kubectl create -n istio-system secret tls cyuliu-credential --key=cyuliu.io.key --cert=cyuliu.io.crt
kubectl apply -f istio-specs.yaml -n securesvc
```
### check ingress ip
```
k get svc -nistio-system
istio-ingressgateway   LoadBalancer   $INGRESS_IP
```
![image](https://user-images.githubusercontent.com/8605741/147476706-89b76136-888f-4eb2-9a6f-5d9a2a6acfed.png)

### access the httpserver via ingress
```
curl --resolve service.cyuliu.io:443:10.108.251.207 https://service.cyuliu.io/service0 -v -k
```
![image](https://user-images.githubusercontent.com/8605741/147497742-ad5831a0-3862-4462-bf3f-05a144f43472.png)
