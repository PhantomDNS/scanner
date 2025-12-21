package scanner

import (
	"context"
	"fmt"

	"github.com/PhantomDNS/scanner/internal/checks"
	"github.com/PhantomDNS/scanner/internal/detection"
)

type Scanner struct {
	Detector detection.ResolverDetection
	Checks   []checks.Check
}

func (s *Scanner) Scan(ctx context.Context) ([]checks.Result, error) {
	fmt.Print("Scanning ...")
	return nil, nil
}
