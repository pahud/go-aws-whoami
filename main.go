package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func whoami(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, stsGetCallerIdentity()+"\n")
}

func stsGetCallerIdentity() string {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := sts.New(cfg)
	input := &sts.GetCallerIdentityInput{}

	req := svc.GetCallerIdentityRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return "error"
	}
	jsondata, _ := json.Marshal(result)
	return string(jsondata)
}

func main() {
	http.HandleFunc("/", whoami)
	http.HandleFunc("/whoami", whoami)
	http.ListenAndServe(":8080", nil)
}
