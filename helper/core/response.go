package coffeego

import (
	"fmt"
	"strings"
)

// Response Should be Exported
func Response(data []string) []byte {
	ResponseString := "{"
	for i := 0; i < len(data); i++ {
		key, value := splitData(data[i])
		splitstr := strings.Split(value, "|")
		switch splitstr[0] {
		case "number":
			ResponseString += "\"" + key + "\":" + splitstr[1]

		case "string":
			ResponseString += "\"" + key + "\":\"" + splitstr[1] + "\""

		}
		if i < len(data)-1 {
			ResponseString += ","
		}
	}
	ResponseString += "}"
	return []byte(fmt.Sprintf(ResponseString))
}

func splitData(data string) (string, string) {
	splitstr := strings.Split(data, ":")
	return splitstr[0], splitstr[1]
}
