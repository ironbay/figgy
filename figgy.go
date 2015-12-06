package figgy

import (
	"os"
	"strings"
)

type Figgy map[string]string

func Get(defaults Figgy, prefix ...string) Figgy {
	joined := strings.Join(prefix, "_")
	for _, kv := range os.Environ() {
		splits := strings.Split(kv, "=")
		key, value := splits[0], splits[1]
		if !strings.HasPrefix(key, joined) {
			continue
		}
		next := strings.Split(key, "_")[len(prefix):]
		defaults[strings.Join(next, "_")] = value
	}
	return defaults
}
