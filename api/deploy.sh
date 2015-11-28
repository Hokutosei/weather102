#!/bin/bash

set -e

print () {
  echo "-->> $1"
}

print "building.."
make

print "stopping weather102"
kubectl stop rc weather102

print "deploying..."
kubectl create -f kubernetes.yaml
