package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/golf1-sprint6-final/pkg/morse"
)

func Convert(data string) (string, error) {
	data = strings.TrimSpace(data)
	if data == "" {
		return "", errors.New("empty input")
	}

	if strings.ContainsAny(data, ".-") {
		return morse.ToText(data), nil
	}

	return morse.ToMorse(data), nil
}
