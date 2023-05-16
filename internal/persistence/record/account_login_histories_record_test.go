package record

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountLoginHistoriesRecord_TableName(t *testing.T) {
	t.Run("Test Account History Record TableName", func(t *testing.T) {
		var accountHistory AccountLoginHistoriesRecord
		assert.NotNil(t, accountHistory)
		assert.Equal(t, "account_login_histories", accountHistory.TableName())
	})
}
