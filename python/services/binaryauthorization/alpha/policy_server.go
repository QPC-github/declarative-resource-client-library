// Copyright 2022 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"context"
	"errors"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/binaryauthorization/alpha/binaryauthorization_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization/alpha"
)

// Server implements the gRPC interface for Policy.
type PolicyServer struct{}

// ProtoToPolicyAdmissionRuleEvaluationModeEnum converts a PolicyAdmissionRuleEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationAlphaPolicyAdmissionRuleEvaluationModeEnum(e alphapb.BinaryauthorizationAlphaPolicyAdmissionRuleEvaluationModeEnum) *alpha.PolicyAdmissionRuleEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BinaryauthorizationAlphaPolicyAdmissionRuleEvaluationModeEnum_name[int32(e)]; ok {
		e := alpha.PolicyAdmissionRuleEvaluationModeEnum(n[len("BinaryauthorizationAlphaPolicyAdmissionRuleEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyAdmissionRuleEnforcementModeEnum converts a PolicyAdmissionRuleEnforcementModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationAlphaPolicyAdmissionRuleEnforcementModeEnum(e alphapb.BinaryauthorizationAlphaPolicyAdmissionRuleEnforcementModeEnum) *alpha.PolicyAdmissionRuleEnforcementModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BinaryauthorizationAlphaPolicyAdmissionRuleEnforcementModeEnum_name[int32(e)]; ok {
		e := alpha.PolicyAdmissionRuleEnforcementModeEnum(n[len("BinaryauthorizationAlphaPolicyAdmissionRuleEnforcementModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyGlobalPolicyEvaluationModeEnum converts a PolicyGlobalPolicyEvaluationModeEnum enum from its proto representation.
func ProtoToBinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnum(e alphapb.BinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnum) *alpha.PolicyGlobalPolicyEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnum_name[int32(e)]; ok {
		e := alpha.PolicyGlobalPolicyEvaluationModeEnum(n[len("BinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToPolicyAdmissionWhitelistPatterns converts a PolicyAdmissionWhitelistPatterns object from its proto representation.
func ProtoToBinaryauthorizationAlphaPolicyAdmissionWhitelistPatterns(p *alphapb.BinaryauthorizationAlphaPolicyAdmissionWhitelistPatterns) *alpha.PolicyAdmissionWhitelistPatterns {
	if p == nil {
		return nil
	}
	obj := &alpha.PolicyAdmissionWhitelistPatterns{
		NamePattern: dcl.StringOrNil(p.GetNamePattern()),
	}
	return obj
}

// ProtoToPolicyAdmissionRule converts a PolicyAdmissionRule object from its proto representation.
func ProtoToBinaryauthorizationAlphaPolicyAdmissionRule(p *alphapb.BinaryauthorizationAlphaPolicyAdmissionRule) *alpha.PolicyAdmissionRule {
	if p == nil {
		return nil
	}
	obj := &alpha.PolicyAdmissionRule{
		EvaluationMode:  ProtoToBinaryauthorizationAlphaPolicyAdmissionRuleEvaluationModeEnum(p.GetEvaluationMode()),
		EnforcementMode: ProtoToBinaryauthorizationAlphaPolicyAdmissionRuleEnforcementModeEnum(p.GetEnforcementMode()),
	}
	for _, r := range p.GetRequireAttestationsBy() {
		obj.RequireAttestationsBy = append(obj.RequireAttestationsBy, r)
	}
	return obj
}

// ProtoToPolicy converts a Policy resource from its proto representation.
func ProtoToPolicy(p *alphapb.BinaryauthorizationAlphaPolicy) *alpha.Policy {
	obj := &alpha.Policy{
		DefaultAdmissionRule:       ProtoToBinaryauthorizationAlphaPolicyAdmissionRule(p.GetDefaultAdmissionRule()),
		Description:                dcl.StringOrNil(p.GetDescription()),
		GlobalPolicyEvaluationMode: ProtoToBinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnum(p.GetGlobalPolicyEvaluationMode()),
		SelfLink:                   dcl.StringOrNil(p.GetSelfLink()),
		Project:                    dcl.StringOrNil(p.GetProject()),
		UpdateTime:                 dcl.StringOrNil(p.GetUpdateTime()),
	}
	for _, r := range p.GetAdmissionWhitelistPatterns() {
		obj.AdmissionWhitelistPatterns = append(obj.AdmissionWhitelistPatterns, *ProtoToBinaryauthorizationAlphaPolicyAdmissionWhitelistPatterns(r))
	}
	return obj
}

// PolicyAdmissionRuleEvaluationModeEnumToProto converts a PolicyAdmissionRuleEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationAlphaPolicyAdmissionRuleEvaluationModeEnumToProto(e *alpha.PolicyAdmissionRuleEvaluationModeEnum) alphapb.BinaryauthorizationAlphaPolicyAdmissionRuleEvaluationModeEnum {
	if e == nil {
		return alphapb.BinaryauthorizationAlphaPolicyAdmissionRuleEvaluationModeEnum(0)
	}
	if v, ok := alphapb.BinaryauthorizationAlphaPolicyAdmissionRuleEvaluationModeEnum_value["PolicyAdmissionRuleEvaluationModeEnum"+string(*e)]; ok {
		return alphapb.BinaryauthorizationAlphaPolicyAdmissionRuleEvaluationModeEnum(v)
	}
	return alphapb.BinaryauthorizationAlphaPolicyAdmissionRuleEvaluationModeEnum(0)
}

// PolicyAdmissionRuleEnforcementModeEnumToProto converts a PolicyAdmissionRuleEnforcementModeEnum enum to its proto representation.
func BinaryauthorizationAlphaPolicyAdmissionRuleEnforcementModeEnumToProto(e *alpha.PolicyAdmissionRuleEnforcementModeEnum) alphapb.BinaryauthorizationAlphaPolicyAdmissionRuleEnforcementModeEnum {
	if e == nil {
		return alphapb.BinaryauthorizationAlphaPolicyAdmissionRuleEnforcementModeEnum(0)
	}
	if v, ok := alphapb.BinaryauthorizationAlphaPolicyAdmissionRuleEnforcementModeEnum_value["PolicyAdmissionRuleEnforcementModeEnum"+string(*e)]; ok {
		return alphapb.BinaryauthorizationAlphaPolicyAdmissionRuleEnforcementModeEnum(v)
	}
	return alphapb.BinaryauthorizationAlphaPolicyAdmissionRuleEnforcementModeEnum(0)
}

// PolicyGlobalPolicyEvaluationModeEnumToProto converts a PolicyGlobalPolicyEvaluationModeEnum enum to its proto representation.
func BinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnumToProto(e *alpha.PolicyGlobalPolicyEvaluationModeEnum) alphapb.BinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnum {
	if e == nil {
		return alphapb.BinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnum(0)
	}
	if v, ok := alphapb.BinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnum_value["PolicyGlobalPolicyEvaluationModeEnum"+string(*e)]; ok {
		return alphapb.BinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnum(v)
	}
	return alphapb.BinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnum(0)
}

// PolicyAdmissionWhitelistPatternsToProto converts a PolicyAdmissionWhitelistPatterns object to its proto representation.
func BinaryauthorizationAlphaPolicyAdmissionWhitelistPatternsToProto(o *alpha.PolicyAdmissionWhitelistPatterns) *alphapb.BinaryauthorizationAlphaPolicyAdmissionWhitelistPatterns {
	if o == nil {
		return nil
	}
	p := &alphapb.BinaryauthorizationAlphaPolicyAdmissionWhitelistPatterns{}
	p.SetNamePattern(dcl.ValueOrEmptyString(o.NamePattern))
	return p
}

// PolicyAdmissionRuleToProto converts a PolicyAdmissionRule object to its proto representation.
func BinaryauthorizationAlphaPolicyAdmissionRuleToProto(o *alpha.PolicyAdmissionRule) *alphapb.BinaryauthorizationAlphaPolicyAdmissionRule {
	if o == nil {
		return nil
	}
	p := &alphapb.BinaryauthorizationAlphaPolicyAdmissionRule{}
	p.SetEvaluationMode(BinaryauthorizationAlphaPolicyAdmissionRuleEvaluationModeEnumToProto(o.EvaluationMode))
	p.SetEnforcementMode(BinaryauthorizationAlphaPolicyAdmissionRuleEnforcementModeEnumToProto(o.EnforcementMode))
	sRequireAttestationsBy := make([]string, len(o.RequireAttestationsBy))
	for i, r := range o.RequireAttestationsBy {
		sRequireAttestationsBy[i] = r
	}
	p.SetRequireAttestationsBy(sRequireAttestationsBy)
	return p
}

// PolicyToProto converts a Policy resource to its proto representation.
func PolicyToProto(resource *alpha.Policy) *alphapb.BinaryauthorizationAlphaPolicy {
	p := &alphapb.BinaryauthorizationAlphaPolicy{}
	p.SetDefaultAdmissionRule(BinaryauthorizationAlphaPolicyAdmissionRuleToProto(resource.DefaultAdmissionRule))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetGlobalPolicyEvaluationMode(BinaryauthorizationAlphaPolicyGlobalPolicyEvaluationModeEnumToProto(resource.GlobalPolicyEvaluationMode))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	sAdmissionWhitelistPatterns := make([]*alphapb.BinaryauthorizationAlphaPolicyAdmissionWhitelistPatterns, len(resource.AdmissionWhitelistPatterns))
	for i, r := range resource.AdmissionWhitelistPatterns {
		sAdmissionWhitelistPatterns[i] = BinaryauthorizationAlphaPolicyAdmissionWhitelistPatternsToProto(&r)
	}
	p.SetAdmissionWhitelistPatterns(sAdmissionWhitelistPatterns)
	mClusterAdmissionRules := make(map[string]*alphapb.BinaryauthorizationAlphaPolicyAdmissionRule, len(resource.ClusterAdmissionRules))
	for k, r := range resource.ClusterAdmissionRules {
		mClusterAdmissionRules[k] = BinaryauthorizationAlphaPolicyAdmissionRuleToProto(&r)
	}
	p.SetClusterAdmissionRules(mClusterAdmissionRules)
	mKubernetesNamespaceAdmissionRules := make(map[string]*alphapb.BinaryauthorizationAlphaPolicyAdmissionRule, len(resource.KubernetesNamespaceAdmissionRules))
	for k, r := range resource.KubernetesNamespaceAdmissionRules {
		mKubernetesNamespaceAdmissionRules[k] = BinaryauthorizationAlphaPolicyAdmissionRuleToProto(&r)
	}
	p.SetKubernetesNamespaceAdmissionRules(mKubernetesNamespaceAdmissionRules)
	mKubernetesServiceAccountAdmissionRules := make(map[string]*alphapb.BinaryauthorizationAlphaPolicyAdmissionRule, len(resource.KubernetesServiceAccountAdmissionRules))
	for k, r := range resource.KubernetesServiceAccountAdmissionRules {
		mKubernetesServiceAccountAdmissionRules[k] = BinaryauthorizationAlphaPolicyAdmissionRuleToProto(&r)
	}
	p.SetKubernetesServiceAccountAdmissionRules(mKubernetesServiceAccountAdmissionRules)
	mIstioServiceIdentityAdmissionRules := make(map[string]*alphapb.BinaryauthorizationAlphaPolicyAdmissionRule, len(resource.IstioServiceIdentityAdmissionRules))
	for k, r := range resource.IstioServiceIdentityAdmissionRules {
		mIstioServiceIdentityAdmissionRules[k] = BinaryauthorizationAlphaPolicyAdmissionRuleToProto(&r)
	}
	p.SetIstioServiceIdentityAdmissionRules(mIstioServiceIdentityAdmissionRules)

	return p
}

// applyPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) applyPolicy(ctx context.Context, c *alpha.Client, request *alphapb.ApplyBinaryauthorizationAlphaPolicyRequest) (*alphapb.BinaryauthorizationAlphaPolicy, error) {
	p := ProtoToPolicy(request.GetResource())
	res, err := c.ApplyPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := PolicyToProto(res)
	return r, nil
}

// applyBinaryauthorizationAlphaPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) ApplyBinaryauthorizationAlphaPolicy(ctx context.Context, request *alphapb.ApplyBinaryauthorizationAlphaPolicyRequest) (*alphapb.BinaryauthorizationAlphaPolicy, error) {
	cl, err := createConfigPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyPolicy(ctx, cl, request)
}

// DeletePolicy handles the gRPC request by passing it to the underlying Policy Delete() method.
func (s *PolicyServer) DeleteBinaryauthorizationAlphaPolicy(ctx context.Context, request *alphapb.DeleteBinaryauthorizationAlphaPolicyRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for Policy")

}

// ListBinaryauthorizationAlphaPolicy is a no-op method because Policy has no list method.
func (s *PolicyServer) ListBinaryauthorizationAlphaPolicy(_ context.Context, _ *alphapb.ListBinaryauthorizationAlphaPolicyRequest) (*alphapb.ListBinaryauthorizationAlphaPolicyResponse, error) {
	return nil, nil
}

func createConfigPolicy(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
