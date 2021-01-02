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

findIP() {
  VM_NAME="$1"
  macAddress=`vboxmanage showvminfo $VM_NAME | grep "NIC 1" | awk -F'MAC: ' '{print $2}' | awk -F',' '{print $1}'`
  sleep 60
  ip=`vboxmanage dhcpserver findlease --interface vboxnet0 --mac-address=$macAddress | grep "IP" | awk -F'Address:  ' '{print $2}'`
  echo $ip
}
chmod 400 ssh-default.pem

createVM "server"
findIP "server"
ssh -o StrictHostKeyChecking=no -i ./ssh-default.pem rancher@$ip "sudo bash scripts/configure_k3s_server.sh thisistoken $ip"
ServerIP=$ip

createVM "node"
findIP "node"
ssh -o StrictHostKeyChecking=no -i ./ssh-default.pem rancher@$ip "sudo bash scripts/configure_k3s_node.sh thisistoken $ServerIP $ip"

# TODO: install rancher
# helm install rancher rancher-stable/rancher --version 2.5.3 --set tls=external