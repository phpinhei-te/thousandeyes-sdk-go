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
	"time"
)

// checks if the ClusterMember type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ClusterMember{}

// ClusterMember struct for ClusterMember
type ClusterMember struct {
	// Array of private IP addresses.
	IpAddresses []string `json:"ipAddresses,omitempty"`
	// Array of public IP addresses.
	PublicIpAddresses []string `json:"publicIpAddresses,omitempty"`
	// Network (including ASN) of agent’s public IP.
	Network *string `json:"network,omitempty"`
	// Unique ID of the cluster member
	MemberId *string `json:"memberId,omitempty"`
	// Name of the cluster member
	Name *string `json:"name,omitempty"`
	// If an enterprise agent or a cluster member presents at least one error, the errors will be shown as an array of entries in the errorDetails field (Enterprise Agents and Enterprise Cluster members only)
	ErrorDetails []ErrorDetail `json:"errorDetails,omitempty"`
	// UTC last seen date (ISO date-time format).
	LastSeen   *time.Time            `json:"lastSeen,omitempty"`
	AgentState *EnterpriseAgentState `json:"agentState,omitempty"`
	// Test target IP address.
	TargetForTests *string `json:"targetForTests,omitempty"`
	// Shows overall utilization percentage (online Enterprise Agents and Enterprise Clusters only).
	Utilization *int32 `json:"utilization,omitempty"`
}

// NewClusterMember instantiates a new ClusterMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterMember() *ClusterMember {
	this := ClusterMember{}
	return &this
}

// NewClusterMemberWithDefaults instantiates a new ClusterMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterMemberWithDefaults() *ClusterMember {
	this := ClusterMember{}
	return &this
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *ClusterMember) GetIpAddresses() []string {
	if o == nil || utils.IsNil(o.IpAddresses) {
		var ret []string
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterMember) GetIpAddressesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.IpAddresses) {
		return nil, false
	}
	return o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *ClusterMember) HasIpAddresses() bool {
	if o != nil && !utils.IsNil(o.IpAddresses) {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *ClusterMember) SetIpAddresses(v []string) {
	o.IpAddresses = v
}

// GetPublicIpAddresses returns the PublicIpAddresses field value if set, zero value otherwise.
func (o *ClusterMember) GetPublicIpAddresses() []string {
	if o == nil || utils.IsNil(o.PublicIpAddresses) {
		var ret []string
		return ret
	}
	return o.PublicIpAddresses
}

// GetPublicIpAddressesOk returns a tuple with the PublicIpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterMember) GetPublicIpAddressesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.PublicIpAddresses) {
		return nil, false
	}
	return o.PublicIpAddresses, true
}

// HasPublicIpAddresses returns a boolean if a field has been set.
func (o *ClusterMember) HasPublicIpAddresses() bool {
	if o != nil && !utils.IsNil(o.PublicIpAddresses) {
		return true
	}

	return false
}

// SetPublicIpAddresses gets a reference to the given []string and assigns it to the PublicIpAddresses field.
func (o *ClusterMember) SetPublicIpAddresses(v []string) {
	o.PublicIpAddresses = v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *ClusterMember) GetNetwork() string {
	if o == nil || utils.IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterMember) GetNetworkOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *ClusterMember) HasNetwork() bool {
	if o != nil && !utils.IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *ClusterMember) SetNetwork(v string) {
	o.Network = &v
}

// GetMemberId returns the MemberId field value if set, zero value otherwise.
func (o *ClusterMember) GetMemberId() string {
	if o == nil || utils.IsNil(o.MemberId) {
		var ret string
		return ret
	}
	return *o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterMember) GetMemberIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MemberId) {
		return nil, false
	}
	return o.MemberId, true
}

// HasMemberId returns a boolean if a field has been set.
func (o *ClusterMember) HasMemberId() bool {
	if o != nil && !utils.IsNil(o.MemberId) {
		return true
	}

	return false
}

