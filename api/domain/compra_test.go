package domain_test

import (
	"api/domain"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewCompra(t *testing.T) {
	compra, err := domain.NewCompra(time.Now())
	require.NotNil(t, compra)
	require.Nil(t, err)
}
