#!/bin/sh

packer build packer/vagrant.json 
vagrant box add k3os_virtualbox.box --name k3os --force