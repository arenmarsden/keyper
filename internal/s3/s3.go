package s3

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/arenmarsden/keyper/internal/config"
)

type Client struct {
	client *minio.Client
}

func NewClient() (*Client, error) {
	config, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: config.UseSSL,
		Region: config.Region,
	})
	if err != nil {
		return nil, err
	}

	return &Client{client: client}, nil
}

func (c *Client) Validate(ctx context.Context) error {
	// We only care about the error here, minio will not return a
	// BucketInfo if credentials failed.
	_, err := c.client.ListBuckets(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Upload(ctx context.Context, bucket, key string, reader io.Reader, size int64) error {
	_, err := c.client.PutObject(ctx, bucket, key, reader, size, minio.PutObjectOptions{})
	return err
}

func (c *Client) Download(ctx context.Context, bucket, key string) error {
	_, err := c.client.GetObject(ctx, bucket, key, minio.GetObjectOptions{})
	return err
}

func (c *Client) List(ctx context.Context, bucket, prefix string) ([]string, error) {
	var objects []string
	for obj := range c.client.ListObjects(ctx, bucket, minio.ListObjectsOptions{Prefix: prefix}) {
		if obj.Err != nil {
			return nil, obj.Err
		}
		objects = append(objects, obj.Key)
	}
	return objects, nil
}

func (c *Client) Delete(ctx context.Context, bucket, key string) error {
	return c.client.RemoveObject(ctx, bucket, key, minio.RemoveObjectOptions{})
}
