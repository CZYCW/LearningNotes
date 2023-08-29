package main

import (
	"fmt"
	"strings"
)

const PvcURIPrefix = "pvc://"

func parsePvcURI(srcURI string) (pvcName string, pvcPath string, err error) {
	parts := strings.Split(strings.TrimPrefix(srcURI, PvcURIPrefix), "/")
	if len(parts) > 1 {
		pvcName = parts[0]
		pvcPath = strings.Join(parts[1:], "/")
	} else if len(parts) == 1 {
		pvcName = parts[0]
		pvcPath = ""
	} else {
		return "", "", fmt.Errorf("Invalid URI must be pvc://<pvcname>/[path]: %s", srcURI)
	}

	return pvcName, pvcPath, nil
}

func main() {
	pvcName, pvcPath, err := parsePvcURI("pvc://pvc-model-64b7a0a2ad6077101ecb1ab2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pvcName)
	fmt.Println(pvcPath)
}