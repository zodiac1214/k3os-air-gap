# k3os Machine Image for air-gap system

Pack k8s applications with k3os on air-gap system. Please note that we only produce vagrant box image for now.

## Quick Start

Build hello-world example:

```
./buildVagrantBox.sh --extraImagesList example/images.list
```
Run example* 

```
./example/demo.sh
```

\* in progress, you will have to manually update kube config for now (see issue [#1](/../../issues/1)). The script pause and you will need to update kube config manually before you proceed to example app deployment.

Get kube Config

```
ssh -o StrictHostKeyChecking=no  -i packer/vagrant-ssh-default.pem rancher@192.168.33.10 "cat /etc/rancher/k3s/k3s.yaml"
```

Continue in `example/demo.sh`
