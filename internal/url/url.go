package url

import "strings"

func Join(parts ...string) string {
	var out string = "/"

	for _, part := range parts {
		if !strings.HasSuffix(out, "/") {
			out += "/"
		}

		if strings.HasPrefix(part, "/") {
			part = part[1:]
		}

		out += part
	}

	if strings.HasSuffix(out, "/") {
		out = out[:len(out)-1]
	}

	return out
}
