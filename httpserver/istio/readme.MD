### deploy httpserver
```
kubectl create ns securesvc
kubectl label ns securesvc istio-injection=enabled

```
```
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=cyuliu Inc./CN=*.cyuliu.io' -keyout cyuliu.io.key -out cyuliu.io.crt
kubectl create -n istio-system secret tls cyuliu-credential --key=cyuliu.io.key --cert=cyuliu.io.crt