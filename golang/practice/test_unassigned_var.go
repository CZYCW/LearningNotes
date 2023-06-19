package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	ProjectId, TemplateId, DatasetId, err := test()
	fmt.Println(ProjectId, TemplateId, DatasetId, err)
	fmt.Printf("%s:template/%v", "S3BUCKET", TemplateId)
}

func test() (int64, string, int64, error) {
	var ProjectId, TemplateId, DatasetId string = "", "123", "4"
	if ProjectId == "" && TemplateId == "" {
		return -1, "", -1, errors.New("text string")
	}
	var projectId, datasetId int64
	var err error
	if ProjectId != "" {
		projectId, err = ParseInt64FromString(ProjectId)
		if err != nil {
			return -1, "", -1, err
		}
	}
	datasetId, err = ParseInt64FromString(DatasetId)
	if err != nil {
		return -1, "", -1, err
	}
	return projectId, TemplateId, datasetId, nil
}

func ParseInt64FromString(s string) (i int64, err error) {
	convertedResult, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return convertedResult, err
	}
	return convertedResult, nil
}
