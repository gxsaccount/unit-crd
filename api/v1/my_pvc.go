package v1

import (
	v1 "k8s.io/api/core/v1"
)

// pvc声明信息
type OwnPVC struct {
	Spec v1.PersistentVolumeClaimSpec ` json:"spec"`
}