// SetMemberId gets a reference to the given string and assigns it to the MemberId field.
func (o *ClusterMember) SetMemberId(v string) {
	o.MemberId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClusterMember) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterMember) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClusterMember) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClusterMember) SetName(v string) {
	o.Name = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *ClusterMember) GetErrorDetails() []ErrorDetail {
	if o == nil || utils.IsNil(o.ErrorDetails) {
		var ret []ErrorDetail
		return ret
	}
	return o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterMember) GetErrorDetailsOk() ([]ErrorDetail, bool) {
	if o == nil || utils.IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *ClusterMember) HasErrorDetails() bool {
	if o != nil && !utils.IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given []ErrorDetail and assigns it to the ErrorDetails field.
func (o *ClusterMember) SetErrorDetails(v []ErrorDetail) {
	o.ErrorDetails = v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *ClusterMember) GetLastSeen() time.Time {
	if o == nil || utils.IsNil(o.LastSeen) {
		var ret time.Time
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterMember) GetLastSeenOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.LastSeen) {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *ClusterMember) HasLastSeen() bool {
	if o != nil && !utils.IsNil(o.LastSeen) {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given time.Time and assigns it to the LastSeen field.
func (o *ClusterMember) SetLastSeen(v time.Time) {
	o.LastSeen = &v
}

// GetAgentState returns the AgentState field value if set, zero value otherwise.
func (o *ClusterMember) GetAgentState() EnterpriseAgentState {
	if o == nil || utils.IsNil(o.AgentState) {
		var ret EnterpriseAgentState
		return ret
	}
	return *o.AgentState
}

// GetAgentStateOk returns a tuple with the AgentState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterMember) GetAgentStateOk() (*EnterpriseAgentState, bool) {
	if o == nil || utils.IsNil(o.AgentState) {
		return nil, false
	}
	return o.AgentState, true
}

// HasAgentState returns a boolean if a field has been set.
func (o *ClusterMember) HasAgentState() bool {
	if o != nil && !utils.IsNil(o.AgentState) {
		return true
	}

	return false
}

// SetAgentState gets a reference to the given EnterpriseAgentState and assigns it to the AgentState field.
func (o *ClusterMember) SetAgentState(v EnterpriseAgentState) {
	o.AgentState = &v
}

// GetTargetForTests returns the TargetForTests field value if set, zero value otherwise.
func (o *ClusterMember) GetTargetForTests() string {
	if o == nil || utils.IsNil(o.TargetForTests) {
		var ret string
		return ret
	}
	return *o.TargetForTests
}

// GetTargetForTestsOk returns a tuple with the TargetForTests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterMember) GetTargetForTestsOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TargetForTests) {
		return nil, false
	}
	return o.TargetForTests, true
}

// HasTargetForTests returns a boolean if a field has been set.
func (o *ClusterMember) HasTargetForTests() bool {
	if o != nil && !utils.IsNil(o.TargetForTests) {
		return true
	}

	return false
}

// SetTargetForTests gets a reference to the given string and assigns it to the TargetForTests field.
func (o *ClusterMember) SetTargetForTests(v string) {
	o.TargetForTests = &v
}

// GetUtilization returns the Utilization field value if set, zero value otherwise.
func (o *ClusterMember) GetUtilization() int32 {
	if o == nil || utils.IsNil(o.Utilization) {
		var ret int32
		return ret
	}
	return *o.Utilization
}

// GetUtilizationOk returns a tuple with the Utilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterMember) GetUtilizationOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Utilization) {
		return nil, false
	}
	return o.Utilization, true
}

// HasUtilization returns a boolean if a field has been set.
func (o *ClusterMember) HasUtilization() bool {
	if o != nil && !utils.IsNil(o.Utilization) {
		return true
	}

	return false
}

// SetUtilization gets a reference to the given int32 and assigns it to the Utilization field.
func (o *ClusterMember) SetUtilization(v int32) {
	o.Utilization = &v
}

func (o ClusterMember) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterMember) ToMap() (map[string]interface{}, error) {
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
	if !utils.IsNil(o.MemberId) {
		toSerialize["memberId"] = o.MemberId
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	if !utils.IsNil(o.LastSeen) {
		toSerialize["lastSeen"] = o.LastSeen
	}
	if !utils.IsNil(o.AgentState) {
		toSerialize["agentState"] = o.AgentState
	}
	if !utils.IsNil(o.TargetForTests) {
		toSerialize["targetForTests"] = o.TargetForTests
	}
	if !utils.IsNil(o.Utilization) {
		toSerialize["utilization"] = o.Utilization
	}
	return toSerialize, nil
}

type NullableClusterMember struct {
	value *ClusterMember
	isSet bool
}

func (v NullableClusterMember) Get() *ClusterMember {
	return v.value
}

func (v *NullableClusterMember) Set(val *ClusterMember) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterMember) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterMember(val *ClusterMember) *NullableClusterMember {
	return &NullableClusterMember{value: val, isSet: true}
}

func (v NullableClusterMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
