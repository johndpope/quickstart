/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.31.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// NumbersEFTNullable struct for NumbersEFTNullable
type NumbersEFTNullable struct {
	// The Plaid account ID associated with the account numbers
	AccountId string `json:"account_id"`
	// The EFT account number for the account
	Account string `json:"account"`
	// The EFT institution number for the account
	Institution string `json:"institution"`
	// The EFT branch number for the account
	Branch string `json:"branch"`
}

// NewNumbersEFTNullable instantiates a new NumbersEFTNullable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumbersEFTNullable(accountId string, account string, institution string, branch string) *NumbersEFTNullable {
	this := NumbersEFTNullable{}
	this.AccountId = accountId
	this.Account = account
	this.Institution = institution
	this.Branch = branch
	return &this
}

// NewNumbersEFTNullableWithDefaults instantiates a new NumbersEFTNullable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumbersEFTNullableWithDefaults() *NumbersEFTNullable {
	this := NumbersEFTNullable{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *NumbersEFTNullable) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *NumbersEFTNullable) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *NumbersEFTNullable) SetAccountId(v string) {
	o.AccountId = v
}

// GetAccount returns the Account field value
func (o *NumbersEFTNullable) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *NumbersEFTNullable) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *NumbersEFTNullable) SetAccount(v string) {
	o.Account = v
}

// GetInstitution returns the Institution field value
func (o *NumbersEFTNullable) GetInstitution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Institution
}

// GetInstitutionOk returns a tuple with the Institution field value
// and a boolean to check if the value has been set.
func (o *NumbersEFTNullable) GetInstitutionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Institution, true
}

// SetInstitution sets field value
func (o *NumbersEFTNullable) SetInstitution(v string) {
	o.Institution = v
}

// GetBranch returns the Branch field value
func (o *NumbersEFTNullable) GetBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Branch
}

// GetBranchOk returns a tuple with the Branch field value
// and a boolean to check if the value has been set.
func (o *NumbersEFTNullable) GetBranchOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Branch, true
}

// SetBranch sets field value
func (o *NumbersEFTNullable) SetBranch(v string) {
	o.Branch = v
}

func (o NumbersEFTNullable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["account"] = o.Account
	}
	if true {
		toSerialize["institution"] = o.Institution
	}
	if true {
		toSerialize["branch"] = o.Branch
	}
	return json.Marshal(toSerialize)
}

type NullableNumbersEFTNullable struct {
	value *NumbersEFTNullable
	isSet bool
}

func (v NullableNumbersEFTNullable) Get() *NumbersEFTNullable {
	return v.value
}

func (v *NullableNumbersEFTNullable) Set(val *NumbersEFTNullable) {
	v.value = val
	v.isSet = true
}

func (v NullableNumbersEFTNullable) IsSet() bool {
	return v.isSet
}

func (v *NullableNumbersEFTNullable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumbersEFTNullable(val *NumbersEFTNullable) *NullableNumbersEFTNullable {
	return &NullableNumbersEFTNullable{value: val, isSet: true}
}

func (v NullableNumbersEFTNullable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumbersEFTNullable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

