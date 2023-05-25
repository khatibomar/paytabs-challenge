package parser

import (
	"encoding/json"
	"os"
)

type Parser struct {
}

func New() *Parser {
	return &Parser{}
}

type RawAccounts []struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Balance string `json:"balance"`
}

func (p *Parser) ParseJson(path string) (RawAccounts, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var rawAccounts RawAccounts

	err = json.Unmarshal(content, &rawAccounts)
	if err != nil {
		return nil, err
	}
	return rawAccounts, nil
}
