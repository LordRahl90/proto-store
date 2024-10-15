// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: v1/invoicer.proto

package invoicerconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	invoicer "github.com/LordRahl90/invoicer"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// InvoicerServiceName is the fully-qualified name of the InvoicerService service.
	InvoicerServiceName = "invoicer.v1.InvoicerService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// InvoicerServiceCreateProcedure is the fully-qualified name of the InvoicerService's Create RPC.
	InvoicerServiceCreateProcedure = "/invoicer.v1.InvoicerService/Create"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	invoicerServiceServiceDescriptor      = invoicer.File_v1_invoicer_proto.Services().ByName("InvoicerService")
	invoicerServiceCreateMethodDescriptor = invoicerServiceServiceDescriptor.Methods().ByName("Create")
)

// InvoicerServiceClient is a client for the invoicer.v1.InvoicerService service.
type InvoicerServiceClient interface {
	Create(context.Context, *connect.Request[invoicer.CreateRequest]) (*connect.Response[invoicer.CreateResponse], error)
}

// NewInvoicerServiceClient constructs a client for the invoicer.v1.InvoicerService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewInvoicerServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) InvoicerServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &invoicerServiceClient{
		create: connect.NewClient[invoicer.CreateRequest, invoicer.CreateResponse](
			httpClient,
			baseURL+InvoicerServiceCreateProcedure,
			connect.WithSchema(invoicerServiceCreateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// invoicerServiceClient implements InvoicerServiceClient.
type invoicerServiceClient struct {
	create *connect.Client[invoicer.CreateRequest, invoicer.CreateResponse]
}

// Create calls invoicer.v1.InvoicerService.Create.
func (c *invoicerServiceClient) Create(ctx context.Context, req *connect.Request[invoicer.CreateRequest]) (*connect.Response[invoicer.CreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// InvoicerServiceHandler is an implementation of the invoicer.v1.InvoicerService service.
type InvoicerServiceHandler interface {
	Create(context.Context, *connect.Request[invoicer.CreateRequest]) (*connect.Response[invoicer.CreateResponse], error)
}

// NewInvoicerServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewInvoicerServiceHandler(svc InvoicerServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	invoicerServiceCreateHandler := connect.NewUnaryHandler(
		InvoicerServiceCreateProcedure,
		svc.Create,
		connect.WithSchema(invoicerServiceCreateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/invoicer.v1.InvoicerService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case InvoicerServiceCreateProcedure:
			invoicerServiceCreateHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedInvoicerServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedInvoicerServiceHandler struct{}

func (UnimplementedInvoicerServiceHandler) Create(context.Context, *connect.Request[invoicer.CreateRequest]) (*connect.Response[invoicer.CreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("invoicer.v1.InvoicerService.Create is not implemented"))
}
