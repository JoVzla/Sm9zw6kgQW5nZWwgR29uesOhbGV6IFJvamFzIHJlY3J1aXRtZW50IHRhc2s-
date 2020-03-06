package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitLogger(t *testing.T) {
	t.Run("Don't Panic", func(t *testing.T) {

		assert.NotPanics(t, func() {
			InitLogger()
		})

		assert.NotNil(t, Log)
	})
}