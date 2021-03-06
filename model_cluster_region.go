/*
Camunda Cloud Management API

Manage Camunda Cloud Clusters and API Clients via API.

API version: 1.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ClusterRegion The data center where your Camunda Cloud cluster is running.
type ClusterRegion struct {
	Name string `json:"name"`
	Uuid string `json:"uuid"`
}

// NewClusterRegion instantiates a new ClusterRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterRegion(name string, uuid string) *ClusterRegion {
	this := ClusterRegion{}
	this.Name = name
	this.Uuid = uuid
	return &this
}

// NewClusterRegionWithDefaults instantiates a new ClusterRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterRegionWithDefaults() *ClusterRegion {
	this := ClusterRegion{}
	return &this
}

// GetName returns the Name field value
func (o *ClusterRegion) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterRegion) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClusterRegion) SetName(v string) {
	o.Name = v
}

// GetUuid returns the Uuid field value
func (o *ClusterRegion) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *ClusterRegion) GetUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *ClusterRegion) SetUuid(v string) {
	o.Uuid = v
}

func (o ClusterRegion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableClusterRegion struct {
	value *ClusterRegion
	isSet bool
}

func (v NullableClusterRegion) Get() *ClusterRegion {
	return v.value
}

func (v *NullableClusterRegion) Set(val *ClusterRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterRegion(val *ClusterRegion) *NullableClusterRegion {
	return &NullableClusterRegion{value: val, isSet: true}
}

func (v NullableClusterRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


