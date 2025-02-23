// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/aws/client (interfaces: WafClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	waf "github.com/aws/aws-sdk-go-v2/service/waf"
	gomock "github.com/golang/mock/gomock"
)

// MockWafClient is a mock of WafClient interface.
type MockWafClient struct {
	ctrl     *gomock.Controller
	recorder *MockWafClientMockRecorder
}

// MockWafClientMockRecorder is the mock recorder for MockWafClient.
type MockWafClientMockRecorder struct {
	mock *MockWafClient
}

// NewMockWafClient creates a new mock instance.
func NewMockWafClient(ctrl *gomock.Controller) *MockWafClient {
	mock := &MockWafClient{ctrl: ctrl}
	mock.recorder = &MockWafClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWafClient) EXPECT() *MockWafClientMockRecorder {
	return m.recorder
}

// GetLoggingConfiguration mocks base method.
func (m *MockWafClient) GetLoggingConfiguration(arg0 context.Context, arg1 *waf.GetLoggingConfigurationInput, arg2 ...func(*waf.Options)) (*waf.GetLoggingConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLoggingConfiguration", varargs...)
	ret0, _ := ret[0].(*waf.GetLoggingConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoggingConfiguration indicates an expected call of GetLoggingConfiguration.
func (mr *MockWafClientMockRecorder) GetLoggingConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoggingConfiguration", reflect.TypeOf((*MockWafClient)(nil).GetLoggingConfiguration), varargs...)
}

// GetRule mocks base method.
func (m *MockWafClient) GetRule(arg0 context.Context, arg1 *waf.GetRuleInput, arg2 ...func(*waf.Options)) (*waf.GetRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRule", varargs...)
	ret0, _ := ret[0].(*waf.GetRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRule indicates an expected call of GetRule.
func (mr *MockWafClientMockRecorder) GetRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRule", reflect.TypeOf((*MockWafClient)(nil).GetRule), varargs...)
}

// GetRuleGroup mocks base method.
func (m *MockWafClient) GetRuleGroup(arg0 context.Context, arg1 *waf.GetRuleGroupInput, arg2 ...func(*waf.Options)) (*waf.GetRuleGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRuleGroup", varargs...)
	ret0, _ := ret[0].(*waf.GetRuleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRuleGroup indicates an expected call of GetRuleGroup.
func (mr *MockWafClientMockRecorder) GetRuleGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRuleGroup", reflect.TypeOf((*MockWafClient)(nil).GetRuleGroup), varargs...)
}

// GetWebACL mocks base method.
func (m *MockWafClient) GetWebACL(arg0 context.Context, arg1 *waf.GetWebACLInput, arg2 ...func(*waf.Options)) (*waf.GetWebACLOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWebACL", varargs...)
	ret0, _ := ret[0].(*waf.GetWebACLOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWebACL indicates an expected call of GetWebACL.
func (mr *MockWafClientMockRecorder) GetWebACL(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWebACL", reflect.TypeOf((*MockWafClient)(nil).GetWebACL), varargs...)
}

// ListActivatedRulesInRuleGroup mocks base method.
func (m *MockWafClient) ListActivatedRulesInRuleGroup(arg0 context.Context, arg1 *waf.ListActivatedRulesInRuleGroupInput, arg2 ...func(*waf.Options)) (*waf.ListActivatedRulesInRuleGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListActivatedRulesInRuleGroup", varargs...)
	ret0, _ := ret[0].(*waf.ListActivatedRulesInRuleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListActivatedRulesInRuleGroup indicates an expected call of ListActivatedRulesInRuleGroup.
func (mr *MockWafClientMockRecorder) ListActivatedRulesInRuleGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListActivatedRulesInRuleGroup", reflect.TypeOf((*MockWafClient)(nil).ListActivatedRulesInRuleGroup), varargs...)
}

// ListRuleGroups mocks base method.
func (m *MockWafClient) ListRuleGroups(arg0 context.Context, arg1 *waf.ListRuleGroupsInput, arg2 ...func(*waf.Options)) (*waf.ListRuleGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRuleGroups", varargs...)
	ret0, _ := ret[0].(*waf.ListRuleGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRuleGroups indicates an expected call of ListRuleGroups.
func (mr *MockWafClientMockRecorder) ListRuleGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRuleGroups", reflect.TypeOf((*MockWafClient)(nil).ListRuleGroups), varargs...)
}

// ListRules mocks base method.
func (m *MockWafClient) ListRules(arg0 context.Context, arg1 *waf.ListRulesInput, arg2 ...func(*waf.Options)) (*waf.ListRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRules", varargs...)
	ret0, _ := ret[0].(*waf.ListRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRules indicates an expected call of ListRules.
func (mr *MockWafClientMockRecorder) ListRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRules", reflect.TypeOf((*MockWafClient)(nil).ListRules), varargs...)
}

// ListSubscribedRuleGroups mocks base method.
func (m *MockWafClient) ListSubscribedRuleGroups(arg0 context.Context, arg1 *waf.ListSubscribedRuleGroupsInput, arg2 ...func(*waf.Options)) (*waf.ListSubscribedRuleGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSubscribedRuleGroups", varargs...)
	ret0, _ := ret[0].(*waf.ListSubscribedRuleGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubscribedRuleGroups indicates an expected call of ListSubscribedRuleGroups.
func (mr *MockWafClientMockRecorder) ListSubscribedRuleGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubscribedRuleGroups", reflect.TypeOf((*MockWafClient)(nil).ListSubscribedRuleGroups), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockWafClient) ListTagsForResource(arg0 context.Context, arg1 *waf.ListTagsForResourceInput, arg2 ...func(*waf.Options)) (*waf.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*waf.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockWafClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockWafClient)(nil).ListTagsForResource), varargs...)
}

// ListWebACLs mocks base method.
func (m *MockWafClient) ListWebACLs(arg0 context.Context, arg1 *waf.ListWebACLsInput, arg2 ...func(*waf.Options)) (*waf.ListWebACLsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListWebACLs", varargs...)
	ret0, _ := ret[0].(*waf.ListWebACLsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWebACLs indicates an expected call of ListWebACLs.
func (mr *MockWafClientMockRecorder) ListWebACLs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWebACLs", reflect.TypeOf((*MockWafClient)(nil).ListWebACLs), varargs...)
}
