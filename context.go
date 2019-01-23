package featureflags

import (
	"context"

	"github.com/coreos/go-semver/semver"
)

// key is an unexported type for keys defined in this package. This
// prevents collisions with keys defined in other packages.
type key string

// releaseKey is the key for HTTP Authorization release values in
// context.Context. Clients use release.NewReleaseContext and
// release.ReleaseFromContext instead of using this key directly.
var releaseKey key = "release"

// NewReleaseContext returns a new context.Context that carries value r.
func NewReleaseContext(ctx context.Context, r semver.Version) context.Context {
	return context.WithValue(ctx, releaseKey, r)
}

// ReleaseFromContext returns the HTTP Authorization release value stored in
// ctx, if any.
func ReleaseFromContext(ctx context.Context) (semver.Version, bool) {
	r, ok := ctx.Value(releaseKey).(semver.Version)
	return r, ok
}
