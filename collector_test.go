// +build go1.11,!go1.15

package sqlstats

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

func TestMultipleDB(t *testing.T) {
	db1, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	db2, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)

	label1 := map[string]string{"name": "db1"}
	label2 := map[string]string{"name": "db2"}

	c1 := NewStatsCollector(label1, db1)
	c2 := NewStatsCollector(label2, db2)

	assert.NoError(t, prometheus.Register(c1))
	assert.NoError(t, prometheus.Register(c2))
}
