package oasis

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetBalance(t *testing.T) {
	c, err := NewClient("02DiuDkfxbUyQCDbAMscAxT996N9Xw8YZHJhXe8Qi2J6g", "https://rosetta.oasis.dev")
	assert.NoError(t, err)

	ret, err := c.GetBalance(context.Background(), "oasis1qpg2xuz46g53737343r20yxeddhlvc2ldqsjh70p")
	assert.NoError(t, err)
	fmt.Println(ret)
}

func TestClient_GetNonce(t *testing.T) {
	c, err := NewClient("02DiuDkfxbUyQCDbAMscAxT996N9Xw8YZHJhXe8Qi2J6g", "https://rosetta.oasis.dev")
	assert.NoError(t, err)

	ret, err := c.GetNonce(context.Background(), "oasis1qpg2xuz46g53737343r20yxeddhlvc2ldqsjh70p")
	assert.NoError(t, err)
	fmt.Println(ret)
}

func TestClient_GetSupportNetwork(t *testing.T) {
	c, err := NewClient("02DiuDkfxbUyQCDbAMscAxT996N9Xw8YZHJhXe8Qi2J6g", "https://rosetta.oasis.dev")
	assert.NoError(t, err)

	ret, err := c.GetSupportNetwork(context.Background())
	assert.NoError(t, err)
	fmt.Println(ret)
}
