# Install guide

kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.14.5/config/manifests/metallb-native.yaml

l2 config
```
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ConfigMap
metadata:
 namespace: metallb-system
 name: config
data:
 config: |
  address-pools:
  - name: default
    protocol: layer2
    addresses:
    - 192.168.1.50-192.168.1.100
EOF
```
bgp config 
```
```
