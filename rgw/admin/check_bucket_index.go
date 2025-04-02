//go:build ceph_preview

package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// CheckBucketIndexRequest contains the input parameters.
type CheckBucketIndexRequest struct {
	Bucket       string `url:"bucket" json:"bucket"`
	CheckObjects bool   `url:"check-objects" json:"check-objects"`
	Fix          bool   `url:"fix" json:"fix"`
}

// CheckBucketIndexResponse contains the response.
type CheckBucketIndexResponse struct {
	Index string `json:"index"`
}

// CheckBucketIndex checks the index of an existing bucket.
// NOTE: to check multipart object accounting with check-objects, fix must be set to True.
func (api *API) CheckBucketIndex(ctx context.Context, input CheckBucketIndexRequest) (CheckBucketIndexResponse, error) {
	if input.Bucket == "" {
		return CheckBucketIndexResponse{}, errMissingBucket
	}
	body, err := api.call(ctx, http.MethodGet, "/bucket?index", valueToURLParams(input, []string{"bucket", "check-objects", "fix"}))
	if err != nil {
		return CheckBucketIndexResponse{}, err
	}

	resp := CheckBucketIndexResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return CheckBucketIndexResponse{}, fmt.Errorf("%s. %s. %w", unmarshalError, string(body), err)
	}

	return resp, nil
}
