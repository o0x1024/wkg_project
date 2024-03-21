package ksubdomain

import (
	"testing"
	"wkg-agent/internal/assetCollect/ksubdomain/core"
)

func TestLocalKSubdomain(t *testing.T) {

	options := &core.Options{Domain: []string{"sf-express.com"}}

	core.Start(options)
}
