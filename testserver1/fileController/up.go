package fileController

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"

	//"google.golang.org/api/oauth2/v1"
	"google.golang.org/api/option"
)

func UploadFile(token *oauth2.Token, googleOauthConfig *oauth2.Config, filename string, i int) {

	ctx := context.Background()
	driveService, err := drive.NewService(ctx, option.WithTokenSource(googleOauthConfig.TokenSource(ctx, token)))
	if err != nil {
		fmt.Println(err.Error())
	}
	//var driveFile *drive.FilesCreateCall
	if i == 0 {
		f := &drive.File{Name: filename}
		driveFile := driveService.Files.Create(f)
		localFile, err := os.Open(fmt.Sprintf(filename+"%d", i))
		if err != nil {
			fmt.Println(err.Error())
		}
		defer localFile.Close()
		_, err = driveFile.Media(localFile).Do()
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		f := &drive.File{Name: filename}
		driveFile := driveService.Files.Update(f.Id, f)
		localFile, err := os.Open(fmt.Sprintf(filename+"%d", i))
		if err != nil {
			fmt.Println(err.Error())
		}
		defer localFile.Close()
		_, err = driveFile.Media(localFile).Do()
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}
