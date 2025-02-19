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
)

// EventState Indicates the state of the event, whether it is ongoing (active) or has been resolved.
type EventState string

// List of EventState
const (
	EVENTSTATE_ACTIVE   EventState = "active"
	EVENTSTATE_RESOLVED EventState = "resolved"
)

// All allowed values of EventState enum
var AllowedEventStateEnumValues = []EventState{
	"active",
	"resolved",
}

func (v *EventState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventState(value)
	for _, existing := range AllowedEventStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventState", value)
}

// NewEventStateFromValue returns a pointer to a valid EventState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventStateFromValue(v string) (*EventState, error) {
	ev := EventState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventState: valid values are %v", v, AllowedEventStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventState) IsValid() bool {
	for _, existing := range AllowedEventStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventState value
func (v EventState) Ptr() *EventState {
	return &v
}

type NullableEventState struct {
	value *EventState
	isSet bool
}

func (v NullableEventState) Get() *EventState {
	return v.value
}

func (v *NullableEventState) Set(val *EventState) {
	v.value = val
	v.isSet = true
}

func (v NullableEventState) IsSet() bool {
	return v.isSet
}

func (v *NullableEventState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventState(val *EventState) *NullableEventState {
	return &NullableEventState{value: val, isSet: true}
}

func (v NullableEventState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
