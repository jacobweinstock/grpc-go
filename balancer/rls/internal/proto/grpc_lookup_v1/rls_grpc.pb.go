// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc_lookup_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// RouteLookupServiceClient is the client API for RouteLookupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteLookupServiceClient interface {
	// Lookup returns a target for a single key.
	RouteLookup(ctx context.Context, in *RouteLookupRequest, opts ...grpc.CallOption) (*RouteLookupResponse, error)
}

type routeLookupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteLookupServiceClient(cc grpc.ClientConnInterface) RouteLookupServiceClient {
	return &routeLookupServiceClient{cc}
}

var routeLookupServiceRouteLookupStreamDesc = &grpc.StreamDesc{
	StreamName: "RouteLookup",
}

func (c *routeLookupServiceClient) RouteLookup(ctx context.Context, in *RouteLookupRequest, opts ...grpc.CallOption) (*RouteLookupResponse, error) {
	out := new(RouteLookupResponse)
	err := c.cc.Invoke(ctx, "/grpc.lookup.v1.RouteLookupService/RouteLookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteLookupServiceService is the service API for RouteLookupService service.
// Fields should be assigned to their respective handler implementations only before
// RegisterRouteLookupServiceService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type RouteLookupServiceService struct {
	// Lookup returns a target for a single key.
	RouteLookup func(context.Context, *RouteLookupRequest) (*RouteLookupResponse, error)
}

func (s *RouteLookupServiceService) routeLookup(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteLookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.RouteLookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/grpc.lookup.v1.RouteLookupService/RouteLookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.RouteLookup(ctx, req.(*RouteLookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterRouteLookupServiceService registers a service implementation with a gRPC server.
func RegisterRouteLookupServiceService(s grpc.ServiceRegistrar, srv *RouteLookupServiceService) {
	srvCopy := *srv
	if srvCopy.RouteLookup == nil {
		srvCopy.RouteLookup = func(context.Context, *RouteLookupRequest) (*RouteLookupResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method RouteLookup not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "grpc.lookup.v1.RouteLookupService",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "RouteLookup",
				Handler:    srvCopy.routeLookup,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "grpc/lookup/v1/rls.proto",
	}

	s.RegisterService(&sd, nil)
}

// RouteLookupServiceServer is the service API for RouteLookupService service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended unless you own the service definition.
type RouteLookupServiceServer interface {
	// Lookup returns a target for a single key.
	RouteLookup(context.Context, *RouteLookupRequest) (*RouteLookupResponse, error)
}

// UnimplementedRouteLookupServiceServer can be embedded to have forward compatible implementations of
// RouteLookupServiceServer
type UnimplementedRouteLookupServiceServer struct {
}

func (UnimplementedRouteLookupServiceServer) RouteLookup(context.Context, *RouteLookupRequest) (*RouteLookupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RouteLookup not implemented")
}

// RegisterRouteLookupServiceServer registers a service implementation with a gRPC server.
func RegisterRouteLookupServiceServer(s grpc.ServiceRegistrar, srv RouteLookupServiceServer) {
	str := &RouteLookupServiceService{
		RouteLookup: srv.RouteLookup,
	}
	RegisterRouteLookupServiceService(s, str)
}
