package services

import (
	"log"
	"regexp"
	"strings"
	"time"
	"unicode"

	"cardap.in/lambda/model"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type ImageServices struct {
}

func (*ImageServices) GeneratePresignedUrlToPut(fileRequest model.FileRequest) (string, error) {

	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	companyName, _, _ := transform.String(t, fileRequest.CompanyName)
	companyName = reg.ReplaceAllString(companyName, "-")
	fileName := reg.ReplaceAllString(fileRequest.FileName, "")
	s3FileName := companyName + "-" + fileName + "-" + time.Now().Format("20060102150405") + "." + fileRequest.Extension

	s3FileName = strings.ReplaceAll(strings.ToLower(s3FileName), " ", "-")
	s3Service := S3Services{}
	url, err := s3Service.GeneratePresignedUrlToPut(s3FileName)
	if err != nil {
		log.Printf(err.Error())
	}
	jsonToReturn := "{\"url\":\"" + url + "\",\"fileName\":\"" + s3FileName + "\"}"
	return jsonToReturn, err
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}
