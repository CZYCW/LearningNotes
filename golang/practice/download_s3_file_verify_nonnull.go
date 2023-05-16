package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	modelUrl := "https://luchen-storage.s3.amazonaws.com/models/model.zip?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAZUOMWBQ5GSFPF2WF%2F20230515%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20230515T085924Z&X-Amz-Expires=604800&X-Amz-SignedHeaders=host&X-Amz-Signature=566252011115cdd4e8410df05b2e18699b7c8d0b751b8598e34f12ace24944c8"
	modelFile, err := http.Get(modelUrl)
	if err != nil {
		fmt.Println("Failed to download file:", err)
		return
	}
	// Create the local file to save the downloaded model.zip
	file, err := os.Create("model.zip")
	if err != nil {
		fmt.Println("Failed to create local file:", err)
		return
	}
	defer file.Close()
	defer modelFile.Body.Close()
	_, err = io.Copy(file, modelFile.Body)
	if err != nil {
		fmt.Println("Failed to save file:", err)
		return
	}

	// Check if the downloaded file is not null
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Failed to get file information:", err)
		return
	}

	if fileInfo.Size() == 0 {
		fmt.Println("Downloaded file is empty")
		return
	}

	fmt.Println("File downloaded successfully!")
	os.Remove("model.zip")
}
