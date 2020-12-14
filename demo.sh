#!/bin/sh
read -p "turn off wifi!!!!"
./cleanRestartVagrant.sh
vagrant ssh k3os-server -c 'sudo /home/rancher/scripts/configure_k3s_server.sh not4you2see! 192.168.33.10'
vagrant ssh k3os-1 -c 'sudo /home/rancher/scripts/configure_k3s_node.sh not4you2see! 192.168.33.10 192.168.33.11'

echo "get kube config, deploy hello world"
./update-kube-config.sh