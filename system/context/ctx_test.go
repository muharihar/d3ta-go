package context

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCtxBehaviour(t *testing.T) {
	c := NewCtx(SystemCtx)

	ip4 := c.RealIP()
	assert.NotEmpty(t, ip4, "Invalid IP Address (ipv4)")
	t.Logf("IPv4: %s", ip4)

	ua := c.Request().UserAgent()
	assert.NotEmpty(t, ua, "Invalid UserAgent")
	t.Logf("UserAgent: %s", ua)

	hn := c.Request().HostName()
	assert.NotEmpty(t, hn, "Invalid HostName")
	t.Logf("Hostname: %s", hn)
}
