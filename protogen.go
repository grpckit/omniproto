package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const space = " "

// GenProtoString generates a protoc command
// from a config
func GenProtoString(config *Config) string {
	var b strings.Builder

	b.WriteString("protoc")
	b.WriteString(space)
	b.WriteString(fmt.Sprintf("-I%v%v", config.RootDir, space))

	for _, plugin := range config.Plugins {
		b.WriteString(fmt.Sprintf("--%v_out=%v%v%v", plugin.Name, getArgs(plugin), getOutputForPlugin(plugin, config), space))
	}

	if config.Descriptors.Enabled {
		if len(config.Descriptors.Output) > 0 {
			b.WriteString(fmt.Sprintf("--descriptor_set_out=%v%v", config.Descriptors.Output, space))
		} else {
			b.WriteString(fmt.Sprintf("--descriptor_set_out=%v/descriptors.pb%v", config.Output, space))
		}
		if config.Descriptors.IncludeImports {
			b.WriteString("--include_imports")
			b.WriteString(space)
		}
		if config.Descriptors.IncludeSourceInfo {
			b.WriteString("--include_source_info")
			b.WriteString(space)
		}
	}
	writeSources(config, &b)

	return b.String()
}

func getArgs(plugin Plugin) string {
	if len(plugin.Args) > 0 {
		return fmt.Sprintf("%v:", plugin.Args)
	}
	return ""
}

func writeSources(config *Config, b *strings.Builder) {
	for _, source := range config.Sources {
		if strings.HasSuffix(source, ".proto") {
			b.WriteString(source)
			b.WriteString(space)
		} else {
			expandSource(config.RootDir, source, b)
		}
	}
}

func expandSource(root, source string, b *strings.Builder) {
	root = fmt.Sprintf("%v/", root)
	err := filepath.Walk(fmt.Sprintf("%v/%v", root, source), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(info.Name(), ".proto") {
			b.WriteString(strings.TrimPrefix(path, root))
			b.WriteString(space)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func getOutputForPlugin(plugin Plugin, config *Config) string {
	if len(plugin.Output) > 0 {
		return plugin.Output
	}
	return config.Output
}
