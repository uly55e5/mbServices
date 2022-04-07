package minio

import (
	"context"
	min "github.com/minio/minio-go/v7"
)
import cred "github.com/minio/minio-go/v7/pkg/credentials"

var client *min.Client

func GetNotifications() error {
	endpoint := "minio:9000"
	username := "massbank"
	password := "massbank"
	useSSL := false

	c := cred.NewStaticV4(username, password, "")
	options := min.Options{
		Creds:        c,
		Secure:       useSSL,
		Transport:    nil,
		Region:       "",
		BucketLookup: 0,
		CustomMD5:    nil,
		CustomSHA256: nil,
	}
	var err error
	client, err = min.New(endpoint, &options)
	if err != nil {
		return err
	}
	for info := range client.ListenBucketNotification(context.Background(), "massbank", "", ".txt", []string{
		"s3:ObjectCreated:*",
		"s3:ObjectAccessed:*",
		"s3:ObjectRemoved:*",
	}) {
		println(info.Err)
		println(info.Records)
	}
	return nil
}
