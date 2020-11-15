// Code generated by protoc-gen-goclay. DO NOT EDIT.
// source: accounts.proto

/*
Package accounts is a self-registering gRPC and JSON+Swagger service definition.

It conforms to the github.com/utrack/clay/v2/transport Service interface.
*/
package accounts

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-openapi/spec"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"github.com/pkg/errors"
	"github.com/utrack/clay/v2/transport"
	"github.com/utrack/clay/v2/transport/httpclient"
	"github.com/utrack/clay/v2/transport/httpruntime"
	"github.com/utrack/clay/v2/transport/httpruntime/httpmw"
	"github.com/utrack/clay/v2/transport/httptransport"
	"github.com/utrack/clay/v2/transport/swagger"
	"google.golang.org/grpc"
)

// Update your shared lib or downgrade generator to v1 if there's an error
var _ = transport.IsVersion2

var _ = ioutil.Discard
var _ chi.Router
var _ runtime.Marshaler
var _ bytes.Buffer
var _ context.Context
var _ fmt.Formatter
var _ strings.Reader
var _ errors.Frame
var _ httpruntime.Marshaler
var _ http.Handler
var _ url.Values
var _ base64.Encoding
var _ httptransport.MarshalerError
var _ utilities.DoubleArray

// AccountsDesc is a descriptor/registrator for the AccountsServer.
type AccountsDesc struct {
	svc  AccountsServer
	opts httptransport.DescOptions
}

// NewAccountsServiceDesc creates new registrator for the AccountsServer.
// It implements httptransport.ConfigurableServiceDesc as well.
func NewAccountsServiceDesc(svc AccountsServer) *AccountsDesc {
	return &AccountsDesc{
		svc: svc,
	}
}

// RegisterGRPC implements service registrator interface.
func (d *AccountsDesc) RegisterGRPC(s *grpc.Server) {
	RegisterAccountsServer(s, d.svc)
}

// Apply applies passed options.
func (d *AccountsDesc) Apply(oo ...transport.DescOption) {
	for _, o := range oo {
		o.Apply(&d.opts)
	}
}

// SwaggerDef returns this file's Swagger definition.
func (d *AccountsDesc) SwaggerDef(options ...swagger.Option) (result []byte) {
	if len(options) > 0 || len(d.opts.SwaggerDefaultOpts) > 0 {
		var err error
		var s = &spec.Swagger{}
		if err = s.UnmarshalJSON(_swaggerDef_accounts_proto); err != nil {
			panic("Bad swagger definition: " + err.Error())
		}

		for _, o := range d.opts.SwaggerDefaultOpts {
			o(s)
		}
		for _, o := range options {
			o(s)
		}
		if result, err = s.MarshalJSON(); err != nil {
			panic("Failed marshal spec.Swagger definition: " + err.Error())
		}
	} else {
		result = _swaggerDef_accounts_proto
	}
	return result
}

