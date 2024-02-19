package config

import (
	"errors"

	"github.com/spf13/viper"
)

type DBConfig struct {
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
}

type ServerConfig struct {
	Port                string `env:"PORT"`
	Token               string `env:"API_TOKEN"`
	TextToImage         string `env:"URL_TEXT_TO_IMAGE"`
	ImageToText         string `env:"URL_IMAGE_TO_TEXT"`
	ImageClassification string `env:"URL_IMAGE_CLASSIFICATION"`
	TextSummerization   string `env:"URL_TEXT_SUMMERIZATION"`
	TextClassification  string `env:"URL_TEXT_CLASSIFICATION"`
}

type Config struct {
	DB     DBConfig
	Server ServerConfig
}

func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()

	viper.SetDefault("DB.Host", "localhost")
	viper.SetDefault("DB.Port", "5432")
	viper.SetDefault("Server.Port", "8080")

	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	name := viper.GetString("DB_NAME")
	serverPort := viper.GetString("PORT")
	token := viper.GetString("API_TOKEN")
	text_to_image := viper.GetString("URL_TEXT_TO_IMAGE")
	image_to_text := viper.GetString("URL_IMAGE_TO_TEXT")
	image_classification := viper.GetString("URL_IMAGE_CLASSIFICATION")
	text_summerization := viper.GetString("URL_TEXT_SUMMARIZATION")
	text_classification := viper.GetString("URL_TEXT_CLASSIFICATION")

	if host == "" || port == "" || user == "" || password == "" || name == "" || serverPort == "" || token == "" || text_to_image == "" || image_to_text == "" || image_classification == "" || text_classification == "" || text_summerization == "" {
		return nil, errors.New("missing required environment variables")
	}

	db := &Config{
		DB: DBConfig{
			Host:     host,
			Port:     port,
			User:     user,
			Password: password,
			Name:     name,
		},
		Server: ServerConfig{
			Port: serverPort,
		},
	}

	return db, nil
}
