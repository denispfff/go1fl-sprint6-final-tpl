package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(text string) bool {
	for _, char := range text {
		if !strings.ContainsRune("-. ", char) {
			return false
		}
	}
	return true
}

func ReverseMorse(text string) string {
	if isMorse(text) {
		return morse.ToText(text)
	}
	return morse.ToMorse(text)
}
