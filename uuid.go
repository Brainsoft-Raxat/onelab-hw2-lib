package testlibrary

import "github.com/hart87/goflake/generator"

func GenerateUUID() string {
	return generator.GenerateIdentifier()
}
