## Overview

### Types of namespaces
- kube-publick: kube-public's existence is to provide space where we can create objects that should be visible throughout the whole cluster.
- kube-system: 

### Practice
- list existing namespaces `kubectl get ns`
- create a new namespace: `kubectl create ns testing`
- set context for ctl: `kubectl config set-context testing \
    --namespace testing \
    --cluster k3d-mycluster \
    --user admin@k3d-mycluster`
- view the context: `kubectl config view`
- switch context: `kubectl config use-context testing`