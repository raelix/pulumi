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

// StackSummary struct for StackSummary
type StackSummary struct {
	OrgName *string `json:"orgName,omitempty"`
	ProjectName *string `json:"projectName,omitempty"`
	StackName *string `json:"stackName,omitempty"`
	LastUpdate *int64 `json:"lastUpdate,omitempty"`
	ResourceCount *int32 `json:"resourceCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StackSummary StackSummary

// NewStackSummary instantiates a new StackSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStackSummary() *StackSummary {
	this := StackSummary{}
	return &this
}

// NewStackSummaryWithDefaults instantiates a new StackSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStackSummaryWithDefaults() *StackSummary {
	this := StackSummary{}
	return &this
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *StackSummary) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackSummary) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *StackSummary) HasOrgName() bool {
	if o != nil && o.OrgName != nil {
		return true
	}

	return false
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *StackSummary) SetOrgName(v string) {
	o.OrgName = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise.
func (o *StackSummary) GetProjectName() string {
	if o == nil || o.ProjectName == nil {
		var ret string
		return ret
	}
	return *o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackSummary) GetProjectNameOk() (*string, bool) {
	if o == nil || o.ProjectName == nil {
		return nil, false
	}
	return o.ProjectName, true
}

// HasProjectName returns a boolean if a field has been set.
func (o *StackSummary) HasProjectName() bool {
	if o != nil && o.ProjectName != nil {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given string and assigns it to the ProjectName field.
func (o *StackSummary) SetProjectName(v string) {
	o.ProjectName = &v
}

// GetStackName returns the StackName field value if set, zero value otherwise.
func (o *StackSummary) GetStackName() string {
	if o == nil || o.StackName == nil {
		var ret string
		return ret
	}
	return *o.StackName
}

// GetStackNameOk returns a tuple with the StackName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackSummary) GetStackNameOk() (*string, bool) {
	if o == nil || o.StackName == nil {
		return nil, false
	}
	return o.StackName, true
}

// HasStackName returns a boolean if a field has been set.
func (o *StackSummary) HasStackName() bool {
	if o != nil && o.StackName != nil {
		return true
	}

	return false
}

// SetStackName gets a reference to the given string and assigns it to the StackName field.
func (o *StackSummary) SetStackName(v string) {
	o.StackName = &v
}

// GetLastUpdate returns the LastUpdate field value if set, zero value otherwise.
func (o *StackSummary) GetLastUpdate() int64 {
	if o == nil || o.LastUpdate == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackSummary) GetLastUpdateOk() (*int64, bool) {
	if o == nil || o.LastUpdate == nil {
		return nil, false
	}
	return o.LastUpdate, true
}

// HasLastUpdate returns a boolean if a field has been set.
func (o *StackSummary) HasLastUpdate() bool {
	if o != nil && o.LastUpdate != nil {
		return true
	}

	return false
}

// SetLastUpdate gets a reference to the given int64 and assigns it to the LastUpdate field.
func (o *StackSummary) SetLastUpdate(v int64) {
	o.LastUpdate = &v
}

// GetResourceCount returns the ResourceCount field value if set, zero value otherwise.
func (o *StackSummary) GetResourceCount() int32 {
	if o == nil || o.ResourceCount == nil {
		var ret int32
		return ret
	}
	return *o.ResourceCount
}

// GetResourceCountOk returns a tuple with the ResourceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackSummary) GetResourceCountOk() (*int32, bool) {
	if o == nil || o.ResourceCount == nil {
		return nil, false
	}
	return o.ResourceCount, true
}

// HasResourceCount returns a boolean if a field has been set.
func (o *StackSummary) HasResourceCount() bool {
	if o != nil && o.ResourceCount != nil {
		return true
	}

	return false
}

// SetResourceCount gets a reference to the given int32 and assigns it to the ResourceCount field.
func (o *StackSummary) SetResourceCount(v int32) {
	o.ResourceCount = &v
}

func (o StackSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.ProjectName != nil {
		toSerialize["projectName"] = o.ProjectName
	}
	if o.StackName != nil {
		toSerialize["stackName"] = o.StackName
	}
	if o.LastUpdate != nil {
		toSerialize["lastUpdate"] = o.LastUpdate
	}
	if o.ResourceCount != nil {
		toSerialize["resourceCount"] = o.ResourceCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StackSummary) UnmarshalJSON(bytes []byte) (err error) {
	varStackSummary := _StackSummary{}

	if err = json.Unmarshal(bytes, &varStackSummary); err == nil {
		*o = StackSummary(varStackSummary)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "orgName")
		delete(additionalProperties, "projectName")
		delete(additionalProperties, "stackName")
		delete(additionalProperties, "lastUpdate")
		delete(additionalProperties, "resourceCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStackSummary struct {
	value *StackSummary
	isSet bool
}

func (v NullableStackSummary) Get() *StackSummary {
	return v.value
}

func (v *NullableStackSummary) Set(val *StackSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableStackSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableStackSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStackSummary(val *StackSummary) *NullableStackSummary {
	return &NullableStackSummary{value: val, isSet: true}
}

func (v NullableStackSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStackSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

