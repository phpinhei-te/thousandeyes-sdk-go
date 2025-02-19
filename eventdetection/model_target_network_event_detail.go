/*
Event Detection API

 Event detection occurs when ThousandEyes identifies that error signals related to a component (proxy, network node, AS, server etc) have deviated from the baselines established by events. * To determine this, ThousandEyes takes the test results from all accounts groups within an organization, and analyzes that data. * Noisy test results (those that have too many errors in a short window) are removed until they stabilize, and the rest of the results are tagged with the components associated with that test result (for example, proxy, network, or server). * Next, any increase in failures from the test results and each component helps in determining the problem domain and which component may be at fault. * When this failure rate increases beyond a pre-defined threshold (set by the algorithm), an event is triggered and an email notification is sent to the user (if they've enabled email alerts).  With the Events API, you can perform the following tasks on the ThousandEyes platform: * **Retrieve Events**: Obtain a list of events and detailed information for each event. For more information about events, see [Event Detection](https://docs.thousandeyes.com/product-documentation/event-detection).

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eventdetection

import (
	"encoding/json"
	"fmt"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"time"
)

// checks if the TargetNetworkEventDetail type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TargetNetworkEventDetail{}

// TargetNetworkEventDetail struct for TargetNetworkEventDetail
type TargetNetworkEventDetail struct {
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
	// A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint.
	Aid *string `json:"aid,omitempty"`
	// A brief summary describing the cause of the event.
	Summary         *string          `json:"summary,omitempty"`
	AffectedTests   *AffectedTests   `json:"affectedTests,omitempty"`
	AffectedTargets *AffectedTargets `json:"affectedTargets,omitempty"`
	AffectedAgents  *AffectedAgents  `json:"affectedAgents,omitempty"`
	// The cause of the error.
	Cause []string   `json:"cause,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
	// Target network event type.
	Type     string                      `json:"type"`
	Grouping *TargetNetworkEventGrouping `json:"grouping,omitempty"`
}

type _TargetNetworkEventDetail TargetNetworkEventDetail

// NewTargetNetworkEventDetail instantiates a new TargetNetworkEventDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetNetworkEventDetail(type_ string) *TargetNetworkEventDetail {
	this := TargetNetworkEventDetail{}
	this.Type = type_
	return &this
}

// NewTargetNetworkEventDetailWithDefaults instantiates a new TargetNetworkEventDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetNetworkEventDetailWithDefaults() *TargetNetworkEventDetail {
	this := TargetNetworkEventDetail{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TargetNetworkEventDetail) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetNetworkEventDetail) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TargetNetworkEventDetail) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TargetNetworkEventDetail) SetId(v string) {
	o.Id = &v
}

// GetTypeName returns the TypeName field value if set, zero value otherwise.
func (o *TargetNetworkEventDetail) GetTypeName() string {
	if o == nil || utils.IsNil(o.TypeName) {
		var ret string
		return ret
	}
	return *o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetNetworkEventDetail) GetTypeNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TypeName) {
		return nil, false
	}
	return o.TypeName, true
}

// HasTypeName returns a boolean if a field has been set.
func (o *TargetNetworkEventDetail) HasTypeName() bool {
	if o != nil && !utils.IsNil(o.TypeName) {
		return true
	}

	return false
}

// SetTypeName gets a reference to the given string and assigns it to the TypeName field.
func (o *TargetNetworkEventDetail) SetTypeName(v string) {
	o.TypeName = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *TargetNetworkEventDetail) GetState() EventState {
	if o == nil || utils.IsNil(o.State) {
		var ret EventState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetNetworkEventDetail) GetStateOk() (*EventState, bool) {
	if o == nil || utils.IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *TargetNetworkEventDetail) HasState() bool {
	if o != nil && !utils.IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given EventState and assigns it to the State field.
func (o *TargetNetworkEventDetail) SetState(v EventState) {
	o.State = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *TargetNetworkEventDetail) GetStartDate() time.Time {
	if o == nil || utils.IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetNetworkEventDetail) GetStartDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *TargetNetworkEventDetail) HasStartDate() bool {
	if o != nil && !utils.IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *TargetNetworkEventDetail) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *TargetNetworkEventDetail) GetEndDate() time.Time {
	if o == nil || utils.IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetNetworkEventDetail) GetEndDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *TargetNetworkEventDetail) HasEndDate() bool {
	if o != nil && !utils.IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *TargetNetworkEventDetail) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *TargetNetworkEventDetail) GetSeverity() EventAlertSeverity {
	if o == nil || utils.IsNil(o.Severity) {
		var ret EventAlertSeverity
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetNetworkEventDetail) GetSeverityOk() (*EventAlertSeverity, bool) {
	if o == nil || utils.IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *TargetNetworkEventDetail) HasSeverity() bool {
	if o != nil && !utils.IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given EventAlertSeverity and assigns it to the Severity field.
func (o *TargetNetworkEventDetail) SetSeverity(v EventAlertSeverity) {
	o.Severity = &v
}

// GetAid returns the Aid field value if set, zero value otherwise.
func (o *TargetNetworkEventDetail) GetAid() string {
	if o == nil || utils.IsNil(o.Aid) {
		var ret string
		return ret
	}
	return *o.Aid
}

// GetAidOk returns a tuple with the Aid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetNetworkEventDetail) GetAidOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Aid) {
		return nil, false
	}
	return o.Aid, true
}

// HasAid returns a boolean if a field has been set.
func (o *TargetNetworkEventDetail) HasAid() bool {
	if o != nil && !utils.IsNil(o.Aid) {
		return true
	}

	return false
}

// SetAid gets a reference to the given string and assigns it to the Aid field.
func (o *TargetNetworkEventDetail) SetAid(v string) {
	o.Aid = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *TargetNetworkEventDetail) GetSummary() string {
	if o == nil || utils.IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetNetworkEventDetail) GetSummaryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *TargetNetworkEventDetail) HasSummary() bool {
	if o != nil && !utils.IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *TargetNetworkEventDetail) SetSummary(v string) {
	o.Summary = &v
}

// GetAffectedTests returns the AffectedTests field value if set, zero value otherwise.
func (o *TargetNetworkEventDetail) GetAffectedTests() AffectedTests {
	if o == nil || utils.IsNil(o.AffectedTests) {
		var ret AffectedTests
		return ret
	}
	return *o.AffectedTests
}

// GetAffectedTestsOk returns a tuple with the AffectedTests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetNetworkEventDetail) GetAffectedTestsOk() (*AffectedTests, bool) {
	if o == nil || utils.IsNil(o.AffectedTests) {
		return nil, false
	}
	return o.AffectedTests, true
}

// HasAffectedTests returns a boolean if a field has been set.
func (o *TargetNetworkEventDetail) HasAffectedTests() bool {
	if o != nil && !utils.IsNil(o.AffectedTests) {
		return true
	}

	return false
}

// SetAffectedTests gets a reference to the given AffectedTests and assigns it to the AffectedTests field.
func (o *TargetNetworkEventDetail) SetAffectedTests(v AffectedTests) {
	o.AffectedTests = &v
}

// GetAffectedTargets returns the AffectedTargets field value if set, zero value otherwise.
func (o *TargetNetworkEventDetail) GetAffectedTargets() AffectedTargets {
	if o == nil || utils.IsNil(o.AffectedTargets) {
		var ret AffectedTargets
		return ret
	}
	return *o.AffectedTargets
}

// GetAffectedTargetsOk returns a tuple with the AffectedTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetNetworkEventDetail) GetAffectedTargetsOk() (*AffectedTargets, bool) {
	if o == nil || utils.IsNil(o.AffectedTargets) {
		return nil, false
	}
	return o.AffectedTargets, true
}

// HasAffectedTargets returns a boolean if a field has been set.
func (o *TargetNetworkEventDetail) HasAffectedTargets() bool {
	if o != nil && !utils.IsNil(o.AffectedTargets) {
		return true
	}

	return false
}

// SetAffectedTargets gets a reference to the given AffectedTargets and assigns it to the AffectedTargets field.
func (o *TargetNetworkEventDetail) SetAffectedTargets(v AffectedTargets) {
	o.AffectedTargets = &v
}

// GetAffectedAgents returns the AffectedAgents field value if set, zero value otherwise.
func (o *TargetNetworkEventDetail) GetAffectedAgents() AffectedAgents {
	if o == nil || utils.IsNil(o.AffectedAgents) {
		var ret AffectedAgents
		return ret
	}
	return *o.AffectedAgents
}

// GetAffectedAgentsOk returns a tuple with the AffectedAgents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetNetworkEventDetail) GetAffectedAgentsOk() (*AffectedAgents, bool) {
	if o == nil || utils.IsNil(o.AffectedAgents) {
		return nil, false
	}
	return o.AffectedAgents, true
}

// HasAffectedAgents returns a boolean if a field has been set.
func (o *TargetNetworkEventDetail) HasAffectedAgents() bool {
	if o != nil && !utils.IsNil(o.AffectedAgents) {
		return true
	}

	return false
}

// SetAffectedAgents gets a reference to the given AffectedAgents and assigns it to the AffectedAgents field.
func (o *TargetNetworkEventDetail) SetAffectedAgents(v AffectedAgents) {
	o.AffectedAgents = &v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *TargetNetworkEventDetail) GetCause() []string {
	if o == nil || utils.IsNil(o.Cause) {
		var ret []string
		return ret
	}
	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetNetworkEventDetail) GetCauseOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *TargetNetworkEventDetail) HasCause() bool {
	if o != nil && !utils.IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given []string and assigns it to the Cause field.
func (o *TargetNetworkEventDetail) SetCause(v []string) {
	o.Cause = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TargetNetworkEventDetail) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetNetworkEventDetail) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TargetNetworkEventDetail) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *TargetNetworkEventDetail) SetLinks(v SelfLinks) {
	o.Links = &v
}

// GetType returns the Type field value
func (o *TargetNetworkEventDetail) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TargetNetworkEventDetail) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TargetNetworkEventDetail) SetType(v string) {
	o.Type = v
}

// GetGrouping returns the Grouping field value if set, zero value otherwise.
func (o *TargetNetworkEventDetail) GetGrouping() TargetNetworkEventGrouping {
	if o == nil || utils.IsNil(o.Grouping) {
		var ret TargetNetworkEventGrouping
		return ret
	}
	return *o.Grouping
}

// GetGroupingOk returns a tuple with the Grouping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetNetworkEventDetail) GetGroupingOk() (*TargetNetworkEventGrouping, bool) {
	if o == nil || utils.IsNil(o.Grouping) {
		return nil, false
	}
	return o.Grouping, true
}

// HasGrouping returns a boolean if a field has been set.
func (o *TargetNetworkEventDetail) HasGrouping() bool {
	if o != nil && !utils.IsNil(o.Grouping) {
		return true
	}

	return false
}

// SetGrouping gets a reference to the given TargetNetworkEventGrouping and assigns it to the Grouping field.
func (o *TargetNetworkEventDetail) SetGrouping(v TargetNetworkEventGrouping) {
	o.Grouping = &v
}

func (o TargetNetworkEventDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetNetworkEventDetail) ToMap() (map[string]interface{}, error) {
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
	if !utils.IsNil(o.Aid) {
		toSerialize["aid"] = o.Aid
	}
	if !utils.IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !utils.IsNil(o.AffectedTests) {
		toSerialize["affectedTests"] = o.AffectedTests
	}
	if !utils.IsNil(o.AffectedTargets) {
		toSerialize["affectedTargets"] = o.AffectedTargets
	}
	if !utils.IsNil(o.AffectedAgents) {
		toSerialize["affectedAgents"] = o.AffectedAgents
	}
	if !utils.IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	toSerialize["type"] = o.Type
	if !utils.IsNil(o.Grouping) {
		toSerialize["grouping"] = o.Grouping
	}
	return toSerialize, nil
}

func (o *TargetNetworkEventDetail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTargetNetworkEventDetail := _TargetNetworkEventDetail{}

	err = json.Unmarshal(data, &varTargetNetworkEventDetail)

	if err != nil {
		return err
	}

	*o = TargetNetworkEventDetail(varTargetNetworkEventDetail)

	return err
}

type NullableTargetNetworkEventDetail struct {
	value *TargetNetworkEventDetail
	isSet bool
}

func (v NullableTargetNetworkEventDetail) Get() *TargetNetworkEventDetail {
	return v.value
}

func (v *NullableTargetNetworkEventDetail) Set(val *TargetNetworkEventDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetNetworkEventDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetNetworkEventDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetNetworkEventDetail(val *TargetNetworkEventDetail) *NullableTargetNetworkEventDetail {
	return &NullableTargetNetworkEventDetail{value: val, isSet: true}
}

func (v NullableTargetNetworkEventDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetNetworkEventDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
