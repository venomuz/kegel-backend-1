package service

import (
	"context"
	"fmt"
	"github.com/venomuz/kegel-backend/config"
	"github.com/venomuz/kegel-backend/internal/models"
	"io"
	"os"
	"regexp"
	"strconv"
	"time"
)

type FilesService struct {
	cfg config.Config
}

func NewFilesService(cfg config.Config) *FilesService {
	return &FilesService{cfg: cfg}
}

func (f *FilesService) Save(ctx context.Context, file models.File) (string, error) {
	src, err := file.File.Open()
	if err != nil {
		fmt.Println("ee")
		return "", err
	}

	src.Close()

	if _, err = os.Stat(f.cfg.StaticFilePath); os.IsNotExist(err) {
		err = os.Mkdir(f.cfg.StaticFilePath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	createPath := f.cfg.StaticFilePath + file.Path

	if _, err = os.Stat(createPath); os.IsNotExist(err) {
		err = os.Mkdir(createPath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	pattern := regexp.MustCompile("\\.[0-9a-z]+$")

	extension := pattern.FindString(file.File.Filename)
	if extension == "" {
		return "", models.ErrFileName
	}

	now := strconv.FormatInt(time.Now().UnixNano(), 10)

	newName := now + extension

	dst := createPath + "/" + newName

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(out, src)
	if err != nil {
		return "", err
	}

	return newName, nil
}

func (f *FilesService) DeleteByName(ctx context.Context, path, filename string) error {

	err := os.Remove(f.cfg.StaticFilePath + path + "/" + filename)
	if err != nil {
		return err
	}
	return nil
}
