// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippets

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SnippetService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSnippetService] method instead.
type SnippetService struct {
	Options []option.RequestOption
	Content *ContentService
	Rules   *RuleService
}

// NewSnippetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSnippetService(opts ...option.RequestOption) (r *SnippetService) {
	r = &SnippetService{}
	r.Options = opts
	r.Content = NewContentService(opts...)
	r.Rules = NewRuleService(opts...)
	return
}
