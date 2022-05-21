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

// ClusterChannel The software channel your cluster is running on.
type ClusterChannel struct {
	Name string `json:"name"`
	Uuid string `json:"uuid"`
}

// NewClusterChannel instantiates a new ClusterChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterChannel(name string, uuid string) *ClusterChannel {
	this := ClusterChannel{}
	this.Name = name
	this.Uuid = uuid
	return &this
}

// NewClusterChannelWithDefaults instantiates a new ClusterChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterChannelWithDefaults() *ClusterChannel {
	this := ClusterChannel{}
	return &this
}

// GetName returns the Name field value
func (o *ClusterChannel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterChannel) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClusterChannel) SetName(v string) {
	o.Name = v
}

// GetUuid returns the Uuid field value
func (o *ClusterChannel) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *ClusterChannel) GetUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *ClusterChannel) SetUuid(v string) {
	o.Uuid = v
}

func (o ClusterChannel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableClusterChannel struct {
	value *ClusterChannel
	isSet bool
}

func (v NullableClusterChannel) Get() *ClusterChannel {
	return v.value
}

func (v *NullableClusterChannel) Set(val *ClusterChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterChannel(val *ClusterChannel) *NullableClusterChannel {
	return &NullableClusterChannel{value: val, isSet: true}
}

func (v NullableClusterChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

