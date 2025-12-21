package detection

import (
	"context"
	"net"
)

type ResolverDetection interface {
	Detect(ctx context.Context) (net.IP, error)
}
