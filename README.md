# k3os Machine Image for air-gap system

Pack k8s applications with k3os on air-gap system. All you need is to provide a list of images your application needs. Please note that we only produce vagrant box image for now.

## Prerequisites
* [packer](https://www.packer.io/)
packer by HashiCorp is used to pack machine images.
* [vagrant*](https://www.vagrantup.com/)
* [virtual box*](https://www.virtualbox.org/)

\* only required if you want to use vagrant to run [example](example) or test packed images locally.
## Quick Start (example)
The example project demonstrate how to pack [hello-app](https://github.com/GoogleCloudPlatform/kubernetes-engine-samples/tree/master/hello-app) into air gap machine image.
Build hello-world example:
```
./pack.sh --extraImagesList example/images.list --builders=vagrant
```

Run example
```
cd example
./demo.sh
```
