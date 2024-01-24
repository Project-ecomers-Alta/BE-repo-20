// aws/session.go
package aws

import (
	"BE-REPO-20/app/configs"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Session struct {
	S3Session *session.Session
}

func NewSession() (*Session, error) {
	// Load configurations

	credential := credentials.NewStaticCredentials(
		configs.AWS_AccessKeyID,
		configs.AWS_SecretAccessKey,
		"",
	)

	awsConfig := aws.Config{
		Region:      aws.String(configs.AWS_Region),
		Credentials: credential,
	}

	s, err := session.NewSession(&awsConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create S3 session: %s", err.Error())
	}

	return &Session{s}, nil
}

// Upload mengunggah file ke S3 dan mengembalikan URL file yang diunggah
func (s Session) Upload(fileHeader *multipart.FileHeader, fileReader io.Reader) (string, error) {
	fileBuffer, err := ioutil.ReadAll(fileReader)
	if err != nil {
		log.Printf("failed to read uploaded file, error: %v", err.Error())
		return "", err
	}

	// Generate a unique filename
	uniqueFilename := generateUniqueFilename(fileHeader.Filename)

	// uploading to S3
	_, err = s3.New(s.S3Session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(configs.AWS_BucketName),
		Key:                  aws.String(uniqueFilename),
		Body:                 bytes.NewReader(fileBuffer),
		ContentLength:        aws.Int64(int64(len(fileBuffer))),
		ContentType:          aws.String(http.DetectContentType(fileBuffer)),
		ContentDisposition:   aws.String("inline"),
		ServerSideEncryption: aws.String("AES256"),
	})

	if err != nil {
		log.Printf("failed to upload %s to S3, error: %v", uniqueFilename, err.Error())
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", configs.AWS_BucketName, uniqueFilename)
	fmt.Printf("Uploaded File Url %s\n", url)

	return url, nil
}

// generateUniqueFilename generates a unique filename based on the original filename.
func generateUniqueFilename(originalFilename string) string {
	// Use timestamp and random string to create a unique filename
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	randomString := randString(6) // Adjust the length of the random string as needed
	return fmt.Sprintf("%d_%s", timestamp, randomString)
}

// randString generates a random string of a specified length.
func randString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
