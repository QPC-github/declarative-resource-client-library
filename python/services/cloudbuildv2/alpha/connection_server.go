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
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudbuildv2/alpha/cloudbuildv2_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuildv2/alpha"
)

// ConnectionServer implements the gRPC interface for Connection.
type ConnectionServer struct{}

// ProtoToConnectionInstallationStateStageEnum converts a ConnectionInstallationStateStageEnum enum from its proto representation.
func ProtoToCloudbuildv2AlphaConnectionInstallationStateStageEnum(e alphapb.Cloudbuildv2AlphaConnectionInstallationStateStageEnum) *alpha.ConnectionInstallationStateStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Cloudbuildv2AlphaConnectionInstallationStateStageEnum_name[int32(e)]; ok {
		e := alpha.ConnectionInstallationStateStageEnum(n[len("Cloudbuildv2AlphaConnectionInstallationStateStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToConnectionGithubConfig converts a ConnectionGithubConfig object from its proto representation.
func ProtoToCloudbuildv2AlphaConnectionGithubConfig(p *alphapb.Cloudbuildv2AlphaConnectionGithubConfig) *alpha.ConnectionGithubConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ConnectionGithubConfig{
		AuthorizerCredential: ProtoToCloudbuildv2AlphaConnectionGithubConfigAuthorizerCredential(p.GetAuthorizerCredential()),
		AppInstallationId:    dcl.Int64OrNil(p.GetAppInstallationId()),
	}
	return obj
}

// ProtoToConnectionGithubConfigAuthorizerCredential converts a ConnectionGithubConfigAuthorizerCredential object from its proto representation.
func ProtoToCloudbuildv2AlphaConnectionGithubConfigAuthorizerCredential(p *alphapb.Cloudbuildv2AlphaConnectionGithubConfigAuthorizerCredential) *alpha.ConnectionGithubConfigAuthorizerCredential {
	if p == nil {
		return nil
	}
	obj := &alpha.ConnectionGithubConfigAuthorizerCredential{
		OAuthTokenSecretVersion: dcl.StringOrNil(p.GetOauthTokenSecretVersion()),
		Username:                dcl.StringOrNil(p.GetUsername()),
	}
	return obj
}

// ProtoToConnectionGithubEnterpriseConfig converts a ConnectionGithubEnterpriseConfig object from its proto representation.
func ProtoToCloudbuildv2AlphaConnectionGithubEnterpriseConfig(p *alphapb.Cloudbuildv2AlphaConnectionGithubEnterpriseConfig) *alpha.ConnectionGithubEnterpriseConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ConnectionGithubEnterpriseConfig{
		HostUri:                    dcl.StringOrNil(p.GetHostUri()),
		AppId:                      dcl.Int64OrNil(p.GetAppId()),
		AppSlug:                    dcl.StringOrNil(p.GetAppSlug()),
		PrivateKeySecretVersion:    dcl.StringOrNil(p.GetPrivateKeySecretVersion()),
		WebhookSecretSecretVersion: dcl.StringOrNil(p.GetWebhookSecretSecretVersion()),
		AppInstallationId:          dcl.Int64OrNil(p.GetAppInstallationId()),
		ServiceDirectoryConfig:     ProtoToCloudbuildv2AlphaConnectionGithubEnterpriseConfigServiceDirectoryConfig(p.GetServiceDirectoryConfig()),
		SslCa:                      dcl.StringOrNil(p.GetSslCa()),
	}
	return obj
}

// ProtoToConnectionGithubEnterpriseConfigServiceDirectoryConfig converts a ConnectionGithubEnterpriseConfigServiceDirectoryConfig object from its proto representation.
func ProtoToCloudbuildv2AlphaConnectionGithubEnterpriseConfigServiceDirectoryConfig(p *alphapb.Cloudbuildv2AlphaConnectionGithubEnterpriseConfigServiceDirectoryConfig) *alpha.ConnectionGithubEnterpriseConfigServiceDirectoryConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ConnectionGithubEnterpriseConfigServiceDirectoryConfig{
		Service: dcl.StringOrNil(p.GetService()),
	}
	return obj
}

// ProtoToConnectionInstallationState converts a ConnectionInstallationState object from its proto representation.
func ProtoToCloudbuildv2AlphaConnectionInstallationState(p *alphapb.Cloudbuildv2AlphaConnectionInstallationState) *alpha.ConnectionInstallationState {
	if p == nil {
		return nil
	}
	obj := &alpha.ConnectionInstallationState{
		Stage:     ProtoToCloudbuildv2AlphaConnectionInstallationStateStageEnum(p.GetStage()),
		Message:   dcl.StringOrNil(p.GetMessage()),
		ActionUri: dcl.StringOrNil(p.GetActionUri()),
	}
	return obj
}

// ProtoToConnection converts a Connection resource from its proto representation.
func ProtoToConnection(p *alphapb.Cloudbuildv2AlphaConnection) *alpha.Connection {
	obj := &alpha.Connection{
		Name:                   dcl.StringOrNil(p.GetName()),
		CreateTime:             dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:             dcl.StringOrNil(p.GetUpdateTime()),
		GithubConfig:           ProtoToCloudbuildv2AlphaConnectionGithubConfig(p.GetGithubConfig()),
		GithubEnterpriseConfig: ProtoToCloudbuildv2AlphaConnectionGithubEnterpriseConfig(p.GetGithubEnterpriseConfig()),
		InstallationState:      ProtoToCloudbuildv2AlphaConnectionInstallationState(p.GetInstallationState()),
		Disabled:               dcl.Bool(p.GetDisabled()),
		Reconciling:            dcl.Bool(p.GetReconciling()),
		Etag:                   dcl.StringOrNil(p.GetEtag()),
		Project:                dcl.StringOrNil(p.GetProject()),
		Location:               dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ConnectionInstallationStateStageEnumToProto converts a ConnectionInstallationStateStageEnum enum to its proto representation.
func Cloudbuildv2AlphaConnectionInstallationStateStageEnumToProto(e *alpha.ConnectionInstallationStateStageEnum) alphapb.Cloudbuildv2AlphaConnectionInstallationStateStageEnum {
	if e == nil {
		return alphapb.Cloudbuildv2AlphaConnectionInstallationStateStageEnum(0)
	}
	if v, ok := alphapb.Cloudbuildv2AlphaConnectionInstallationStateStageEnum_value["ConnectionInstallationStateStageEnum"+string(*e)]; ok {
		return alphapb.Cloudbuildv2AlphaConnectionInstallationStateStageEnum(v)
	}
	return alphapb.Cloudbuildv2AlphaConnectionInstallationStateStageEnum(0)
}

// ConnectionGithubConfigToProto converts a ConnectionGithubConfig object to its proto representation.
func Cloudbuildv2AlphaConnectionGithubConfigToProto(o *alpha.ConnectionGithubConfig) *alphapb.Cloudbuildv2AlphaConnectionGithubConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Cloudbuildv2AlphaConnectionGithubConfig{}
	p.SetAuthorizerCredential(Cloudbuildv2AlphaConnectionGithubConfigAuthorizerCredentialToProto(o.AuthorizerCredential))
	p.SetAppInstallationId(dcl.ValueOrEmptyInt64(o.AppInstallationId))
	return p
}

// ConnectionGithubConfigAuthorizerCredentialToProto converts a ConnectionGithubConfigAuthorizerCredential object to its proto representation.
func Cloudbuildv2AlphaConnectionGithubConfigAuthorizerCredentialToProto(o *alpha.ConnectionGithubConfigAuthorizerCredential) *alphapb.Cloudbuildv2AlphaConnectionGithubConfigAuthorizerCredential {
	if o == nil {
		return nil
	}
	p := &alphapb.Cloudbuildv2AlphaConnectionGithubConfigAuthorizerCredential{}
	p.SetOauthTokenSecretVersion(dcl.ValueOrEmptyString(o.OAuthTokenSecretVersion))
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	return p
}

// ConnectionGithubEnterpriseConfigToProto converts a ConnectionGithubEnterpriseConfig object to its proto representation.
func Cloudbuildv2AlphaConnectionGithubEnterpriseConfigToProto(o *alpha.ConnectionGithubEnterpriseConfig) *alphapb.Cloudbuildv2AlphaConnectionGithubEnterpriseConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Cloudbuildv2AlphaConnectionGithubEnterpriseConfig{}
	p.SetHostUri(dcl.ValueOrEmptyString(o.HostUri))
	p.SetAppId(dcl.ValueOrEmptyInt64(o.AppId))
	p.SetAppSlug(dcl.ValueOrEmptyString(o.AppSlug))
	p.SetPrivateKeySecretVersion(dcl.ValueOrEmptyString(o.PrivateKeySecretVersion))
	p.SetWebhookSecretSecretVersion(dcl.ValueOrEmptyString(o.WebhookSecretSecretVersion))
	p.SetAppInstallationId(dcl.ValueOrEmptyInt64(o.AppInstallationId))
	p.SetServiceDirectoryConfig(Cloudbuildv2AlphaConnectionGithubEnterpriseConfigServiceDirectoryConfigToProto(o.ServiceDirectoryConfig))
	p.SetSslCa(dcl.ValueOrEmptyString(o.SslCa))
	return p
}

// ConnectionGithubEnterpriseConfigServiceDirectoryConfigToProto converts a ConnectionGithubEnterpriseConfigServiceDirectoryConfig object to its proto representation.
func Cloudbuildv2AlphaConnectionGithubEnterpriseConfigServiceDirectoryConfigToProto(o *alpha.ConnectionGithubEnterpriseConfigServiceDirectoryConfig) *alphapb.Cloudbuildv2AlphaConnectionGithubEnterpriseConfigServiceDirectoryConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Cloudbuildv2AlphaConnectionGithubEnterpriseConfigServiceDirectoryConfig{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	return p
}

// ConnectionInstallationStateToProto converts a ConnectionInstallationState object to its proto representation.
func Cloudbuildv2AlphaConnectionInstallationStateToProto(o *alpha.ConnectionInstallationState) *alphapb.Cloudbuildv2AlphaConnectionInstallationState {
	if o == nil {
		return nil
	}
	p := &alphapb.Cloudbuildv2AlphaConnectionInstallationState{}
	p.SetStage(Cloudbuildv2AlphaConnectionInstallationStateStageEnumToProto(o.Stage))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	p.SetActionUri(dcl.ValueOrEmptyString(o.ActionUri))
	return p
}

// ConnectionToProto converts a Connection resource to its proto representation.
func ConnectionToProto(resource *alpha.Connection) *alphapb.Cloudbuildv2AlphaConnection {
	p := &alphapb.Cloudbuildv2AlphaConnection{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetGithubConfig(Cloudbuildv2AlphaConnectionGithubConfigToProto(resource.GithubConfig))
	p.SetGithubEnterpriseConfig(Cloudbuildv2AlphaConnectionGithubEnterpriseConfigToProto(resource.GithubEnterpriseConfig))
	p.SetInstallationState(Cloudbuildv2AlphaConnectionInstallationStateToProto(resource.InstallationState))
	p.SetDisabled(dcl.ValueOrEmptyBool(resource.Disabled))
	p.SetReconciling(dcl.ValueOrEmptyBool(resource.Reconciling))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mAnnotations := make(map[string]string, len(resource.Annotations))
	for k, r := range resource.Annotations {
		mAnnotations[k] = r
	}
	p.SetAnnotations(mAnnotations)

	return p
}

// applyConnection handles the gRPC request by passing it to the underlying Connection Apply() method.
func (s *ConnectionServer) applyConnection(ctx context.Context, c *alpha.Client, request *alphapb.ApplyCloudbuildv2AlphaConnectionRequest) (*alphapb.Cloudbuildv2AlphaConnection, error) {
	p := ProtoToConnection(request.GetResource())
	res, err := c.ApplyConnection(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ConnectionToProto(res)
	return r, nil
}

// applyCloudbuildv2AlphaConnection handles the gRPC request by passing it to the underlying Connection Apply() method.
func (s *ConnectionServer) ApplyCloudbuildv2AlphaConnection(ctx context.Context, request *alphapb.ApplyCloudbuildv2AlphaConnectionRequest) (*alphapb.Cloudbuildv2AlphaConnection, error) {
	cl, err := createConfigConnection(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyConnection(ctx, cl, request)
}

// DeleteConnection handles the gRPC request by passing it to the underlying Connection Delete() method.
func (s *ConnectionServer) DeleteCloudbuildv2AlphaConnection(ctx context.Context, request *alphapb.DeleteCloudbuildv2AlphaConnectionRequest) (*emptypb.Empty, error) {

	cl, err := createConfigConnection(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteConnection(ctx, ProtoToConnection(request.GetResource()))

}

// ListCloudbuildv2AlphaConnection handles the gRPC request by passing it to the underlying ConnectionList() method.
func (s *ConnectionServer) ListCloudbuildv2AlphaConnection(ctx context.Context, request *alphapb.ListCloudbuildv2AlphaConnectionRequest) (*alphapb.ListCloudbuildv2AlphaConnectionResponse, error) {
	cl, err := createConfigConnection(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListConnection(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.Cloudbuildv2AlphaConnection
	for _, r := range resources.Items {
		rp := ConnectionToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListCloudbuildv2AlphaConnectionResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigConnection(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
