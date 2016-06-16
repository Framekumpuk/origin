package v1

import (
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/runtime"
)

const GroupName = "quota.openshift.io"

var SchemeGroupVersion = unversioned.GroupVersion{Group: GroupName, Version: "v1"}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) unversioned.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns back a Group qualified GroupResource
func Resource(resource string) unversioned.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func AddToScheme(scheme *runtime.Scheme) {
	addKnownTypes(scheme)
}

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&kapi.ListOptions{},
		&ClusterResourceQuota{},
		&ClusterResourceQuotaList{},
	)
}

func (obj *ClusterResourceQuotaList) GetObjectKind() unversioned.ObjectKind { return &obj.TypeMeta }
func (obj *ClusterResourceQuota) GetObjectKind() unversioned.ObjectKind     { return &obj.TypeMeta }
