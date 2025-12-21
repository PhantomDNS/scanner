package main

import (
	"context"

	"github.com/PhantomDNS/scanner/internal/scanner"
)

func main() {
	s := &scanner.Scanner{}
	_, _ = s.Scan(context.Background())
}
