# k3os Machine Image for air-gap system

Pack k8s applications with k3os on air-gap system. Please note that we only produce vagrant box image for now.

## Prerequisites
* [packer](https://www.packer.io/)
packer by HashiCorp is used to pack machine images.
* [vagrant*](https://www.vagrantup.com/)
* [virtual box*](https://www.virtualbox.org/)

\* only required if you want to use vagrant to run [example](example) or test packed images locally.
## Quick Start (example)

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
