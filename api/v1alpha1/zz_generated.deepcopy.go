//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKey) DeepCopyInto(out *APIKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKey.
func (in *APIKey) DeepCopy() *APIKey {
	if in == nil {
		return nil
	}
	out := new(APIKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeyList) DeepCopyInto(out *APIKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeyList.
func (in *APIKeyList) DeepCopy() *APIKeyList {
	if in == nil {
		return nil
	}
	out := new(APIKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeySpec) DeepCopyInto(out *APIKeySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeySpec.
func (in *APIKeySpec) DeepCopy() *APIKeySpec {
	if in == nil {
		return nil
	}
	out := new(APIKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeyStatus) DeepCopyInto(out *APIKeyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeyStatus.
func (in *APIKeyStatus) DeepCopy() *APIKeyStatus {
	if in == nil {
		return nil
	}
	out := new(APIKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertManager) DeepCopyInto(out *AlertManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertManager.
func (in *AlertManager) DeepCopy() *AlertManager {
	if in == nil {
		return nil
	}
	out := new(AlertManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertManagerList) DeepCopyInto(out *AlertManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlertManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertManagerList.
func (in *AlertManagerList) DeepCopy() *AlertManagerList {
	if in == nil {
		return nil
	}
	out := new(AlertManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertManagerSpec) DeepCopyInto(out *AlertManagerSpec) {
	*out = *in
	if in.ContactPoints != nil {
		in, out := &in.ContactPoints, &out.ContactPoints
		*out = make([]ContactPoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Routing != nil {
		in, out := &in.Routing, &out.Routing
		*out = make([]RoutingPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertManagerSpec.
func (in *AlertManagerSpec) DeepCopy() *AlertManagerSpec {
	if in == nil {
		return nil
	}
	out := new(AlertManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertManagerStatus) DeepCopyInto(out *AlertManagerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertManagerStatus.
func (in *AlertManagerStatus) DeepCopy() *AlertManagerStatus {
	if in == nil {
		return nil
	}
	out := new(AlertManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicAuth) DeepCopyInto(out *BasicAuth) {
	*out = *in
	in.Username.DeepCopyInto(&out.Username)
	in.Password.DeepCopyInto(&out.Password)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicAuth.
func (in *BasicAuth) DeepCopy() *BasicAuth {
	if in == nil {
		return nil
	}
	out := new(BasicAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactPoint) DeepCopyInto(out *ContactPoint) {
	*out = *in
	if in.Contacts != nil {
		in, out := &in.Contacts, &out.Contacts
		*out = make([]ContactPointType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactPoint.
func (in *ContactPoint) DeepCopy() *ContactPoint {
	if in == nil {
		return nil
	}
	out := new(ContactPoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactPointType) DeepCopyInto(out *ContactPointType) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(EmailContactType)
		(*in).DeepCopyInto(*out)
	}
	if in.Slack != nil {
		in, out := &in.Slack, &out.Slack
		*out = new(SlackContactType)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactPointType.
func (in *ContactPointType) DeepCopy() *ContactPointType {
	if in == nil {
		return nil
	}
	out := new(ContactPointType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Datasource) DeepCopyInto(out *Datasource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Datasource.
func (in *Datasource) DeepCopy() *Datasource {
	if in == nil {
		return nil
	}
	out := new(Datasource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Datasource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasourceList) DeepCopyInto(out *DatasourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Datasource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasourceList.
func (in *DatasourceList) DeepCopy() *DatasourceList {
	if in == nil {
		return nil
	}
	out := new(DatasourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatasourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasourceSpec) DeepCopyInto(out *DatasourceSpec) {
	*out = *in
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(PrometheusDatasource)
		(*in).DeepCopyInto(*out)
	}
	if in.Stackdriver != nil {
		in, out := &in.Stackdriver, &out.Stackdriver
		*out = new(StackdriverDatasource)
		(*in).DeepCopyInto(*out)
	}
	if in.Jaeger != nil {
		in, out := &in.Jaeger, &out.Jaeger
		*out = new(JaegerDatasource)
		(*in).DeepCopyInto(*out)
	}
	if in.Loki != nil {
		in, out := &in.Loki, &out.Loki
		*out = new(LokiDatasource)
		(*in).DeepCopyInto(*out)
	}
	if in.Tempo != nil {
		in, out := &in.Tempo, &out.Tempo
		*out = new(TempoDatasource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasourceSpec.
func (in *DatasourceSpec) DeepCopy() *DatasourceSpec {
	if in == nil {
		return nil
	}
	out := new(DatasourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasourceStatus) DeepCopyInto(out *DatasourceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasourceStatus.
func (in *DatasourceStatus) DeepCopy() *DatasourceStatus {
	if in == nil {
		return nil
	}
	out := new(DatasourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailContactType) DeepCopyInto(out *EmailContactType) {
	*out = *in
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailContactType.
func (in *EmailContactType) DeepCopy() *EmailContactType {
	if in == nil {
		return nil
	}
	out := new(EmailContactType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerDatasource) DeepCopyInto(out *JaegerDatasource) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(bool)
		**out = **in
	}
	if in.ForwardOauth != nil {
		in, out := &in.ForwardOauth, &out.ForwardOauth
		*out = new(bool)
		**out = **in
	}
	if in.ForwardCredentials != nil {
		in, out := &in.ForwardCredentials, &out.ForwardCredentials
		*out = new(bool)
		**out = **in
	}
	if in.SkipTLSVerify != nil {
		in, out := &in.SkipTLSVerify, &out.SkipTLSVerify
		*out = new(bool)
		**out = **in
	}
	if in.ForwardCookies != nil {
		in, out := &in.ForwardCookies, &out.ForwardCookies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		*out = new(BasicAuth)
		(*in).DeepCopyInto(*out)
	}
	if in.CACertificate != nil {
		in, out := &in.CACertificate, &out.CACertificate
		*out = new(ValueOrRef)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeGraph != nil {
		in, out := &in.NodeGraph, &out.NodeGraph
		*out = new(bool)
		**out = **in
	}
	if in.TraceToLogs != nil {
		in, out := &in.TraceToLogs, &out.TraceToLogs
		*out = new(TraceToLogs)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerDatasource.
func (in *JaegerDatasource) DeepCopy() *JaegerDatasource {
	if in == nil {
		return nil
	}
	out := new(JaegerDatasource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelsMatchingRule) DeepCopyInto(out *LabelsMatchingRule) {
	*out = *in
	if in.Eq != nil {
		in, out := &in.Eq, &out.Eq
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Neq != nil {
		in, out := &in.Neq, &out.Neq
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Matches != nil {
		in, out := &in.Matches, &out.Matches
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NotMatches != nil {
		in, out := &in.NotMatches, &out.NotMatches
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelsMatchingRule.
func (in *LabelsMatchingRule) DeepCopy() *LabelsMatchingRule {
	if in == nil {
		return nil
	}
	out := new(LabelsMatchingRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LokiDatasource) DeepCopyInto(out *LokiDatasource) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(bool)
		**out = **in
	}
	if in.ForwardOauth != nil {
		in, out := &in.ForwardOauth, &out.ForwardOauth
		*out = new(bool)
		**out = **in
	}
	if in.ForwardCredentials != nil {
		in, out := &in.ForwardCredentials, &out.ForwardCredentials
		*out = new(bool)
		**out = **in
	}
	if in.SkipTLSVerify != nil {
		in, out := &in.SkipTLSVerify, &out.SkipTLSVerify
		*out = new(bool)
		**out = **in
	}
	if in.ForwardCookies != nil {
		in, out := &in.ForwardCookies, &out.ForwardCookies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		*out = new(BasicAuth)
		(*in).DeepCopyInto(*out)
	}
	if in.CACertificate != nil {
		in, out := &in.CACertificate, &out.CACertificate
		*out = new(ValueOrRef)
		(*in).DeepCopyInto(*out)
	}
	if in.MaximumLines != nil {
		in, out := &in.MaximumLines, &out.MaximumLines
		*out = new(int)
		**out = **in
	}
	if in.DerivedFields != nil {
		in, out := &in.DerivedFields, &out.DerivedFields
		*out = make([]LokiDerivedField, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LokiDatasource.
func (in *LokiDatasource) DeepCopy() *LokiDatasource {
	if in == nil {
		return nil
	}
	out := new(LokiDatasource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LokiDerivedField) DeepCopyInto(out *LokiDerivedField) {
	*out = *in
	if in.Datasource != nil {
		in, out := &in.Datasource, &out.Datasource
		*out = new(ValueOrDatasourceRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LokiDerivedField.
func (in *LokiDerivedField) DeepCopy() *LokiDerivedField {
	if in == nil {
		return nil
	}
	out := new(LokiDerivedField)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusDatasource) DeepCopyInto(out *PrometheusDatasource) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(bool)
		**out = **in
	}
	if in.ForwardOauth != nil {
		in, out := &in.ForwardOauth, &out.ForwardOauth
		*out = new(bool)
		**out = **in
	}
	if in.ForwardCredentials != nil {
		in, out := &in.ForwardCredentials, &out.ForwardCredentials
		*out = new(bool)
		**out = **in
	}
	if in.SkipTLSVerify != nil {
		in, out := &in.SkipTLSVerify, &out.SkipTLSVerify
		*out = new(bool)
		**out = **in
	}
	if in.ForwardCookies != nil {
		in, out := &in.ForwardCookies, &out.ForwardCookies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		*out = new(BasicAuth)
		(*in).DeepCopyInto(*out)
	}
	if in.CACertificate != nil {
		in, out := &in.CACertificate, &out.CACertificate
		*out = new(ValueOrRef)
		(*in).DeepCopyInto(*out)
	}
	if in.Exemplars != nil {
		in, out := &in.Exemplars, &out.Exemplars
		*out = make([]PrometheusExemplar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusDatasource.
func (in *PrometheusDatasource) DeepCopy() *PrometheusDatasource {
	if in == nil {
		return nil
	}
	out := new(PrometheusDatasource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusExemplar) DeepCopyInto(out *PrometheusExemplar) {
	*out = *in
	if in.Datasource != nil {
		in, out := &in.Datasource, &out.Datasource
		*out = new(ValueOrDatasourceRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusExemplar.
func (in *PrometheusExemplar) DeepCopy() *PrometheusExemplar {
	if in == nil {
		return nil
	}
	out := new(PrometheusExemplar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingPolicy) DeepCopyInto(out *RoutingPolicy) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]LabelsMatchingRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingPolicy.
func (in *RoutingPolicy) DeepCopy() *RoutingPolicy {
	if in == nil {
		return nil
	}
	out := new(RoutingPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackContactType) DeepCopyInto(out *SlackContactType) {
	*out = *in
	in.Webhook.DeepCopyInto(&out.Webhook)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackContactType.
func (in *SlackContactType) DeepCopy() *SlackContactType {
	if in == nil {
		return nil
	}
	out := new(SlackContactType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackdriverDatasource) DeepCopyInto(out *StackdriverDatasource) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(bool)
		**out = **in
	}
	if in.JWTAuthentication != nil {
		in, out := &in.JWTAuthentication, &out.JWTAuthentication
		*out = new(ValueOrRef)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackdriverDatasource.
func (in *StackdriverDatasource) DeepCopy() *StackdriverDatasource {
	if in == nil {
		return nil
	}
	out := new(StackdriverDatasource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TempoDatasource) DeepCopyInto(out *TempoDatasource) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(bool)
		**out = **in
	}
	if in.ForwardOauth != nil {
		in, out := &in.ForwardOauth, &out.ForwardOauth
		*out = new(bool)
		**out = **in
	}
	if in.ForwardCredentials != nil {
		in, out := &in.ForwardCredentials, &out.ForwardCredentials
		*out = new(bool)
		**out = **in
	}
	if in.SkipTLSVerify != nil {
		in, out := &in.SkipTLSVerify, &out.SkipTLSVerify
		*out = new(bool)
		**out = **in
	}
	if in.ForwardCookies != nil {
		in, out := &in.ForwardCookies, &out.ForwardCookies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		*out = new(BasicAuth)
		(*in).DeepCopyInto(*out)
	}
	if in.CACertificate != nil {
		in, out := &in.CACertificate, &out.CACertificate
		*out = new(ValueOrRef)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeGraph != nil {
		in, out := &in.NodeGraph, &out.NodeGraph
		*out = new(bool)
		**out = **in
	}
	if in.TraceToLogs != nil {
		in, out := &in.TraceToLogs, &out.TraceToLogs
		*out = new(TraceToLogs)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TempoDatasource.
func (in *TempoDatasource) DeepCopy() *TempoDatasource {
	if in == nil {
		return nil
	}
	out := new(TempoDatasource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TraceToLogs) DeepCopyInto(out *TraceToLogs) {
	*out = *in
	out.Datasource = in.Datasource
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FilterByTrace != nil {
		in, out := &in.FilterByTrace, &out.FilterByTrace
		*out = new(bool)
		**out = **in
	}
	if in.FilterBySpan != nil {
		in, out := &in.FilterBySpan, &out.FilterBySpan
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraceToLogs.
func (in *TraceToLogs) DeepCopy() *TraceToLogs {
	if in == nil {
		return nil
	}
	out := new(TraceToLogs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueOrDatasourceRef) DeepCopyInto(out *ValueOrDatasourceRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueOrDatasourceRef.
func (in *ValueOrDatasourceRef) DeepCopy() *ValueOrDatasourceRef {
	if in == nil {
		return nil
	}
	out := new(ValueOrDatasourceRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueOrRef) DeepCopyInto(out *ValueOrRef) {
	*out = *in
	if in.ValueRef != nil {
		in, out := &in.ValueRef, &out.ValueRef
		*out = new(ValueRef)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueOrRef.
func (in *ValueOrRef) DeepCopy() *ValueOrRef {
	if in == nil {
		return nil
	}
	out := new(ValueOrRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueRef) DeepCopyInto(out *ValueRef) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueRef.
func (in *ValueRef) DeepCopy() *ValueRef {
	if in == nil {
		return nil
	}
	out := new(ValueRef)
	in.DeepCopyInto(out)
	return out
}
