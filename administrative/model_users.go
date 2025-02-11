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

// checks if the Users type satisfies the MappedNullable interface at compile time
var _ core.MappedNullable = &Users{}

// Users struct for Users
type Users struct {
	Users []ExtendedUser `json:"users,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
}

// NewUsers instantiates a new Users object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsers() *Users {
	this := Users{}
	return &this
}

// NewUsersWithDefaults instantiates a new Users object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersWithDefaults() *Users {
	this := Users{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *Users) GetUsers() []ExtendedUser {
	if o == nil || core.IsNil(o.Users) {
		var ret []ExtendedUser
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Users) GetUsersOk() ([]ExtendedUser, bool) {
	if o == nil || core.IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *Users) HasUsers() bool {
	if o != nil && !core.IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []ExtendedUser and assigns it to the Users field.
func (o *Users) SetUsers(v []ExtendedUser) {
	o.Users = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Users) GetLinks() SelfLinks {
	if o == nil || core.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Users) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || core.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Users) HasLinks() bool {
	if o != nil && !core.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *Users) SetLinks(v SelfLinks) {
	o.Links = &v
}

func (o Users) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Users) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !core.IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !core.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableUsers struct {
	value *Users
	isSet bool
}

func (v NullableUsers) Get() *Users {
	return v.value
}

func (v *NullableUsers) Set(val *Users) {
	v.value = val
	v.isSet = true
}

func (v NullableUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsers(val *Users) *NullableUsers {
	return &NullableUsers{value: val, isSet: true}
}

func (v NullableUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


