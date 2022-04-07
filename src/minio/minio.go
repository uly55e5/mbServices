package minio

import (
	"context"
	min "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/notification"
	"github.com/uly55e5/mbServices/src/openapi"
	"github.com/uly55e5/mbServices/src/redisstore"
	"log"
)
import cred "github.com/minio/minio-go/v7/pkg/credentials"

var client *min.Client

func StartAllNotifications() error {
	cs, err := redisstore.GetConnections()
	if err != nil {
		return err
	}
	cc, err := openapi.ConvertToConnection(cs)
	for _, c := range cc {
		go GetNotifications(&c)
	}
	return nil

}

func GetNotifications(connection *openapi.Connection) error {
	endpoint := connection.Minio.Endpoint
	username := connection.Minio.Username
	password := connection.Minio.Password
	useSSL := connection.Minio.UseSsl

	options := min.Options{
		Creds:  cred.NewStaticV4(username, password, ""),
		Secure: useSSL,
	}
	var err error
	client, err = min.New(endpoint, &options)
	if err != nil {
		return err
	}
	for info := range client.ListenBucketNotification(context.Background(), connection.Filter.Bucket, connection.Filter.Prefix, connection.Filter.Suffix, []string{
		"s3:ObjectCreated:*",
		"s3:ObjectAccessed:*",
		"s3:ObjectRemoved:*",
	}) {
		handleNotification(&info, connection)
	}
	return nil
}

func handleNotification(info *notification.Info, connection *openapi.Connection) {
	log.Println(info.Err)
	log.Println(info.Records)
}
