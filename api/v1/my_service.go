package v1

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// svc的端口映射关系
type ServicePort struct {
	Name       string             `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Protocol   string             `json:"protocol,omitempty" protobuf:"bytes,2,opt,name=protocol,casttype=Protocol"`
	Port       int32              `json:"port" protobuf:"varint,3,opt,name=port"`
	TargetPort intstr.IntOrString `json:"targetPort,omitempty" protobuf:"bytes,4,opt,name=targetPort"`
	NodePort   int32              `json:"nodePort,omitempty" protobuf:"varint,5,opt,name=nodePort"`
}

type OwnService struct {
	Ports     []v1.ServicePort `json:"ports,omitempty" patchStrategy:"merge" patchMergeKey:"port" protobuf:"bytes,1,rep,name=ports"`
	ClusterIP string           `json:"clusterIP,omitempty" protobuf:"bytes,3,opt,name=clusterIP"`
}

type ServicePortStatus struct {
	v1.ServicePort `json:"servicePort,omitempty"`
	// 检查此端口连通性
	Health bool `json:"health,omitempty"`
}

type UnitRelationServiceStatus struct {
	Type            v1.ServiceType      `json:"type,omitempty"`
	ClusterIP       string              `json:"clusterIP,omitempty"`
	Ports           []ServicePortStatus `json:"ports,omitempty"`
	SessionAffinity v1.ServiceAffinity  `json:"sessionAffinity,omitempty"`
}

type UnitRelationEndpointStatus struct {
	PodName  string `json:"podName"`
	PodIP    string `json:"podIP"`
	NodeName string `json:"nodeName"`
}
