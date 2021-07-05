package array

import "testing"

func TestStringEncode(t *testing.T) {
	s := "中国"

	t.Log(len(s))
	t.Log(s)

	c := []rune(s)
	t.Log(len(c))
	t.Log(c)
	t.Logf("utf8 %x", s)
}

func TestStrings(t *testing.T) {

}

func TestStringConv(t *testing.T) {

}
