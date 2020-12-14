# -*- mode: ruby -*-
# vi: set ft=ruby :

require 'ipaddr'
require 'yaml'

x = YAML.load_file('config.yaml')

Vagrant.configure(2) do |config|
  config.vbguest.auto_update = false if Vagrant.has_plugin?("vagrant-vbguest") # disable conflicting plugin

  config.vm.define "k3os-server" do |server|
      c = x.fetch('server')
      server.vm.box= "k3os"
      server.vm.guest = :linux
      server.vm.provider "virtualbox" do |v|
        v.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
        v.customize ['modifyvm', :id, '--natnet1', '192.168.222.0/24']
        v.cpus = c.fetch('cpus')
        v.linked_clone = true if Gem::Version.new(Vagrant::VERSION) >= Gem::Version.new('1.8.0') and x.fetch('linked_clones')
        v.memory = c.fetch('memory')
      end
      config.vm.synced_folder '.', '/vagrant', disabled: true
      server_ip = x.fetch('ip').fetch('server')
      server.vm.network x.fetch('net').fetch('network_type'), ip: server_ip, auto_config: false
      #server.vm.provision "shell", path: "/home/rancher/scripts/configure_k3s_server.sh", args: [x.fetch('k3s_token'), server_ip]
  end

  node_ip_base = IPAddr.new(x.fetch('ip').fetch('node'))
  (1..x.fetch('node').fetch('count')).each do |i|
    c = x.fetch('node')
    hostname = "k3os-%d" % i
    config.vm.define hostname do |node|
      node.vm.box = "k3os"
      node.vm.guest = :linux
      node.vm.provider "virtualbox" do |v|
        v.customize ['modifyvm', :id, '--natnet1', '192.168.222.0/24']
        v.cpus = c.fetch('cpus')
        v.linked_clone = true if Gem::Version.new(Vagrant::VERSION) >= Gem::Version.new('1.8.0') and x.fetch('linked_clones')
        v.memory = c.fetch('memory')
      end
      config.vm.synced_folder '.', '/vagrant', disabled: true
      node_ip = IPAddr.new(node_ip_base.to_i + i - 1, Socket::AF_INET).to_s 
      node.vm.network x.fetch('net').fetch('network_type'), ip: node_ip, auto_config: false
      #node.vm.provision "shell", path: "/home/rancher/scripts/configure_k3s_node.sh", args: [x.fetch('k3s_token'), x.fetch('ip').fetch('server'), node_ip, i]
    end
  end
end
