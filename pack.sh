#!/bin/sh
set -e
usage() {
  echo "$0: Build your vagrant vbox image for air gapped installation"
  echo
  echo "Usage:"
  echo "$0 --extraImagesList path"
  echo "$0 --builders name1,name2"
  echo "Available options: "
  echo -e "\t      --extraImagesList      - path to a file has list of extra images"
  echo -e "\t      --builders             - comma separated names of packer builder [vagrant/vsphere]"
  exit 1
}

# get command options
while [[ $# > 0 ]]; do
  case "${1}" in
    --extraImagesList)
      extraImagesList=${2}
      shift
      ;;
    --builders)
      builders=${2}
      shift
      ;;
    -h|--help)
      usage
      ;;
    *)
      echo "Unknown option ${1}"
      usage
  esac
  shift
done

if [[ ! -f "k3s-airgap-images-amd64.tar" ]]; then
    echo "download air gap system images"
    curl -L https://github.com/k3s-io/k3s/releases/download/v1.18.9%2Bk3s1/k3s-airgap-images-amd64.tar > k3s-airgap-images-amd64.tar
else 
    echo "air gap system images exist"
fi

rm -rf images
mkdir -p images
while IFS= read -r line
do
  echo "  pack system image: $line"
  imageSha256=`docker pull $line | grep Digest | awk -F':' '{print $3}'`
  docker save $line > images/$imageSha256.tar
done < "system-images.list"
echo

if [[ ! -z $extraImagesList ]]; then
    echo "Pack extra images ..."
    while IFS='' read -r line2
    do
    echo "  pack extra image: $line2"
    imageSha256=`docker pull $line2 | grep Digest | awk -F':' '{print $3}'`
    docker save $line2 > images/$imageSha256.tar
    done < $extraImagesList
fi

if [[ -z $builders ]]; then
    echo "Error: builders is not spcified.(default: vagrant)"
    builders="vagrant"
fi

packer build --force --only=$builders packer/vagrant.json 