// RegisterHTTP registers this service's HTTP handlers/bindings.
func (d *AccountsDesc) RegisterHTTP(mux transport.Router) {
	chiMux, isChi := mux.(chi.Router)

	{
		// Handler for List, binding: GET /v1/accounts
		var h http.HandlerFunc
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()

			unmFunc := unmarshaler_goclay_Accounts_List_0(r)
			rsp, err := _Accounts_List_Handler(d.svc, r.Context(), unmFunc, d.opts.UnaryInterceptor)

			if err != nil {
				if err, ok := err.(httptransport.MarshalerError); ok {
					httpruntime.SetError(r.Context(), r, w, errors.Wrap(err.Err, "couldn't parse request"))
					return
				}
				httpruntime.SetError(r.Context(), r, w, err)
				return
			}

			if ctxErr := r.Context().Err(); ctxErr != nil && ctxErr == context.Canceled {
				w.WriteHeader(499) // Client Closed Request
				return
			}

			_, outbound := httpruntime.MarshalerForRequest(r)
			w.Header().Set("Content-Type", outbound.ContentType())
			err = outbound.Marshal(w, rsp)
			if err != nil {
				httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "couldn't write response"))
				return
			}
		})

		h = httpmw.DefaultChain(h)

		if isChi {
			chiMux.Method("GET", pattern_goclay_Accounts_List_0, h)
		} else {
			mux.Handle(pattern_goclay_Accounts_List_0, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "GET" {
					w.WriteHeader(http.StatusMethodNotAllowed)
					return
				}
				h(w, r)
			}))
		}
	}

	{
		// Handler for Get, binding: GET /v1/accounts/{ID}
		var h http.HandlerFunc
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()

			unmFunc := unmarshaler_goclay_Accounts_Get_0(r)
			rsp, err := _Accounts_Get_Handler(d.svc, r.Context(), unmFunc, d.opts.UnaryInterceptor)

			if err != nil {
				if err, ok := err.(httptransport.MarshalerError); ok {
					httpruntime.SetError(r.Context(), r, w, errors.Wrap(err.Err, "couldn't parse request"))
					return
				}
				httpruntime.SetError(r.Context(), r, w, err)
				return
			}

			if ctxErr := r.Context().Err(); ctxErr != nil && ctxErr == context.Canceled {
				w.WriteHeader(499) // Client Closed Request
				return
			}

			_, outbound := httpruntime.MarshalerForRequest(r)
			w.Header().Set("Content-Type", outbound.ContentType())
			err = outbound.Marshal(w, rsp)
			if err != nil {
				httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "couldn't write response"))
				return
			}
		})

		h = httpmw.DefaultChain(h)

		if isChi {
			chiMux.Method("GET", pattern_goclay_Accounts_Get_0, h)
		} else {
			panic("query URI params supported only for chi.Router")
		}
	}

	{
		// Handler for Create, binding: POST /v1/accounts
		var h http.HandlerFunc
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()

			unmFunc := unmarshaler_goclay_Accounts_Create_0(r)
			rsp, err := _Accounts_Create_Handler(d.svc, r.Context(), unmFunc, d.opts.UnaryInterceptor)

			if err != nil {
				if err, ok := err.(httptransport.MarshalerError); ok {
					httpruntime.SetError(r.Context(), r, w, errors.Wrap(err.Err, "couldn't parse request"))
					return
				}
				httpruntime.SetError(r.Context(), r, w, err)
				return
			}

			if ctxErr := r.Context().Err(); ctxErr != nil && ctxErr == context.Canceled {
				w.WriteHeader(499) // Client Closed Request
				return
			}

			_, outbound := httpruntime.MarshalerForRequest(r)
			w.Header().Set("Content-Type", outbound.ContentType())
			err = outbound.Marshal(w, rsp)
			if err != nil {
				httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "couldn't write response"))
				return
			}
		})

		h = httpmw.DefaultChain(h)

		if isChi {
			chiMux.Method("POST", pattern_goclay_Accounts_Create_0, h)
		} else {
			mux.Handle(pattern_goclay_Accounts_Create_0, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "POST" {
					w.WriteHeader(http.StatusMethodNotAllowed)
					return
				}
				h(w, r)
			}))
		}
	}

	{
		// Handler for Update, binding: PUT /v1/accounts/{ID}
		var h http.HandlerFunc
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()

			unmFunc := unmarshaler_goclay_Accounts_Update_0(r)
			rsp, err := _Accounts_Update_Handler(d.svc, r.Context(), unmFunc, d.opts.UnaryInterceptor)

			if err != nil {
				if err, ok := err.(httptransport.MarshalerError); ok {
					httpruntime.SetError(r.Context(), r, w, errors.Wrap(err.Err, "couldn't parse request"))
					return
				}
				httpruntime.SetError(r.Context(), r, w, err)
				return
			}

			if ctxErr := r.Context().Err(); ctxErr != nil && ctxErr == context.Canceled {
				w.WriteHeader(499) // Client Closed Request
				return
			}

			_, outbound := httpruntime.MarshalerForRequest(r)
			w.Header().Set("Content-Type", outbound.ContentType())
			err = outbound.Marshal(w, rsp)
			if err != nil {
				httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "couldn't write response"))
				return
			}
		})

		h = httpmw.DefaultChain(h)

		if isChi {
			chiMux.Method("PUT", pattern_goclay_Accounts_Update_0, h)
		} else {
			panic("query URI params supported only for chi.Router")
		}
	}

	{
		// Handler for Delete, binding: DELETE /v1/accounts/{ID}
		var h http.HandlerFunc
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()

			unmFunc := unmarshaler_goclay_Accounts_Delete_0(r)
			rsp, err := _Accounts_Delete_Handler(d.svc, r.Context(), unmFunc, d.opts.UnaryInterceptor)

			if err != nil {
				if err, ok := err.(httptransport.MarshalerError); ok {
					httpruntime.SetError(r.Context(), r, w, errors.Wrap(err.Err, "couldn't parse request"))
					return
				}
				httpruntime.SetError(r.Context(), r, w, err)
				return
			}

			if ctxErr := r.Context().Err(); ctxErr != nil && ctxErr == context.Canceled {
				w.WriteHeader(499) // Client Closed Request
				return
			}

			_, outbound := httpruntime.MarshalerForRequest(r)
			w.Header().Set("Content-Type", outbound.ContentType())
			err = outbound.Marshal(w, rsp)
			if err != nil {
				httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "couldn't write response"))
				return
			}
		})

		h = httpmw.DefaultChain(h)

		if isChi {
			chiMux.Method("DELETE", pattern_goclay_Accounts_Delete_0, h)
		} else {
			panic("query URI params supported only for chi.Router")
		}
	}

}

