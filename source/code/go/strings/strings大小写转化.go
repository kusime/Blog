package main

import (
	"fmt"
	"strings"
)
func main() {
	字符串 := "Ming Cloud"
	a1:=strings.ToLower(字符串)
	a2:=strings.ToUpper(字符串)
	fmt.Printf("%s\n%s\n",a1,a2)
}
// ─── INPUT ──────────────────────────────────────────────────────────────────────

/* 
ming cloud
MING CLOUD
*/