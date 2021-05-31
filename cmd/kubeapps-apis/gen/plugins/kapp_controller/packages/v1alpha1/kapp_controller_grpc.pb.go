// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	v1alpha1 "github.com/kubeapps/kubeapps/cmd/kubeapps-apis/gen/core/packages/v1alpha1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// KappControllerPackagesServiceClient is the client API for KappControllerPackagesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KappControllerPackagesServiceClient interface {
	// GetAvailablePackages returns the available packages managed by the 'kapp_controller' plugin
	GetAvailablePackages(ctx context.Context, in *v1alpha1.GetAvailablePackagesRequest, opts ...grpc.CallOption) (*v1alpha1.GetAvailablePackagesResponse, error)
	// GetPackageRepositories returns the repositories managed by the 'kapp_controller' plugin
	GetPackageRepositories(ctx context.Context, in *v1alpha1.GetPackageRepositoriesRequest, opts ...grpc.CallOption) (*v1alpha1.GetPackageRepositoriesResponse, error)
}

type kappControllerPackagesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKappControllerPackagesServiceClient(cc grpc.ClientConnInterface) KappControllerPackagesServiceClient {
	return &kappControllerPackagesServiceClient{cc}
}

func (c *kappControllerPackagesServiceClient) GetAvailablePackages(ctx context.Context, in *v1alpha1.GetAvailablePackagesRequest, opts ...grpc.CallOption) (*v1alpha1.GetAvailablePackagesResponse, error) {
	out := new(v1alpha1.GetAvailablePackagesResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.kapp_controller.packages.v1alpha1.KappControllerPackagesService/GetAvailablePackages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kappControllerPackagesServiceClient) GetPackageRepositories(ctx context.Context, in *v1alpha1.GetPackageRepositoriesRequest, opts ...grpc.CallOption) (*v1alpha1.GetPackageRepositoriesResponse, error) {
	out := new(v1alpha1.GetPackageRepositoriesResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.kapp_controller.packages.v1alpha1.KappControllerPackagesService/GetPackageRepositories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KappControllerPackagesServiceServer is the server API for KappControllerPackagesService service.
// All implementations must embed UnimplementedKappControllerPackagesServiceServer
// for forward compatibility
type KappControllerPackagesServiceServer interface {
	// GetAvailablePackages returns the available packages managed by the 'kapp_controller' plugin
	GetAvailablePackages(context.Context, *v1alpha1.GetAvailablePackagesRequest) (*v1alpha1.GetAvailablePackagesResponse, error)
	// GetPackageRepositories returns the repositories managed by the 'kapp_controller' plugin
	GetPackageRepositories(context.Context, *v1alpha1.GetPackageRepositoriesRequest) (*v1alpha1.GetPackageRepositoriesResponse, error)
	mustEmbedUnimplementedKappControllerPackagesServiceServer()
}

// UnimplementedKappControllerPackagesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKappControllerPackagesServiceServer struct {
}

func (UnimplementedKappControllerPackagesServiceServer) GetAvailablePackages(context.Context, *v1alpha1.GetAvailablePackagesRequest) (*v1alpha1.GetAvailablePackagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailablePackages not implemented")
}
func (UnimplementedKappControllerPackagesServiceServer) GetPackageRepositories(context.Context, *v1alpha1.GetPackageRepositoriesRequest) (*v1alpha1.GetPackageRepositoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPackageRepositories not implemented")
}
func (UnimplementedKappControllerPackagesServiceServer) mustEmbedUnimplementedKappControllerPackagesServiceServer() {
}

// UnsafeKappControllerPackagesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KappControllerPackagesServiceServer will
// result in compilation errors.
type UnsafeKappControllerPackagesServiceServer interface {
	mustEmbedUnimplementedKappControllerPackagesServiceServer()
}

func RegisterKappControllerPackagesServiceServer(s grpc.ServiceRegistrar, srv KappControllerPackagesServiceServer) {
	s.RegisterService(&KappControllerPackagesService_ServiceDesc, srv)
}

func _KappControllerPackagesService_GetAvailablePackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.GetAvailablePackagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KappControllerPackagesServiceServer).GetAvailablePackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.kapp_controller.packages.v1alpha1.KappControllerPackagesService/GetAvailablePackages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KappControllerPackagesServiceServer).GetAvailablePackages(ctx, req.(*v1alpha1.GetAvailablePackagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KappControllerPackagesService_GetPackageRepositories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.GetPackageRepositoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KappControllerPackagesServiceServer).GetPackageRepositories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.kapp_controller.packages.v1alpha1.KappControllerPackagesService/GetPackageRepositories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KappControllerPackagesServiceServer).GetPackageRepositories(ctx, req.(*v1alpha1.GetPackageRepositoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KappControllerPackagesService_ServiceDesc is the grpc.ServiceDesc for KappControllerPackagesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KappControllerPackagesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kubeappsapis.plugins.kapp_controller.packages.v1alpha1.KappControllerPackagesService",
	HandlerType: (*KappControllerPackagesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAvailablePackages",
			Handler:    _KappControllerPackagesService_GetAvailablePackages_Handler,
		},
		{
			MethodName: "GetPackageRepositories",
			Handler:    _KappControllerPackagesService_GetPackageRepositories_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kubeappsapis/plugins/kapp_controller/packages/v1alpha1/kapp_controller.proto",
}