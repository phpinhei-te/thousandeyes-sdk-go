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

// checks if the UserDetail type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &UserDetail{}

// UserDetail struct for UserDetail
type UserDetail struct {
	// User's display name.
	Name *string `json:"name,omitempty"`
	// User's email address.
	Email *string `json:"email,omitempty"`
	// Unique ID of the user.
	Uid *string `json:"uid,omitempty"`
	// UTC date the user registered their account (ISO date-time format).
	DateRegistered    *time.Time    `json:"dateRegistered,omitempty"`
	LoginAccountGroup *AccountGroup `json:"loginAccountGroup,omitempty"`
	// UTC last login of the user (ISO date-time format).
	LastLogin            *time.Time         `json:"lastLogin,omitempty"`
	AccountGroupRoles    []AccountGroupRole `json:"accountGroupRoles,omitempty"`
	AllAccountGroupRoles []Role             `json:"allAccountGroupRoles,omitempty"`
	Links                *SelfLinks         `json:"_links,omitempty"`
}

// NewUserDetail instantiates a new UserDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDetail() *UserDetail {
	this := UserDetail{}
	return &this
}

// NewUserDetailWithDefaults instantiates a new UserDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDetailWithDefaults() *UserDetail {
	this := UserDetail{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserDetail) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDetail) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserDetail) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserDetail) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserDetail) GetEmail() string {
	if o == nil || utils.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDetail) GetEmailOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserDetail) HasEmail() bool {
	if o != nil && !utils.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserDetail) SetEmail(v string) {
	o.Email = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *UserDetail) GetUid() string {
	if o == nil || utils.IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDetail) GetUidOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *UserDetail) HasUid() bool {
	if o != nil && !utils.IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *UserDetail) SetUid(v string) {
	o.Uid = &v
}

// GetDateRegistered returns the DateRegistered field value if set, zero value otherwise.
func (o *UserDetail) GetDateRegistered() time.Time {
	if o == nil || utils.IsNil(o.DateRegistered) {
		var ret time.Time
		return ret
	}
	return *o.DateRegistered
}

// GetDateRegisteredOk returns a tuple with the DateRegistered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDetail) GetDateRegisteredOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.DateRegistered) {
		return nil, false
	}
	return o.DateRegistered, true
}

// HasDateRegistered returns a boolean if a field has been set.
func (o *UserDetail) HasDateRegistered() bool {
	if o != nil && !utils.IsNil(o.DateRegistered) {
		return true
	}

	return false
}

// SetDateRegistered gets a reference to the given time.Time and assigns it to the DateRegistered field.
func (o *UserDetail) SetDateRegistered(v time.Time) {
	o.DateRegistered = &v
}

// GetLoginAccountGroup returns the LoginAccountGroup field value if set, zero value otherwise.
func (o *UserDetail) GetLoginAccountGroup() AccountGroup {
	if o == nil || utils.IsNil(o.LoginAccountGroup) {
		var ret AccountGroup
		return ret
	}
	return *o.LoginAccountGroup
}

// GetLoginAccountGroupOk returns a tuple with the LoginAccountGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDetail) GetLoginAccountGroupOk() (*AccountGroup, bool) {
	if o == nil || utils.IsNil(o.LoginAccountGroup) {
		return nil, false
	}
	return o.LoginAccountGroup, true
}

// HasLoginAccountGroup returns a boolean if a field has been set.
func (o *UserDetail) HasLoginAccountGroup() bool {
	if o != nil && !utils.IsNil(o.LoginAccountGroup) {
		return true
	}

	return false
}

// SetLoginAccountGroup gets a reference to the given AccountGroup and assigns it to the LoginAccountGroup field.
func (o *UserDetail) SetLoginAccountGroup(v AccountGroup) {
	o.LoginAccountGroup = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise.
func (o *UserDetail) GetLastLogin() time.Time {
	if o == nil || utils.IsNil(o.LastLogin) {
		var ret time.Time
		return ret
	}
	return *o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDetail) GetLastLoginOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.LastLogin) {
		return nil, false
	}
	return o.LastLogin, true
}

