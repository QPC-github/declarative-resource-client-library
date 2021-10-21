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
// gen_go_data -package beta -var YAML_project blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudresourcemanager/beta/project.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudresourcemanager/beta/project.yaml
var YAML_project = []byte("info:\n  title: CloudResourceManager/Project\n  description: The CloudResourceManager Project resource\n  x-dcl-has-iam: true\npaths:\n  get:\n    description: The function used to get information about a Project\n    parameters:\n    - name: Project\n      required: true\n      description: A full instance of a Project\n  apply:\n    description: The function used to apply information about a Project\n    parameters:\n    - name: Project\n      required: true\n      description: A full instance of a Project\n  delete:\n    description: The function used to delete a Project\n    parameters:\n    - name: Project\n      required: true\n      description: A full instance of a Project\n  deleteAll:\n    description: The function used to delete all Project\n    parameters: []\n  list:\n    description: The function used to list information about many Project\n    parameters: []\ncomponents:\n  schemas:\n    Project:\n      title: Project\n      x-dcl-id: projects/{{name}}\n      x-dcl-labels: labels\n      type: object\n      properties:\n        displayname:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: 'The optional user-assigned display name of the Project. When\n            present it must be between 4 to 30 characters. Allowed characters are:\n            lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote,\n            space, and exclamation point. Example: `My Project` Read-write.'\n          x-kubernetes-immutable: true\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: User-specified labels.\n        lifecycleState:\n          type: string\n          x-dcl-go-name: LifecycleState\n          x-dcl-go-type: ProjectLifecycleStateEnum\n          readOnly: true\n          description: 'The Project lifecycle state. Read-only. Possible values: LIFECYCLE_STATE_UNSPECIFIED,\n            ACTIVE, DELETE_REQUESTED, DELETE_IN_PROGRESS'\n          x-kubernetes-immutable: true\n          enum:\n          - LIFECYCLE_STATE_UNSPECIFIED\n          - ACTIVE\n          - DELETE_REQUESTED\n          - DELETE_IN_PROGRESS\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'The unique, user-assigned ID of the Project. It must be 6\n            to 30 lowercase letters, digits, or hyphens. It must start with a letter.\n            Trailing hyphens are prohibited. Example: `tokyo-rain-123` Read-only after\n            creation.'\n          x-kubernetes-immutable: true\n        parent:\n          type: string\n          x-dcl-go-name: Parent\n          description: 'An optional reference to a parent Resource. Supported values\n            include organizations/<org_id> and folders/<folder_id>. Once set, the\n            parent cannot be cleared. The `parent` can be set on creation or using\n            the `UpdateProject` method; the end user must have the `resourcemanager.projects.create`\n            permission on the parent. Read-write. '\n          x-kubernetes-immutable: true\n        projectNumber:\n          type: integer\n          format: int64\n          x-dcl-go-name: ProjectNumber\n          readOnly: true\n          description: 'The number uniquely identifying the project. Example: `415104041262`\n            Read-only. '\n          x-kubernetes-immutable: true\n")

// 3336 bytes
// MD5: da524ba682125fb35afb0613086e6063
