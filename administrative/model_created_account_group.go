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

// checks if the CreatedAccountGroup type satisfies the MappedNullable interface at compile time
var _ core.MappedNullable = &CreatedAccountGroup{}

// CreatedAccountGroup struct for CreatedAccountGroup
type CreatedAccountGroup struct {
	// A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint.
	Aid *string `json:"aid,omitempty"`
	// Account group name
	AccountGroupName *string `json:"accountGroupName,omitempty"`
	// Indicates whether the requested aid is the context of the current account.
	IsCurrentAccountGroup *bool `json:"isCurrentAccountGroup,omitempty"`
	// Indicates whether the aid is the default one for the requesting user.
	IsDefaultAccountGroup *bool `json:"isDefaultAccountGroup,omitempty"`
	// (Optional) The name of the organization associated with the account group.
	OrganizationName *string `json:"organizationName,omitempty"`
	// (Optional) The ID for the organization associated with the account group.
	OrgId *string `json:"orgId,omitempty"`
	Users []UserAccountGroup `json:"users,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
}

// NewCreatedAccountGroup instantiates a new CreatedAccountGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatedAccountGroup() *CreatedAccountGroup {
	this := CreatedAccountGroup{}
	return &this
}

// NewCreatedAccountGroupWithDefaults instantiates a new CreatedAccountGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatedAccountGroupWithDefaults() *CreatedAccountGroup {
	this := CreatedAccountGroup{}
	return &this
}

// GetAid returns the Aid field value if set, zero value otherwise.
func (o *CreatedAccountGroup) GetAid() string {
	if o == nil || core.IsNil(o.Aid) {
		var ret string
		return ret
	}
	return *o.Aid
}

// GetAidOk returns a tuple with the Aid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedAccountGroup) GetAidOk() (*string, bool) {
	if o == nil || core.IsNil(o.Aid) {
		return nil, false
	}
	return o.Aid, true
}

// HasAid returns a boolean if a field has been set.
func (o *CreatedAccountGroup) HasAid() bool {
	if o != nil && !core.IsNil(o.Aid) {
		return true
	}

	return false
}

// SetAid gets a reference to the given string and assigns it to the Aid field.
func (o *CreatedAccountGroup) SetAid(v string) {
	o.Aid = &v
}

// GetAccountGroupName returns the AccountGroupName field value if set, zero value otherwise.
func (o *CreatedAccountGroup) GetAccountGroupName() string {
	if o == nil || core.IsNil(o.AccountGroupName) {
		var ret string
		return ret
	}
	return *o.AccountGroupName
}

// GetAccountGroupNameOk returns a tuple with the AccountGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedAccountGroup) GetAccountGroupNameOk() (*string, bool) {
	if o == nil || core.IsNil(o.AccountGroupName) {
		return nil, false
	}
	return o.AccountGroupName, true
}

// HasAccountGroupName returns a boolean if a field has been set.
func (o *CreatedAccountGroup) HasAccountGroupName() bool {
	if o != nil && !core.IsNil(o.AccountGroupName) {
		return true
	}

	return false
}

// SetAccountGroupName gets a reference to the given string and assigns it to the AccountGroupName field.
func (o *CreatedAccountGroup) SetAccountGroupName(v string) {
	o.AccountGroupName = &v
}

// GetIsCurrentAccountGroup returns the IsCurrentAccountGroup field value if set, zero value otherwise.
func (o *CreatedAccountGroup) GetIsCurrentAccountGroup() bool {
	if o == nil || core.IsNil(o.IsCurrentAccountGroup) {
		var ret bool
		return ret
	}
	return *o.IsCurrentAccountGroup
}

// GetIsCurrentAccountGroupOk returns a tuple with the IsCurrentAccountGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedAccountGroup) GetIsCurrentAccountGroupOk() (*bool, bool) {
	if o == nil || core.IsNil(o.IsCurrentAccountGroup) {
		return nil, false
	}
	return o.IsCurrentAccountGroup, true
}

// HasIsCurrentAccountGroup returns a boolean if a field has been set.
func (o *CreatedAccountGroup) HasIsCurrentAccountGroup() bool {
	if o != nil && !core.IsNil(o.IsCurrentAccountGroup) {
		return true
	}

	return false
}

// SetIsCurrentAccountGroup gets a reference to the given bool and assigns it to the IsCurrentAccountGroup field.
func (o *CreatedAccountGroup) SetIsCurrentAccountGroup(v bool) {
	o.IsCurrentAccountGroup = &v
}

// GetIsDefaultAccountGroup returns the IsDefaultAccountGroup field value if set, zero value otherwise.
func (o *CreatedAccountGroup) GetIsDefaultAccountGroup() bool {
	if o == nil || core.IsNil(o.IsDefaultAccountGroup) {
		var ret bool
		return ret
	}
	return *o.IsDefaultAccountGroup
}

// GetIsDefaultAccountGroupOk returns a tuple with the IsDefaultAccountGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedAccountGroup) GetIsDefaultAccountGroupOk() (*bool, bool) {
	if o == nil || core.IsNil(o.IsDefaultAccountGroup) {
		return nil, false
	}
	return o.IsDefaultAccountGroup, true
}

// HasIsDefaultAccountGroup returns a boolean if a field has been set.
func (o *CreatedAccountGroup) HasIsDefaultAccountGroup() bool {
	if o != nil && !core.IsNil(o.IsDefaultAccountGroup) {
		return true
	}

	return false
}

// SetIsDefaultAccountGroup gets a reference to the given bool and assigns it to the IsDefaultAccountGroup field.
func (o *CreatedAccountGroup) SetIsDefaultAccountGroup(v bool) {
	o.IsDefaultAccountGroup = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise.
func (o *CreatedAccountGroup) GetOrganizationName() string {
	if o == nil || core.IsNil(o.OrganizationName) {
		var ret string
		return ret
	}
	return *o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedAccountGroup) GetOrganizationNameOk() (*string, bool) {
	if o == nil || core.IsNil(o.OrganizationName) {
		return nil, false
	}
	return o.OrganizationName, true
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *CreatedAccountGroup) HasOrganizationName() bool {
	if o != nil && !core.IsNil(o.OrganizationName) {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.
func (o *CreatedAccountGroup) SetOrganizationName(v string) {
	o.OrganizationName = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *CreatedAccountGroup) GetOrgId() string {
	if o == nil || core.IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedAccountGroup) GetOrgIdOk() (*string, bool) {
	if o == nil || core.IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *CreatedAccountGroup) HasOrgId() bool {
	if o != nil && !core.IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *CreatedAccountGroup) SetOrgId(v string) {
	o.OrgId = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *CreatedAccountGroup) GetUsers() []UserAccountGroup {
	if o == nil || core.IsNil(o.Users) {
		var ret []UserAccountGroup
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedAccountGroup) GetUsersOk() ([]UserAccountGroup, bool) {
	if o == nil || core.IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *CreatedAccountGroup) HasUsers() bool {
	if o != nil && !core.IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []UserAccountGroup and assigns it to the Users field.
func (o *CreatedAccountGroup) SetUsers(v []UserAccountGroup) {
	o.Users = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CreatedAccountGroup) GetLinks() SelfLinks {
	if o == nil || core.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedAccountGroup) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || core.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CreatedAccountGroup) HasLinks() bool {
	if o != nil && !core.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *CreatedAccountGroup) SetLinks(v SelfLinks) {
	o.Links = &v
}

func (o CreatedAccountGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatedAccountGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !core.IsNil(o.Aid) {
		toSerialize["aid"] = o.Aid
	}
	if !core.IsNil(o.AccountGroupName) {
		toSerialize["accountGroupName"] = o.AccountGroupName
	}
	if !core.IsNil(o.IsCurrentAccountGroup) {
		toSerialize["isCurrentAccountGroup"] = o.IsCurrentAccountGroup
	}
	if !core.IsNil(o.IsDefaultAccountGroup) {
		toSerialize["isDefaultAccountGroup"] = o.IsDefaultAccountGroup
	}
	if !core.IsNil(o.OrganizationName) {
		toSerialize["organizationName"] = o.OrganizationName
	}
	if !core.IsNil(o.OrgId) {
		toSerialize["orgId"] = o.OrgId
	}
	if !core.IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !core.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableCreatedAccountGroup struct {
	value *CreatedAccountGroup
	isSet bool
}

func (v NullableCreatedAccountGroup) Get() *CreatedAccountGroup {
	return v.value
}

func (v *NullableCreatedAccountGroup) Set(val *CreatedAccountGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatedAccountGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatedAccountGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatedAccountGroup(val *CreatedAccountGroup) *NullableCreatedAccountGroup {
	return &NullableCreatedAccountGroup{value: val, isSet: true}
}

func (v NullableCreatedAccountGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatedAccountGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


