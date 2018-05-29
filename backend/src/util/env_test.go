package util

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MainFolderPath_Unit(t *testing.T) {

	joined := filepath.Join(MainFolderPath(), "/config/config.properties")
	_, err := os.Stat(joined)

	assert.Nil(t, err)
}