type Accounts_httpClient struct {
	c    *http.Client
	host string
}

// NewAccountsHTTPClient creates new HTTP client for AccountsServer.
// Pass addr in format "http://host[:port]".
func NewAccountsHTTPClient(c *http.Client, addr string) *Accounts_httpClient {
	if strings.HasSuffix(addr, "/") {
		addr = addr[:len(addr)-1]
	}
	return &Accounts_httpClient{c: c, host: addr}
}

func (c *Accounts_httpClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	mw, err := httpclient.NewMiddlewareGRPC(opts)
	if err != nil {
		return nil, err
	}

	path := pattern_goclay_Accounts_List_0_builder(in)

	buf := bytes.NewBuffer(nil)

	m := httpruntime.DefaultMarshaler(nil)

	req, err := http.NewRequest("GET", c.host+path, buf)
	if err != nil {
		return nil, errors.Wrap(err, "can't initiate HTTP request")
	}
	req = req.WithContext(ctx)

	req.Header.Add("Accept", m.ContentType())

	req, err = mw.ProcessRequest(req)
	if err != nil {
		return nil, err
	}
	rsp, err := c.c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error from client")
	}
	defer rsp.Body.Close()

	rsp, err = mw.ProcessResponse(rsp)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(rsp.Body)
		return nil, errors.Errorf("%v %v: server returned HTTP %v: '%v'", req.Method, req.URL.String(), rsp.StatusCode, string(b))
	}

	ret := ListResponse{}

	err = m.Unmarshal(rsp.Body, &ret)

	return &ret, errors.Wrap(err, "can't unmarshal response")
}

func (c *Accounts_httpClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	mw, err := httpclient.NewMiddlewareGRPC(opts)
	if err != nil {
		return nil, err
	}

	path := pattern_goclay_Accounts_Get_0_builder(in)

	buf := bytes.NewBuffer(nil)

	m := httpruntime.DefaultMarshaler(nil)

	req, err := http.NewRequest("GET", c.host+path, buf)
	if err != nil {
		return nil, errors.Wrap(err, "can't initiate HTTP request")
	}
	req = req.WithContext(ctx)

	req.Header.Add("Accept", m.ContentType())

	req, err = mw.ProcessRequest(req)
	if err != nil {
		return nil, err
	}
	rsp, err := c.c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error from client")
	}
	defer rsp.Body.Close()

	rsp, err = mw.ProcessResponse(rsp)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(rsp.Body)
		return nil, errors.Errorf("%v %v: server returned HTTP %v: '%v'", req.Method, req.URL.String(), rsp.StatusCode, string(b))
	}

	ret := GetResponse{}

	err = m.Unmarshal(rsp.Body, &ret)

	return &ret, errors.Wrap(err, "can't unmarshal response")
}

func (c *Accounts_httpClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	mw, err := httpclient.NewMiddlewareGRPC(opts)
	if err != nil {
		return nil, err
	}

	path := pattern_goclay_Accounts_Create_0_builder(in)

	buf := bytes.NewBuffer(nil)

	m := httpruntime.DefaultMarshaler(nil)

	if err = m.Marshal(buf, in); err != nil {
		return nil, errors.Wrap(err, "can't marshal request")
	}

	req, err := http.NewRequest("POST", c.host+path, buf)
	if err != nil {
		return nil, errors.Wrap(err, "can't initiate HTTP request")
	}
	req = req.WithContext(ctx)

	req.Header.Add("Accept", m.ContentType())

	req, err = mw.ProcessRequest(req)
	if err != nil {
		return nil, err
	}
	rsp, err := c.c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error from client")
	}
	defer rsp.Body.Close()

	rsp, err = mw.ProcessResponse(rsp)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(rsp.Body)
		return nil, errors.Errorf("%v %v: server returned HTTP %v: '%v'", req.Method, req.URL.String(), rsp.StatusCode, string(b))
	}

	ret := CreateResponse{}

	err = m.Unmarshal(rsp.Body, &ret)

	return &ret, errors.Wrap(err, "can't unmarshal response")
}

