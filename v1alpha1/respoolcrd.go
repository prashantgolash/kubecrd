package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"time"
)

type Respool struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata"`
	Spec               RespoolSpec   `json:"spec"`
	Status             RespoolStatus `json:"status,omitempty"`
}

func (in *Respool) DeepCopyInto(out *Respool) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = RespoolSpec {
		RespoolID: in.Spec.RespoolID,
		RespoolConfig: in.Spec.RespoolConfig,
		Owner: in.Spec.Owner,
		CreationTime: in.Spec.CreationTime,
		UpdateTime: in.Spec.UpdateTime,
	}
	out.Status = in.Status
}

func (in* Respool) DeepCopyObject() runtime.Object {
	out := Respool{}
	in.DeepCopyInto(&out)

	return &out
}

func (in* RespoolList) DeepCopyObject() runtime.Object {
	out := RespoolList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]Respool, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}
	return &out
}

type RespoolSpec struct {
	RespoolID     string `protobuf:"bytes,1,opt,name=respoolId" json:"changeLog,omitempty"`
	RespoolConfig string `protobuf:"bytes,2,opt,name=respoolConfig" json:"respoolConfig,omitempty"`
	Owner         string `protobuf:"bytes,3,opt,name=owner" json:"owner,omitempty"`
	CreationTime  time.Time `protobuf:"bytes,4,opt,name=creationTime" json:"creationTime,omitempty"`
	UpdateTime    time.Time `protobuf:"bytes,5,opt,name=updateTime" json:"updateTime,omitempty"`
}

type RespoolStatus struct {
	State   string `json:"state,omitempty"`
	Message string `json:"message,omitempty"`
}

type RespoolList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`
	Items            []Respool `json:"items"`
}

