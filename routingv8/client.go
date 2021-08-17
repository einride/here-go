package routingv8

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	userAgent = "einride/here-go"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// MatrixService handles communication with the matrix-related methods of the v7 HERE API.
type MatrixService service

type Client struct {
	// HTTP client used to communicate with the API.
	client HTTPClient

	UserAgent string

	// Matrix service.
	Matrix *MatrixService
}

type service struct {
	// URL for service API requests
	URL    *url.URL
	Client *Client
}

// An errorResponse reports the error caused by an API request.
type errorResponse struct {
	// HTTP response that caused this error
	Response *HereErrorResponse
}

func (r *errorResponse) Error() string {
	return fmt.Sprintf(
		"Title: %v, Status: %d, Code: %v, Cause: %v, Action: %v",
		r.Response.Title,
		r.Response.Status,
		r.Response.Code,
		r.Response.Cause,
		r.Response.Action,
	)
}

// NewClient returns a new HERE API Client. If a nil httpClient is
// provided, a new http.Client will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient HTTPClient) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	c := &Client{client: httpClient, UserAgent: userAgent}
	matrixURL, _ := url.Parse("https://matrix.router.hereapi.com/v8/")
	c.Matrix = &MatrixService{URL: matrixURL, Client: c}
	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr, which will be resolved to the
// BaseURL of the Client. Relative URLS should always be specified without a preceding slash. If specified, the
// value pointed to by body is JSON encoded and included in as the request body.
// A raw query string can be specified by rawQuery.
func (c *Client) NewRequest(
	ctx context.Context,
	u *url.URL,
	method string,
	rawQuery string,
	body []byte,
) (*http.Request, error) {
	if len(rawQuery) > 0 {
		u.RawQuery = rawQuery
	}
	var r io.Reader
	if len(body) > 0 {
		r = bytes.NewReader(body)
	}
	req, err := http.NewRequestWithContext(ctx, method, u.String(), r)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Do sends an API request and returns the API response. The API response is JSON decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, v interface{}) error {
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()
	err = checkResponse(resp)
	if err != nil {
		return err
	}
	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				return err
			}
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return err
			}
		}
	}
	return err
}

// checkResponse checks the API response for errors, and returns them if present. A response is considered an
// error if it has a status code outside the 200 range.
func checkResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}
	var response HereErrorResponse
	err := json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		return err
	}
	errorResponse := &errorResponse{Response: &response}
	return errorResponse
}
