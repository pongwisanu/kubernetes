# For basic testing

kubectl create deployment nginx --image=nginx

kubectl expose deployment nginx --port=80 --type=LoadBalancer
