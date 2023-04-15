package utils

import (
	"fmt"
	"testing"
)

func TestGetGetUUID(t *testing.T) {
	uuid := NewUUID()
	fmt.Println(uuid)
	fmt.Println(len(uuid))
}
