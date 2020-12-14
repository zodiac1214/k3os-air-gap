#!/bin/sh

usage() {
  echo "$0: Build your vagrant vbox image for air gapped installation"
  echo
  echo "Usage:"
  echo "$0 --extraImagesList path"
  echo "Available options: "
  echo -e "\t      --extraImagesList      - path to a file has list of extra images"
  exit 1
}

# get command options
while [[ $# > 0 ]]; do
  case "${1}" in
    --extraImagesList)
      extraImagesList=${2}
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

systemImagePath=`dirname "$0"`/system-images.list
echo "Pack system images ..."
rm -rf images
mkdir -p images
while IFS= read -r line
do
  echo "  pack system image: $line"
  imageSha256=`docker pull $line | grep Digest | awk -F':' '{print $3}'`
  docker save $line > images/$imageSha256.tar
done < "$systemImagePath"
echo

echo "Pack extra images ..."
while IFS='' read -r line2
do
  echo "  pack extra image: $line2"
  imageSha256=`docker pull $line2 | grep Digest | awk -F':' '{print $3}'`
  docker save $line2 > images/$imageSha256.tar
done < $extraImagesList

packer build packer/vagrant.json 
vagrant box add k3os_virtualbox.box --name k3os --force