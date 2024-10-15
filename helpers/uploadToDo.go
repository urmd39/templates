package helpers

import (
	"bufio"
	"bytes"
	"os"
	"path"
	"templates/infrastructure"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

var AWSExcelContentType = "application/vnd.ms-excel"

func UploadToDOSpace(localFolder, destFolder, fileName, contentType string) (err error) {
	localFile := path.Join(localFolder, fileName)
	// create buffer
	buffer, err := retrieveROM(localFile)
	if err != nil {
		infrastructure.ErrLog.Printf("Cannot get buffer from file error: %+v\n", err)
	}

	// convert buffer to reader
	fileBytes := bytes.NewReader(buffer)

	path := path.Join(destFolder, fileName)
	object := &s3.PutObjectInput{
		ACL:          aws.String("public-read"),
		Bucket:       aws.String("anvita"),
		Key:          aws.String(path),
		Body:         fileBytes,
		ContentType:  aws.String(contentType),
		CacheControl: aws.String("no-cache"),
		Expires:      aws.Time(time.Now()),
	}

	resp, err := infrastructure.DOSpaceConn.S3Client.PutObject(object)
	if err != nil {
		infrastructure.ErrLog.Printf("Put object to S3 error: %+v\n", err)
	}

	infrastructure.InfoLog.Printf("Response %+v\n", resp)
	return nil
}

func retrieveROM(filename string) ([]byte, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	return bytes, err
}
