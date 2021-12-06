package models

type Provider1Config struct {
	Url string `yaml:"url"`
}

type Provider2Config struct {
	Url string `yaml:"url"`
}

type Provider3Config struct {
	Url string `yaml:"url"`
}

type ProviderConfig struct {
	Provider1 Provider1Config `yaml:"provider1"`
	Provider2 Provider2Config `yaml:"provider2"`
	Provider3 Provider3Config `yaml:"provider3"`
}
