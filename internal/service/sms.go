package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/venomuz/kegel-backend/internal/models"
	"github.com/venomuz/kegel-backend/internal/storage/mysql"
	"github.com/venomuz/kegel-backend/pkg/gen"
	"github.com/venomuz/kegel-backend/pkg/logger"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
)

type SmsService struct {
	settingsRepo mysql.Settings
	generator    gen.Generator
	log          logger.Logger
}

func NewSmsService(settingsRepo mysql.Settings, generator gen.Generator, log logger.Logger) *SmsService {
	return &SmsService{
		settingsRepo: settingsRepo,
		generator:    generator,
		log:          log,
	}
}

func (s *SmsService) SendVerificationCode(ctx context.Context, phone string) (string, error) {
	token, err := s.eskizAuth(ctx)
	if err != nil {
		return "", err
	}

	randNum := s.generator.RandomNumber(10000, 99999)

	randNumSt := strconv.FormatInt(int64(randNum), 10)

	var body bytes.Buffer

	multi := multipart.NewWriter(&body)

	mobilePhoneFiled, err := multi.CreateFormField("mobile_phone")
	if err != nil {
		return "", err
	}

	_, err = mobilePhoneFiled.Write([]byte(phone))
	if err != nil {
		return "", err
	}

	messageField, err := multi.CreateFormField("message")
	if err != nil {
		return "", err
	}

	_, err = messageField.Write([]byte("code: " + randNumSt))
	if err != nil {
		return "", err
	}

	fromField, err := multi.CreateFormField("from")
	if err != nil {
		return "", err
	}

	_, err = fromField.Write([]byte("4546"))
	if err != nil {
		return "", err
	}

	_ = multi.Close()

	req, err := http.NewRequest("POST", models.EskizBaseUrl+"/message/sms/send", &body)
	if err != nil {
		s.log.Error("error while make request eskiz/auth/login", logger.Error(err))
		return "", err
	}

	req.Header.Set("Content-Type", multi.FormDataContentType())

	req.Header.Set("Authorization", "Bearer "+token.Data.Token)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		s.log.Error("error while read body to eskiz service auth", logger.Error(err))
		return "", err
	}
	fmt.Println(string(respBody))
	return randNumSt, err
}

func (s *SmsService) eskizAuth(ctx context.Context) (models.EskizToken, error) {

	setting, err := s.settingsRepo.GetByKey(ctx, "eskiz-key")
	if err != nil {
		s.log.Error("error while get setting eskiz password", logger.Error(err))
		return models.EskizToken{}, err
	}

	var body bytes.Buffer

	multi := multipart.NewWriter(&body)

	emailField, err := multi.CreateFormField("email")
	if err != nil {
		return models.EskizToken{}, err
	}

	_, err = emailField.Write([]byte(setting.Key))
	if err != nil {
		return models.EskizToken{}, err
	}

	passwordField, err := multi.CreateFormField("password")
	if err != nil {
		return models.EskizToken{}, err
	}

	_, err = passwordField.Write([]byte(setting.Value))
	if err != nil {
		return models.EskizToken{}, err
	}

	_ = multi.Close()

	req, err := http.NewRequest("POST", models.EskizBaseUrl+"/auth/login", &body)
	if err != nil {
		s.log.Error("error while make request eskiz/auth/login", logger.Error(err))
		return models.EskizToken{}, err
	}

	req.Header.Set("Content-Type", multi.FormDataContentType())

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return models.EskizToken{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		s.log.Error("error while read body to eskiz service auth", logger.Error(err))
		return models.EskizToken{}, err
	}

	var eskiztToken models.EskizToken

	err = json.Unmarshal(respBody, &eskiztToken)
	if err != nil {
		s.log.Error("error while unmarshal to struct eskiz service auth", logger.Error(err))
		return models.EskizToken{}, err
	}

	return eskiztToken, nil
}
