package fwdshim

import (
	"github.com/forward-mcp/internal/config"
	fwd "github.com/forward-mcp/internal/forward"
)

type ClientInterface = fwd.ClientInterface
type Config = config.ForwardConfig

func NewClient(cfg *Config) ClientInterface {
	return fwd.NewClient(cfg)
}
