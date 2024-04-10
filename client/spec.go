package client

type Spec struct {
	Field string `json:"field,omitempty" yaml:"field,omitempty"`
}

func (s *Spec) Validate() error {
	// TODO: implement
	return nil
}

func (s *Spec) SetDefaults() {
}
