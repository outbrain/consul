package api

type SamenessGroupConfigEntry struct {
	Kind        string
	Name        string
	Partition   string `json:",omitempty"`
	IsDefault   bool   `json:",omitempty" alias:"is_default"`
	Members     []SamenessGroupMember
	Meta        map[string]string `json:",omitempty"`
	CreateIndex uint64
	ModifyIndex uint64
}

type SamenessGroupMember struct {
	Partition string
	Peer      string
}

func (s *SamenessGroupConfigEntry) GetKind() string            { return s.Kind }
func (s *SamenessGroupConfigEntry) GetName() string            { return s.Name }
func (s *SamenessGroupConfigEntry) GetPartition() string       { return s.Partition }
func (s *SamenessGroupConfigEntry) GetNamespace() string       { return "" }
func (s *SamenessGroupConfigEntry) GetCreateIndex() uint64     { return s.CreateIndex }
func (s *SamenessGroupConfigEntry) GetModifyIndex() uint64     { return s.ModifyIndex }
func (s *SamenessGroupConfigEntry) GetMeta() map[string]string { return s.Meta }
