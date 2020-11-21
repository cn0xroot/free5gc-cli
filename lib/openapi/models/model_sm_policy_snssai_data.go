/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains the SM policy data for a given subscriber and S-NSSAI.
type SmPolicySnssaiData struct {
	Snssai          *Snssai                    `json:"snssai" bson:"snssai" yaml:"snssai"`
	SmPolicyDnnData map[string]SmPolicyDnnData `json:"smPolicyDnnData,omitempty" bson:"smPolicyDnnData" yaml:"smPolicyDnnData,omitempty"`
}
