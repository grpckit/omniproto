package main

import (
	"bytes"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestCanGenerateFromExampleConfig(t *testing.T) {
	config, err := ReadConfig()

	assert.NoError(t, err)

	result := GenProtoString(config)
	assert.NotEmpty(t, result)

	assert.Contains(t, result, "-Iprotorepo-example")
	assert.Contains(t, result, "--go_out=gen")
	assert.Contains(t, result, "--go_opt=paths=source_relative")
	assert.Contains(t, result, "--include_imports")
	assert.Contains(t, result, "grpckit/accounts/user.proto")
	assert.Contains(t, result, "grpckit/base.proto")
	assert.Contains(t, result, "--descriptor_set_out")
	assert.Contains(t, result, "--validate_out=lang=go:gen")
}

func genConfig(t *testing.T, input []byte) *Config {
	var config Config

	viper.ReadConfig(bytes.NewBuffer((input)))

	if err := viper.Unmarshal(&config); err != nil {
		t.Fatal("could not read yaml config")
	}

	return &config
}
