package files

import (
	"encoding/json"
	"sync"
	"testing"

	"github.com/simon-co/fltr-cli/internal/apperr"
)

func TestEmebeddedPaths(t *testing.T) {
	var pathMap map[string]string
	sb, err := json.Marshal(EmbeddedFsPaths)
	if err != nil {
		t.Fatal(apperr.Parse(err))
	}
	if err := json.Unmarshal(sb, &pathMap); err != nil {
		t.Fatal(apperr.Parse(err))
	}
	wg := sync.WaitGroup{}
	for k, p := range pathMap {
		wg.Add(1)
		go func(key string, path string) {
			defer wg.Done()
			_, err := efs.Open(path)
			if err != nil {
				t.Logf("\nkey: %s\npath: %s\n", key, path)
				t.Errorf(apperr.Parse(err).Error())
			}
		}(k, p)
		wg.Wait()
	}
}
