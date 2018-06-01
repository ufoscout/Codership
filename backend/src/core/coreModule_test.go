package core

import (
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/ufoscout/Codership/backend/src/core/configuration"
	"github.com/ufoscout/Codership/backend/src/util"
)

func Test(t *testing.T) {

	config := configuration.LoadConfig(path.Join(util.MainFolderPath(), configuration.CONFIG_FILE_NAME))
	config.Server.Port = 0

	module := CoreModule(&config)
	go module.Start()

	time.Sleep(300 * time.Millisecond)

	assert.True(t, module.ServerPort() > 0)

	module.Stop()

}