// HasLastLogin returns a boolean if a field has been set.
func (o *UserDetail) HasLastLogin() bool {
	if o != nil && !utils.IsNil(o.LastLogin) {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given time.Time and assigns it to the LastLogin field.
func (o *UserDetail) SetLastLogin(v time.Time) {
	o.LastLogin = &v
}

// GetAccountGroupRoles returns the AccountGroupRoles field value if set, zero value otherwise.
func (o *UserDetail) GetAccountGroupRoles() []AccountGroupRole {
	if o == nil || utils.IsNil(o.AccountGroupRoles) {
		var ret []AccountGroupRole
		return ret
	}
	return o.AccountGroupRoles
}

// GetAccountGroupRolesOk returns a tuple with the AccountGroupRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDetail) GetAccountGroupRolesOk() ([]AccountGroupRole, bool) {
	if o == nil || utils.IsNil(o.AccountGroupRoles) {
		return nil, false
	}
	return o.AccountGroupRoles, true
}

// HasAccountGroupRoles returns a boolean if a field has been set.
func (o *UserDetail) HasAccountGroupRoles() bool {
	if o != nil && !utils.IsNil(o.AccountGroupRoles) {
		return true
	}

	return false
}

// SetAccountGroupRoles gets a reference to the given []AccountGroupRole and assigns it to the AccountGroupRoles field.
func (o *UserDetail) SetAccountGroupRoles(v []AccountGroupRole) {
	o.AccountGroupRoles = v
}

// GetAllAccountGroupRoles returns the AllAccountGroupRoles field value if set, zero value otherwise.
func (o *UserDetail) GetAllAccountGroupRoles() []Role {
	if o == nil || utils.IsNil(o.AllAccountGroupRoles) {
		var ret []Role
		return ret
	}
	return o.AllAccountGroupRoles
}

// GetAllAccountGroupRolesOk returns a tuple with the AllAccountGroupRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDetail) GetAllAccountGroupRolesOk() ([]Role, bool) {
	if o == nil || utils.IsNil(o.AllAccountGroupRoles) {
		return nil, false
	}
	return o.AllAccountGroupRoles, true
}

// HasAllAccountGroupRoles returns a boolean if a field has been set.
func (o *UserDetail) HasAllAccountGroupRoles() bool {
	if o != nil && !utils.IsNil(o.AllAccountGroupRoles) {
		return true
	}

	return false
}

// SetAllAccountGroupRoles gets a reference to the given []Role and assigns it to the AllAccountGroupRoles field.
func (o *UserDetail) SetAllAccountGroupRoles(v []Role) {
	o.AllAccountGroupRoles = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UserDetail) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDetail) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UserDetail) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *UserDetail) SetLinks(v SelfLinks) {
	o.Links = &v
}

func (o UserDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !utils.IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	if !utils.IsNil(o.DateRegistered) {
		toSerialize["dateRegistered"] = o.DateRegistered
	}
	if !utils.IsNil(o.LoginAccountGroup) {
		toSerialize["loginAccountGroup"] = o.LoginAccountGroup
	}
	if !utils.IsNil(o.LastLogin) {
		toSerialize["lastLogin"] = o.LastLogin
	}
	if !utils.IsNil(o.AccountGroupRoles) {
		toSerialize["accountGroupRoles"] = o.AccountGroupRoles
	}
	if !utils.IsNil(o.AllAccountGroupRoles) {
		toSerialize["allAccountGroupRoles"] = o.AllAccountGroupRoles
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableUserDetail struct {
	value *UserDetail
	isSet bool
}

func (v NullableUserDetail) Get() *UserDetail {
	return v.value
}

func (v *NullableUserDetail) Set(val *UserDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDetail(val *UserDetail) *NullableUserDetail {
	return &NullableUserDetail{value: val, isSet: true}
}

func (v NullableUserDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
