package utils

import (
	"bytes"
	"encoding/json"
)

func FormatJSON(data []byte) string {
	var output bytes.Buffer

	err := json.Indent(&output, data, "", "\t")
	if err != nil {
		return ""
	}
	return output.String()
}
