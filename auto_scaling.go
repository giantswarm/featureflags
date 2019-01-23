package featureflags

import (
	"github.com/coreos/go-semver/semver"
)

//
//     if featureflag.IsAutoScalingEnabled(featureflag.ReleaseFromContext(ctx)) {
//       ...
//     }
//

var AutoScalingVersion = semver.Version{
	Major: 6,
	Minor: 3,
	Patch: 6,
}

// IsAutoScalingEnabled compares given versions against
//
//     * Auto Scaling is NOT enabled for lower versions.
//     * Auto Scaling is enabled for equal versions.
//     * Auto Scaling is enabled for higher versions.
//
// TODO the provider is important as well. Maybe even more.
//
func IsAutoScalingEnabled(v semver.Version) bool {
	switch AutoScalingVersion.Compare(v) {
	case -1:
		return false
	case 0:
		return true
	case 1:
		return true
	default:
		panic(versionComparisonError)
	}
}
