#!/bin/sh
#TODO: remove vagrant dep
read -p "turn off wifi!!!! (press any key to continue ...)"
vagrant destroy -f
vagrant box remove k3os --force
vagrant box add ../../k3os_virtualbox.box --name k3os --force

vagrant plugin install vagrant-disksize
vagrant up

vagrant ssh k3os-server -c 'sudo bash /home/rancher/scripts/configure_k3s_server.sh not4you2see! 192.168.33.10'
vagrant ssh k3os-1 -c 'sudo bash /home/rancher/scripts/configure_k3s_node.sh not4you2see! 192.168.33.10 192.168.33.11'

sleep 60

echo "==== Get Kube Config ===="
bash ./update-kube-config.sh

echo "==== Show k8s nodes ===="
KUBECONFIG=kube.config kubectl get nodes

cd $pwd