func (c *Accounts_httpClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	mw, err := httpclient.NewMiddlewareGRPC(opts)
	if err != nil {
		return nil, err
	}

	path := pattern_goclay_Accounts_Update_0_builder(in)

	buf := bytes.NewBuffer(nil)

	m := httpruntime.DefaultMarshaler(nil)

	if err = m.Marshal(buf, in); err != nil {
		return nil, errors.Wrap(err, "can't marshal request")
	}

	req, err := http.NewRequest("PUT", c.host+path, buf)
	if err != nil {
		return nil, errors.Wrap(err, "can't initiate HTTP request")
	}
	req = req.WithContext(ctx)

	req.Header.Add("Accept", m.ContentType())

	req, err = mw.ProcessRequest(req)
	if err != nil {
		return nil, err
	}
	rsp, err := c.c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error from client")
	}
	defer rsp.Body.Close()

	rsp, err = mw.ProcessResponse(rsp)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(rsp.Body)
		return nil, errors.Errorf("%v %v: server returned HTTP %v: '%v'", req.Method, req.URL.String(), rsp.StatusCode, string(b))
	}

	ret := UpdateResponse{}

	err = m.Unmarshal(rsp.Body, &ret)

	return &ret, errors.Wrap(err, "can't unmarshal response")
}

func (c *Accounts_httpClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	mw, err := httpclient.NewMiddlewareGRPC(opts)
	if err != nil {
		return nil, err
	}

	path := pattern_goclay_Accounts_Delete_0_builder(in)

	buf := bytes.NewBuffer(nil)

	m := httpruntime.DefaultMarshaler(nil)

	req, err := http.NewRequest("DELETE", c.host+path, buf)
	if err != nil {
		return nil, errors.Wrap(err, "can't initiate HTTP request")
	}
	req = req.WithContext(ctx)

	req.Header.Add("Accept", m.ContentType())

	req, err = mw.ProcessRequest(req)
	if err != nil {
		return nil, err
	}
	rsp, err := c.c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error from client")
	}
	defer rsp.Body.Close()

	rsp, err = mw.ProcessResponse(rsp)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(rsp.Body)
		return nil, errors.Errorf("%v %v: server returned HTTP %v: '%v'", req.Method, req.URL.String(), rsp.StatusCode, string(b))
	}

	ret := DeleteResponse{}

	err = m.Unmarshal(rsp.Body, &ret)

	return &ret, errors.Wrap(err, "can't unmarshal response")
}

// patterns for Accounts
var (
	pattern_goclay_Accounts_List_0 = "/v1/accounts"

	pattern_goclay_Accounts_List_0_builder = func(in *ListRequest) string {
		values := url.Values{}

		u := url.URL{
			Path:     fmt.Sprintf("/v1/accounts"),
			RawQuery: values.Encode(),
		}
		return u.String()
	}

	unmarshaler_goclay_Accounts_List_0_boundParams = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}

	pattern_goclay_Accounts_Get_0 = "/v1/accounts/{ID}"

	pattern_goclay_Accounts_Get_0_builder = func(in *GetRequest) string {
		values := url.Values{}

		u := url.URL{
			Path:     fmt.Sprintf("/v1/accounts/%v", in.ID),
			RawQuery: values.Encode(),
		}
		return u.String()
	}

	unmarshaler_goclay_Accounts_Get_0_boundParams = &utilities.DoubleArray{Encoding: map[string]int{"ID": 0}, Base: []int{1, 1, 0}, Check: []int{0, 1, 2}}

	pattern_goclay_Accounts_Create_0 = "/v1/accounts"

	pattern_goclay_Accounts_Create_0_builder = func(in *CreateRequest) string {
		values := url.Values{}

		u := url.URL{
			Path:     fmt.Sprintf("/v1/accounts"),
			RawQuery: values.Encode(),
		}
		return u.String()
	}

	unmarshaler_goclay_Accounts_Create_0_boundParams = &utilities.DoubleArray{Encoding: map[string]int{"": 0}, Base: []int{1, 1, 0}, Check: []int{0, 1, 2}}

	pattern_goclay_Accounts_Update_0 = "/v1/accounts/{ID}"

	pattern_goclay_Accounts_Update_0_builder = func(in *UpdateRequest) string {
		values := url.Values{}

		u := url.URL{
			Path:     fmt.Sprintf("/v1/accounts/%v", in.ID),
			RawQuery: values.Encode(),
		}
		return u.String()
	}

	unmarshaler_goclay_Accounts_Update_0_boundParams = &utilities.DoubleArray{Encoding: map[string]int{"": 0, "ID": 1}, Base: []int{1, 1, 2, 0, 0}, Check: []int{0, 1, 1, 2, 3}}

	pattern_goclay_Accounts_Delete_0 = "/v1/accounts/{ID}"

	pattern_goclay_Accounts_Delete_0_builder = func(in *DeleteRequest) string {
		values := url.Values{}

		u := url.URL{
			Path:     fmt.Sprintf("/v1/accounts/%v", in.ID),
			RawQuery: values.Encode(),
		}
		return u.String()
	}

	unmarshaler_goclay_Accounts_Delete_0_boundParams = &utilities.DoubleArray{Encoding: map[string]int{"ID": 0}, Base: []int{1, 1, 0}, Check: []int{0, 1, 2}}
)

