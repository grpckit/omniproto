package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestCanGenerateFromExampleConfig(t *testing.T) {
	config, err := ReadConfig()

	assert.NoError(t, err)

	result := GenProtoString(config)
	assert.NotEmpty(t, result)

	fmt.Println(result)

}

func genConfig(t *testing.T, input []byte) *Config {
	var config Config

	viper.ReadConfig(bytes.NewBuffer((input)))

	if err := viper.Unmarshal(&config); err != nil {
		t.Fatal("could not read yaml config")
	}

	return &config
}
