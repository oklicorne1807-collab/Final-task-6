package service

import (
	"errors"
	"strings"

	"your_project/pkg/morse"
)

func Convert(data string) (string, error) {
	data = strings.TrimSpace(data)
	if data == "" {
		return "", errors.New("empty input")
	}

	// если содержит точки и тире — считаем Морзе
	if strings.ContainsAny(data, ".-") {
		return morse.ToText(data), nil
	}

	return morse.ToMorse(data), nil
}
