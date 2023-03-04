package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/pincheng0101/go-linebot-server-template/config"
	"github.com/skip2/go-qrcode"
	"google.golang.org/api/option"
)

type Storage struct {
	ctx    context.Context
	bucket *storage.BucketHandle
}

func NewStorage() (*Storage, error) {
	cfg, _ := config.LoadConfig()

	ctx := context.Background()

	config := &firebase.Config{
		StorageBucket: cfg.Firebase.StorageBucket,
	}

	opt := option.WithCredentialsFile(cfg.Firebase.ServiceAccountFile)

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return nil, fmt.Errorf("Error initializing Firebase app: %v", err)
	}

	client, err := app.Storage(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error initializing Firebase Storage client: %v", err)
	}

	bucket, err := client.Bucket(cfg.Firebase.StorageBucket)
	if err != nil {
		return nil, fmt.Errorf("Error getting bucket handle: %v", err)
	}

	return &Storage{ctx: ctx, bucket: bucket}, nil
}

func (s *Storage) UploadQrcode(userID string) error {
	object := s.bucket.Object(userID + ".png")
	attrs, err := object.Attrs(s.ctx)

	if err != nil {
		return fmt.Errorf("Error checking object attributes: %v", err)
	}

	// If the file already exists, return success
	if attrs != nil {
		return nil
	}

	qrcodeBytes, err := qrcode.Encode(userID, qrcode.Medium, 256)
	if err != nil {
		return fmt.Errorf("Error generating QR code: %v", err)
	}

	writer := object.NewWriter(s.ctx)
	if _, err = io.Copy(writer, bytes.NewReader(qrcodeBytes)); err != nil {
		writer.Close()
		return fmt.Errorf("Error copying QR code bytes to writer: %v", err)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("Error closing writer: %v", err)
	}

	return nil
}
