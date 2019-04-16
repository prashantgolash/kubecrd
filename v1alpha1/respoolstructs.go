package v1alpha1


type ResourcePoolConfig struct {
	// Change log entry of the Resource Pool config
	// TODO use peloton.Changelog
	ChangeLog *ChangeLog `protobuf:"bytes,1,opt,name=changeLog" json:"changeLog,omitempty"`
	// Name of the resource pool
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Owning team of the pool
	OwningTeam string `protobuf:"bytes,3,opt,name=owningTeam" json:"owningTeam,omitempty"`
	// LDAP groups of the pool
	LdapGroups []string `protobuf:"bytes,4,rep,name=ldapGroups" json:"ldapGroups,omitempty"`
	// Description of the resource pool
	Description string `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	// Resource config of the Resource Pool
	Resources []*ResourceConfig `protobuf:"bytes,6,rep,name=resources" json:"resources,omitempty"`
	// Resource Pool's parent
	Parent *ResourcePoolID `protobuf:"bytes,7,opt,name=parent" json:"parent,omitempty"`
	// Task Scheduling policy
	Policy SchedulingPolicy `protobuf:"varint,8,opt,name=policy,enum=peloton.api.v0.respool.SchedulingPolicy" json:"policy,omitempty"`
	// The controller limit for this resource pool
	ControllerLimit *ControllerLimit `protobuf:"bytes,9,opt,name=controllerLimit" json:"controllerLimit,omitempty"`
	// Cap on max non-slack resources[mem,disk] in percentage
	// that can be used by revocable task.
	SlackLimit *SlackLimit `protobuf:"bytes,10,opt,name=slackLimit" json:"slackLimit,omitempty"`
}

// *
//  Resource configuration for a resource
type ResourceConfig struct {
	// Type of the resource
	Kind string `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	// Reservation/min of the resource
	Reservation float64 `protobuf:"fixed64,2,opt,name=reservation" json:"reservation,omitempty"`
	// Limit of the resource
	Limit float64 `protobuf:"fixed64,3,opt,name=limit" json:"limit,omitempty"`
	// Share on the resource pool
	Share float64 `protobuf:"fixed64,4,opt,name=share" json:"share,omitempty"`
	// ReservationType indicates the the type of reservation
	// There are two kind of reservation
	// 1. ELASTIC
	// 2. STATIC
	Type ReservationType `protobuf:"varint,5,opt,name=type,enum=peloton.api.v0.respool.ReservationType" json:"type,omitempty"`
}

type ReservationType int32

// For cpu, it will use revocable resources.
type SlackLimit struct {
	MaxPercent float64 `protobuf:"fixed64,1,opt,name=maxPercent" json:"maxPercent,omitempty"`
}

type ControllerLimit struct {
	MaxPercent float64 `protobuf:"fixed64,1,opt,name=maxPercent" json:"maxPercent,omitempty"`
}

// *
//  A unique ID assigned to a Resource Pool. This is a UUID in RFC4122 format.
type ResourcePoolID struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

type ChangeLog struct {
	// *
	//  Version number of the entity info which is monotonically increasing.
	//  Clients can use this to guide against race conditions using MVCC.
	Version uint64 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	// The timestamp when the entity info is created
	CreatedAt uint64 `protobuf:"varint,2,opt,name=createdAt" json:"createdAt,omitempty"`
	// The timestamp when the entity info is updated
	UpdatedAt uint64 `protobuf:"varint,3,opt,name=updatedAt" json:"updatedAt,omitempty"`
	// The user or service that updated the entity info
	UpdatedBy string `protobuf:"bytes,4,opt,name=updatedBy" json:"updatedBy,omitempty"`
}

type SchedulingPolicy int32
