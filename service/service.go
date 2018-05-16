package service

const NodeAddress = "http://128.199.57.221"

type Result struct {
	Success          *bool   `protobuf:"varint,1,req,name=success" json:"success,omitempty"`
	Errcode          *int32  `protobuf:"varint,2,opt,name=errcode" json:"errcode,omitempty"`
	Reason           *string `protobuf:"bytes,3,opt,name=reason" json:"reason,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Result) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *Result) GetReason() string {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return ""
}
