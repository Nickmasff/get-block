package main

import (
	"get-block/application/balance"
	getBlock "get-block/tests/mocks"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

const mostChangedAddress = "0x4ecde565958dd14ac37dcc9d3d08125f17c7aaaf"

func TestGetMostChangedBalanceDto(t *testing.T) {
	mockClient := getBlock.NewClient(getBlock.NewConfig("https://example.com", "secret-key"))
	service := balance.NewService(mockClient)

	address, err := service.GetMostChangedBalanceAddress()
	if err != nil {
		log.Fatal(err)
	}

	assert.True(t, address.Address == mostChangedAddress)

}
