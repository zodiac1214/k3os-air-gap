#!/bin/sh
read -p "turn off wifi!!!! (press any key to continue ...)"
vagrant destroy -f
vagrant box remove k3os --force
vagrant box add ../k3os_virtualbox.box --name k3os --force

vagrant up

vagrant ssh k3os-server -c 'sudo /home/rancher/scripts/configure_k3s_server.sh not4you2see! 192.168.33.10'
vagrant ssh k3os-1 -c 'sudo /home/rancher/scripts/configure_k3s_node.sh not4you2see! 192.168.33.10 192.168.33.11'

echo "==== Get Kube Config ===="
./update-kube-config.sh

echo "==== Show k8s nodes ===="
KUBECONFIG=kube.config kubectl get nodes

echo "==== Deploy Hello-world ===="
KUBECONFIG=kube.config kubectl apply -f hello-world.yaml
KUBECONFIG=kube.config kubectl expose deployment hello-world --type=NodePort --name=example-service

port=$(KUBECONFIG=kube.config kubectl describe services example-service | grep "NodePort")

read -p "go to http://192.168.33.11:${port}"
cd $pwd