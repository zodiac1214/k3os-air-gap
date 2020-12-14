#!/bin/sh

rm -rf kube.config
SERVER_IP=$(cat config.yaml | grep " server" | cut -d " " -f 4)

bash -c 'while [[ "$(curl -k -s -o /dev/null -w ''%{http_code}'' https://${SERVER_IP}:6443)" != "401" ]]; do sleep 5; done'

vagrant ssh k3os-server -c "cat /etc/rancher/k3s/k3s.yaml | sed -e 's/127.0.0.1/${SERVER_IP}/g'" | tail -n 19 > kube.config
cat kube.config