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

package group

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/crossplane/provider-aws/apis/cognitoidentityprovider/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateGetGroupInput returns input for read
// operation.
func GenerateGetGroupInput(cr *svcapitypes.Group) *svcsdk.GetGroupInput {
	res := &svcsdk.GetGroupInput{}

	if cr.Status.AtProvider.GroupName != nil {
		res.SetGroupName(*cr.Status.AtProvider.GroupName)
	}
	if cr.Status.AtProvider.UserPoolID != nil {
		res.SetUserPoolId(*cr.Status.AtProvider.UserPoolID)
	}

	return res
}

// GenerateGroup returns the current state in the form of *svcapitypes.Group.
func GenerateGroup(resp *svcsdk.GetGroupOutput) *svcapitypes.Group {
	cr := &svcapitypes.Group{}

	if resp.Group.CreationDate != nil {
		cr.Status.AtProvider.CreationDate = &metav1.Time{*resp.Group.CreationDate}
	} else {
		cr.Status.AtProvider.CreationDate = nil
	}
	if resp.Group.Description != nil {
		cr.Spec.ForProvider.Description = resp.Group.Description
	} else {
		cr.Spec.ForProvider.Description = nil
	}
	if resp.Group.GroupName != nil {
		cr.Status.AtProvider.GroupName = resp.Group.GroupName
	} else {
		cr.Status.AtProvider.GroupName = nil
	}
	if resp.Group.LastModifiedDate != nil {
		cr.Status.AtProvider.LastModifiedDate = &metav1.Time{*resp.Group.LastModifiedDate}
	} else {
		cr.Status.AtProvider.LastModifiedDate = nil
	}
	if resp.Group.Precedence != nil {
		cr.Spec.ForProvider.Precedence = resp.Group.Precedence
	} else {
		cr.Spec.ForProvider.Precedence = nil
	}
	if resp.Group.RoleArn != nil {
		cr.Status.AtProvider.RoleARN = resp.Group.RoleArn
	} else {
		cr.Status.AtProvider.RoleARN = nil
	}
	if resp.Group.UserPoolId != nil {
		cr.Status.AtProvider.UserPoolID = resp.Group.UserPoolId
	} else {
		cr.Status.AtProvider.UserPoolID = nil
	}

	return cr
}

// GenerateCreateGroupInput returns a create input.
func GenerateCreateGroupInput(cr *svcapitypes.Group) *svcsdk.CreateGroupInput {
	res := &svcsdk.CreateGroupInput{}

	if cr.Spec.ForProvider.Description != nil {
		res.SetDescription(*cr.Spec.ForProvider.Description)
	}
	if cr.Spec.ForProvider.Precedence != nil {
		res.SetPrecedence(*cr.Spec.ForProvider.Precedence)
	}

	return res
}

// GenerateUpdateGroupInput returns an update input.
func GenerateUpdateGroupInput(cr *svcapitypes.Group) *svcsdk.UpdateGroupInput {
	res := &svcsdk.UpdateGroupInput{}

	if cr.Spec.ForProvider.Description != nil {
		res.SetDescription(*cr.Spec.ForProvider.Description)
	}
	if cr.Status.AtProvider.GroupName != nil {
		res.SetGroupName(*cr.Status.AtProvider.GroupName)
	}
	if cr.Spec.ForProvider.Precedence != nil {
		res.SetPrecedence(*cr.Spec.ForProvider.Precedence)
	}
	if cr.Status.AtProvider.RoleARN != nil {
		res.SetRoleArn(*cr.Status.AtProvider.RoleARN)
	}
	if cr.Status.AtProvider.UserPoolID != nil {
		res.SetUserPoolId(*cr.Status.AtProvider.UserPoolID)
	}

	return res
}

// GenerateDeleteGroupInput returns a deletion input.
func GenerateDeleteGroupInput(cr *svcapitypes.Group) *svcsdk.DeleteGroupInput {
	res := &svcsdk.DeleteGroupInput{}

	if cr.Status.AtProvider.GroupName != nil {
		res.SetGroupName(*cr.Status.AtProvider.GroupName)
	}
	if cr.Status.AtProvider.UserPoolID != nil {
		res.SetUserPoolId(*cr.Status.AtProvider.UserPoolID)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "ResourceNotFoundException"
}
