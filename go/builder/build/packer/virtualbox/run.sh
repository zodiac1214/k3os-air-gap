#!/bin/sh
set -e
createVM () {
  VM_NAME="$1"
  VBoxManage controlvm $VM_NAME acpipowerbutton || true
  VBoxManage unregistervm --delete $VM_NAME || true
  VBoxManage import output-virtualbox/packer-virtualbox-*.ovf  \
    --vsys 0 \
    --ostype "Linux_64" \
    --vmname $VM_NAME \
    --group "/" \
    --cpus 2 \
    --memory 2048
  VBoxManage modifyvm $VM_NAME --nic1 hostonly --hostonlyadapter1 vboxnet0
  VBoxManage startvm $VM_NAME --type headless
}

chmod 400 ssh-default.pem

createVM "server"
read -p "server Ip: " ServerIP
ssh -o StrictHostKeyChecking=no -i ./ssh-default.pem rancher@$ServerIP "sudo bash scripts/configure_k3s_server.sh thisistoken $ServerIP"


createVM "node1"
read -p "server Ip: " NodeIP
ssh -o StrictHostKeyChecking=no -i ./ssh-default.pem rancher@$NodeIP "sudo bash scripts/configure_k3s_node.sh thisistoken $ServerIP $NodeIP 1"

createVM "node2"
read -p "server Ip: " NodeIP2
ssh -o StrictHostKeyChecking=no -i ./ssh-default.pem rancher@$NodeIP2 "sudo bash scripts/configure_k3s_node.sh thisistoken $ServerIP $NodeIP2 2"

ssh -o  StrictHostKeyChecking=no -i ./ssh-default.pem rancher@$ServerIP "cat /etc/rancher/k3s/k3s.yaml" > kube.config
KUBECONFIG=kube.config helm install rancher rancher-stable/rancher --version 2.5.3 --set tls=external