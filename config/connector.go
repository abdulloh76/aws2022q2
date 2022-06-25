package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection(envs *ConfigVariables) (db *gorm.DB, err error) {
	dsn := envs.DOCKER_POSTGRES_URI

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
