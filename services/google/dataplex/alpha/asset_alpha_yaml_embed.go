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
// gen_go_data -package alpha -var YAML_asset blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/dataplex/alpha/asset.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/dataplex/alpha/asset.yaml
var YAML_asset = []byte("info:\n  title: Dataplex/Asset\n  description: The Dataplex Asset resource\n  x-dcl-struct-name: Asset\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Asset\n    parameters:\n    - name: Asset\n      required: true\n      description: A full instance of a Asset\n  apply:\n    description: The function used to apply information about a Asset\n    parameters:\n    - name: Asset\n      required: true\n      description: A full instance of a Asset\n  delete:\n    description: The function used to delete a Asset\n    parameters:\n    - name: Asset\n      required: true\n      description: A full instance of a Asset\n  deleteAll:\n    description: The function used to delete all Asset\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: zone\n      required: true\n      schema:\n        type: string\n    - name: lake\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Asset\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: zone\n      required: true\n      schema:\n        type: string\n    - name: lake\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Asset:\n      title: Asset\n      x-dcl-id: projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{zone}}/assets/{{name}}\n      x-dcl-locations:\n      - zone\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - resourceSpec\n      - discoverySpec\n      - project\n      - location\n      - lake\n      - zone\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The time when the asset was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. Description of the asset.\n        discoverySpec:\n          type: object\n          x-dcl-go-name: DiscoverySpec\n          x-dcl-go-type: AssetDiscoverySpec\n          description: Optional. Specification of the discovery feature applied to\n            data referenced by this asset. When this spec is left unset, the asset\n            will use the spec set on the parent zone.\n          properties:\n            csvOptions:\n              type: object\n              x-dcl-go-name: CsvOptions\n              x-dcl-go-type: AssetDiscoverySpecCsvOptions\n              description: Optional. Configuration for CSV data.\n              properties:\n                delimiter:\n                  type: string\n                  x-dcl-go-name: Delimiter\n                  description: Optional. The delimiter being used to separate values.\n                    This defaults to ','.\n                disableTypeInference:\n                  type: boolean\n                  x-dcl-go-name: DisableTypeInference\n                  description: Optional. Whether to disable the inference of data\n                    type for CSV data. If true, all columns will be registered as\n                    strings.\n                encoding:\n                  type: string\n                  x-dcl-go-name: Encoding\n                  description: Optional. The character encoding of the data. The default\n                    is UTF-8.\n                headerRows:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: HeaderRows\n                  description: Optional. The number of rows to interpret as header\n                    rows that should be skipped when reading data rows.\n            enabled:\n              type: boolean\n              x-dcl-go-name: Enabled\n              description: Optional. Whether discovery is enabled.\n            excludePatterns:\n              type: array\n              x-dcl-go-name: ExcludePatterns\n              description: Optional. The list of patterns to apply for selecting data\n                to exclude during discovery. For Cloud Storage bucket assets, these\n                are interpreted as glob patterns used to match object names. For BigQuery\n                dataset assets, these are interpreted as patterns to match table names.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            includePatterns:\n              type: array\n              x-dcl-go-name: IncludePatterns\n              description: Optional. The list of patterns to apply for selecting data\n                to include during discovery if only a subset of the data should considered.\n                For Cloud Storage bucket assets, these are interpreted as glob patterns\n                used to match object names. For BigQuery dataset assets, these are\n                interpreted as patterns to match table names.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            jsonOptions:\n              type: object\n              x-dcl-go-name: JsonOptions\n              x-dcl-go-type: AssetDiscoverySpecJsonOptions\n              description: Optional. Configuration for Json data.\n              properties:\n                disableTypeInference:\n                  type: boolean\n                  x-dcl-go-name: DisableTypeInference\n                  description: Optional. Whether to disable the inference of data\n                    type for Json data. If true, all columns will be registered as\n                    their primitive types (strings, number or boolean).\n                encoding:\n                  type: string\n                  x-dcl-go-name: Encoding\n                  description: Optional. The character encoding of the data. The default\n                    is UTF-8.\n            schedule:\n              type: string\n              x-dcl-go-name: Schedule\n              description: 'Optional. Cron schedule (https://en.wikipedia.org/wiki/Cron)\n                for running discovery periodically. Successive discovery runs must\n                be scheduled at least 60 minutes apart. The default value is to run\n                discovery every 60 minutes. To explicitly set a timezone to the cron\n                tab, apply a prefix in the cron tab: \"CRON_TZ=${IANA_TIME_ZONE}\" or\n                TZ=${IANA_TIME_ZONE}\". The ${IANA_TIME_ZONE} may only be a valid string\n                from IANA time zone database. For example, \"CRON_TZ=America/New_York\n                1 * * * *\", or \"TZ=America/New_York 1 * * * *\".'\n        discoveryStatus:\n          type: object\n          x-dcl-go-name: DiscoveryStatus\n          x-dcl-go-type: AssetDiscoveryStatus\n          readOnly: true\n          description: Output only. Status of the discovery feature applied to data\n            referenced by this asset.\n          properties:\n            lastRunDuration:\n              type: string\n              x-dcl-go-name: LastRunDuration\n              description: The duration of the last discovery run.\n            lastRunTime:\n              type: string\n              format: date-time\n              x-dcl-go-name: LastRunTime\n              description: The start time of the last discovery run.\n            message:\n              type: string\n              x-dcl-go-name: Message\n              description: Additional information about the current state.\n            state:\n              type: string\n              x-dcl-go-name: State\n              x-dcl-go-type: AssetDiscoveryStatusStateEnum\n              description: 'The current status of the discovery feature. Possible\n                values: STATE_UNSPECIFIED, SCHEDULED, IN_PROGRESS, PAUSED, DISABLED'\n              enum:\n              - STATE_UNSPECIFIED\n              - SCHEDULED\n              - IN_PROGRESS\n              - PAUSED\n              - DISABLED\n            stats:\n              type: object\n              x-dcl-go-name: Stats\n              x-dcl-go-type: AssetDiscoveryStatusStats\n              description: Data Stats of the asset reported by discovery.\n              properties:\n                dataItems:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: DataItems\n                  description: The count of data items within the referenced resource.\n                dataSize:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: DataSize\n                  description: The number of stored data bytes within the referenced\n                    resource.\n                filesets:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: Filesets\n                  description: The count of fileset entities within the referenced\n                    resource.\n                tables:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: Tables\n                  description: The count of table entities within the referenced resource.\n            updateTime:\n              type: string\n              format: date-time\n              x-dcl-go-name: UpdateTime\n              description: Last update time of the status.\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: Optional. User friendly display name.\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Optional. User defined labels for the asset.\n        lake:\n          type: string\n          x-dcl-go-name: Lake\n          description: The lake for the resource\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The name of the asset.\n          x-dcl-references:\n          - resource: Dataplex/Asset\n            field: selfLink\n            parent: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        resourceSpec:\n          type: object\n          x-dcl-go-name: ResourceSpec\n          x-dcl-go-type: AssetResourceSpec\n          description: Required. Specification of the resource that is referenced\n            by this asset.\n          required:\n          - type\n          properties:\n            name:\n              type: string\n              x-dcl-go-name: Name\n              description: 'Immutable. Relative name of the cloud resource that contains\n                the data that is being managed within a lake. For example: `projects/{project_number}/buckets/{bucket_id}`\n                `projects/{project_number}/datasets/{dataset_id}`'\n            type:\n              type: string\n              x-dcl-go-name: Type\n              x-dcl-go-type: AssetResourceSpecTypeEnum\n              description: 'Required. Immutable. Type of resource. Possible values:\n                TYPE_UNSPECIFIED, STORAGE_BUCKET, BIGQUERY_DATASET'\n              enum:\n              - TYPE_UNSPECIFIED\n              - STORAGE_BUCKET\n              - BIGQUERY_DATASET\n        resourceStatus:\n          type: object\n          x-dcl-go-name: ResourceStatus\n          x-dcl-go-type: AssetResourceStatus\n          readOnly: true\n          description: Output only. Status of the resource referenced by this asset.\n          properties:\n            message:\n              type: string\n              x-dcl-go-name: Message\n              description: Additional information about the current state.\n            state:\n              type: string\n              x-dcl-go-name: State\n              x-dcl-go-type: AssetResourceStatusStateEnum\n              description: 'The current state of the managed resource. Possible values:\n                STATE_UNSPECIFIED, READY, ERROR'\n              enum:\n              - STATE_UNSPECIFIED\n              - READY\n              - ERROR\n            updateTime:\n              type: string\n              format: date-time\n              x-dcl-go-name: UpdateTime\n              description: Last update time of the status.\n        securityStatus:\n          type: object\n          x-dcl-go-name: SecurityStatus\n          x-dcl-go-type: AssetSecurityStatus\n          readOnly: true\n          description: Output only. Status of the security policy applied to resource\n            referenced by this asset.\n          properties:\n            message:\n              type: string\n              x-dcl-go-name: Message\n              description: Additional information about the current state.\n            state:\n              type: string\n              x-dcl-go-name: State\n              x-dcl-go-type: AssetSecurityStatusStateEnum\n              description: 'The current state of the security policy applied to the\n                attached resource. Possible values: STATE_UNSPECIFIED, READY, APPLYING,\n                ERROR'\n              enum:\n              - STATE_UNSPECIFIED\n              - READY\n              - APPLYING\n              - ERROR\n            updateTime:\n              type: string\n              format: date-time\n              x-dcl-go-name: UpdateTime\n              description: Last update time of the status.\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: AssetStateEnum\n          readOnly: true\n          description: 'Output only. Current state of the asset. Possible values:\n            STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - ACTIVE\n          - CREATING\n          - DELETING\n          - ACTION_REQUIRED\n        uid:\n          type: string\n          x-dcl-go-name: Uid\n          readOnly: true\n          description: Output only. System generated globally unique ID for the asset.\n            This ID will be different if the asset is deleted and re-created with\n            the same name.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The time when the asset was last updated.\n          x-kubernetes-immutable: true\n        zone:\n          type: string\n          x-dcl-go-name: Zone\n          description: The zone for the resource\n          x-kubernetes-immutable: true\n")

// 15183 bytes
// MD5: a9b376d63ada800a2afaae7e51bf79ac
