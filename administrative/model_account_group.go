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
)

// checks if the AccountGroup type satisfies the MappedNullable interface at compile time
var _ core.MappedNullable = &AccountGroup{}

// AccountGroup struct for AccountGroup
type AccountGroup struct {
	// A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint.
	Aid *string `json:"aid,omitempty"`
	// Account group name
	AccountGroupName *string `json:"accountGroupName,omitempty"`
}

// NewAccountGroup instantiates a new AccountGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountGroup() *AccountGroup {
	this := AccountGroup{}
	return &this
}

// NewAccountGroupWithDefaults instantiates a new AccountGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountGroupWithDefaults() *AccountGroup {
	this := AccountGroup{}
	return &this
}

// GetAid returns the Aid field value if set, zero value otherwise.
func (o *AccountGroup) GetAid() string {
	if o == nil || core.IsNil(o.Aid) {
		var ret string
		return ret
	}
	return *o.Aid
}

// GetAidOk returns a tuple with the Aid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountGroup) GetAidOk() (*string, bool) {
	if o == nil || core.IsNil(o.Aid) {
		return nil, false
	}
	return o.Aid, true
}

// HasAid returns a boolean if a field has been set.
func (o *AccountGroup) HasAid() bool {
	if o != nil && !core.IsNil(o.Aid) {
		return true
	}

	return false
}

// SetAid gets a reference to the given string and assigns it to the Aid field.
func (o *AccountGroup) SetAid(v string) {
	o.Aid = &v
}

// GetAccountGroupName returns the AccountGroupName field value if set, zero value otherwise.
func (o *AccountGroup) GetAccountGroupName() string {
	if o == nil || core.IsNil(o.AccountGroupName) {
		var ret string
		return ret
	}
	return *o.AccountGroupName
}

// GetAccountGroupNameOk returns a tuple with the AccountGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountGroup) GetAccountGroupNameOk() (*string, bool) {
	if o == nil || core.IsNil(o.AccountGroupName) {
		return nil, false
	}
	return o.AccountGroupName, true
}

// HasAccountGroupName returns a boolean if a field has been set.
func (o *AccountGroup) HasAccountGroupName() bool {
	if o != nil && !core.IsNil(o.AccountGroupName) {
		return true
	}

	return false
}

// SetAccountGroupName gets a reference to the given string and assigns it to the AccountGroupName field.
func (o *AccountGroup) SetAccountGroupName(v string) {
	o.AccountGroupName = &v
}

func (o AccountGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !core.IsNil(o.Aid) {
		toSerialize["aid"] = o.Aid
	}
	if !core.IsNil(o.AccountGroupName) {
		toSerialize["accountGroupName"] = o.AccountGroupName
	}
	return toSerialize, nil
}

type NullableAccountGroup struct {
	value *AccountGroup
	isSet bool
}

func (v NullableAccountGroup) Get() *AccountGroup {
	return v.value
}

func (v *NullableAccountGroup) Set(val *AccountGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountGroup(val *AccountGroup) *NullableAccountGroup {
	return &NullableAccountGroup{value: val, isSet: true}
}

func (v NullableAccountGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


