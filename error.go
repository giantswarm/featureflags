package featureflags

import (
	"github.com/giantswarm/microerror"
)

var versionComparisonError = &microerror.Error{
	Kind: "versionComparisonError",
}

// IsVersionComparison asserts versionComparisonError.
func IsVersionComparison(err error) bool {
	return microerror.Cause(err) == versionComparisonError
}
