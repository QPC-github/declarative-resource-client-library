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
// gen_go_data -package iam -var YAML_service_account blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/iam/service_account.yaml

package iam

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/iam/service_account.yaml
var YAML_service_account = []byte("info:\n  title: Iam/ServiceAccount\n  description: The Iam ServiceAccount resource\n  x-dcl-has-iam: true\npaths:\n  get:\n    description: The function used to get information about a ServiceAccount\n    parameters:\n    - name: ServiceAccount\n      required: true\n      description: A full instance of a ServiceAccount\n  apply:\n    description: The function used to apply information about a ServiceAccount\n    parameters:\n    - name: ServiceAccount\n      required: true\n      description: A full instance of a ServiceAccount\n  delete:\n    description: The function used to delete a ServiceAccount\n    parameters:\n    - name: ServiceAccount\n      required: true\n      description: A full instance of a ServiceAccount\n  deleteAll:\n    description: The function used to delete all ServiceAccount\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many ServiceAccount\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    ServiceAccount:\n      title: ServiceAccount\n      x-dcl-id: projects/{{project}}/serviceAccounts/{{name}}@{{project}}.iam.gserviceaccount.com\n      x-dcl-parent-container: project\n      type: object\n      properties:\n        actasResources:\n          type: object\n          x-dcl-go-name: ActasResources\n          x-dcl-go-type: ServiceAccountActasResources\n          description: Optional.\n          x-kubernetes-immutable: true\n          properties:\n            resources:\n              type: array\n              x-dcl-go-name: Resources\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: ServiceAccountActasResourcesResources\n                properties:\n                  fullResourceName:\n                    type: string\n                    x-dcl-go-name: FullResourceName\n                    x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. A user-specified, human-readable description of the\n            service account. The maximum length is 256 UTF-8 bytes.\n        disabled:\n          type: boolean\n          x-dcl-go-name: Disabled\n          readOnly: true\n          description: Output only. Whether the service account is disabled.\n          x-kubernetes-immutable: true\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: Optional. A user-specified, human-readable name for the service\n            account. The maximum length is 100 UTF-8 bytes.\n        email:\n          type: string\n          x-dcl-go-name: Email\n          readOnly: true\n          description: Output only. The email address of the service account.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'The resource name of the service account. Use one of the following\n            formats: * `projects/{PROJECT_ID}/serviceAccounts/{EMAIL_ADDRESS}` * `projects/{PROJECT_ID}/serviceAccounts/{UNIQUE_ID}`\n            As an alternative, you can use the `-` wildcard character instead of the\n            project ID: * `projects/-/serviceAccounts/{EMAIL_ADDRESS}` * `projects/-/serviceAccounts/{UNIQUE_ID}`\n            When possible, avoid using the `-` wildcard character, because it can\n            cause response messages to contain misleading error codes. For example,\n            if you try to get the service account `projects/-/serviceAccounts/fake@example.com`,\n            which does not exist, the response contains an HTTP `403 Forbidden` error\n            instead of a `404 Not Found` error.'\n          x-kubernetes-immutable: true\n        oauth2ClientId:\n          type: string\n          x-dcl-go-name: OAuth2ClientId\n          readOnly: true\n          description: Output only. The OAuth 2.0 client ID for the service account.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          readOnly: true\n          description: Output only. The ID of the project that owns the service account.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        uniqueId:\n          type: string\n          x-dcl-go-name: UniqueId\n          readOnly: true\n          description: Output only. The unique, stable numeric ID for the service\n            account. Each service account retains its unique ID even if you delete\n            the service account. For example, if you delete a service account, then\n            create a new service account with the same name, the new service account\n            has a different unique ID than the deleted service account.\n          x-kubernetes-immutable: true\n")

// 5013 bytes
// MD5: cc8f4448004159ac6ebff498cc1c8554