// marshalers for Accounts
var (
	unmarshaler_goclay_Accounts_List_0 = func(r *http.Request) func(interface{}) error {
		return func(rif interface{}) error {
			req := rif.(*ListRequest)

			if err := errors.Wrap(runtime.PopulateQueryParameters(req, r.URL.Query(), unmarshaler_goclay_Accounts_List_0_boundParams), "couldn't populate query parameters"); err != nil {
				return httpruntime.TransformUnmarshalerError(err)
			}

			return nil
		}
	}

	unmarshaler_goclay_Accounts_Get_0 = func(r *http.Request) func(interface{}) error {
		return func(rif interface{}) error {
			req := rif.(*GetRequest)

			if err := errors.Wrap(runtime.PopulateQueryParameters(req, r.URL.Query(), unmarshaler_goclay_Accounts_Get_0_boundParams), "couldn't populate query parameters"); err != nil {
				return httpruntime.TransformUnmarshalerError(err)
			}

			rctx := chi.RouteContext(r.Context())
			if rctx == nil {
				panic("Only chi router is supported for GETs atm")
			}
			for pos, k := range rctx.URLParams.Keys {
				if err := errors.Wrapf(runtime.PopulateFieldFromPath(req, k, rctx.URLParams.Values[pos]), "can't read '%v' from path", k); err != nil {
					return httptransport.NewMarshalerError(httpruntime.TransformUnmarshalerError(err))
				}
			}

			return nil
		}
	}

	unmarshaler_goclay_Accounts_Create_0 = func(r *http.Request) func(interface{}) error {
		return func(rif interface{}) error {
			req := rif.(*CreateRequest)

			if err := errors.Wrap(runtime.PopulateQueryParameters(req, r.URL.Query(), unmarshaler_goclay_Accounts_Create_0_boundParams), "couldn't populate query parameters"); err != nil {
				return httpruntime.TransformUnmarshalerError(err)
			}

			inbound, _ := httpruntime.MarshalerForRequest(r)
			if err := errors.Wrap(inbound.Unmarshal(r.Body, &req), "couldn't read request JSON"); err != nil {
				return httptransport.NewMarshalerError(httpruntime.TransformUnmarshalerError(err))
			}
			return nil
		}
	}

	unmarshaler_goclay_Accounts_Update_0 = func(r *http.Request) func(interface{}) error {
		return func(rif interface{}) error {
			req := rif.(*UpdateRequest)

			if err := errors.Wrap(runtime.PopulateQueryParameters(req, r.URL.Query(), unmarshaler_goclay_Accounts_Update_0_boundParams), "couldn't populate query parameters"); err != nil {
				return httpruntime.TransformUnmarshalerError(err)
			}

			inbound, _ := httpruntime.MarshalerForRequest(r)
			if err := errors.Wrap(inbound.Unmarshal(r.Body, &req), "couldn't read request JSON"); err != nil {
				return httptransport.NewMarshalerError(httpruntime.TransformUnmarshalerError(err))
			}
			rctx := chi.RouteContext(r.Context())
			if rctx == nil {
				panic("Only chi router is supported for GETs atm")
			}
			for pos, k := range rctx.URLParams.Keys {
				if err := errors.Wrapf(runtime.PopulateFieldFromPath(req, k, rctx.URLParams.Values[pos]), "can't read '%v' from path", k); err != nil {
					return httptransport.NewMarshalerError(httpruntime.TransformUnmarshalerError(err))
				}
			}

			return nil
		}
	}

	unmarshaler_goclay_Accounts_Delete_0 = func(r *http.Request) func(interface{}) error {
		return func(rif interface{}) error {
			req := rif.(*DeleteRequest)

			if err := errors.Wrap(runtime.PopulateQueryParameters(req, r.URL.Query(), unmarshaler_goclay_Accounts_Delete_0_boundParams), "couldn't populate query parameters"); err != nil {
				return httpruntime.TransformUnmarshalerError(err)
			}

			rctx := chi.RouteContext(r.Context())
			if rctx == nil {
				panic("Only chi router is supported for GETs atm")
			}
			for pos, k := range rctx.URLParams.Keys {
				if err := errors.Wrapf(runtime.PopulateFieldFromPath(req, k, rctx.URLParams.Values[pos]), "can't read '%v' from path", k); err != nil {
					return httptransport.NewMarshalerError(httpruntime.TransformUnmarshalerError(err))
				}
			}

			return nil
		}
	}
)

