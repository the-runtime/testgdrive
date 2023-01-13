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

var fileid = ""

func UploadFile(token *oauth2.Token, googleOauthConfig *oauth2.Config, filename string, i int) {

	fmt.Println("up runnning %d", i)
	ctx := context.Background()
	driveService, err := drive.NewService(ctx, option.WithTokenSource(googleOauthConfig.TokenSource(ctx, token)))
	if err != nil {
		fmt.Println(err.Error())
	}
	//var driveFile *drive.FilesCreateCall
	if i == 0 {
		ids, err := driveService.Files.GenerateIds().Space("drive").Count(1).Do()
		if err != nil {
			fmt.Println(err.Error())
		}
		id := ids.Ids[0]
		f := &drive.File{Name: filename, Id: id}
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
		fileid = f.Id
		fmt.Println(f.Id)
	} else {

		f := &drive.File{Name: filename, Id: fileid}
		driveFile := driveService.Files.Update(fileid, f)
		//driveFile := driveService.Files.Create(f)

		localFile, err := os.Open(fmt.Sprintf(filename+"%d", i))
		if err != nil {
			fmt.Println(err.Error())
		}
		defer localFile.Close()
		_, err = driveFile.Media(localFile).Do()
		if err != nil {
			fmt.Println(err.Error())
		}

		// f := &drive.File{Name: filename}
		// fmt.Println("File id is ", fileid)
		// driveFile := driveService.Files.Update(fileid, f)
		// localFile, err := os.Open(fmt.Sprintf(filename+"%d", i))
		// if err != nil {
		// 	fmt.Println(err.Error())
		// }
		// defer localFile.Close()
		// _, err = driveFile.Media(localFile).Do()
		// if err != nil {
		// 	fmt.Println("Problem in uploading 2nd part")
		// 	fmt.Println(err.Error())
		// }
	}

}
