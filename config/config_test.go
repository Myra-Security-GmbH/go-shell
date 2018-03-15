// +build !testing

package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadValidConfigFile(t *testing.T) {
	filepath := "./valid-config.yml"

	cfg, err := ReadConfigFile(filepath)

	require.Nil(t, err, "Config [%s] file is valid there should be no errors", filepath)
	require.NotNil(t, cfg, "Expected != nil for config but null returned")
	require.NotNil(t, cfg.Login, "No users found")
	require.Equal(t, "https://beta.myracloud.com", cfg.Endpoint)
	require.Equal(t, "en", cfg.Language)
}

func TestReadNotExistingFile(t *testing.T) {
	cfg, err := ReadConfigFile("does-not-exist")

	require.NotNil(t, err, "Should return an error that file cannot be opened")
	require.Nil(t, cfg, "Invalid return on file error")
}

func TestInvalidFile(t *testing.T) {
	cfg, err := ReadConfigFile("./invalid-config.yml")

	require.NotNil(t, err, "Should return an error that file cannot be opened")
	require.Nil(t, cfg, "Invalid return on file error")
}
