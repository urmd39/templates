package infrastructure

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type DOSpaceConnection struct {
	Session  *session.Session
	S3Client *s3.S3
}

func setupDOSpaceConnection() DOSpaceConnection {
	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(SpacesKey, SpacesSecret, ""),
		Endpoint:    aws.String("sgp1.digitaloceanspaces.com"),
		Region:      aws.String("sgp1"),
	}

	newSession, err := session.NewSession(s3Config)
	if err != nil {
		ErrLog.Printf("Can not connection to DO Space: %+v\n", err)
		ErrLog.Fatal(err)
	}

	s3Client := s3.New(newSession)

	return DOSpaceConnection{
		Session:  newSession,
		S3Client: s3Client,
	}
}
