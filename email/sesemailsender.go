package email

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	"cardap.in/lambda/awsenvironment"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	Sender  = "Cardapin <oi@cardap.in>"
	CharSet = "UTF-8"
)

type Email struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Comment     string `json:"comment"`
	Email       string `json:"email"`
	CompanyName string `json:"companyName"`
}

func Send(to string, ccToMe bool, templateBody string, subject string, mailInfo Email) bool {
	log.Printf("Get aws connection")
	sess := awsenvironment.ConnectAWS()
	svc := ses.New(sess)
	buf := &bytes.Buffer{}
	htmlBody, err := template.New("body").Parse(templateBody)
	htmlBody.Execute(buf, mailInfo)
	templateAsString := buf.String()
	ccAddresses := make([]*string, 0)
	if ccToMe {
		ccAddresses = append(ccAddresses, aws.String(Sender))
	}
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: ccAddresses,
			ToAddresses: []*string{
				aws.String(to),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(templateAsString),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(Sender),
	}

	log.Printf("Sending email to: " + to)

	result, err := svc.SendEmail(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}

		return false
	}

	fmt.Println("Email Sent to address: " + to)
	fmt.Println(result)
	return true
}
