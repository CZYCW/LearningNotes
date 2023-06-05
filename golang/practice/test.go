package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var s string = "{\"pretrain\":\"bigscience/bloom-560m\",\"model\":\"bloom\",\"strategy\":\"colossalai_zero2_cpu\",\"log_interval\":10,\"dataset\":\"/mnt/dataset/data.json\",\"batch_size\":4,\"accumulation_steps\":8,\"lr\":0.00002,\"max_datasets_size\":512,\"max_epochs\":1}"
	var hyperParameters map[string]interface{}
	err := json.Unmarshal([]byte(s), &hyperParameters)

	if err != nil {
		fmt.Print(err)
	}
	var cmd string
	if hyperParameters == nil {
		cmd = ""
	} else {
		cmd = "export "
		for key, value := range hyperParameters {
			cmd += fmt.Sprintf("%s=%v ", key, value)
		}
		cmd += ";"
	}
	fmt.Printf("final command :%s", cmd)

}
