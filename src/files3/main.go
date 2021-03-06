package main

import (
	"bytes"
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Event struct {
	Txt string
}

type Response struct {
	T string
	E string
}

func handler(c context.Context, e Event) (Response, error) {
	sess := session.Must(session.NewSession())
	s3c := s3.New(sess)
	_, err := s3c.PutObjectWithContext(c, &s3.PutObjectInput{
		Bucket: aws.String("files3-tests"),
		Key:    aws.String("file1.txt"),
		Body:   bytes.NewReader([]byte("hello to you " + e.Txt)),
	})

	res := Response{
		T: "saving " + e.Txt,
	}

	if err != nil {
		res.E = err.Error()
	}
	return res, nil
}

func main() {
	lambda.Start(handler)
}

// make build
// sls invoke local -f files3 -d '{"Txt": " content of the file +++"}'
// sam local invoke MyTestGoServicesDevFiles3 -e examples/files3/payload.json  <<< "{}"
//
