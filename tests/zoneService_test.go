package main

import (
	"testing"

	. "github.com/fazzani/cfcli/service"
	"github.com/fazzani/cfcli/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetZoneOk(t *testing.T) {
	s := HttpZoneService{ Config: domain.Configuration{ API: domain.API{ BaseURL: "https://api.cloudflare.com/client/v4/" } }}
	zone, err := s.Get("4e7521218955d660919420f8c0e16cd6")
	assert.Nil(t, err, "err must be nil")
	assert.NotNil(t, zone, "zone must be not nil")
}
