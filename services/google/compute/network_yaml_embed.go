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
// gen_go_data -package compute -var YAML_network blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/network.yaml

package compute

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/network.yaml
var YAML_network = []byte("info:\n  title: Compute/Network\n  description: The Compute Network resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Network\n    parameters:\n    - name: Network\n      required: true\n      description: A full instance of a Network\n  apply:\n    description: The function used to apply information about a Network\n    parameters:\n    - name: Network\n      required: true\n      description: A full instance of a Network\n  delete:\n    description: The function used to delete a Network\n    parameters:\n    - name: Network\n      required: true\n      description: A full instance of a Network\n  deleteAll:\n    description: The function used to delete all Network\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Network\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Network:\n      title: Network\n      x-dcl-id: projects/{{project}}/global/networks/{{name}}\n      x-dcl-locations:\n      - global\n      x-dcl-parent-container: project\n      type: object\n      required:\n      - name\n      - project\n      properties:\n        autoCreateSubnetworks:\n          type: boolean\n          x-dcl-go-name: AutoCreateSubnetworks\n          description: 'When set to `true`, the network is created in \"auto subnet\n            mode\" and it will create a subnet for each region automatically across\n            the `10.128.0.0/9` address range.  When set to `false`, the network is\n            created in \"custom subnet mode\" so the user can explicitly connect subnetwork\n            resources. '\n          x-kubernetes-immutable: true\n          default: true\n          x-dcl-send-empty: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: 'An optional description of this resource. The resource must\n            be recreated to modify this field. '\n          x-kubernetes-immutable: true\n        gatewayIPv4:\n          type: string\n          x-dcl-go-name: GatewayIPv4\n          readOnly: true\n          description: 'The gateway address for default routing out of the network.\n            This value is selected by GCP. '\n          x-kubernetes-immutable: true\n        mtu:\n          type: integer\n          format: int64\n          x-dcl-go-name: Mtu\n          description: Maximum Transmission Unit in bytes. The minimum value for this\n            field is 1460 and the maximum value is 1500 bytes.\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'Name of the resource. Provided by the client when the resource\n            is created. The name must be 1-63 characters long, and comply with RFC1035.\n            Specifically, the name must be 1-63 characters long and match the regular\n            expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character\n            must be a lowercase letter, and all following characters must be a dash,\n            lowercase letter, or digit, except the last character, which cannot be\n            a dash. '\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project id of the resource.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        routingConfig:\n          type: object\n          x-dcl-go-name: RoutingConfig\n          x-dcl-go-type: NetworkRoutingConfig\n          description: 'The network-level routing configuration for this network.\n            Used by Cloud Router to determine what type of network-wide routing behavior\n            to enforce. '\n          properties:\n            routingMode:\n              type: string\n              x-dcl-go-name: RoutingMode\n              x-dcl-go-type: NetworkRoutingConfigRoutingModeEnum\n              description: 'The network-wide routing mode to use. If set to `REGIONAL`,\n                this network''s cloud routers will only advertise routes with subnetworks\n                of this network in the same region as the router. If set to `GLOBAL`,\n                this network''s cloud routers will advertise routes with all subnetworks\n                of this network, across regions. '\n              x-dcl-server-default: true\n              enum:\n              - REGIONAL\n              - GLOBAL\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: Server-defined URL for the resource.\n          x-kubernetes-immutable: true\n        selfLinkWithId:\n          type: string\n          x-dcl-go-name: SelfLinkWithId\n          readOnly: true\n          description: Server-defined URL for the resource containing the network\n            ID.\n          x-kubernetes-immutable: true\n")

// 4989 bytes
// MD5: 5a38473b861b905391c763e99854092a
