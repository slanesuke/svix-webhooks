/*
 * Svix API
 *
 * Welcome to the Svix API documentation!  Useful links: [Homepage](https://www.svix.com) | [Support email](mailto:support+docs@svix.com) | [Slack Community](https://www.svix.com/slack/)  # Introduction  This is the reference documentation and schemas for the Svix API. For tutorials and other documentation please refer to [the documentation](https://docs.svix.com).  ## Main concepts  In Svix you have four important entities you will be interacting with:  - `messages`: these are the webhooks being sent. They can have contents and a few other properties. - `application`: this is where `messages` are sent to. Usually you want to create one application for each of your users. - `endpoint`: endpoints are the URLs messages will be sent to. Each application can have multiple `endpoints` and each message sent to that application will be sent to all of them (unless they are not subscribed to the sent event type). - `event-type`: event types are identifiers denoting the type of the message being sent. Event types are primarily used to decide which events are sent to which endpoint.   ## Authentication  Get your authentication token (`AUTH_TOKEN`) from the [Svix dashboard](https://dashboard.svix.com) and use it as part of the `Authorization` header as such: `Authorization: Bearer ${AUTH_TOKEN}`.  <SecurityDefinitions />   ## Code samples  The code samples assume you already have the respective libraries installed and you know how to use them. For the latest information on how to do that, please refer to [the documentation](https://docs.svix.com/).   ## Cross-Origin Resource Sharing  This API features Cross-Origin Resource Sharing (CORS) implemented in compliance with [W3C spec](https://www.w3.org/TR/cors/). And that allows cross-domain communication from the browser. All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site. 
 *
 * API version: 1.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// MessageEndpointOut struct for MessageEndpointOut
type MessageEndpointOut struct {
	Url string `json:"url"`
	Version int32 `json:"version"`
	Description *string `json:"description,omitempty"`
	FilterTypes *[]string `json:"filterTypes,omitempty"`
	Id string `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Status MessageStatus `json:"status"`
}

// NewMessageEndpointOut instantiates a new MessageEndpointOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageEndpointOut(url string, version int32, id string, createdAt time.Time, status MessageStatus, ) *MessageEndpointOut {
	this := MessageEndpointOut{}
	this.Url = url
	this.Version = version
	var description string = ""
	this.Description = &description
	this.Id = id
	this.CreatedAt = createdAt
	this.Status = status
	return &this
}

// NewMessageEndpointOutWithDefaults instantiates a new MessageEndpointOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageEndpointOutWithDefaults() *MessageEndpointOut {
	this := MessageEndpointOut{}
	var description string = ""
	this.Description = &description
	return &this
}

// GetUrl returns the Url field value
func (o *MessageEndpointOut) GetUrl() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *MessageEndpointOut) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *MessageEndpointOut) SetUrl(v string) {
	o.Url = v
}

// GetVersion returns the Version field value
func (o *MessageEndpointOut) GetVersion() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *MessageEndpointOut) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *MessageEndpointOut) SetVersion(v int32) {
	o.Version = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MessageEndpointOut) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEndpointOut) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MessageEndpointOut) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MessageEndpointOut) SetDescription(v string) {
	o.Description = &v
}

// GetFilterTypes returns the FilterTypes field value if set, zero value otherwise.
func (o *MessageEndpointOut) GetFilterTypes() []string {
	if o == nil || o.FilterTypes == nil {
		var ret []string
		return ret
	}
	return *o.FilterTypes
}

// GetFilterTypesOk returns a tuple with the FilterTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEndpointOut) GetFilterTypesOk() (*[]string, bool) {
	if o == nil || o.FilterTypes == nil {
		return nil, false
	}
	return o.FilterTypes, true
}

// HasFilterTypes returns a boolean if a field has been set.
func (o *MessageEndpointOut) HasFilterTypes() bool {
	if o != nil && o.FilterTypes != nil {
		return true
	}

	return false
}

// SetFilterTypes gets a reference to the given []string and assigns it to the FilterTypes field.
func (o *MessageEndpointOut) SetFilterTypes(v []string) {
	o.FilterTypes = &v
}

// GetId returns the Id field value
func (o *MessageEndpointOut) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MessageEndpointOut) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MessageEndpointOut) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *MessageEndpointOut) GetCreatedAt() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *MessageEndpointOut) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *MessageEndpointOut) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetStatus returns the Status field value
func (o *MessageEndpointOut) GetStatus() MessageStatus {
	if o == nil  {
		var ret MessageStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *MessageEndpointOut) GetStatusOk() (*MessageStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *MessageEndpointOut) SetStatus(v MessageStatus) {
	o.Status = v
}

func (o MessageEndpointOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.FilterTypes != nil {
		toSerialize["filterTypes"] = o.FilterTypes
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableMessageEndpointOut struct {
	value *MessageEndpointOut
	isSet bool
}

func (v NullableMessageEndpointOut) Get() *MessageEndpointOut {
	return v.value
}

func (v *NullableMessageEndpointOut) Set(val *MessageEndpointOut) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageEndpointOut) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageEndpointOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageEndpointOut(val *MessageEndpointOut) *NullableMessageEndpointOut {
	return &NullableMessageEndpointOut{value: val, isSet: true}
}

func (v NullableMessageEndpointOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageEndpointOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


