//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright © 2020 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package output

import (
	"github.com/banzaicloud/operator-tools/pkg/secret"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskBuffer) DeepCopyInto(out *DiskBuffer) {
	*out = *in
	if in.Compaction != nil {
		in, out := &in.Compaction, &out.Compaction
		*out = new(bool)
		**out = **in
	}
	if in.MemBufLength != nil {
		in, out := &in.MemBufLength, &out.MemBufLength
		*out = new(int64)
		**out = **in
	}
	if in.MemBufSize != nil {
		in, out := &in.MemBufSize, &out.MemBufSize
		*out = new(int64)
		**out = **in
	}
	if in.QOutSize != nil {
		in, out := &in.QOutSize, &out.QOutSize
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskBuffer.
func (in *DiskBuffer) DeepCopy() *DiskBuffer {
	if in == nil {
		return nil
	}
	out := new(DiskBuffer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileOutput) DeepCopyInto(out *FileOutput) {
	*out = *in
	if in.DiskBuffer != nil {
		in, out := &in.DiskBuffer, &out.DiskBuffer
		*out = new(DiskBuffer)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileOutput.
func (in *FileOutput) DeepCopy() *FileOutput {
	if in == nil {
		return nil
	}
	out := new(FileOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPOutput) DeepCopyInto(out *HTTPOutput) {
	*out = *in
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLS)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskBuffer != nil {
		in, out := &in.DiskBuffer, &out.DiskBuffer
		*out = new(DiskBuffer)
		(*in).DeepCopyInto(*out)
	}
	out.Batch = in.Batch
	in.Password.DeepCopyInto(&out.Password)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPOutput.
func (in *HTTPOutput) DeepCopy() *HTTPOutput {
	if in == nil {
		return nil
	}
	out := new(HTTPOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Loggly) DeepCopyInto(out *Loggly) {
	*out = *in
	if in.Token != nil {
		in, out := &in.Token, &out.Token
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	in.SyslogOutput.DeepCopyInto(&out.SyslogOutput)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Loggly.
func (in *Loggly) DeepCopy() *Loggly {
	if in == nil {
		return nil
	}
	out := new(Loggly)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SumologicHTTPOutput) DeepCopyInto(out *SumologicHTTPOutput) {
	*out = *in
	if in.Collector != nil {
		in, out := &in.Collector, &out.Collector
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLS)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskBuffer != nil {
		in, out := &in.DiskBuffer, &out.DiskBuffer
		*out = new(DiskBuffer)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SumologicHTTPOutput.
func (in *SumologicHTTPOutput) DeepCopy() *SumologicHTTPOutput {
	if in == nil {
		return nil
	}
	out := new(SumologicHTTPOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SumologicSyslogOutput) DeepCopyInto(out *SumologicSyslogOutput) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLS)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskBuffer != nil {
		in, out := &in.DiskBuffer, &out.DiskBuffer
		*out = new(DiskBuffer)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SumologicSyslogOutput.
func (in *SumologicSyslogOutput) DeepCopy() *SumologicSyslogOutput {
	if in == nil {
		return nil
	}
	out := new(SumologicSyslogOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyslogOutput) DeepCopyInto(out *SyslogOutput) {
	*out = *in
	if in.CloseOnInput != nil {
		in, out := &in.CloseOnInput, &out.CloseOnInput
		*out = new(bool)
		**out = **in
	}
	if in.Flags != nil {
		in, out := &in.Flags, &out.Flags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SoKeepalive != nil {
		in, out := &in.SoKeepalive, &out.SoKeepalive
		*out = new(bool)
		**out = **in
	}
	if in.TemplateEscape != nil {
		in, out := &in.TemplateEscape, &out.TemplateEscape
		*out = new(bool)
		**out = **in
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLS)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskBuffer != nil {
		in, out := &in.DiskBuffer, &out.DiskBuffer
		*out = new(DiskBuffer)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyslogOutput.
func (in *SyslogOutput) DeepCopy() *SyslogOutput {
	if in == nil {
		return nil
	}
	out := new(SyslogOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLS) DeepCopyInto(out *TLS) {
	*out = *in
	if in.CaDir != nil {
		in, out := &in.CaDir, &out.CaDir
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.CaFile != nil {
		in, out := &in.CaFile, &out.CaFile
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyFile != nil {
		in, out := &in.KeyFile, &out.KeyFile
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.CertFile != nil {
		in, out := &in.CertFile, &out.CertFile
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.UseSystemCertStore != nil {
		in, out := &in.UseSystemCertStore, &out.UseSystemCertStore
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLS.
func (in *TLS) DeepCopy() *TLS {
	if in == nil {
		return nil
	}
	out := new(TLS)
	in.DeepCopyInto(out)
	return out
}
