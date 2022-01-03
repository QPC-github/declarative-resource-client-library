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
// gen_go_data -package logging -var YAML_log_bucket blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/logging/log_bucket.yaml

package logging

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/logging/log_bucket.yaml
var YAML_log_bucket = []byte("info:\n  title: Logging/LogBucket\n  description: The Logging LogBucket resource\n  x-dcl-struct-name: LogBucket\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a LogBucket\n    parameters:\n    - name: LogBucket\n      required: true\n      description: A full instance of a LogBucket\n  apply:\n    description: The function used to apply information about a LogBucket\n    parameters:\n    - name: LogBucket\n      required: true\n      description: A full instance of a LogBucket\n  delete:\n    description: The function used to delete a LogBucket\n    parameters:\n    - name: LogBucket\n      required: true\n      description: A full instance of a LogBucket\n  deleteAll:\n    description: The function used to delete all LogBucket\n    parameters:\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: parent\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many LogBucket\n    parameters:\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: parent\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    LogBucket:\n      title: LogBucket\n      x-dcl-id: '{{parent}}/locations/{{location}}/buckets/{{name}}'\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - parent\n      - location\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The creation timestamp of the bucket. This is\n            not set for any of the default buckets.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Describes this bucket.\n        lifecycleState:\n          type: string\n          x-dcl-go-name: LifecycleState\n          x-dcl-go-type: LogBucketLifecycleStateEnum\n          readOnly: true\n          description: 'Output only. The bucket lifecycle state. Possible values:\n            LIFECYCLE_STATE_UNSPECIFIED, ACTIVE, DELETE_REQUESTED'\n          x-kubernetes-immutable: true\n          enum:\n          - LIFECYCLE_STATE_UNSPECIFIED\n          - ACTIVE\n          - DELETE_REQUESTED\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: 'The location of the resource. The supported locations are:\n            global, us-central1, us-east1, us-west1, asia-east1, europe-west1.'\n          x-kubernetes-immutable: true\n        locked:\n          type: boolean\n          x-dcl-go-name: Locked\n          description: Whether the bucket has been locked. The retention period on\n            a locked bucket may not be changed. Locked buckets may only be deleted\n            if they are empty.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'The resource name of the bucket. For example: \"projects/my-project-id/locations/my-location/buckets/my-bucket-id\"\n            The supported locations are: `global`, `us-central1`, `us-east1`, `us-west1`,\n            `asia-east1`, `europe-west1`. For the location of `global` it is unspecified\n            where logs are actually stored. Once a bucket has been created, the location\n            can not be changed.'\n          x-kubernetes-immutable: true\n        parent:\n          type: string\n          x-dcl-go-name: Parent\n          description: The parent of the resource.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/BillingAccount\n            field: name\n            parent: true\n          - resource: Cloudresourcemanager/Folder\n            field: name\n            parent: true\n          - resource: Cloudresourcemanager/Organization\n            field: name\n            parent: true\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        retentionDays:\n          type: integer\n          format: int64\n          x-dcl-go-name: RetentionDays\n          description: Logs will be retained by default for this amount of time, after\n            which they will automatically be deleted. The minimum retention period\n            is 1 day. If this value is set to zero at bucket creation time, the default\n            time of 30 days will be used.\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The last update timestamp of the bucket.\n          x-kubernetes-immutable: true\n")

// 4720 bytes
// MD5: 83026e2b05656742634ac521ad096477
