package checks

import (
	"context"
	"net"
)

type Check interface {
	Name() string
	Run(ctx context.Context, resolver net.IP) (Result, error)
}

type Result struct {
	Status   string // e.g., "pass", "fail"
	Evidence map[string]interface{}
}
