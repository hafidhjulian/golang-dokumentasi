package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func S3services() *s3.S3 {
	awsAccessKey := "<ACCESS_KEY>"
	awsSecretKey := "<SECRET_KEY>"

	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, "")
	_, err := creds.Get()

	if err != nil {
		fmt.Println(err)
	}

	s3Region := "us-east-1"
	endpoint := "is3.cloudhost.id"

	cfg := aws.NewConfig().WithRegion(s3Region).WithCredentials(creds).WithEndpoint(endpoint)

	s3Connection := s3.New(session.New(), cfg)
	return s3Connection

}

func S3upload(base64File string, objectKey string) error {
	decode, err := base64.StdEncoding.DecodeString(base64File)

	if err != nil {
		return err
	}

	awsSession := S3services()

	uploadParams := &s3.PutObjectInput{
		Bucket:      aws.String("<BUCKET_NAME>"),
		Key:         aws.String(objectKey),
		Body:        bytes.NewReader(decode),
		ContentType: aws.String(http.DetectContentType(decode)),
	}

	_, err = awsSession.PutObject(uploadParams)

	return err
}

func S3downloadtransaksiblumb(filepath string) (data string, err error) {
	awsSession := S3services()
	// h := md5.New()
	r, _ := awsSession.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("<BUCKET_NAME>"),
		Key:    aws.String(filepath),
	})

	// md5s := base64.StdEncoding.EncodeToString(h.Sum(nil))
	// r.HTTPRequest.Header.Get("Content-Encoding")
	url, header, err := r.PresignRequest(10 * time.Hour)
	if err != nil {
		fmt.Println("error presigning request", err)
		return "", err
	}
	fmt.Println("Header", header)
	fmt.Println("URL", url)
	return url, err
}
