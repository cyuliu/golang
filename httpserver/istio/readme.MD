### deploy httpserver
```
kubectl create ns securesvc
kubectl label ns securesvc istio-injection=enabled
kubectl create -f httpserver.yaml -n securesvc
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
curl --resolve httpsserver.cncamp.io:443:$INGRESS_IP https://httpsserver.cncamp.io/healthz -v -k
curl --resolve httpsserver.cyuliu.io:443:10.108.251.207 https://httpsserver.cyuliu.io/healthz -v -k
```
![image](https://user-images.githubusercontent.com/8605741/147476818-56433be3-88e9-4111-b98c-a78d026c8fa4.png)

