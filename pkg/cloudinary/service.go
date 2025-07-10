package cloudinaryfx

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/oliwallywonka/alpaca_backend/settings"
)

type Service interface {
	UploadImage(file multipart.File) (*uploader.UploadResult, error)
	DeleteImage(imageID string) error
}

type cloudinaryService struct {
	settings *settings.Settings
	ctx      context.Context
}

func NewService(s *settings.Settings, ctx context.Context) Service {
	return &cloudinaryService{
		settings: s,
		ctx:      ctx,
	}
}

func (s *cloudinaryService) UploadImage(file multipart.File) (*uploader.UploadResult, error) {
	cld, err := cloudinary.NewFromURL(s.settings.CloudinaryURL)
	if err != nil {
		fmt.Println(err)
		return nil, ConnectionError
	}

	uploadResult, err := cld.Upload.Upload(s.ctx, file, uploader.UploadParams{})
	if err != nil {
		fmt.Println(err)
		return nil, UploadError
	}
	return uploadResult, nil
}

func (s *cloudinaryService) DeleteImage(imageID string) error {
	cld, err := cloudinary.NewFromURL(s.settings.CloudinaryURL)
	if err != nil {
		fmt.Println(err)
		return ConnectionError
	}
	_, err = cld.Upload.Destroy(s.ctx, uploader.DestroyParams{PublicID: imageID})
	if err != nil {
		fmt.Println(err)
		return DeleteError
	}
	return nil
}
