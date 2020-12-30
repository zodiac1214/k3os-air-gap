#!/bin/sh

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
  macAddress=`vboxmanage showvminfo $VM_NAME | grep "NIC 1" | awk -F'MAC: ' '{print $2}' | awk -F',' '{print $1}'`
  sleep 60
  ip=`vboxmanage dhcpserver findlease --interface vboxnet0 --mac-address=$macAddress | grep "IP" | awk -F'Address:  ' '{print $2}'`
  chmod 400 ssh-default.pem
  ssh -o StrictHostKeyChecking=no -i ./ssh-default.pem rancher@$ServerIp "$2"
}
createVM "server" "ls"


#vagrant ssh k3os-server -c 'sudo bash /home/rancher/scripts/configure_k3s_server.sh not4you2see! 192.168.33.10'
#vagrant ssh k3os-1 -c 'sudo bash /home/rancher/scripts/configure_k3s_node.sh not4you2see! 192.168.33.10 192.168.33.11'
#
#sleep 60
#
#echo "==== Get Kube Config ===="
#bash ./update-kube-config.sh
#
#echo "==== Show k8s nodes ===="
#KUBECONFIG=kube.config kubectl get nodes
#
#cd $pwd