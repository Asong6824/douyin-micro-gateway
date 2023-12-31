package setting

type EngineSetting struct {
	WithHostPorts  string
	WithStreamBody bool
	Registry       RegistryConfig
}

type RegistryConfig struct {
    ServiceName string
    Addr        string
    Weight      int
    Tags        map[string]string
}

type MinioSetting struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}