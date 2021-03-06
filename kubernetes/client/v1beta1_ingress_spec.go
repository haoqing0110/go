/*
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// IngressSpec describes the Ingress the user wishes to exist.
type V1beta1IngressSpec struct {

	// A default backend capable of servicing requests that don't match any rule. At least one of 'backend' or 'rules' must be specified. This field is optional to allow the loadbalancer controller or defaulting logic to specify a global default.
	Backend *V1beta1IngressBackend `json:"backend,omitempty"`

	// A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.
	Rules []V1beta1IngressRule `json:"rules,omitempty"`

	// TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.
	Tls []V1beta1IngressTls `json:"tls,omitempty"`
}
