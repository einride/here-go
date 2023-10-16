package geocodingsearchv7

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
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

// GeocodingService handles communication with geocoding-related methods of the v7 HERE API.
type GeocodingService service

// ReverseGeocodingService handles communication with reverse geocoding-related methods of the v7 HERE API.
type ReverseGeocodingService service

// BatchGeocodingService handles communication with batch geocoder-related methods of the v7 HERE API.
type BatchGeocodingService service

type Client struct {
	// HTTP client used to communicate with the API.
	client HTTPClient

	UserAgent string

	// Geocoding service
	Geocoding *GeocodingService
	// ReverseGeocoding service
	ReverseGeocoding *ReverseGeocodingService
	// BatchGeocoding service
	BatchGeocoding *BatchGeocodingService
}

type service struct {
	// URL for service API requests
	URL    *url.URL
	Client *Client
}

// A ResponseError reports the error caused by an API request.
type ResponseError struct {
	// HTTP response that caused this error
	Response *HereErrorResponse
	// The HTTP body of the error response
	HTTPBody string
	// The HTTP status code of the response
	HTTPStatusCode int
}

func (r *ResponseError) Error() string {
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
	geocodingURL, _ := url.Parse("https://geocode.search.hereapi.com/v1/")
	c.Geocoding = &GeocodingService{URL: geocodingURL, Client: c}
	reverseGeocodingURL, _ := url.Parse("https://revgeocode.search.hereapi.com/v1/")
	c.ReverseGeocoding = &ReverseGeocodingService{URL: reverseGeocodingURL, Client: c}
	batchGeocoderURL, _ := url.Parse("https://batch.geocoder.ls.hereapi.com/6.2/")
	c.BatchGeocoding = &BatchGeocodingService{URL: batchGeocoderURL, Client: c}
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
				return fmt.Errorf("failed to json decode: %w", err)
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
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, r.Body)
	if err != nil {
		return err
	}
	var response HereErrorResponse
	err = json.Unmarshal(buf.Bytes(), &response)
	if err != nil {
		return err
	}
	return &ResponseError{
		Response:       &response,
		HTTPBody:       buf.String(),
		HTTPStatusCode: r.StatusCode,
	}
}

// DoXML sends an API request and returns the API response. The API response is XML decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) DoXML(req *http.Request, v interface{}) error {
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()
	err = checkResponseXML(resp)
	if err != nil {
		return fmt.Errorf("checkResponse failed: %v, %v", resp.StatusCode, err)
	}
	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				return fmt.Errorf("copy body failed: %v", err)
			}
		} else {
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("read body: %v", err)
			}
			err = xml.Unmarshal(data, v)
			if err != nil {
				return fmt.Errorf("xml unmarshal failed: %w", err)
			}
		}
	}
	return err
}

// checkResponseXML checks the API response for errors, and returns them if present. A response is considered an
// error if it has a status code outside the 200 range. It decodes the error as XML.
func checkResponseXML(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, r.Body)
	if err != nil {
		return err
	}
	response := HereErrorResponse{}
	err = xml.Unmarshal(buf.Bytes(), &response)
	if err != nil {
		return fmt.Errorf("failed unmarshal error: %w", err)
	}
	return &ResponseError{
		Response:       &response,
		HTTPBody:       buf.String(),
		HTTPStatusCode: r.StatusCode,
	}
}
