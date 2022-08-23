/*
Pulumi Service API

The Pulumi Service API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// OperationStatus OperationStatus describes the state of an operation being performed on a Pulumi stack.
type OperationStatus struct {
	Kind *UpdateKind `json:"kind,omitempty"`
	Author *string `json:"author,omitempty"`
	Started *int64 `json:"started,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OperationStatus OperationStatus

// NewOperationStatus instantiates a new OperationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationStatus() *OperationStatus {
	this := OperationStatus{}
	return &this
}

// NewOperationStatusWithDefaults instantiates a new OperationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationStatusWithDefaults() *OperationStatus {
	this := OperationStatus{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *OperationStatus) GetKind() UpdateKind {
	if o == nil || o.Kind == nil {
		var ret UpdateKind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationStatus) GetKindOk() (*UpdateKind, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *OperationStatus) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given UpdateKind and assigns it to the Kind field.
func (o *OperationStatus) SetKind(v UpdateKind) {
	o.Kind = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *OperationStatus) GetAuthor() string {
	if o == nil || o.Author == nil {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationStatus) GetAuthorOk() (*string, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *OperationStatus) HasAuthor() bool {
	if o != nil && o.Author != nil {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *OperationStatus) SetAuthor(v string) {
	o.Author = &v
}

// GetStarted returns the Started field value if set, zero value otherwise.
func (o *OperationStatus) GetStarted() int64 {
	if o == nil || o.Started == nil {
		var ret int64
		return ret
	}
	return *o.Started
}

// GetStartedOk returns a tuple with the Started field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationStatus) GetStartedOk() (*int64, bool) {
	if o == nil || o.Started == nil {
		return nil, false
	}
	return o.Started, true
}

// HasStarted returns a boolean if a field has been set.
func (o *OperationStatus) HasStarted() bool {
	if o != nil && o.Started != nil {
		return true
	}

	return false
}

// SetStarted gets a reference to the given int64 and assigns it to the Started field.
func (o *OperationStatus) SetStarted(v int64) {
	o.Started = &v
}

func (o OperationStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if o.Started != nil {
		toSerialize["started"] = o.Started
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OperationStatus) UnmarshalJSON(bytes []byte) (err error) {
	varOperationStatus := _OperationStatus{}

	if err = json.Unmarshal(bytes, &varOperationStatus); err == nil {
		*o = OperationStatus(varOperationStatus)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "kind")
		delete(additionalProperties, "author")
		delete(additionalProperties, "started")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOperationStatus struct {
	value *OperationStatus
	isSet bool
}

func (v NullableOperationStatus) Get() *OperationStatus {
	return v.value
}

func (v *NullableOperationStatus) Set(val *OperationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationStatus(val *OperationStatus) *NullableOperationStatus {
	return &NullableOperationStatus{value: val, isSet: true}
}

func (v NullableOperationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

