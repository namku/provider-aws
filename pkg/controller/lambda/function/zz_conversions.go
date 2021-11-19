/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package function

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/lambda"

	svcapitypes "github.com/crossplane/provider-aws/apis/lambda/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateGetFunctionInput returns input for read
// operation.
func GenerateGetFunctionInput(cr *svcapitypes.Function) *svcsdk.GetFunctionInput {
	res := &svcsdk.GetFunctionInput{}

	return res
}

// GenerateFunction returns the current state in the form of *svcapitypes.Function.
func GenerateFunction(resp *svcsdk.GetFunctionOutput) *svcapitypes.Function {
	cr := &svcapitypes.Function{}

	return cr
}

// GenerateCreateFunctionInput returns a create input.
func GenerateCreateFunctionInput(cr *svcapitypes.Function) *svcsdk.CreateFunctionInput {
	res := &svcsdk.CreateFunctionInput{}

	if cr.Spec.ForProvider.CodeSigningConfigARN != nil {
		res.SetCodeSigningConfigArn(*cr.Spec.ForProvider.CodeSigningConfigARN)
	}
	if cr.Spec.ForProvider.DeadLetterConfig != nil {
		f1 := &svcsdk.DeadLetterConfig{}
		if cr.Spec.ForProvider.DeadLetterConfig.TargetARN != nil {
			f1.SetTargetArn(*cr.Spec.ForProvider.DeadLetterConfig.TargetARN)
		}
		res.SetDeadLetterConfig(f1)
	}
	if cr.Spec.ForProvider.Description != nil {
		res.SetDescription(*cr.Spec.ForProvider.Description)
	}
	if cr.Spec.ForProvider.Environment != nil {
		f3 := &svcsdk.Environment{}
		if cr.Spec.ForProvider.Environment.Variables != nil {
			f3f0 := map[string]*string{}
			for f3f0key, f3f0valiter := range cr.Spec.ForProvider.Environment.Variables {
				var f3f0val string
				f3f0val = *f3f0valiter
				f3f0[f3f0key] = &f3f0val
			}
			f3.SetVariables(f3f0)
		}
		res.SetEnvironment(f3)
	}
	if cr.Spec.ForProvider.FileSystemConfigs != nil {
		f4 := []*svcsdk.FileSystemConfig{}
		for _, f4iter := range cr.Spec.ForProvider.FileSystemConfigs {
			f4elem := &svcsdk.FileSystemConfig{}
			if f4iter.ARN != nil {
				f4elem.SetArn(*f4iter.ARN)
			}
			if f4iter.LocalMountPath != nil {
				f4elem.SetLocalMountPath(*f4iter.LocalMountPath)
			}
			f4 = append(f4, f4elem)
		}
		res.SetFileSystemConfigs(f4)
	}
	if cr.Spec.ForProvider.Handler != nil {
		res.SetHandler(*cr.Spec.ForProvider.Handler)
	}
	if cr.Spec.ForProvider.ImageConfig != nil {
		f6 := &svcsdk.ImageConfig{}
		if cr.Spec.ForProvider.ImageConfig.Command != nil {
			f6f0 := []*string{}
			for _, f6f0iter := range cr.Spec.ForProvider.ImageConfig.Command {
				var f6f0elem string
				f6f0elem = *f6f0iter
				f6f0 = append(f6f0, &f6f0elem)
			}
			f6.SetCommand(f6f0)
		}
		if cr.Spec.ForProvider.ImageConfig.EntryPoint != nil {
			f6f1 := []*string{}
			for _, f6f1iter := range cr.Spec.ForProvider.ImageConfig.EntryPoint {
				var f6f1elem string
				f6f1elem = *f6f1iter
				f6f1 = append(f6f1, &f6f1elem)
			}
			f6.SetEntryPoint(f6f1)
		}
		if cr.Spec.ForProvider.ImageConfig.WorkingDirectory != nil {
			f6.SetWorkingDirectory(*cr.Spec.ForProvider.ImageConfig.WorkingDirectory)
		}
		res.SetImageConfig(f6)
	}
	if cr.Spec.ForProvider.KMSKeyARN != nil {
		res.SetKMSKeyArn(*cr.Spec.ForProvider.KMSKeyARN)
	}
	if cr.Spec.ForProvider.Layers != nil {
		f8 := []*string{}
		for _, f8iter := range cr.Spec.ForProvider.Layers {
			var f8elem string
			f8elem = *f8iter
			f8 = append(f8, &f8elem)
		}
		res.SetLayers(f8)
	}
	if cr.Spec.ForProvider.MemorySize != nil {
		res.SetMemorySize(*cr.Spec.ForProvider.MemorySize)
	}
	if cr.Spec.ForProvider.PackageType != nil {
		res.SetPackageType(*cr.Spec.ForProvider.PackageType)
	}
	if cr.Spec.ForProvider.Publish != nil {
		res.SetPublish(*cr.Spec.ForProvider.Publish)
	}
	if cr.Spec.ForProvider.Runtime != nil {
		res.SetRuntime(*cr.Spec.ForProvider.Runtime)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f13 := map[string]*string{}
		for f13key, f13valiter := range cr.Spec.ForProvider.Tags {
			var f13val string
			f13val = *f13valiter
			f13[f13key] = &f13val
		}
		res.SetTags(f13)
	}
	if cr.Spec.ForProvider.Timeout != nil {
		res.SetTimeout(*cr.Spec.ForProvider.Timeout)
	}
	if cr.Spec.ForProvider.TracingConfig != nil {
		f15 := &svcsdk.TracingConfig{}
		if cr.Spec.ForProvider.TracingConfig.Mode != nil {
			f15.SetMode(*cr.Spec.ForProvider.TracingConfig.Mode)
		}
		res.SetTracingConfig(f15)
	}

	return res
}

// GenerateDeleteFunctionInput returns a deletion input.
func GenerateDeleteFunctionInput(cr *svcapitypes.Function) *svcsdk.DeleteFunctionInput {
	res := &svcsdk.DeleteFunctionInput{}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "ResourceNotFoundException"
}