package handlers

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func uploadFile(code string) error {
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	ctx := context.Background()
	driveService, err := drive.NewService(ctx, option.WithTokenSource(googleOauthConfig.TokenSource(ctx, token)))
	if err != nil {
		fmt.Println(err.Error())
	}

	f := &drive.File{Name: "Hellodrive.pdf"}
	driveFile := driveService.Files.Create(f)

	localFile, err := os.Open("hello.pdf")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer localFile.Close()
	_, err = driveFile.Media(localFile).Do()
	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
