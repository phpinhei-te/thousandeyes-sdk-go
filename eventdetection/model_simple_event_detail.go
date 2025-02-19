/*
Event Detection API

 Event detection occurs when ThousandEyes identifies that error signals related to a component (proxy, network node, AS, server etc) have deviated from the baselines established by events. * To determine this, ThousandEyes takes the test results from all accounts groups within an organization, and analyzes that data. * Noisy test results (those that have too many errors in a short window) are removed until they stabilize, and the rest of the results are tagged with the components associated with that test result (for example, proxy, network, or server). * Next, any increase in failures from the test results and each component helps in determining the problem domain and which component may be at fault. * When this failure rate increases beyond a pre-defined threshold (set by the algorithm), an event is triggered and an email notification is sent to the user (if they've enabled email alerts).  With the Events API, you can perform the following tasks on the ThousandEyes platform: * **Retrieve Events**: Obtain a list of events and detailed information for each event. For more information about events, see [Event Detection](https://docs.thousandeyes.com/product-documentation/event-detection).

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eventdetection

import (
	"encoding/json"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"time"
)

// checks if the SimpleEventDetail type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SimpleEventDetail{}

// SimpleEventDetail struct for SimpleEventDetail
type SimpleEventDetail struct {
	// A unique ID for each event.
	Id *string `json:"id,omitempty"`
	// Event type name. Examples include, Local Agent Issue, Network Path Issue, Network Outage, DNS Issue, Server Issue, and Proxy Issue.
	TypeName *string     `json:"typeName,omitempty"`
	State    *EventState `json:"state,omitempty"`
	// The start date and time (in UTC, ISO 8601 format) when the event was first detected.
	StartDate *time.Time `json:"startDate,omitempty"`
	// The end date and time (in UTC, ISO 8601 format) when the event was resolved (due to timeout). This value is populated for \"ongoing\" events.
	EndDate  *time.Time          `json:"endDate,omitempty"`
	Severity *EventAlertSeverity `json:"severity,omitempty"`
}

// NewSimpleEventDetail instantiates a new SimpleEventDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleEventDetail() *SimpleEventDetail {
	this := SimpleEventDetail{}
	return &this
}

// NewSimpleEventDetailWithDefaults instantiates a new SimpleEventDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleEventDetailWithDefaults() *SimpleEventDetail {
	this := SimpleEventDetail{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SimpleEventDetail) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleEventDetail) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SimpleEventDetail) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SimpleEventDetail) SetId(v string) {
	o.Id = &v
}

// GetTypeName returns the TypeName field value if set, zero value otherwise.
func (o *SimpleEventDetail) GetTypeName() string {
	if o == nil || utils.IsNil(o.TypeName) {
		var ret string
		return ret
	}
	return *o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleEventDetail) GetTypeNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TypeName) {
		return nil, false
	}
	return o.TypeName, true
}

// HasTypeName returns a boolean if a field has been set.
func (o *SimpleEventDetail) HasTypeName() bool {
	if o != nil && !utils.IsNil(o.TypeName) {
		return true
	}

	return false
}

// SetTypeName gets a reference to the given string and assigns it to the TypeName field.
func (o *SimpleEventDetail) SetTypeName(v string) {
	o.TypeName = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SimpleEventDetail) GetState() EventState {
	if o == nil || utils.IsNil(o.State) {
		var ret EventState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleEventDetail) GetStateOk() (*EventState, bool) {
	if o == nil || utils.IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SimpleEventDetail) HasState() bool {
	if o != nil && !utils.IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given EventState and assigns it to the State field.
func (o *SimpleEventDetail) SetState(v EventState) {
	o.State = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SimpleEventDetail) GetStartDate() time.Time {
	if o == nil || utils.IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleEventDetail) GetStartDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SimpleEventDetail) HasStartDate() bool {
	if o != nil && !utils.IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *SimpleEventDetail) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *SimpleEventDetail) GetEndDate() time.Time {
	if o == nil || utils.IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleEventDetail) GetEndDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *SimpleEventDetail) HasEndDate() bool {
	if o != nil && !utils.IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *SimpleEventDetail) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *SimpleEventDetail) GetSeverity() EventAlertSeverity {
	if o == nil || utils.IsNil(o.Severity) {
		var ret EventAlertSeverity
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleEventDetail) GetSeverityOk() (*EventAlertSeverity, bool) {
	if o == nil || utils.IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *SimpleEventDetail) HasSeverity() bool {
	if o != nil && !utils.IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given EventAlertSeverity and assigns it to the Severity field.
func (o *SimpleEventDetail) SetSeverity(v EventAlertSeverity) {
	o.Severity = &v
}

func (o SimpleEventDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimpleEventDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.TypeName) {
		toSerialize["typeName"] = o.TypeName
	}
	if !utils.IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !utils.IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !utils.IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !utils.IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	return toSerialize, nil
}

type NullableSimpleEventDetail struct {
	value *SimpleEventDetail
	isSet bool
}

func (v NullableSimpleEventDetail) Get() *SimpleEventDetail {
	return v.value
}

func (v *NullableSimpleEventDetail) Set(val *SimpleEventDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleEventDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleEventDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleEventDetail(val *SimpleEventDetail) *NullableSimpleEventDetail {
	return &NullableSimpleEventDetail{value: val, isSet: true}
}

func (v NullableSimpleEventDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleEventDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
