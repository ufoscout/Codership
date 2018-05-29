package configuration

import (
	"testing"

	"path"

	"github.com/stretchr/testify/assert"
	"github.com/ufoscout/Codership/backend/src/util"
)

func Test_load_Unit(t *testing.T) {
	config := Load(path.Join(util.MainFolderPath(), CONFIG_FILE_NAME))
	assert.Equal(t, 8080, config.Server.Port)
	assert.Equal(t, "../front/dist", config.Frontend.ResourcesPath)
}
