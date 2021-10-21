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
// gen_go_data -package beta -var YAML_policy blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/binaryauthorization/beta/policy.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/binaryauthorization/beta/policy.yaml
var YAML_policy = []byte("info:\n  title: BinaryAuthorization/Policy\n  description: The BinaryAuthorization Policy resource\n  x-dcl-has-iam: true\npaths:\n  get:\n    description: The function used to get information about a Policy\n    parameters:\n    - name: Policy\n      required: true\n      description: A full instance of a Policy\n  apply:\n    description: The function used to apply information about a Policy\n    parameters:\n    - name: Policy\n      required: true\n      description: A full instance of a Policy\ncomponents:\n  schemas:\n    AdmissionRule:\n      type: object\n      x-dcl-go-name: DefaultAdmissionRule\n      x-dcl-go-type: PolicyAdmissionRule\n      description: Required. Default admission rule for a cluster without a per-cluster,\n        per-kubernetes-service-account, or per-istio-service-identity admission rule.\n      required:\n      - evaluationMode\n      - enforcementMode\n      properties:\n        enforcementMode:\n          type: string\n          x-dcl-go-name: EnforcementMode\n          x-dcl-go-type: PolicyAdmissionRuleEnforcementModeEnum\n          description: 'Required. The action when a pod creation is denied by the\n            admission rule. Possible values: ENFORCEMENT_MODE_UNSPECIFIED, ENFORCED_BLOCK_AND_AUDIT_LOG,\n            DRYRUN_AUDIT_LOG_ONLY'\n          enum:\n          - ENFORCEMENT_MODE_UNSPECIFIED\n          - ENFORCED_BLOCK_AND_AUDIT_LOG\n          - DRYRUN_AUDIT_LOG_ONLY\n        evaluationMode:\n          type: string\n          x-dcl-go-name: EvaluationMode\n          x-dcl-go-type: PolicyAdmissionRuleEvaluationModeEnum\n          description: 'Required. How this admission rule will be evaluated. Possible\n            values: ALWAYS_ALLOW, ALWAYS_DENY, REQUIRE_ATTESTATION'\n          enum:\n          - ALWAYS_ALLOW\n          - ALWAYS_DENY\n          - REQUIRE_ATTESTATION\n        requireAttestationsBy:\n          type: array\n          x-dcl-go-name: RequireAttestationsBy\n          description: 'Optional. The resource names of the attestors that must attest\n            to a container image, in the format `projects/*/attestors/*`. Each attestor\n            must exist before a policy can reference it. To add an attestor to a policy\n            the principal issuing the policy change request must be able to read the\n            attestor resource. Note: this field must be non-empty when the evaluation_mode\n            field specifies REQUIRE_ATTESTATION, otherwise it must be empty.'\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n            x-dcl-references:\n            - resource: Binaryauthorization/Attestor\n              field: name\n    Policy:\n      title: Policy\n      x-dcl-id: projects/{{project}}/policy\n      x-dcl-parent-container: project\n      type: object\n      required:\n      - defaultAdmissionRule\n      properties:\n        admissionWhitelistPatterns:\n          type: array\n          x-dcl-go-name: AdmissionWhitelistPatterns\n          description: Optional. Admission policy allowlisting. A matching admission\n            request will always be permitted. This feature is typically used to exclude\n            Google or third-party infrastructure images from Binary Authorization\n            policies.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: PolicyAdmissionWhitelistPatterns\n            properties:\n              namePattern:\n                type: string\n                x-dcl-go-name: NamePattern\n                description: An image name pattern to allowlist, in the form `registry/path/to/image`.\n                  This supports a trailing `*` as a wildcard, but this is allowed\n                  only in text after the `registry/` part.\n        clusterAdmissionRules:\n          type: object\n          additionalProperties:\n            $ref: '#/components/schemas/AdmissionRule'\n          x-dcl-go-name: ClusterAdmissionRules\n          description: 'Optional. Per-cluster admission rules. Cluster spec format:\n            location.clusterId. There can be at most one admission rule per cluster\n            spec. A location is either a compute zone (e.g. us-central1-a) or a region\n            (e.g. us-central1). For clusterId syntax restrictions see https://cloud.google.com/container-engine/reference/rest/v1/projects.zones.clusters.'\n          x-dcl-conflicts:\n          - kubernetesNamespaceAdmissionRules\n          - kubernetesServiceAccountAdmissionRules\n          - istioServiceIdentityAdmissionRules\n        defaultAdmissionRule:\n          $ref: '#/components/schemas/AdmissionRule'\n          x-dcl-go-name: DefaultAdmissionRule\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. A descriptive comment.\n        globalPolicyEvaluationMode:\n          type: string\n          x-dcl-go-name: GlobalPolicyEvaluationMode\n          x-dcl-go-type: PolicyGlobalPolicyEvaluationModeEnum\n          description: 'Optional. Controls the evaluation of a Google-maintained global\n            admission policy for common system-level images. Images not covered by\n            the global policy will be subject to the project admission policy. This\n            setting has no effect when specified inside a global admission policy.\n            Possible values: GLOBAL_POLICY_EVALUATION_MODE_UNSPECIFIED, ENABLE, DISABLE'\n          enum:\n          - GLOBAL_POLICY_EVALUATION_MODE_UNSPECIFIED\n          - ENABLE\n          - DISABLE\n        istioServiceIdentityAdmissionRules:\n          type: object\n          additionalProperties:\n            $ref: '#/components/schemas/AdmissionRule'\n          x-dcl-go-name: IstioServiceIdentityAdmissionRules\n          description: 'Optional. Per-istio-service-identity admission rules. Istio\n            service identity spec format: spiffe:///ns//sa/ or /ns//sa/ e.g. spiffe://example.com/ns/test-ns/sa/default'\n          x-dcl-conflicts:\n          - kubernetesNamespaceAdmissionRules\n          - kubernetesServiceAccountAdmissionRules\n          - clusterAdmissionRules\n        kubernetesNamespaceAdmissionRules:\n          type: object\n          additionalProperties:\n            $ref: '#/components/schemas/AdmissionRule'\n          x-dcl-go-name: KubernetesNamespaceAdmissionRules\n          description: 'Optional. Per-kubernetes-namespace admission rules. K8s namespace\n            spec format: [a-z.-]+, e.g. ''some-namespace'''\n          x-dcl-conflicts:\n          - kubernetesServiceAccountAdmissionRules\n          - istioServiceIdentityAdmissionRules\n          - clusterAdmissionRules\n        kubernetesServiceAccountAdmissionRules:\n          type: object\n          additionalProperties:\n            $ref: '#/components/schemas/AdmissionRule'\n          x-dcl-go-name: KubernetesServiceAccountAdmissionRules\n          description: 'Optional. Per-kubernetes-service-account admission rules.\n            Service account spec format: namespace:serviceaccount. e.g. ''test-ns:default'''\n          x-dcl-conflicts:\n          - kubernetesNamespaceAdmissionRules\n          - istioServiceIdentityAdmissionRules\n          - clusterAdmissionRules\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project of the resource.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: Output only. The resource name, in the format `projects/*/policy`.\n            There is at most one policy per project.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. Time when the policy was last updated.\n          x-kubernetes-immutable: true\n")

// 7994 bytes
// MD5: 6955391f549147870f1a00c919c8cb2f
