/*
Administrative API

Manage users, accounts, and account groups in the ThousandEyes platform using the Administrative API. This API provides the following operations to manage your organization:     * `/account-groups`: Account groups are used to divide an organization into different sections. These operations can be used to create, retrieve, update and delete account groups.   * `/users`: Create, retrieve, update and delete users within an organization.    * `/roles`: Create, retrieve and update roles for the current user.    * `/permissions`: Retrieve all assignable permissions. Used in the context of modifying roles.    * `/audit-user-events`: Retrieve all activity log events.    For more information about the administrative models, see [Account Management](https://docs.thousandeyes.com/product-documentation/user-management).

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package administrative

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/core"
	"bytes"
	"fmt"
)

// checks if the AccountGroupRequest type satisfies the MappedNullable interface at compile time
var _ core.MappedNullable = &AccountGroupRequest{}

// AccountGroupRequest struct for AccountGroupRequest
type AccountGroupRequest struct {
	// The name of the account group
	AccountGroupName string `json:"accountGroupName"`
	// To grant access to enterprise agents, specify the agent list. Note that this is not an additive list - the full list must be specified if changing access to agents.
	Agents []string `json:"agents,omitempty"`
}

type _AccountGroupRequest AccountGroupRequest

// NewAccountGroupRequest instantiates a new AccountGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountGroupRequest(accountGroupName string) *AccountGroupRequest {
	this := AccountGroupRequest{}
	this.AccountGroupName = accountGroupName
	return &this
}

// NewAccountGroupRequestWithDefaults instantiates a new AccountGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountGroupRequestWithDefaults() *AccountGroupRequest {
	this := AccountGroupRequest{}
	return &this
}

// GetAccountGroupName returns the AccountGroupName field value
func (o *AccountGroupRequest) GetAccountGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountGroupName
}

// GetAccountGroupNameOk returns a tuple with the AccountGroupName field value
// and a boolean to check if the value has been set.
func (o *AccountGroupRequest) GetAccountGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountGroupName, true
}

// SetAccountGroupName sets field value
func (o *AccountGroupRequest) SetAccountGroupName(v string) {
	o.AccountGroupName = v
}

// GetAgents returns the Agents field value if set, zero value otherwise.
func (o *AccountGroupRequest) GetAgents() []string {
	if o == nil || core.IsNil(o.Agents) {
		var ret []string
		return ret
	}
	return o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountGroupRequest) GetAgentsOk() ([]string, bool) {
	if o == nil || core.IsNil(o.Agents) {
		return nil, false
	}
	return o.Agents, true
}

// HasAgents returns a boolean if a field has been set.
func (o *AccountGroupRequest) HasAgents() bool {
	if o != nil && !core.IsNil(o.Agents) {
		return true
	}

	return false
}

// SetAgents gets a reference to the given []string and assigns it to the Agents field.
func (o *AccountGroupRequest) SetAgents(v []string) {
	o.Agents = v
}

func (o AccountGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountGroupName"] = o.AccountGroupName
	if !core.IsNil(o.Agents) {
		toSerialize["agents"] = o.Agents
	}
	return toSerialize, nil
}

func (o *AccountGroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accountGroupName",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAccountGroupRequest := _AccountGroupRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountGroupRequest)

	if err != nil {
		return err
	}

	*o = AccountGroupRequest(varAccountGroupRequest)

	return err
}

type NullableAccountGroupRequest struct {
	value *AccountGroupRequest
	isSet bool
}

func (v NullableAccountGroupRequest) Get() *AccountGroupRequest {
	return v.value
}

func (v *NullableAccountGroupRequest) Set(val *AccountGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountGroupRequest(val *AccountGroupRequest) *NullableAccountGroupRequest {
	return &NullableAccountGroupRequest{value: val, isSet: true}
}

func (v NullableAccountGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


