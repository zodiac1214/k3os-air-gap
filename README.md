![build](https://github.com/zodiac1214/kcap/workflows/Go/badge.svg)
# k3os Machine Image for air-gap system

Pack k8s applications with k3os on air-gap system. All you need is to provide a list of images your application needs. Please note that we only produce vagrant box image for now.

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
