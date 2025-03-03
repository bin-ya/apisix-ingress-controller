// +build !ignore_autogenerated

// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v2beta1

import (
	v2beta3 "github.com/apache/apisix-ingress-controller/pkg/kube/apisix/apis/config/v2beta3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRoute) DeepCopyInto(out *ApisixRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRoute.
func (in *ApisixRoute) DeepCopy() *ApisixRoute {
	if in == nil {
		return nil
	}
	out := new(ApisixRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApisixRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteAuthentication) DeepCopyInto(out *ApisixRouteAuthentication) {
	*out = *in
	out.KeyAuth = in.KeyAuth
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteAuthentication.
func (in *ApisixRouteAuthentication) DeepCopy() *ApisixRouteAuthentication {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteAuthenticationKeyAuth) DeepCopyInto(out *ApisixRouteAuthenticationKeyAuth) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteAuthenticationKeyAuth.
func (in *ApisixRouteAuthenticationKeyAuth) DeepCopy() *ApisixRouteAuthenticationKeyAuth {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteAuthenticationKeyAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteHTTP) DeepCopyInto(out *ApisixRouteHTTP) {
	*out = *in
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(UpstreamTimeout)
		**out = **in
	}
	in.Match.DeepCopyInto(&out.Match)
	in.Backend.DeepCopyInto(&out.Backend)
	if in.Backends != nil {
		in, out := &in.Backends, &out.Backends
		*out = make([]v2beta3.ApisixRouteHTTPBackend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make([]ApisixRouteHTTPPlugin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Authentication = in.Authentication
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteHTTP.
func (in *ApisixRouteHTTP) DeepCopy() *ApisixRouteHTTP {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteHTTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteHTTPMatch) DeepCopyInto(out *ApisixRouteHTTPMatch) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Methods != nil {
		in, out := &in.Methods, &out.Methods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RemoteAddrs != nil {
		in, out := &in.RemoteAddrs, &out.RemoteAddrs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NginxVars != nil {
		in, out := &in.NginxVars, &out.NginxVars
		*out = make([]v2beta3.ApisixRouteHTTPMatchExpr, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteHTTPMatch.
func (in *ApisixRouteHTTPMatch) DeepCopy() *ApisixRouteHTTPMatch {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteHTTPMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteHTTPMatchExprSubject) DeepCopyInto(out *ApisixRouteHTTPMatchExprSubject) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteHTTPMatchExprSubject.
func (in *ApisixRouteHTTPMatchExprSubject) DeepCopy() *ApisixRouteHTTPMatchExprSubject {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteHTTPMatchExprSubject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteHTTPPlugin) DeepCopyInto(out *ApisixRouteHTTPPlugin) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteHTTPPlugin.
func (in *ApisixRouteHTTPPlugin) DeepCopy() *ApisixRouteHTTPPlugin {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteHTTPPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteList) DeepCopyInto(out *ApisixRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApisixRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteList.
func (in *ApisixRouteList) DeepCopy() *ApisixRouteList {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApisixRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteSpec) DeepCopyInto(out *ApisixRouteSpec) {
	*out = *in
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = make([]ApisixRouteHTTP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Stream != nil {
		in, out := &in.Stream, &out.Stream
		*out = make([]ApisixRouteStream, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteSpec.
func (in *ApisixRouteSpec) DeepCopy() *ApisixRouteSpec {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteStream) DeepCopyInto(out *ApisixRouteStream) {
	*out = *in
	out.Match = in.Match
	out.Backend = in.Backend
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteStream.
func (in *ApisixRouteStream) DeepCopy() *ApisixRouteStream {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteStream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteStreamBackend) DeepCopyInto(out *ApisixRouteStreamBackend) {
	*out = *in
	out.ServicePort = in.ServicePort
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteStreamBackend.
func (in *ApisixRouteStreamBackend) DeepCopy() *ApisixRouteStreamBackend {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteStreamBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteStreamMatch) DeepCopyInto(out *ApisixRouteStreamMatch) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteStreamMatch.
func (in *ApisixRouteStreamMatch) DeepCopy() *ApisixRouteStreamMatch {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteStreamMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixStatus) DeepCopyInto(out *ApisixStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixStatus.
func (in *ApisixStatus) DeepCopy() *ApisixStatus {
	if in == nil {
		return nil
	}
	out := new(ApisixStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpstreamTimeout) DeepCopyInto(out *UpstreamTimeout) {
	*out = *in
	out.Connect = in.Connect
	out.Send = in.Send
	out.Read = in.Read
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpstreamTimeout.
func (in *UpstreamTimeout) DeepCopy() *UpstreamTimeout {
	if in == nil {
		return nil
	}
	out := new(UpstreamTimeout)
	in.DeepCopyInto(out)
	return out
}
