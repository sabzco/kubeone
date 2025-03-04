/*
Copyright 2019 The KubeOne Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kubeflags

var (
	defaultAdmissionControllersv1227v1228 = []string{
		"NamespaceLifecycle",
		"LimitRanger",
		"ServiceAccount",
		"TaintNodesByCondition",
		"PodSecurity",
		"Priority",
		"DefaultTolerationSeconds",
		"DefaultStorageClass",
		"StorageObjectInUseProtection",
		"PersistentVolumeClaimResize",
		"RuntimeClass",
		"CertificateApproval",
		"CertificateSigning",
		"ClusterTrustBundleAttest",
		"CertificateSubjectRestriction",
		"DefaultIngressClass",
		"MutatingAdmissionWebhook",
		"ValidatingAdmissionPolicy",
		"ValidatingAdmissionWebhook",
		"ResourceQuota",
	}

	defaultAdmissionControllersv1226 = []string{
		"NamespaceLifecycle",
		"LimitRanger",
		"ServiceAccount",
		"TaintNodesByCondition",
		"PodSecurity",
		"Priority",
		"DefaultTolerationSeconds",
		"DefaultStorageClass",
		"StorageObjectInUseProtection",
		"PersistentVolumeClaimResize",
		"RuntimeClass",
		"CertificateApproval",
		"CertificateSigning",
		"CertificateSubjectRestriction",
		"DefaultIngressClass",
		"MutatingAdmissionWebhook",
		"ValidatingAdmissionPolicy",
		"ValidatingAdmissionWebhook",
		"ResourceQuota",
	}

	defaultAdmissionControllersv1225 = []string{
		"NamespaceLifecycle",
		"LimitRanger",
		"ServiceAccount",
		"TaintNodesByCondition",
		"PodSecurity",
		"Priority",
		"DefaultTolerationSeconds",
		"DefaultStorageClass",
		"StorageObjectInUseProtection",
		"PersistentVolumeClaimResize",
		"RuntimeClass",
		"CertificateApproval",
		"CertificateSigning",
		"CertificateSubjectRestriction",
		"DefaultIngressClass",
		"MutatingAdmissionWebhook",
		"ValidatingAdmissionWebhook",
		"ResourceQuota",
	}
)
