// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

import (
	"github.com/cloudflare/cloudflare-go/v3/option"
)

// BucketEventNotificationService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketEventNotificationService] method instead.
type BucketEventNotificationService struct {
	Options       []option.RequestOption
	Configuration *BucketEventNotificationConfigurationService
}

// NewBucketEventNotificationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBucketEventNotificationService(opts ...option.RequestOption) (r *BucketEventNotificationService) {
	r = &BucketEventNotificationService{}
	r.Options = opts
	r.Configuration = NewBucketEventNotificationConfigurationService(opts...)
	return
}
