package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFech(t *testing.T) {

	service := CurencyExchngeRateService{}

	_, err := service.Fetch()

	require.Nil(t, err)
}
