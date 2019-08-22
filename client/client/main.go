package main

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func isAvailable(status int) float64 {
	if status < 200 || status >= 400 {
		return 0
	}
	return 1
}

func getURL() *url.URL {
	url, err := url.Parse(os.Getenv("URL"))
	if url.String() == "" {
		log.Fatal("Environment variable \"URL\" is empty")
	}
	if err != nil || url.String() == "" {
		log.Fatal(err)
	}

	return url
}

func checkHealth(url *url.URL, svc *cloudwatch.CloudWatch, ts *time.Time) {

	var status float64
	var storageResolution = int64(1)
	var client = &http.Client{
		Timeout: time.Second * 3,
	}

	resp, err := client.Get(url.String())
	if err != nil {
		log.Println(err)
		status = 0
	} else {
		status = isAvailable(resp.StatusCode)
	}

	_, err = svc.PutMetricData(&cloudwatch.PutMetricDataInput{
		Namespace: aws.String("Test/K8s/SLI"),
		MetricData: []*cloudwatch.MetricDatum{
			&cloudwatch.MetricDatum{
				MetricName:        aws.String("isIngressAvailable"),
				Unit:              aws.String("Count"),
				Value:             aws.Float64(status),
				Timestamp:         ts,
				StorageResolution: &storageResolution,
				Dimensions: []*cloudwatch.Dimension{
					&cloudwatch.Dimension{
						Name:  aws.String("ClusterName"),
						Value: aws.String(url.Host),
					},
				},
			},
		},
	})
	if err != nil {
		log.Println("Error putting metrics:", err.Error())
		return
	}
	log.Printf("Successfully created CloudWatch data point with status: %v", status)
}

// HandleRequest function
func HandleRequest(ctx context.Context) {
	// func HandleRequest() {

	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := cloudwatch.New(session)
	url := getURL()

	ticker := time.NewTicker(2 * time.Second)
	for elem := range ticker.C {
		checkHealth(url, svc, &elem)
		// fmt.Println(elem)
	}

}

func main() {
	lambda.Start(HandleRequest)
	// HandleRequest()
}
