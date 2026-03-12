# Join control plan
create join command 
```
sudo kubeadm init phase upload-certs --upload-certs

sudo kubeadm token create --print-join-command
```

join as control-plan
```
sudo kubeadm join <API_SERVER_IP>:<PORT> --token <TOKEN> \
  --discovery-token-ca-cert-hash sha256:<CERT_HASH> \
  --control-plane --certificate-key <CERTIFICATE_KEY>
```

join as worker
```
sudo kubeadm join <API_SERVER_IP>:<PORT> --token <TOKEN> \
  --discovery-token-ca-cert-hash sha256:<CERT_HASH> \
```
