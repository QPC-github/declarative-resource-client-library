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
// GENERATED BY gen_go_data.go
// gen_go_data -package alpha -var YAML_role blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/iam/alpha/role.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/iam/alpha/role.yaml
var YAML_role = []byte("info:\n  title: Iam/Role\n  description: The Iam Role resource\n  x-dcl-struct-name: Role\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Role\n    parameters:\n    - name: Role\n      required: true\n      description: A full instance of a Role\n  apply:\n    description: The function used to apply information about a Role\n    parameters:\n    - name: Role\n      required: true\n      description: A full instance of a Role\n  delete:\n    description: The function used to delete a Role\n    parameters:\n    - name: Role\n      required: true\n      description: A full instance of a Role\n  deleteAll:\n    description: The function used to delete all Role\n    parameters:\n    - name: parent\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Role\n    parameters:\n    - name: parent\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Role:\n      title: Role\n      x-dcl-id: '{{parent}}/roles/{{name}}'\n      x-dcl-has-iam: false\n      type: object\n      properties:\n        deleted:\n          type: boolean\n          x-dcl-go-name: Deleted\n          description: The current deleted state of the role. This field is read only.\n            It will be ignored in calls to CreateRole and UpdateRole.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. A human-readable description for the role.\n          x-kubernetes-immutable: true\n        etag:\n          type: string\n          x-dcl-go-name: Etag\n          description: Used to perform a consistent read-modify-write.\n          x-kubernetes-immutable: true\n        groupName:\n          type: string\n          x-dcl-go-name: GroupName\n          x-kubernetes-immutable: true\n        groupTitle:\n          type: string\n          x-dcl-go-name: GroupTitle\n          x-kubernetes-immutable: true\n        includedPermissions:\n          type: array\n          x-dcl-go-name: IncludedPermissions\n          description: The names of the permissions this role grants when bound in\n            an IAM policy.\n          x-kubernetes-immutable: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n        includedRoles:\n          type: array\n          x-dcl-go-name: IncludedRoles\n          x-kubernetes-immutable: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n        lifecyclePhase:\n          type: string\n          x-dcl-go-name: LifecyclePhase\n          x-kubernetes-immutable: true\n        localizedValues:\n          type: object\n          x-dcl-go-name: LocalizedValues\n          x-dcl-go-type: RoleLocalizedValues\n          x-kubernetes-immutable: true\n          properties:\n            localizedDescription:\n              type: string\n              x-dcl-go-name: LocalizedDescription\n              description: Will be English by default or if an error occurred during\n                translation.\n              x-kubernetes-immutable: true\n            localizedTitle:\n              type: string\n              x-dcl-go-name: LocalizedTitle\n              description: Will be English by default or if an error occurred during\n                translation.\n              x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The name of the role. When Role is used in CreateRole, the\n            role name must not be set. When Role is used in output and other input\n            such as UpdateRole, the role name is the complete path, e.g., roles/logging.viewer\n            for predefined roles and organizations/{ORGANIZATION_ID}/roles/logging.viewer\n            for custom roles.\n          x-kubernetes-immutable: true\n        parent:\n          type: string\n          x-dcl-go-name: Parent\n          description: 'The parent parameter''s value depends on the target resource\n            for the request, namely projects or organizations. Each resource type''s\n            parent value format is described below: projects.roles.create(): projects/{PROJECT_ID}.\n            This method creates project-level custom roles. Example request URL: https://iam.googleapis.com/v1/projects/{PROJECT_ID}/roles\n            organizations.roles.create(): organizations/{ORGANIZATION_ID}. This method\n            creates organization-level custom roles. Example request URL: https://iam.googleapis.com/v1/organizations/{ORGANIZATION_ID}/roles\n            Note: Wildcard (*) values are invalid; you must specify a complete project\n            ID or organization ID. Authorization requires the following IAM permission\n            on the specified resource parent: iam.roles.create'\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n          - resource: Cloudresourcemanager/Organization\n            field: name\n            parent: true\n        stage:\n          type: string\n          x-dcl-go-name: Stage\n          x-dcl-go-type: RoleStageEnum\n          description: The current launch stage of the role. If the `ALPHA` launch\n            stage has been selected for a role, the `stage` field will not be included\n            in the returned definition for the role.\n          x-kubernetes-immutable: true\n          enum:\n          - ALPHA\n          - BETA\n          - GA\n          - DEPRECATED\n          - DISABLED\n          - EAP\n        title:\n          type: string\n          x-dcl-go-name: Title\n          description: Optional. A human-readable title for the role. Typically this\n            is limited to 100 UTF-8 bytes.\n          x-kubernetes-immutable: true\n")

// 5924 bytes
// MD5: b9caa6a2a87be18060196b3a235a0618
