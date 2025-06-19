package applications

import (
	"github.com/sdslabs/nymeria/internal/database"
	"github.com/sdslabs/nymeria/internal/database/schema"
	"github.com/sdslabs/nymeria/internal/logger"
)

func GetAllApplications() ([]schema.Application, error) {
	var applications []schema.Application

	result := database.DB.Find(&applications)
	if result.Error != nil {
		logger.Err(result.Error).Msg("failed to get applications")
		return nil, result.Error
	}

	return applications, nil
}

func CreateApplication(application schema.Application) error {
	result := database.DB.Create(&application)
	if result.Error != nil {
		logger.Err(result.Error).Msg("failed to create application")
		return result.Error
	}

	return nil
}

func GetApplicationByID(id string) (schema.Application, error) {
	var application schema.Application

	result := database.DB.First(&application, "id = ?", id)
	if result.Error != nil {
		logger.Err(result.Error).Msg("failed to get application")
		return schema.Application{}, result.Error
	}

	return application, nil
}

func UpdateApplication(id string, application schema.Application) error {
	result := database.DB.Model(&schema.Application{}).Where("id = ?", id).Updates(application)
	if result.Error != nil {
		logger.Err(result.Error).Msg("failed to update application")
		return result.Error
	}

	return nil
}

func DeleteApplication(id string) error {
	result := database.DB.Delete(&schema.Application{}, "id = ?", id)
	if result.Error != nil {
		logger.Err(result.Error).Msg("failed to delete application")
		return result.Error
	}

	return nil
}
