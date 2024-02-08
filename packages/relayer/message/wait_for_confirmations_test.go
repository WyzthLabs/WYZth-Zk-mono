package message

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wyzth_zkevmxyz/wyzth_zkevm-mono/packages/relayer/mock"
)

func Test_waitForConfirmations(t *testing.T) {
	p := newTestProcessor(true)

	err := p.waitForConfirmations(context.TODO(), mock.SucceedTxHash, uint64(mock.BlockNum))
	assert.Nil(t, err)
}
