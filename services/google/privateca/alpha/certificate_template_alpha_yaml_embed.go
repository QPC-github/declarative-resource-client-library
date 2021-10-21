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
// gen_go_data -package alpha -var YAML_certificate_template blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/privateca/alpha/certificate_template.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/privateca/alpha/certificate_template.yaml
var YAML_certificate_template = []byte("info:\n  title: Privateca/CertificateTemplate\n  description: The Privateca CertificateTemplate resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a CertificateTemplate\n    parameters:\n    - name: CertificateTemplate\n      required: true\n      description: A full instance of a CertificateTemplate\n  apply:\n    description: The function used to apply information about a CertificateTemplate\n    parameters:\n    - name: CertificateTemplate\n      required: true\n      description: A full instance of a CertificateTemplate\n  delete:\n    description: The function used to delete a CertificateTemplate\n    parameters:\n    - name: CertificateTemplate\n      required: true\n      description: A full instance of a CertificateTemplate\n  deleteAll:\n    description: The function used to delete all CertificateTemplate\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many CertificateTemplate\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    CertificateTemplate:\n      title: CertificateTemplate\n      x-dcl-id: projects/{{project}}/locations/{{location}}/certificateTemplates/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      type: object\n      required:\n      - name\n      - project\n      - location\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The time at which this CertificateTemplate was\n            created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. A human-readable description of scenarios this template\n            is intended for.\n        identityConstraints:\n          type: object\n          x-dcl-go-name: IdentityConstraints\n          x-dcl-go-type: CertificateTemplateIdentityConstraints\n          description: Optional. Describes constraints on identities that may be appear\n            in Certificates issued using this template. If this is omitted, then this\n            template will not add restrictions on a certificate's identity.\n          required:\n          - allowSubjectPassthrough\n          - allowSubjectAltNamesPassthrough\n          properties:\n            allowSubjectAltNamesPassthrough:\n              type: boolean\n              x-dcl-go-name: AllowSubjectAltNamesPassthrough\n              description: Required. If this is true, the SubjectAltNames extension\n                may be copied from a certificate request into the signed certificate.\n                Otherwise, the requested SubjectAltNames will be discarded.\n            allowSubjectPassthrough:\n              type: boolean\n              x-dcl-go-name: AllowSubjectPassthrough\n              description: Required. If this is true, the Subject field may be copied\n                from a certificate request into the signed certificate. Otherwise,\n                the requested Subject will be discarded.\n            celExpression:\n              type: object\n              x-dcl-go-name: CelExpression\n              x-dcl-go-type: CertificateTemplateIdentityConstraintsCelExpression\n              description: Optional. A CEL expression that may be used to validate\n                the resolved X.509 Subject and/or Subject Alternative Name before\n                a certificate is signed. To see the full allowed syntax and some examples,\n                see https://cloud.google.com/certificate-authority-service/docs/using-cel\n              properties:\n                description:\n                  type: string\n                  x-dcl-go-name: Description\n                  description: Optional. Description of the expression. This is a\n                    longer text which describes the expression, e.g. when hovered\n                    over it in a UI.\n                expression:\n                  type: string\n                  x-dcl-go-name: Expression\n                  description: Textual representation of an expression in Common Expression\n                    Language syntax.\n                location:\n                  type: string\n                  x-dcl-go-name: Location\n                  description: Optional. String indicating the location of the expression\n                    for error reporting, e.g. a file name and a position in the file.\n                title:\n                  type: string\n                  x-dcl-go-name: Title\n                  description: Optional. Title for the expression, i.e. a short string\n                    describing its purpose. This can be used e.g. in UIs which allow\n                    to enter the expression.\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Optional. Labels with user-defined metadata.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The resource name for this CertificateTemplate in the format\n            `projects/*/locations/*/certificateTemplates/*`.\n          x-kubernetes-immutable: true\n        passthroughExtensions:\n          type: object\n          x-dcl-go-name: PassthroughExtensions\n          x-dcl-go-type: CertificateTemplatePassthroughExtensions\n          description: Optional. Describes the set of X.509 extensions that may appear\n            in a Certificate issued using this CertificateTemplate. If a certificate\n            request sets extensions that don't appear in the passthrough_extensions,\n            those extensions will be dropped. If the issuing CaPool's IssuancePolicy\n            defines baseline_values that don't appear here, the certificate issuance\n            request will fail. If this is omitted, then this template will not add\n            restrictions on a certificate's X.509 extensions. These constraints do\n            not apply to X.509 extensions set in this CertificateTemplate's predefined_values.\n          properties:\n            additionalExtensions:\n              type: array\n              x-dcl-go-name: AdditionalExtensions\n              description: Optional. A set of ObjectIds identifying custom X.509 extensions.\n                Will be combined with known_extensions to determine the full set of\n                X.509 extensions.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: CertificateTemplatePassthroughExtensionsAdditionalExtensions\n                required:\n                - objectIdPath\n                properties:\n                  objectIdPath:\n                    type: array\n                    x-dcl-go-name: ObjectIdPath\n                    description: Required. The parts of an OID path. The most significant\n                      parts of the path come first.\n                    x-dcl-send-empty: true\n                    x-dcl-list-type: list\n                    items:\n                      type: integer\n                      format: int64\n                      x-dcl-go-type: int64\n            knownExtensions:\n              type: array\n              x-dcl-go-name: KnownExtensions\n              description: Optional. A set of named X.509 extensions. Will be combined\n                with additional_extensions to determine the full set of X.509 extensions.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: CertificateTemplatePassthroughExtensionsKnownExtensionsEnum\n                enum:\n                - KNOWN_CERTIFICATE_EXTENSION_UNSPECIFIED\n                - BASE_KEY_USAGE\n                - EXTENDED_KEY_USAGE\n                - CA_OPTIONS\n                - POLICY_IDS\n                - AIA_OCSP_SERVERS\n        predefinedValues:\n          type: object\n          x-dcl-go-name: PredefinedValues\n          x-dcl-go-type: CertificateTemplatePredefinedValues\n          description: Optional. A set of X.509 values that will be applied to all\n            issued certificates that use this template. If the certificate request\n            includes conflicting values for the same properties, they will be overwritten\n            by the values defined here. If the issuing CaPool's IssuancePolicy defines\n            conflicting baseline_values for the same properties, the certificate issuance\n            request will fail.\n          properties:\n            additionalExtensions:\n              type: array\n              x-dcl-go-name: AdditionalExtensions\n              description: Optional. Describes custom X.509 extensions.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: CertificateTemplatePredefinedValuesAdditionalExtensions\n                required:\n                - objectId\n                - value\n                properties:\n                  critical:\n                    type: boolean\n                    x-dcl-go-name: Critical\n                    description: Optional. Indicates whether or not this extension\n                      is critical (i.e., if the client does not know how to handle\n                      this extension, the client should consider this to be an error).\n                  objectId:\n                    type: object\n                    x-dcl-go-name: ObjectId\n                    x-dcl-go-type: CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId\n                    description: Required. The OID for this X.509 extension.\n                    required:\n                    - objectIdPath\n                    properties:\n                      objectIdPath:\n                        type: array\n                        x-dcl-go-name: ObjectIdPath\n                        description: Required. The parts of an OID path. The most\n                          significant parts of the path come first.\n                        x-dcl-send-empty: true\n                        x-dcl-list-type: list\n                        items:\n                          type: integer\n                          format: int64\n                          x-dcl-go-type: int64\n                  value:\n                    type: string\n                    x-dcl-go-name: Value\n                    description: Required. The value of this X.509 extension.\n            aiaOcspServers:\n              type: array\n              x-dcl-go-name: AiaOcspServers\n              description: Optional. Describes Online Certificate Status Protocol\n                (OCSP) endpoint addresses that appear in the \"Authority Information\n                Access\" extension in the certificate.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            caOptions:\n              type: object\n              x-dcl-go-name: CaOptions\n              x-dcl-go-type: CertificateTemplatePredefinedValuesCaOptions\n              description: Optional. Describes options in this X509Parameters that\n                are relevant in a CA certificate.\n              properties:\n                isCa:\n                  type: boolean\n                  x-dcl-go-name: IsCa\n                  description: Optional. Refers to the \"CA\" X.509 extension, which\n                    is a boolean value. When this value is missing, the extension\n                    will be omitted from the CA certificate.\n                maxIssuerPathLength:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: MaxIssuerPathLength\n                  description: Optional. Refers to the path length restriction X.509\n                    extension. For a CA certificate, this value describes the depth\n                    of subordinate CA certificates that are allowed. If this value\n                    is less than 0, the request will fail. If this value is missing,\n                    the max path length will be omitted from the CA certificate.\n            keyUsage:\n              type: object\n              x-dcl-go-name: KeyUsage\n              x-dcl-go-type: CertificateTemplatePredefinedValuesKeyUsage\n              description: Optional. Indicates the intended use for keys that correspond\n                to a certificate.\n              properties:\n                baseKeyUsage:\n                  type: object\n                  x-dcl-go-name: BaseKeyUsage\n                  x-dcl-go-type: CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage\n                  description: Describes high-level ways in which a key may be used.\n                  properties:\n                    certSign:\n                      type: boolean\n                      x-dcl-go-name: CertSign\n                      description: The key may be used to sign certificates.\n                    contentCommitment:\n                      type: boolean\n                      x-dcl-go-name: ContentCommitment\n                      description: The key may be used for cryptographic commitments.\n                        Note that this may also be referred to as \"non-repudiation\".\n                    crlSign:\n                      type: boolean\n                      x-dcl-go-name: CrlSign\n                      description: The key may be used sign certificate revocation\n                        lists.\n                    dataEncipherment:\n                      type: boolean\n                      x-dcl-go-name: DataEncipherment\n                      description: The key may be used to encipher data.\n                    decipherOnly:\n                      type: boolean\n                      x-dcl-go-name: DecipherOnly\n                      description: The key may be used to decipher only.\n                    digitalSignature:\n                      type: boolean\n                      x-dcl-go-name: DigitalSignature\n                      description: The key may be used for digital signatures.\n                    encipherOnly:\n                      type: boolean\n                      x-dcl-go-name: EncipherOnly\n                      description: The key may be used to encipher only.\n                    keyAgreement:\n                      type: boolean\n                      x-dcl-go-name: KeyAgreement\n                      description: The key may be used in a key agreement protocol.\n                    keyEncipherment:\n                      type: boolean\n                      x-dcl-go-name: KeyEncipherment\n                      description: The key may be used to encipher other keys.\n                extendedKeyUsage:\n                  type: object\n                  x-dcl-go-name: ExtendedKeyUsage\n                  x-dcl-go-type: CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage\n                  description: Detailed scenarios in which a key may be used.\n                  properties:\n                    clientAuth:\n                      type: boolean\n                      x-dcl-go-name: ClientAuth\n                      description: Corresponds to OID 1.3.6.1.5.5.7.3.2. Officially\n                        described as \"TLS WWW client authentication\", though regularly\n                        used for non-WWW TLS.\n                    codeSigning:\n                      type: boolean\n                      x-dcl-go-name: CodeSigning\n                      description: Corresponds to OID 1.3.6.1.5.5.7.3.3. Officially\n                        described as \"Signing of downloadable executable code client\n                        authentication\".\n                    emailProtection:\n                      type: boolean\n                      x-dcl-go-name: EmailProtection\n                      description: Corresponds to OID 1.3.6.1.5.5.7.3.4. Officially\n                        described as \"Email protection\".\n                    ocspSigning:\n                      type: boolean\n                      x-dcl-go-name: OcspSigning\n                      description: Corresponds to OID 1.3.6.1.5.5.7.3.9. Officially\n                        described as \"Signing OCSP responses\".\n                    serverAuth:\n                      type: boolean\n                      x-dcl-go-name: ServerAuth\n                      description: Corresponds to OID 1.3.6.1.5.5.7.3.1. Officially\n                        described as \"TLS WWW server authentication\", though regularly\n                        used for non-WWW TLS.\n                    timeStamping:\n                      type: boolean\n                      x-dcl-go-name: TimeStamping\n                      description: Corresponds to OID 1.3.6.1.5.5.7.3.8. Officially\n                        described as \"Binding the hash of an object to a time\".\n                unknownExtendedKeyUsages:\n                  type: array\n                  x-dcl-go-name: UnknownExtendedKeyUsages\n                  description: Used to describe extended key usages that are not listed\n                    in the KeyUsage.ExtendedKeyUsageOptions message.\n                  x-dcl-send-empty: true\n                  x-dcl-list-type: list\n                  items:\n                    type: object\n                    x-dcl-go-type: CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages\n                    required:\n                    - objectIdPath\n                    properties:\n                      objectIdPath:\n                        type: array\n                        x-dcl-go-name: ObjectIdPath\n                        description: Required. The parts of an OID path. The most\n                          significant parts of the path come first.\n                        x-dcl-send-empty: true\n                        x-dcl-list-type: list\n                        items:\n                          type: integer\n                          format: int64\n                          x-dcl-go-type: int64\n            policyIds:\n              type: array\n              x-dcl-go-name: PolicyIds\n              description: Optional. Describes the X.509 certificate policy object\n                identifiers, per https://tools.ietf.org/html/rfc5280#section-4.2.1.4.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: CertificateTemplatePredefinedValuesPolicyIds\n                required:\n                - objectIdPath\n                properties:\n                  objectIdPath:\n                    type: array\n                    x-dcl-go-name: ObjectIdPath\n                    description: Required. The parts of an OID path. The most significant\n                      parts of the path come first.\n                    x-dcl-send-empty: true\n                    x-dcl-list-type: list\n                    items:\n                      type: integer\n                      format: int64\n                      x-dcl-go-type: int64\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The time at which this CertificateTemplate was\n            updated.\n          x-kubernetes-immutable: true\n")

// 20027 bytes
// MD5: 7c7b4d4f22df20a36ef306b9818c6f19
