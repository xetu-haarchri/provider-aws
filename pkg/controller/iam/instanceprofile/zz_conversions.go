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

package instanceprofile

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/iam"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/crossplane/provider-aws/apis/iam/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateGetInstanceProfileInput returns input for read
// operation.
func GenerateGetInstanceProfileInput(cr *svcapitypes.InstanceProfile) *svcsdk.GetInstanceProfileInput {
	res := &svcsdk.GetInstanceProfileInput{}

	if cr.Status.AtProvider.InstanceProfileName != nil {
		res.SetInstanceProfileName(*cr.Status.AtProvider.InstanceProfileName)
	}

	return res
}

// GenerateInstanceProfile returns the current state in the form of *svcapitypes.InstanceProfile.
func GenerateInstanceProfile(resp *svcsdk.GetInstanceProfileOutput) *svcapitypes.InstanceProfile {
	cr := &svcapitypes.InstanceProfile{}

	if resp.InstanceProfile.Arn != nil {
		cr.Status.AtProvider.ARN = resp.InstanceProfile.Arn
	} else {
		cr.Status.AtProvider.ARN = nil
	}
	if resp.InstanceProfile.CreateDate != nil {
		cr.Status.AtProvider.CreateDate = &metav1.Time{*resp.InstanceProfile.CreateDate}
	} else {
		cr.Status.AtProvider.CreateDate = nil
	}
	if resp.InstanceProfile.InstanceProfileId != nil {
		cr.Status.AtProvider.InstanceProfileID = resp.InstanceProfile.InstanceProfileId
	} else {
		cr.Status.AtProvider.InstanceProfileID = nil
	}
	if resp.InstanceProfile.InstanceProfileName != nil {
		cr.Status.AtProvider.InstanceProfileName = resp.InstanceProfile.InstanceProfileName
	} else {
		cr.Status.AtProvider.InstanceProfileName = nil
	}
	if resp.InstanceProfile.Path != nil {
		cr.Spec.ForProvider.Path = resp.InstanceProfile.Path
	} else {
		cr.Spec.ForProvider.Path = nil
	}
	if resp.InstanceProfile.Roles != nil {
		f5 := []*svcapitypes.Role{}
		for _, f5iter := range resp.InstanceProfile.Roles {
			f5elem := &svcapitypes.Role{}
			if f5iter.Arn != nil {
				f5elem.ARN = f5iter.Arn
			}
			if f5iter.AssumeRolePolicyDocument != nil {
				f5elem.AssumeRolePolicyDocument = f5iter.AssumeRolePolicyDocument
			}
			if f5iter.CreateDate != nil {
				f5elem.CreateDate = &metav1.Time{*f5iter.CreateDate}
			}
			if f5iter.Description != nil {
				f5elem.Description = f5iter.Description
			}
			if f5iter.MaxSessionDuration != nil {
				f5elem.MaxSessionDuration = f5iter.MaxSessionDuration
			}
			if f5iter.Path != nil {
				f5elem.Path = f5iter.Path
			}
			if f5iter.PermissionsBoundary != nil {
				f5elemf6 := &svcapitypes.AttachedPermissionsBoundary{}
				if f5iter.PermissionsBoundary.PermissionsBoundaryArn != nil {
					f5elemf6.PermissionsBoundaryARN = f5iter.PermissionsBoundary.PermissionsBoundaryArn
				}
				if f5iter.PermissionsBoundary.PermissionsBoundaryType != nil {
					f5elemf6.PermissionsBoundaryType = f5iter.PermissionsBoundary.PermissionsBoundaryType
				}
				f5elem.PermissionsBoundary = f5elemf6
			}
			if f5iter.RoleId != nil {
				f5elem.RoleID = f5iter.RoleId
			}
			if f5iter.RoleLastUsed != nil {
				f5elemf8 := &svcapitypes.RoleLastUsed{}
				if f5iter.RoleLastUsed.LastUsedDate != nil {
					f5elemf8.LastUsedDate = &metav1.Time{*f5iter.RoleLastUsed.LastUsedDate}
				}
				if f5iter.RoleLastUsed.Region != nil {
					f5elemf8.Region = f5iter.RoleLastUsed.Region
				}
				f5elem.RoleLastUsed = f5elemf8
			}
			if f5iter.RoleName != nil {
				f5elem.RoleName = f5iter.RoleName
			}
			if f5iter.Tags != nil {
				f5elemf10 := []*svcapitypes.Tag{}
				for _, f5elemf10iter := range f5iter.Tags {
					f5elemf10elem := &svcapitypes.Tag{}
					if f5elemf10iter.Key != nil {
						f5elemf10elem.Key = f5elemf10iter.Key
					}
					if f5elemf10iter.Value != nil {
						f5elemf10elem.Value = f5elemf10iter.Value
					}
					f5elemf10 = append(f5elemf10, f5elemf10elem)
				}
				f5elem.Tags = f5elemf10
			}
			f5 = append(f5, f5elem)
		}
		cr.Status.AtProvider.Roles = f5
	} else {
		cr.Status.AtProvider.Roles = nil
	}
	if resp.InstanceProfile.Tags != nil {
		f6 := []*svcapitypes.Tag{}
		for _, f6iter := range resp.InstanceProfile.Tags {
			f6elem := &svcapitypes.Tag{}
			if f6iter.Key != nil {
				f6elem.Key = f6iter.Key
			}
			if f6iter.Value != nil {
				f6elem.Value = f6iter.Value
			}
			f6 = append(f6, f6elem)
		}
		cr.Spec.ForProvider.Tags = f6
	} else {
		cr.Spec.ForProvider.Tags = nil
	}

	return cr
}

// GenerateCreateInstanceProfileInput returns a create input.
func GenerateCreateInstanceProfileInput(cr *svcapitypes.InstanceProfile) *svcsdk.CreateInstanceProfileInput {
	res := &svcsdk.CreateInstanceProfileInput{}

	if cr.Spec.ForProvider.Path != nil {
		res.SetPath(*cr.Spec.ForProvider.Path)
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

// GenerateDeleteInstanceProfileInput returns a deletion input.
func GenerateDeleteInstanceProfileInput(cr *svcapitypes.InstanceProfile) *svcsdk.DeleteInstanceProfileInput {
	res := &svcsdk.DeleteInstanceProfileInput{}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "NoSuchEntity"
}
