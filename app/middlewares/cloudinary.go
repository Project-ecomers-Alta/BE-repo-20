package middlewares

import (
	"BE-REPO-20/app/configs"
	"context"
	"fmt"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func CloudinaryUpload(file multipart.File) (*uploader.UploadResult, error) {
	ctx := context.Background()

	cldUrl := fmt.Sprintf("cloudinary://%s:%s@%s", configs.CLOUDINARY_KEY, configs.CLOUDINARY_SECRET, configs.CLOUDINARY_CLOUD_NAME)

	cldService, _ := cloudinary.NewFromURL(cldUrl)
	url, errUpl := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})
	if errUpl != nil {
		return nil, errUpl
	}
	fmt.Println("cloud test")
	return url, nil
}
