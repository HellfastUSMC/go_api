package store

import (
	"fmt"
	"strings"
	"testing"
)

// TestStore ...
func TestStore(t *testing.T, DBURL string) (*Store, func(...string)) {
	t.Helper()

	config := NewConf()
	config.DBURL = DBURL
	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}
		s.Close()
	}
}
