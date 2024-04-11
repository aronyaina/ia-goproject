package config

import (
	"errors"
	"fmt"

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
	TextGeneration      string `env:"URL_TEXT_GENERATION"`
}

type Config struct {
	DB     DBConfig
	Server ServerConfig
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.SetDefault("DB.Host", "localhost")
	viper.SetDefault("DB.Port", "5432")
	viper.SetDefault("Server.Port", "8080")

	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	name := viper.GetString("DB_NAME")
	serverPort := viper.GetString("SERVER_PORT")
	token := viper.GetString("API_TOKEN")
	textToImage := viper.GetString("URL_TEXT_TO_IMAGE")
	imageToText := viper.GetString("URL_IMAGE_TO_TEXT")
	imageClassification := viper.GetString("URL_IMAGE_CLASSIFICATION")
	textSummerization := viper.GetString("URL_TEXT_SUMMERIZATION")
	textClassification := viper.GetString("URL_TEXT_CLASSIFICATION")
	textGeneration := viper.GetString("URL_TEXT_GENERATION")

	if host == "" || port == "" || user == "" || password == "" || name == "" || serverPort == "" || token == "" || textToImage == "" || imageToText == "" || imageClassification == "" || textClassification == "" || textSummerization == "" || textGeneration == "" {
		return nil, errors.New("Missing required environment variables .")
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
			Port:                serverPort,
			Token:               token,
			TextToImage:         textToImage,
			ImageToText:         imageToText,
			ImageClassification: imageClassification,
			TextSummerization:   textSummerization,
			TextClassification:  textClassification,
			TextGeneration:      textGeneration,
		},
	}

	return db, nil
}
