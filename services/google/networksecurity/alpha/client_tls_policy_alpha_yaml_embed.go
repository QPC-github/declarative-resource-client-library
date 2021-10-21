// Copyright 2021 Google LLC. All Rights Reserved.
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
// GENERATED BY gen_go_data.go
// gen_go_data -package alpha -var YAML_client_tls_policy blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networksecurity/alpha/client_tls_policy.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networksecurity/alpha/client_tls_policy.yaml
var YAML_client_tls_policy = []byte("info:\n  title: NetworkSecurity/ClientTlsPolicy\n  description: The NetworkSecurity ClientTlsPolicy resource\n  x-dcl-has-iam: true\npaths:\n  get:\n    description: The function used to get information about a ClientTlsPolicy\n    parameters:\n    - name: ClientTlsPolicy\n      required: true\n      description: A full instance of a ClientTlsPolicy\n  apply:\n    description: The function used to apply information about a ClientTlsPolicy\n    parameters:\n    - name: ClientTlsPolicy\n      required: true\n      description: A full instance of a ClientTlsPolicy\n  delete:\n    description: The function used to delete a ClientTlsPolicy\n    parameters:\n    - name: ClientTlsPolicy\n      required: true\n      description: A full instance of a ClientTlsPolicy\n  deleteAll:\n    description: The function used to delete all ClientTlsPolicy\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many ClientTlsPolicy\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    ClientTlsPolicy:\n      title: ClientTlsPolicy\n      x-dcl-id: projects/{{project}}/locations/{{location}}/clientTlsPolicies/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      type: object\n      required:\n      - name\n      - project\n      - location\n      properties:\n        clientCertificate:\n          type: object\n          x-dcl-go-name: ClientCertificate\n          x-dcl-go-type: ClientTlsPolicyClientCertificate\n          description: Optional. Defines a mechanism to provision client identity\n            (public and private keys) for peer to peer authentication. The presence\n            of this dictates mTLS.\n          properties:\n            certificateProviderInstance:\n              type: object\n              x-dcl-go-name: CertificateProviderInstance\n              x-dcl-go-type: ClientTlsPolicyClientCertificateCertificateProviderInstance\n              description: The certificate provider instance specification that will\n                be passed to the data plane, which will be used to load necessary\n                credential information.\n              required:\n              - pluginInstance\n              properties:\n                pluginInstance:\n                  type: string\n                  x-dcl-go-name: PluginInstance\n                  description: Required. Plugin instance name, used to locate and\n                    load CertificateProvider instance configuration. Set to \"google_cloud_private_spiffe\"\n                    to use Certificate Authority Service certificate provider instance.\n            grpcEndpoint:\n              type: object\n              x-dcl-go-name: GrpcEndpoint\n              x-dcl-go-type: ClientTlsPolicyClientCertificateGrpcEndpoint\n              description: gRPC specific configuration to access the gRPC server to\n                obtain the cert and private key.\n              required:\n              - targetUri\n              properties:\n                targetUri:\n                  type: string\n                  x-dcl-go-name: TargetUri\n                  description: Required. The target URI of the gRPC endpoint. Only\n                    UDS path is supported, and should start with “unix:”.\n            localFilepath:\n              type: object\n              x-dcl-go-name: LocalFilepath\n              x-dcl-go-type: ClientTlsPolicyClientCertificateLocalFilepath\n              description: Obtain certificates and private key from a locally mounted\n                filesystem path.\n              required:\n              - certificatePath\n              - privateKeyPath\n              properties:\n                certificatePath:\n                  type: string\n                  x-dcl-go-name: CertificatePath\n                  description: Required. The path to the file that has the certificate\n                    containing public key.\n                privateKeyPath:\n                  type: string\n                  x-dcl-go-name: PrivateKeyPath\n                  description: Required. The path to the file that has the private\n                    key.\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The timestamp when the resource was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. Free-text description of the resource.\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Optional. Set of label tags associated with the resource.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Required. Name of the ClientTlsPolicy resource.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        serverValidationCa:\n          type: array\n          x-dcl-go-name: ServerValidationCa\n          description: Required. Defines the mechanism to obtain the Certificate Authority\n            certificate to validate the server certificate.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: ClientTlsPolicyServerValidationCa\n            properties:\n              caCertPath:\n                type: string\n                x-dcl-go-name: CaCertPath\n                description: The path to the file holding the CA certificate to validate\n                  the client or server certificate.\n              certificateProviderInstance:\n                type: object\n                x-dcl-go-name: CertificateProviderInstance\n                x-dcl-go-type: ClientTlsPolicyServerValidationCaCertificateProviderInstance\n                description: The certificate provider instance specification that\n                  will be passed to the data plane, which will be used to load necessary\n                  credential information.\n                required:\n                - pluginInstance\n                properties:\n                  pluginInstance:\n                    type: string\n                    x-dcl-go-name: PluginInstance\n                    description: Required. Plugin instance name, used to locate and\n                      load CertificateProvider instance configuration. Set to \"google_cloud_private_spiffe\"\n                      to use Certificate Authority Service certificate provider instance.\n              grpcEndpoint:\n                type: object\n                x-dcl-go-name: GrpcEndpoint\n                x-dcl-go-type: ClientTlsPolicyServerValidationCaGrpcEndpoint\n                description: gRPC specific configuration to access the gRPC server\n                  to obtain the CA certificate.\n                required:\n                - targetUri\n                properties:\n                  targetUri:\n                    type: string\n                    x-dcl-go-name: TargetUri\n                    description: Required. The target URI of the gRPC endpoint. Only\n                      UDS path is supported, and should start with “unix:”.\n        sni:\n          type: string\n          x-dcl-go-name: Sni\n          description: 'Optional. Server Name Indication string to present to the\n            server during TLS handshake. E.g: \"secure.example.com\".'\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The timestamp when the resource was updated.\n          x-kubernetes-immutable: true\n")

// 8323 bytes
// MD5: a336cce409c19c70b731e7257af53e63
