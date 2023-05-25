package parser_test

import (
	"testing"

	parser "github.com/khatibomar/paytabs-challenge/internal/parser"
	"github.com/stretchr/testify/require"
)

func TestParseJsonFile(t *testing.T) {
	p := parser.New()
	path := "../../data/accounts-mock.json"
	rawAccounts, err := p.ParseJson(path)
	require.NoError(t, err)
	require.Equal(t, len(rawAccounts), 500)
}
