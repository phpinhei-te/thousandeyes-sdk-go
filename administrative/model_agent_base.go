/*
Administrative API

Manage users, accounts, and account groups in the ThousandEyes platform using the Administrative API. This API provides the following operations to manage your organization:     * `/account-groups`: Account groups are used to divide an organization into different sections. These operations can be used to create, retrieve, update and delete account groups.   * `/users`: Create, retrieve, update and delete users within an organization.    * `/roles`: Create, retrieve and update roles for the current user.    * `/permissions`: Retrieve all assignable permissions. Used in the context of modifying roles.    * `/audit-user-events`: Retrieve all activity log events.    For more information about the administrative models, see [Account Management](https://docs.thousandeyes.com/product-documentation/user-management).

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package administrative

import (
	"encoding/json"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the AgentBase type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AgentBase{}

// AgentBase struct for AgentBase
type AgentBase struct {
	// Array of private IP addresses.
	IpAddresses []string `json:"ipAddresses,omitempty"`
	// Array of public IP addresses.
	PublicIpAddresses []string `json:"publicIpAddresses,omitempty"`
	// Network (including ASN) of agent’s public IP.
	Network *string `json:"network,omitempty"`
}

// NewAgentBase instantiates a new AgentBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentBase() *AgentBase {
	this := AgentBase{}
	return &this
}

// NewAgentBaseWithDefaults instantiates a new AgentBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentBaseWithDefaults() *AgentBase {
	this := AgentBase{}
	return &this
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *AgentBase) GetIpAddresses() []string {
	if o == nil || utils.IsNil(o.IpAddresses) {
		var ret []string
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentBase) GetIpAddressesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.IpAddresses) {
		return nil, false
	}
	return o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *AgentBase) HasIpAddresses() bool {
	if o != nil && !utils.IsNil(o.IpAddresses) {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *AgentBase) SetIpAddresses(v []string) {
	o.IpAddresses = v
}

// GetPublicIpAddresses returns the PublicIpAddresses field value if set, zero value otherwise.
func (o *AgentBase) GetPublicIpAddresses() []string {
	if o == nil || utils.IsNil(o.PublicIpAddresses) {
		var ret []string
		return ret
	}
	return o.PublicIpAddresses
}

// GetPublicIpAddressesOk returns a tuple with the PublicIpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentBase) GetPublicIpAddressesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.PublicIpAddresses) {
		return nil, false
	}
	return o.PublicIpAddresses, true
}

// HasPublicIpAddresses returns a boolean if a field has been set.
func (o *AgentBase) HasPublicIpAddresses() bool {
	if o != nil && !utils.IsNil(o.PublicIpAddresses) {
		return true
	}

	return false
}

// SetPublicIpAddresses gets a reference to the given []string and assigns it to the PublicIpAddresses field.
func (o *AgentBase) SetPublicIpAddresses(v []string) {
	o.PublicIpAddresses = v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *AgentBase) GetNetwork() string {
	if o == nil || utils.IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentBase) GetNetworkOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *AgentBase) HasNetwork() bool {
	if o != nil && !utils.IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *AgentBase) SetNetwork(v string) {
	o.Network = &v
}

func (o AgentBase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.IpAddresses) {
		toSerialize["ipAddresses"] = o.IpAddresses
	}
	if !utils.IsNil(o.PublicIpAddresses) {
		toSerialize["publicIpAddresses"] = o.PublicIpAddresses
	}
	if !utils.IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	return toSerialize, nil
}

type NullableAgentBase struct {
	value *AgentBase
	isSet bool
}

func (v NullableAgentBase) Get() *AgentBase {
	return v.value
}

func (v *NullableAgentBase) Set(val *AgentBase) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentBase) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentBase(val *AgentBase) *NullableAgentBase {
	return &NullableAgentBase{value: val, isSet: true}
}

func (v NullableAgentBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