var _swaggerDef_accounts_proto = []byte(`{
  "swagger": "2.0",
  "info": {
    "title": "accounts.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/accounts": {
      "get": {
        "operationId": "Accounts_List",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/accountsListResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "isAvailable",
            "in": "query",
            "required": false,
            "type": "boolean",
            "format": "boolean"
          }
        ],
        "tags": [
          "Accounts"
        ]
      },
      "post": {
        "operationId": "Accounts_Create",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/accountsCreateResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/accountsCreateRequest"
            }
          }
        ],
        "tags": [
          "Accounts"
        ]
      }
    },
    "/v1/accounts/{ID}": {
      "get": {
        "operationId": "Accounts_Get",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/accountsGetResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "ID",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "int64"
          }
        ],
        "tags": [
          "Accounts"
        ]
      },
      "delete": {
        "operationId": "Accounts_Delete",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/accountsDeleteResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "ID",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "int64"
          }
        ],
        "tags": [
          "Accounts"
        ]
      },
      "put": {
        "operationId": "Accounts_Update",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/accountsUpdateResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "ID",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "int64"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/accountsUpdateRequest"
            }
          }
        ],
        "tags": [
          "Accounts"
        ]
      }
    }
  },
  "definitions": {
    "accountsAccount": {
      "type": "object",
      "properties": {
        "ID": {
          "type": "string",
          "format": "int64"
        },
        "name": {
          "type": "string"
        },
        "balance": {
          "type": "number",
          "format": "double"
        },
        "currency": {
          "$ref": "#/definitions/accountscurrencyType"
        },
        "isAvailable": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "accountsCreateRequest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "balance": {
          "type": "number",
          "format": "double"
        },
        "currency": {
          "$ref": "#/definitions/accountscurrencyType"
        }
      },
      "title": "create"
    },
    "accountsCreateResponse": {
      "type": "object",
      "properties": {
        "ID": {
          "type": "string",
          "format": "int64"
        }
      }
    },
    "accountsDeleteResponse": {
      "type": "object"
    },
    "accountsGetResponse": {
      "type": "object",
      "properties": {
        "account": {
          "$ref": "#/definitions/accountsAccount"
        }
      }
    },
    "accountsListResponse": {
      "type": "object",
      "properties": {
        "accounts": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/accountsAccount"
          }
        }
      }
    },
    "accountsUpdateRequest": {
      "type": "object",
      "properties": {
        "ID": {
          "type": "string",
          "format": "int64"
        },
        "name": {
          "type": "string"
        },
        "balance": {
          "type": "number",
          "format": "double"
        },
        "currency": {
          "$ref": "#/definitions/accountscurrencyType"
        },
        "isAvailable": {
          "type": "boolean",
          "format": "boolean"
        }
      },
      "title": "update"
    },
    "accountsUpdateResponse": {
      "type": "object"
    },
    "accountscurrencyType": {
      "type": "string",
      "enum": [
        "USD",
        "EUR",
        "RU"
      ],
      "default": "USD"
    }
  }
}

`)
