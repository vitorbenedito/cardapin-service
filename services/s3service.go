package services

import (
	"time"

	"cardap.in/awsenvironment"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

var sess = awsenvironment.ConnectAWS()

type S3Services struct {
}

func (*S3Services) GeneratePresignedUrlToPut(fileName string) (string, error) {
	svc := s3.New(sess)

	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(awsenvironment.AwsS3Bucket),
		Key:    aws.String(fileName),
	})

	url, err := req.Presign(15 * time.Minute)
	return url, err
}
