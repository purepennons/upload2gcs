package main

import (
	"fmt"
	"io"
	"log"
	"os"

	// Imports the Google Cloud Storage client package.
	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
)

func main() {
	//projectID := os.Args[1]
	bucketName := os.Args[1]
	objectName := os.Args[2]
	filename := os.Args[3]

	fmt.Printf("b=%s, o=%s, f=%s\n", bucketName, objectName, filename)

	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	if err := write(client, bucketName, objectName, filename); err != nil {
		log.Fatalf("Failed to update a object", err)
	}

}

func write(client *storage.Client, bucket, object, filename string) error {
	ctx := context.Background()

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		return err
	}

	if err := wc.Close(); err != nil {
		return err
	}

	return nil

}
