package awsenvironment

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"
)

const (
	AwsS3Region = "us-east-2"
	AwsS3Bucket = "cardapin.images"
	AwsUrl      = "https://s3." + AwsS3Region + ".amazonaws.com/" + AwsS3Bucket
)

func GetProfile() string {
	godotenv.Load()
	return os.Getenv("aws_profile")
}

func ConnectAWS() *session.Session {
	profile := GetProfile()
	if profile != "" {
		sess, err := session.NewSessionWithOptions(session.Options{Config: aws.Config{Region: aws.String(AwsS3Region)}, Profile: profile})
		if err != nil {
			log.Panicf("Error to get aws session: " + err.Error())
			panic(err)
		}
		return sess
	}
	sess, err := session.NewSession(&aws.Config{Region: aws.String(AwsS3Region)})
	if err != nil {
		panic(err)
	}
	return sess
}
