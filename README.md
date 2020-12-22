# kcap
[![build](https://github.com/zodiac1214/kcap/workflows/Go/badge.svg)](https://github.com/zodiac1214/kcap/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/zodiac1214/kcap)](https://goreportcard.com/report/github.com/zodiac1214/kcap)

Kcap (mirror of pack) is a tool for packing kubernetes application for air-gap installation. Kcap produces different types of VM images for air-gap/zero internet installation

Use Kcap to:
* Pack/Install k8s ([k3os](https://github.com/rancher/k3os)) into air-gapped environment (zero internet access)
* Pack/Install your k8s applications into air-gap k8s cluster

## Helm chart vs Kubernetes yaml?
You are covered, run ``builder gen`` and drop your existing helm charts or native k8s yaml files in. You are good to go

# * ----- Under construction (below) -----*

## Prerequisites
* [packer](https://www.packer.io/)
packer by HashiCorp is used to pack machine images.
* [vagrant*](https://www.vagrantup.com/)
* [virtual box*](https://www.virtualbox.org/)

\* only required if you want to use vagrant to run [example](example) or test packed images locally.
## Quick Start (example - simple)
The example project demonstrate how to pack [hello-app](https://github.com/GoogleCloudPlatform/kubernetes-engine-samples/tree/master/hello-app) into air gap machine image.
Build hello-world example:
```
./pack.sh --extraImagesList example/simple/images.list --builders vagrant
```

Run simple example
```
cd example/simple
./demo.sh
```

## Generate a new project
```bash
builder gen --force --name=my-new-project
```
## Build and pack 
```bash
builder build --force --path=./my-new-project
```

## Install (vagrant only)
```bash
installer --path=./dist/my-new-project 

```
