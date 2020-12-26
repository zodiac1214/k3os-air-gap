#!/bin/sh
cwd=$(pwd)
SERVER_IP=$(cat vm.config.yaml | grep " server" | cut -d " " -f 4)
rm -rf kube.config
vagrant ssh k3os-server -c "cat /etc/rancher/k3s/k3s.yaml | sed -e 's/127.0.0.1/${SERVER_IP}/g'" | tail -n 19 > kube.config
cat kube.config
cd $cwd