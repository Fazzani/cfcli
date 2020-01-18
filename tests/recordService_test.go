package main

import (
	"testing"

	. "github.com/fazzani/cfcli/service"
	"github.com/stretchr/testify/assert"
)

func TestCreateRecordOk(t *testing.T) {
	s := HttpRecordService{}
	assert.NotNil(t, s, "s not nil")
}
