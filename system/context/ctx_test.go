package context

import "testing"

func TestCtxBehaviour(t *testing.T) {
	c := NewCtx(SystemCtx)

	ip4 := c.RealIP()
	if ip4 == "" {
		t.Error("Invalid IP Address (ipv4)")
	}
	t.Logf("IPv4: %s", ip4)

	ua := c.Request().UserAgent()
	if ua == "" {
		t.Error("Invalid UserAgent")
	}
	t.Logf("UserAgent: %s", ua)

	hn := c.Request().HostName()
	if hn == "" {
		t.Error("Invalid HostName")
	}
	t.Logf("Hostname: %s", hn)
}
