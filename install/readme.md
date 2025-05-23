# source <https://www.itsgeekhead.com/tuts/kubernetes-129-ubuntu-22-04-3/> (copy to my git to prevent website is close)

### Start

```sudo su``` <br>

```printf "\n192.168.15.93 k8s-control\n192.168.15.94 k8s-1\n192.168.15.95 k8s-1\n\n" >> /etc/hosts``` <br>

```printf "overlay\nbr_netfilter\n" >> /etc/modules-load.d/containerd.conf``` <br>

```modprobe overlay``` <br>

```modprobe br_netfilter``` <br>

```printf "net.bridge.bridge-nf-call-iptables = 1\nnet.ipv4.ip_forward = 1\nnet.bridge.bridge-nf-call-ip6tables = 1\n" >> /etc/sysctl.d/99-kubernetes-cri.conf``` <br>

```sysctl --system``` <br>

```wget https://github.com/containerd/containerd/releases/download/v1.7.13/containerd-1.7.13-linux-amd64.tar.gz -P /tmp/``` <br>

```tar Cxzvf /usr/local /tmp/containerd-1.7.13-linux-amd64.tar.gz``` <br>

```wget https://raw.githubusercontent.com/containerd/containerd/main/containerd.service -P /etc/systemd/system/``` <br>

```systemctl daemon-reload``` <br>

```systemctl enable --now containerd``` <br>

```wget https://github.com/opencontainers/runc/releases/download/v1.1.12/runc.amd64 -P /tmp/``` <br>

```install -m 755 /tmp/runc.amd64 /usr/local/sbin/runc``` <br>

```wget https://github.com/containernetworking/plugins/releases/download/v1.4.0/cni-plugins-linux-amd64-v1.4.0.tgz -P /tmp/``` <br>

```mkdir -p /opt/cni/bin``` <br>

```tar Cxzvf /opt/cni/bin /tmp/cni-plugins-linux-amd64-v1.4.0.tgz``` <br>

```mkdir -p /etc/containerd``` <br>

```containerd config default | tee /etc/containerd/config.toml```   <<<<<<<<<<< manually edit and change **SystemdCgroup** to **true** (not systemd_cgroup) <br>

```vi /etc/containerd/config.toml``` <br>

```systemctl restart containerd``` <br>

```swapoff -a```  <<<<<<<< just disable it in /etc/fstab instead <br>

```apt-get update``` <br>

```apt-get install -y apt-transport-https ca-certificates curl gpg``` <br>

```mkdir -p -m 755 /etc/apt/keyrings``` <br>

```curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.29/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg``` <br>

```echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.29/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list``` <br>

```apt-get update``` <br>

```reboot``` <br>

```sudo su``` <br>

```apt-get install -y kubelet=1.29.1-1.1 kubeadm=1.29.1-1.1 kubectl=1.29.1-1.1``` <br>

```apt-mark hold kubelet kubeadm kubectl``` <br>

**check swap config, ensure swap is 0** <br>
```free -m``` <br>


### ONLY ON CONTROL NODE .. control plane install: <br>
```kubeadm init --control-plane-endpoint "loadbalancer:6443" --upload-certs --pod-network-cidr 10.10.0.0/16``` <br>

```export KUBECONFIG=/etc/kubernetes/admin.conf``` <br>

**add Calico 3.27.2 CNI** <br>
```kubectl create -f https://raw.githubusercontent.com/projectcalico/calico/v3.27.2/manifests/tigera-operator.yaml``` <br>
```wget https://raw.githubusercontent.com/projectcalico/calico/v3.27.2/manifests/custom-resources.yaml``` <br>
```vi custom-resources.yaml``` <<<<<< edit the CIDR for pods change ip to 10.10.0.0/16 <br>
```kubectl apply -f custom-resources.yaml``` <br>

get worker node commands to run to join additional nodes into cluster <br>
```kubeadm token create --print-join-command``` <br>


### ONLY ON WORKER nodes
Run the command from the token create output above
