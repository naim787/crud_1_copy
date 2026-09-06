package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/atotto/clipboard"
)

func TestTime(t *testing.T) {
	var sb strings.Builder

	for i := 1; i < 1000; i++ {
		sb.WriteString(fmt.Sprintf("%d\n", i))
	}

	if err := clipboard.WriteAll(sb.String()); err != nil {
		t.Fatal(err)
	}

	t.Log("Copied to clipboard!")
}
