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

// checks if the ErrorDetail type satisfies the MappedNullable interface at compile time
var _ core.MappedNullable = &ErrorDetail{}

// ErrorDetail struct for ErrorDetail
type ErrorDetail struct {
	Code *ErrorDetailCode `json:"code,omitempty"`
	// Description for the agent error.
	Description *string `json:"description,omitempty"`
}

// NewErrorDetail instantiates a new ErrorDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorDetail() *ErrorDetail {
	this := ErrorDetail{}
	return &this
}

// NewErrorDetailWithDefaults instantiates a new ErrorDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorDetailWithDefaults() *ErrorDetail {
	this := ErrorDetail{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ErrorDetail) GetCode() ErrorDetailCode {
	if o == nil || core.IsNil(o.Code) {
		var ret ErrorDetailCode
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorDetail) GetCodeOk() (*ErrorDetailCode, bool) {
	if o == nil || core.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ErrorDetail) HasCode() bool {
	if o != nil && !core.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given ErrorDetailCode and assigns it to the Code field.
func (o *ErrorDetail) SetCode(v ErrorDetailCode) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ErrorDetail) GetDescription() string {
	if o == nil || core.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorDetail) GetDescriptionOk() (*string, bool) {
	if o == nil || core.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ErrorDetail) HasDescription() bool {
	if o != nil && !core.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ErrorDetail) SetDescription(v string) {
	o.Description = &v
}

func (o ErrorDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !core.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !core.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableErrorDetail struct {
	value *ErrorDetail
	isSet bool
}

func (v NullableErrorDetail) Get() *ErrorDetail {
	return v.value
}

func (v *NullableErrorDetail) Set(val *ErrorDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorDetail(val *ErrorDetail) *NullableErrorDetail {
	return &NullableErrorDetail{value: val, isSet: true}
}

func (v NullableErrorDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


