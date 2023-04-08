// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

// AuthorizationPolicySpecApplyConfiguration represents an declarative configuration of the AuthorizationPolicySpec type for use
// with apply.
type AuthorizationPolicySpecApplyConfiguration struct {
	Action    *string                                     `json:"action,omitempty"`
	Rules     []AuthorizationPolicyRuleApplyConfiguration `json:"rules,omitempty"`
	Samples   *float32                                    `json:"samples,omitempty"`
	Order     *float32                                    `json:"order,omitempty"`
	MatchType *string                                     `json:"matchType,omitempty"`
}

// AuthorizationPolicySpecApplyConfiguration constructs an declarative configuration of the AuthorizationPolicySpec type for use with
// apply.
func AuthorizationPolicySpec() *AuthorizationPolicySpecApplyConfiguration {
	return &AuthorizationPolicySpecApplyConfiguration{}
}

// WithAction sets the Action field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Action field is set to the value of the last call.
func (b *AuthorizationPolicySpecApplyConfiguration) WithAction(value string) *AuthorizationPolicySpecApplyConfiguration {
	b.Action = &value
	return b
}

// WithRules adds the given value to the Rules field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Rules field.
func (b *AuthorizationPolicySpecApplyConfiguration) WithRules(values ...*AuthorizationPolicyRuleApplyConfiguration) *AuthorizationPolicySpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithRules")
		}
		b.Rules = append(b.Rules, *values[i])
	}
	return b
}

// WithSamples sets the Samples field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Samples field is set to the value of the last call.
func (b *AuthorizationPolicySpecApplyConfiguration) WithSamples(value float32) *AuthorizationPolicySpecApplyConfiguration {
	b.Samples = &value
	return b
}

// WithOrder sets the Order field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Order field is set to the value of the last call.
func (b *AuthorizationPolicySpecApplyConfiguration) WithOrder(value float32) *AuthorizationPolicySpecApplyConfiguration {
	b.Order = &value
	return b
}

// WithMatchType sets the MatchType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MatchType field is set to the value of the last call.
func (b *AuthorizationPolicySpecApplyConfiguration) WithMatchType(value string) *AuthorizationPolicySpecApplyConfiguration {
	b.MatchType = &value
	return b
}