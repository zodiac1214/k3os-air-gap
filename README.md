<p align="center">
  <img width="230" height="150" src="https://github.com/zodiac1214/kcap/blob/master/logo.png?raw=true">

  <span>Kcap (mirror of pack) is a tool for packing kubernetes application for air-gap installation. Kcap produces different types of VM images for air-gap/zero internet installation</span>
</p>

[![build](https://github.com/zodiac1214/kcap/workflows/Go/badge.svg)](https://github.com/zodiac1214/kcap/actions?query=workflow%3AGo+branch%3Amaster)
[![Go Report Card](https://goreportcard.com/badge/github.com/zodiac1214/kcap)](https://goreportcard.com/report/github.com/zodiac1214/kcap)
[![Coverage Status](https://coveralls.io/repos/github/zodiac1214/kcap/badge.svg?branch=master)](https://coveralls.io/github/zodiac1214/kcap?branch=master)

Use **Kcap** to:
* Pack/Install k8s ([k3os](https://github.com/rancher/k3os)) into air-gapped environment (zero internet access)
* Pack/Install your k8s applications into air-gap k8s cluster

## Installation
#### Prerequisites
* [packer](https://www.packer.io/) by HashiCorp is used to pack machine images.
* [Docker](https://docker.io) oh well, you know ...
* [vagrant*](https://www.vagrantup.com/) by HashiCorp to build/test VM images locally
* [virtual box*](https://www.virtualbox.org/) VM solution that works with vagrant on local machine

\* only required if you want to use vagrant on local.
#### Install
* Download binary for [release](https://github.com/zodiac1214/kcap/releases)
* rename the binart to "kcap"
* make it executable: ``chmod +x kcap``
* move *kcap* to: `/usr/bin/kcap`

## How it works
### Packing
Kubernetes is packed using [k3os](https://github.com/rancher/k3os), all required docker images are exported and saved to VM image. Istio ,rancher dashboard and promethues are also pre-packed into the VM image. For your application, we read your helm charts or vanilla k8s yaml files to extract all docker images that are required. When the VM image is produced, it will have all binaries and images required to run kubernetes as well as your application. The idea is similar to what github enterprise is providing: virtual appliance.    
### Install
[k3os](https://github.com/rancher/k3os) is in charge of build kubernetes cluster. We choose k3s mainly for 3 reasons:
* It is production ready
* k3s removed only cloud integration in the original kubernetes which is a perfect fit for air gap system. where we won't use any cloud features like AWS EBS
* Rancher lab has provided a lot of nice integration with their k3s. You can provide your customer with an out-of-box cluster management/monitoring tool on top of your application

Two scripts are pre-packed into the VM image. You can run one of it to configure a VM as master or worker. When we configure worker nodes, the script also import all docker images using `ctr image import`. That said there is no private registry gets involved. 

We will use helm to deploy applications chart provided.   

## In a nutshell ...
### *NOTE*: Anything marked with ``*`` below is not implemented yet 
* run ``kcap gen --name=cool-kids-project``
* put all your existing helm charts under ``cool-kids-project/charts``
* put all your existing kubernetes yaml files under ``cool-kids-project/kubernetes``\*
* run ``kcap build --path=./cool-kids-project --vm-type=vagrant``

  or ``kcap build --path=./cool-kids-project --vm-type=vsphere``\*
* Create VMs from the generated image on your favourite VMS
* on the master node, ``/home/rancher/scripts/configure_k3s_server.sh <TOKEN> <MASTER_NODE_IP>``
* on the worker node, ``/home/rancher/scripts/configure_k3s_server.sh <TOKEN> <MASTER_NODE_IP> <WORKER_NODE_IP>``

All above will get you:
* applications you defined in helm chart or native k8s yaml files
* [longhorn](https://github.com/longhorn/longhorn) distributed block storage\*
* [rancher](https://github.com/rancher/rancher) dashboard integrated with [istio](https://istio.io/) and monitoring\*

[comment]: <> (Under construction)

[comment]: <> (## Quick Start &#40;example - simple&#41;)

[comment]: <> (The example project demonstrate how to pack [hello-app]&#40;https://github.com/GoogleCloudPlatform/kubernetes-engine-samples/tree/master/hello-app&#41; into air gap machine image.)

[comment]: <> (Build hello-world example:)

[comment]: <> (```)

[comment]: <> (./pack.sh --extraImagesList example/simple/images.list --builders vagrant)

[comment]: <> (```)

[comment]: <> (Run simple example)

[comment]: <> (```)

[comment]: <> (cd example/simple)

[comment]: <> (./demo.sh)

[comment]: <> (```)

[comment]: <> (## Generate a new project)

[comment]: <> (```bash)

[comment]: <> (builder gen --force --name=my-new-project)

[comment]: <> (```)

[comment]: <> (## Build and pack )

[comment]: <> (```bash)

[comment]: <> (builder build --force --path=./my-new-project)

[comment]: <> (```)

[comment]: <> (## Install &#40;vagrant only&#41;)

[comment]: <> (```bash)

[comment]: <> (installer --path=./dist/my-new-project )

[comment]: <> (```)

[comment]: <> (vagrant ssh k3os-server -c 'sudo /home/rancher/scripts/configure_k3s_server.sh not4you2see! 192.168.33.10')

[comment]: <> (vagrant ssh k3os-1 -c 'sudo /home/rancher/scripts/configure_k3s_node.sh not4you2see! 192.168.33.10 192.168.33.11')
