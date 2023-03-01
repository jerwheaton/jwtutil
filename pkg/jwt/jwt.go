package jwt

import (
	"encoding/base64"
	"fmt"
	"strings"
)

const outputFmt = `{"alg": %s, "payload": %s}`

func Decode(rawInput []byte) (string, error) {
	tokenStr := string(TrimEmptyBytes(rawInput))

	segments := strings.Split(tokenStr, ".")
	if len(segments) != 3 {
		return "", fmt.Errorf("invalid number of segments: %d", len(segments))
	}

	header, err := base64.RawStdEncoding.DecodeString(segments[0])
	if err != nil {
		return "", fmt.Errorf("error decoding header: %w", err)
	}
	payload, err := base64.RawStdEncoding.DecodeString(segments[1])
	if err != nil {
		return "", fmt.Errorf("error decoding payload: %w", err)
	}
	return fmt.Sprintf(outputFmt, string(header), string(payload)), nil
}

func TrimEmptyBytes(b []byte) []byte {
	trimmed := make([]byte, 0, len(b))
	for i := range b {
		if b[i] == byte(0) {
			break
		}

		trimmed = append(trimmed, b[i])
	}
	return trimmed
}
