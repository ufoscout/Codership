package configuration

import (
	"testing"

	"path"

	"github.com/stretchr/testify/assert"
	"github.com/ufoscout/Codership/backend/src/util"
)

func Test_load_Unit(t *testing.T) {
	config := LoadConfig(path.Join(util.MainFolderPath(), CONFIG_FILE_NAME))
	assert.Equal(t, 8080, config.Server.Port)
	assert.Equal(t, "../frontend/dist", config.Server.ResourcesPath)
	assert.Contains(t, config.Docker.MariaDbImage, "maria")
}
