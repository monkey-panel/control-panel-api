package utils

import (
	"testing"

	"github.com/monkey-panel/control-panel-utils/utils"
)

func TestBcryptHash(t *testing.T) {
	for i := 0; i < 72; i++ {
		tc := utils.RandomString(i)
		if hash := BcryptHash(tc); !BcryptCheck(tc, hash) {
			t.Errorf("BcryptHash(%s) error", tc)
		}
	}
}
