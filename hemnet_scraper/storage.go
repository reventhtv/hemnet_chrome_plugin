package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
)

type Storage struct {
	BucketName string
	key        string
	session    *session.Session
}

func check(e error) {
	if e != nil {
		log.Fatalf("%d", e)
	}
}

func NewStorage(bucketName, key string) *Storage {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-north-1"),
	})
	// Test upload
	s := &Storage{
		bucketName,
		key,
		sess,
	}
	s.UploadToS3("test_data_full.csv")
	
	check(err)
	return s
}

func (s *Storage) UploadToS3(fileName string) {
	
	f, err := os.Open(fileName)
	check(err)
	finalKey := s.key + "" + fileName
	// Upload input parameters
	upParams := &s3manager.UploadInput{
		Bucket: &s.BucketName,
		Key:    &finalKey,
		Body:   f,
	}
	
	// Perform an upload.
	uploader := s3manager.NewUploader(s.session)
	// Perform upload with options different than the those in the Uploader.
	_, err = uploader.Upload(upParams, func(u *s3manager.Uploader) {
		u.PartSize = 10 * 1024 * 1024 // 10MB part size
		u.LeavePartsOnError = true    // Don't delete the parts if the upload fails.
	})
	check(err)
}
