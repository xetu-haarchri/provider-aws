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

package policy

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/iot"

	svcapitypes "github.com/crossplane/provider-aws/apis/iot/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateGetPolicyInput returns input for read
// operation.
func GenerateGetPolicyInput(cr *svcapitypes.Policy) *svcsdk.GetPolicyInput {
	res := &svcsdk.GetPolicyInput{}

	return res
}

// GeneratePolicy returns the current state in the form of *svcapitypes.Policy.
func GeneratePolicy(resp *svcsdk.GetPolicyOutput) *svcapitypes.Policy {
	cr := &svcapitypes.Policy{}

	if resp.PolicyArn != nil {
		cr.Status.AtProvider.PolicyARN = resp.PolicyArn
	} else {
		cr.Status.AtProvider.PolicyARN = nil
	}
	if resp.PolicyDocument != nil {
		cr.Spec.ForProvider.PolicyDocument = resp.PolicyDocument
	} else {
		cr.Spec.ForProvider.PolicyDocument = nil
	}

	return cr
}

// GenerateCreatePolicyInput returns a create input.
func GenerateCreatePolicyInput(cr *svcapitypes.Policy) *svcsdk.CreatePolicyInput {
	res := &svcsdk.CreatePolicyInput{}

	if cr.Spec.ForProvider.PolicyDocument != nil {
		res.SetPolicyDocument(*cr.Spec.ForProvider.PolicyDocument)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f1 := []*svcsdk.Tag{}
		for _, f1iter := range cr.Spec.ForProvider.Tags {
			f1elem := &svcsdk.Tag{}
			if f1iter.Key != nil {
				f1elem.SetKey(*f1iter.Key)
			}
			if f1iter.Value != nil {
				f1elem.SetValue(*f1iter.Value)
			}
			f1 = append(f1, f1elem)
		}
		res.SetTags(f1)
	}

	return res
}

// GenerateDeletePolicyInput returns a deletion input.
func GenerateDeletePolicyInput(cr *svcapitypes.Policy) *svcsdk.DeletePolicyInput {
	res := &svcsdk.DeletePolicyInput{}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "ResourceNotFoundException"
}
