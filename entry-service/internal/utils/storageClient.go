package utils

import (
	"context"
	"mime/multipart"

	"firebase.google.com/go/v4/storage"
)

func UploadToStorage(ctx context.Context, storageClient *storage.Client, file multipart.File, fileName, bucketName, entityID, versionId, resourceType string) (string, error) {
	objectName := ""

	bucket, err := storageClient.DefaultBucket()
	if err != nil {
		return "", err
	}
	writer := bucket.Object(objectName).NewWriter(ctx)
	writer.ContentType = ""

	var fileURL string
	return fileURL, nil
}
