package utils

import (
	"fmt"
	"testing"
)

const (
	key = "amlhbWlzdXpoYW53ZWlxdw=="
)

func TestAes(t *testing.T) {
	ret := Aes128CBCEncrypt("123456789", key)
	fmt.Println("ret", ret)
}

func TestDeAes(t *testing.T) {
	ret := Aes128CBCDecrypt("GfbiLH8J9m/BRUbTzltPmUcs36cw7fVF4Vfm+jMcRhEVE9HV8syvzCVi61URua5N", key, false)
	fmt.Println("ret", ret)
}
func TestBas(t *testing.T) {
	Bas("jiamisuzhanweiqw")
}
