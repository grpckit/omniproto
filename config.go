package main

import (
	"github.com/spf13/viper"
)

// Descriptors represents configuration for
// FileDescriptor output
type Descriptors struct {
	Enabled           bool
	IncludeImports    bool `mapstructure:"include_imports"`
	IncludeSourceInfo bool `mapstructure:"include_source_info"`
	Output            string
}

// Plugin represents a gRPC plugin config
type Plugin struct {
	Name   string
	Output string
	Args   string
}

// Config represents an omniproto configuration
type Config struct {
	RootDir     string
	Sources     []string
	Output      string
	Descriptors Descriptors
	Plugins     []Plugin
	Debug       bool
	DryRun      bool `mapstructure:"dry_run"`
}

// ReadConfig reads a config file and returns a
// populated Config struct
func ReadConfig() (*Config, error) {
	viper.SetDefault("rootdir", "protorepo")
	viper.SetDefault("output", "gen")
	viper.SetDefault("debug", false)
	viper.SetDefault("dry_run", false)
	viper.SetDefault("descriptors", Descriptors{Enabled: true, IncludeImports: true, IncludeSourceInfo: true})

	viper.SetConfigName("omniproto")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
