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
// gen_go_data -package alpha -var YAML_route blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/alpha/route.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/alpha/route.yaml
var YAML_route = []byte("info:\n  title: Compute/Route\n  description: The Compute Route resource\n  x-dcl-struct-name: Route\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Route\n    parameters:\n    - name: Route\n      required: true\n      description: A full instance of a Route\n  apply:\n    description: The function used to apply information about a Route\n    parameters:\n    - name: Route\n      required: true\n      description: A full instance of a Route\n  delete:\n    description: The function used to delete a Route\n    parameters:\n    - name: Route\n      required: true\n      description: A full instance of a Route\n  deleteAll:\n    description: The function used to delete all Route\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Route\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Route:\n      title: Route\n      x-dcl-id: projects/{{project}}/global/routes/{{name}}\n      x-dcl-locations:\n      - global\n      x-dcl-parent-container: project\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - network\n      - destRange\n      - project\n      properties:\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: |-\n            An optional description of this resource. Provide this field when you\n            create the resource.\n          x-kubernetes-immutable: true\n        destRange:\n          type: string\n          x-dcl-go-name: DestRange\n          description: The destination range of the route.\n          x-kubernetes-immutable: true\n        id:\n          type: integer\n          format: int64\n          x-dcl-go-name: Id\n          readOnly: true\n          description: |-\n            [Output Only] The unique identifier for the resource. This identifier is\n            defined by the server.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: |-\n            Name of the resource. Provided by the client when the resource is created.\n            The name must be 1-63 characters long, and comply with\n            <a href=\"https://www.ietf.org/rfc/rfc1035.txt\" target=\"_blank\">RFC1035</a>.\n            Specifically, the name must be 1-63 characters long and match the regular\n            expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a\n            lowercase letter, and all following characters (except for the last\n            character) must be a dash, lowercase letter, or digit. The last character\n            must be a lowercase letter or digit.\n          x-kubernetes-immutable: true\n        network:\n          type: string\n          x-dcl-go-name: Network\n          description: Fully-qualified URL of the network that this route applies\n            to.\n          x-kubernetes-immutable: true\n        nextHopGateway:\n          type: string\n          x-dcl-go-name: NextHopGateway\n          description: |-\n            The URL to a gateway that should handle matching packets.\n            You can only specify the internet gateway using a full or\n            partial valid URL: </br>\n            <code>projects/<var\n            class=\"apiparam\">project</var>/global/gateways/default-internet-gateway</code>\n          x-kubernetes-immutable: true\n          x-dcl-conflicts:\n          - nextHopVpnTunnel\n          - nextHopIP\n          - nextHopInstance\n          - nextHopIlb\n        nextHopIP:\n          type: string\n          x-dcl-go-name: NextHopIP\n          description: |-\n            The network IP address of an instance that should handle matching packets.\n            Only IPv4 is supported.\n          x-kubernetes-immutable: true\n          x-dcl-conflicts:\n          - nextHopVpnTunnel\n          - nextHopInstance\n          - nextHopGateway\n          - nextHopIlb\n          x-dcl-server-default: true\n        nextHopIlb:\n          type: string\n          x-dcl-go-name: NextHopIlb\n          description: |-\n            The URL to a forwarding rule of type\n            <code>loadBalancingScheme=INTERNAL</code> that should handle matching\n            packets. You can only specify the forwarding rule as a partial or full\n            URL. For example, the following are all valid URLs:\n            <ul>\n               <li><code>https://www.googleapis.com/compute/v1/projects/<var\n               class=\"apiparam\">project</var>/regions/<var\n               class=\"apiparam\">region</var>/forwardingRules/<var\n               class=\"apiparam\">forwardingRule</var></code></li> <li><code>regions/<var\n               class=\"apiparam\">region</var>/forwardingRules/<var\n               class=\"apiparam\">forwardingRule</var></code></li>\n            </ul>\n          x-kubernetes-immutable: true\n          x-dcl-conflicts:\n          - nextHopVpnTunnel\n          - nextHopIP\n          - nextHopInstance\n          - nextHopGateway\n        nextHopInstance:\n          type: string\n          x-dcl-go-name: NextHopInstance\n          description: |-\n            The URL to an instance that should handle matching packets. You can specify\n            this as a full or partial URL.\n            For example: <br />\n            <code>https://www.googleapis.com/compute/v1/projects/<var\n            class=\"apiparam\">project</var>/zones/<var\n            class=\"apiparam\">zone</var>/instances/<instance-name></code>\n          x-kubernetes-immutable: true\n          x-dcl-conflicts:\n          - nextHopVpnTunnel\n          - nextHopIP\n          - nextHopGateway\n          - nextHopIlb\n        nextHopNetwork:\n          type: string\n          x-dcl-go-name: NextHopNetwork\n          readOnly: true\n          description: The URL of the local network if it should handle matching packets.\n          x-kubernetes-immutable: true\n        nextHopPeering:\n          type: string\n          x-dcl-go-name: NextHopPeering\n          readOnly: true\n          description: |-\n            [Output Only] The network peering name that should handle matching packets,\n            which should conform to RFC1035.\n          x-kubernetes-immutable: true\n        nextHopVpnTunnel:\n          type: string\n          x-dcl-go-name: NextHopVpnTunnel\n          description: The URL to a VpnTunnel that should handle matching packets.\n          x-kubernetes-immutable: true\n          x-dcl-conflicts:\n          - nextHopIP\n          - nextHopInstance\n          - nextHopGateway\n          - nextHopIlb\n        priority:\n          type: integer\n          format: int64\n          x-dcl-go-name: Priority\n          description: The priority of the peering route.\n          x-kubernetes-immutable: true\n          default: 1000\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: '[Output Only] Server-defined fully-qualified URL for this\n            resource.'\n          x-kubernetes-immutable: true\n        tag:\n          type: array\n          x-dcl-go-name: Tag\n          description: A list of instance tags to which this route applies.\n          x-kubernetes-immutable: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n        warning:\n          type: array\n          x-dcl-go-name: Warning\n          readOnly: true\n          description: |-\n            [Output Only] If potential misconfigurations are detected for this\n            route, this field will be populated with warning messages.\n          x-kubernetes-immutable: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: RouteWarning\n            properties:\n              code:\n                type: string\n                x-dcl-go-name: Code\n                x-dcl-go-type: RouteWarningCodeEnum\n                readOnly: true\n                description: |-\n                  [Output Only] A warning code, if applicable. For example, Compute\n                  Engine returns <code>NO_RESULTS_ON_PAGE</code> if there\n                  are no results in the response. Possible values: BAD_REQUEST, FORBIDDEN, NOT_FOUND, CONFLICT, GONE, PRECONDITION_FAILED, INTERNAL_ERROR, SERVICE_UNAVAILABLE\n                x-kubernetes-immutable: true\n                enum:\n                - BAD_REQUEST\n                - FORBIDDEN\n                - NOT_FOUND\n                - CONFLICT\n                - GONE\n                - PRECONDITION_FAILED\n                - INTERNAL_ERROR\n                - SERVICE_UNAVAILABLE\n              data:\n                type: object\n                additionalProperties:\n                  type: string\n                x-dcl-go-name: Data\n                readOnly: true\n                description: |-\n                  [Output Only] Metadata about this warning in <code class=\"lang-html\">key:\n                  value</code> format. For example:\n                  <pre class=\"lang-html prettyprint\">\"data\": [\n                   : {\n                     \"key\": \"scope\",\n                     \"value\": \"zones/us-east1-d\"\n                    }</pre>\n                x-kubernetes-immutable: true\n              message:\n                type: string\n                x-dcl-go-name: Message\n                readOnly: true\n                description: '[Output Only] A human-readable description of the warning\n                  code.'\n                x-kubernetes-immutable: true\n")

// 9833 bytes
// MD5: 899dad2330aec4276bb00def9da0a2df
