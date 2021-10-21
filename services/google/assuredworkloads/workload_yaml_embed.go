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
// gen_go_data -package assuredworkloads -var YAML_workload blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/assuredworkloads/workload.yaml

package assuredworkloads

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/assuredworkloads/workload.yaml
var YAML_workload = []byte("info:\n  title: AssuredWorkloads/Workload\n  description: The AssuredWorkloads Workload resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Workload\n    parameters:\n    - name: Workload\n      required: true\n      description: A full instance of a Workload\n  apply:\n    description: The function used to apply information about a Workload\n    parameters:\n    - name: Workload\n      required: true\n      description: A full instance of a Workload\n  delete:\n    description: The function used to delete a Workload\n    parameters:\n    - name: Workload\n      required: true\n      description: A full instance of a Workload\n  deleteAll:\n    description: The function used to delete all Workload\n    parameters:\n    - name: organization\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Workload\n    parameters:\n    - name: organization\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Workload:\n      title: Workload\n      x-dcl-id: organizations/{{organization}}/locations/{{location}}/workloads/{{name}}\n      x-dcl-uses-state-hint: true\n      x-dcl-parent-container: organization\n      x-dcl-labels: labels\n      type: object\n      required:\n      - displayName\n      - complianceRegime\n      - billingAccount\n      - organization\n      - location\n      properties:\n        billingAccount:\n          type: string\n          x-dcl-go-name: BillingAccount\n          description: Required. Input only. The billing account used for the resources\n            which are direct children of workload. This billing account is initially\n            associated with the resources created as part of Workload creation. After\n            the initial creation of these resources, the customer can change the assigned\n            billing account. The resource name has the form `billingAccounts/{billing_account_id}`.\n            For example, 'billingAccounts/012345-567890-ABCDEF`.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/BillingAccount\n            field: name\n          x-dcl-mutable-unreadable: true\n        complianceRegime:\n          type: string\n          x-dcl-go-name: ComplianceRegime\n          x-dcl-go-type: WorkloadComplianceRegimeEnum\n          description: 'Required. Immutable. Compliance Regime associated with this\n            workload. Possible values: COMPLIANCE_REGIME_UNSPECIFIED, IL4, CJIS, FEDRAMP_HIGH,\n            FEDRAMP_MODERATE, US_REGIONAL_ACCESS'\n          x-kubernetes-immutable: true\n          enum:\n          - COMPLIANCE_REGIME_UNSPECIFIED\n          - IL4\n          - CJIS\n          - FEDRAMP_HIGH\n          - FEDRAMP_MODERATE\n          - US_REGIONAL_ACCESS\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. Immutable. The Workload creation timestamp.\n          x-kubernetes-immutable: true\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: 'Required. The user-assigned display name of the Workload.\n            When present it must be between 4 to 30 characters. Allowed characters\n            are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example:\n            My Workload'\n        kmsSettings:\n          type: object\n          x-dcl-go-name: KmsSettings\n          x-dcl-go-type: WorkloadKmsSettings\n          description: Input only. Settings used to create a CMEK crypto key. When\n            set a project with a KMS CMEK key is provisioned. This field is mandatory\n            for a subset of Compliance Regimes.\n          x-kubernetes-immutable: true\n          x-dcl-mutable-unreadable: true\n          required:\n          - nextRotationTime\n          - rotationPeriod\n          properties:\n            nextRotationTime:\n              type: string\n              format: date-time\n              x-dcl-go-name: NextRotationTime\n              description: Required. Input only. Immutable. The time at which the\n                Key Management Service will automatically create a new version of\n                the crypto key and mark it as the primary.\n              x-kubernetes-immutable: true\n            rotationPeriod:\n              type: string\n              x-dcl-go-name: RotationPeriod\n              description: Required. Input only. Immutable. will be advanced by this\n                period when the Key Management Service automatically rotates a key.\n                Must be at least 24 hours and at most 876,000 hours.\n              x-kubernetes-immutable: true\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Optional. Labels applied to the workload.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Output only. The resource name of the workload.\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n        organization:\n          type: string\n          x-dcl-go-name: Organization\n          description: The organization for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Organization\n            field: name\n            parent: true\n        provisionedResourcesParent:\n          type: string\n          x-dcl-go-name: ProvisionedResourcesParent\n          description: 'Input only. The parent resource for the resources managed\n            by this Assured Workload. May be either an organization or a folder. Must\n            be the same or a child of the Workload parent. If not specified all resources\n            are created under the Workload parent. Formats: folders/{folder_id}, organizations/{organization_id}'\n          x-kubernetes-immutable: true\n          x-dcl-mutable-unreadable: true\n        resourceSettings:\n          type: array\n          x-dcl-go-name: ResourceSettings\n          description: Input only. Resource properties that are used to customize\n            workload resources. These properties (such as custom project id) will\n            be used to create workload resources if possible. This field is optional.\n          x-kubernetes-immutable: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: WorkloadResourceSettings\n            properties:\n              resourceId:\n                type: string\n                x-dcl-go-name: ResourceId\n                description: Resource identifier. For a project this represents project_number.\n                  If the project is already taken, the workload creation will fail.\n                x-kubernetes-immutable: true\n              resourceType:\n                type: string\n                x-dcl-go-name: ResourceType\n                x-dcl-go-type: WorkloadResourceSettingsResourceTypeEnum\n                description: 'Indicates the type of resource. This field should be\n                  specified to correspond the id to the right project type (CONSUMER_PROJECT\n                  or ENCRYPTION_KEYS_PROJECT) Possible values: RESOURCE_TYPE_UNSPECIFIED,\n                  CONSUMER_PROJECT, ENCRYPTION_KEYS_PROJECT, KEYRING, CONSUMER_FOLDER'\n                x-kubernetes-immutable: true\n                enum:\n                - RESOURCE_TYPE_UNSPECIFIED\n                - CONSUMER_PROJECT\n                - ENCRYPTION_KEYS_PROJECT\n                - KEYRING\n                - CONSUMER_FOLDER\n          x-dcl-mutable-unreadable: true\n        resources:\n          type: array\n          x-dcl-go-name: Resources\n          readOnly: true\n          description: Output only. The resources associated with this workload. These\n            resources will be created when creating the workload. If any of the projects\n            already exist, the workload creation will fail. Always read only.\n          x-kubernetes-immutable: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: WorkloadResources\n            properties:\n              resourceId:\n                type: integer\n                format: int64\n                x-dcl-go-name: ResourceId\n                description: Resource identifier. For a project this represents project_number.\n                x-kubernetes-immutable: true\n              resourceType:\n                type: string\n                x-dcl-go-name: ResourceType\n                x-dcl-go-type: WorkloadResourcesResourceTypeEnum\n                description: 'Indicates the type of resource. Possible values: RESOURCE_TYPE_UNSPECIFIED,\n                  CONSUMER_PROJECT, ENCRYPTION_KEYS_PROJECT, KEYRING, CONSUMER_FOLDER'\n                x-kubernetes-immutable: true\n                enum:\n                - RESOURCE_TYPE_UNSPECIFIED\n                - CONSUMER_PROJECT\n                - ENCRYPTION_KEYS_PROJECT\n                - KEYRING\n                - CONSUMER_FOLDER\n")

// 9450 bytes
// MD5: b583ef7959167b8ef0b893972bc740ca
