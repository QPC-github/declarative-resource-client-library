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
// gen_go_data -package logging -var YAML_log_view blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/logging/log_view.yaml

package logging

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/logging/log_view.yaml
var YAML_log_view = []byte("info:\n  title: Logging/LogView\n  description: The Logging LogView resource\n  x-dcl-struct-name: LogView\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a LogView\n    parameters:\n    - name: LogView\n      required: true\n      description: A full instance of a LogView\n  apply:\n    description: The function used to apply information about a LogView\n    parameters:\n    - name: LogView\n      required: true\n      description: A full instance of a LogView\n  delete:\n    description: The function used to delete a LogView\n    parameters:\n    - name: LogView\n      required: true\n      description: A full instance of a LogView\n  deleteAll:\n    description: The function used to delete all LogView\n    parameters:\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: bucket\n      required: true\n      schema:\n        type: string\n    - name: parent\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many LogView\n    parameters:\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: bucket\n      required: true\n      schema:\n        type: string\n    - name: parent\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    LogView:\n      title: LogView\n      x-dcl-id: '{{parent}}/locations/{{location}}/buckets/{{bucket}}/views/{{name}}'\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - parent\n      - location\n      - bucket\n      properties:\n        bucket:\n          type: string\n          x-dcl-go-name: Bucket\n          description: The bucket of the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Logging/LogBucket\n            field: name\n            parent: true\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The creation timestamp of the view.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Describes this view.\n        filter:\n          type: string\n          x-dcl-go-name: Filter\n          description: 'Filter that restricts which log entries in a bucket are visible\n            in this view. Filters are restricted to be a logical AND of ==/!= of any\n            of the following: - originating project/folder/organization/billing account.\n            - resource type - log id For example: SOURCE(\"projects/myproject\") AND\n            resource.type = \"gce_instance\" AND LOG_ID(\"stdout\")'\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: 'The location of the resource. The supported locations are:\n            global, us-central1, us-east1, us-west1, asia-east1, europe-west1.'\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'The resource name of the view. For example: `projects/my-project/locations/global/buckets/my-bucket/views/my-view`'\n          x-kubernetes-immutable: true\n        parent:\n          type: string\n          x-dcl-go-name: Parent\n          description: The parent of the resource.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/BillingAccount\n            field: name\n            parent: true\n          - resource: Cloudresourcemanager/Folder\n            field: name\n            parent: true\n          - resource: Cloudresourcemanager/Organization\n            field: name\n            parent: true\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The last update timestamp of the view.\n          x-kubernetes-immutable: true\n")

// 4088 bytes
// MD5: 52ed710d2897d7d6a2615becef5b0716
