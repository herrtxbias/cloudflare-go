// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// PcapDownloadService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPcapDownloadService] method
// instead.
type PcapDownloadService struct {
	Options []option.RequestOption
}

// NewPcapDownloadService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPcapDownloadService(opts ...option.RequestOption) (r *PcapDownloadService) {
	r = &PcapDownloadService{}
	r.Options = opts
	return
}

// Download PCAP information into a file. Response is a binary PCAP file.
func (r *PcapDownloadService) Get(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.tcpdump.pcap")}, opts...)
	path := fmt.Sprintf("accounts/%s/pcaps/%s/download", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
