package main

import ctrl "sigs.k8s.io/controller-runtime"

var kubeconfig = ctrl.GetConfigOrDie()
