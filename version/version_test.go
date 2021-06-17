package version

import (
	"testing"

	"github.com/gocolly/colly"
	collyV2 "github.com/gocolly/colly/v2"
)

func testColly(t *testing.T) {
	colly.AllowURLRevisit()
	t.Log("hello version")
}

func testCollyV2(t *testing.T) {
	collyV2.AllowURLRevisit()
	t.Log("hello version")
}
