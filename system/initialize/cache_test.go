package initialize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenAllCacheConnection(t *testing.T) {
	h, err := newHandler(t)
	if assert.NoError(t, err, "Error while creating handler: newHandler") {
		if !assert.NotNil(t, h) {
			return
		}
	}

	if assert.NoError(t, OpenAllCacheConnection(h), "Error while opening all cache connection: OpenAllCacheConnection") {

		cfg, err := h.GetConfig()
		if !assert.NoError(t, err, "Error while getting config: h.GetConfig") {
			return
		}

		ceCon, err := h.GetCacher(cfg.Caches.SessionCache.ConnectionName)
		if assert.NoError(t, err, "Error while getting Cacher Engine Connection: h.GetCacher(cfg.Caches.sessionCache.ConnectionName)") {

			if assert.NoError(t, ceCon.Put("test-cache-key", "test-cache-value", 100)) {
				assert.NotNil(t, ceCon.Get("test-cache-key"))
			}
		}
	}
}
