package jobsql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// need to install mysql
func TestDBSchema(t *testing.T) {
	_, err := InitDB(DefaultDBOption)
	assert.Equal(t, nil, err)
}
