#!/usr/bin/env bash

kubectl create -f db-service.yaml,db-deployment.yaml,scalableapi-service.yaml,scalableapi-claim0-persistentvolumeclaim.yaml,scalableapi-deployment.yaml
