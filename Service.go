package googledrive

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	google "github.com/leapforce-libraries/go_google"
	bigquery "github.com/leapforce-libraries/go_google/bigquery"
)

const (
	apiName string = "GoogleDrive"
	apiURL  string = "https://www.googleapis.com/drive/v3"
)

// Service stores Service configuration
//
type Service struct {
	googleService *google.Service
}

type ServiceConfig struct {
	ClientID     string
	ClientSecret string
	Scope        string
}

func NewService(serviceConfig *ServiceConfig, bigQueryService *bigquery.Service) (*Service, *errortools.Error) {
	if serviceConfig == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	if serviceConfig.ClientID == "" {
		return nil, errortools.ErrorMessage("ClientID not provided")
	}

	if serviceConfig.ClientSecret == "" {
		return nil, errortools.ErrorMessage("ClientSecret not provided")
	}

	googleServiceConfig := google.ServiceConfig{
		APIName:      apiName,
		ClientID:     serviceConfig.ClientID,
		ClientSecret: serviceConfig.ClientSecret,
		Scope:        serviceConfig.Scope,
	}

	googleService := google.NewService(googleServiceConfig, bigQueryService)

	return &Service{
		googleService,
	}, nil
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", apiURL, path)
}

func (service *Service) InitToken() *errortools.Error {
	return service.googleService.InitToken()
}
