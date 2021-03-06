package cfg

import "golang.org/x/oauth2"

type (
	Provider struct {
		oauth2.Config
		TokenRedirect string `yaml:"post_auth_redirect,omitempty"`
		Enabled       bool   `yaml:"enabled"`
		Name          string `yaml:"name"`
	}

	Security struct {
		PublicKeyPath  string   `yaml:"public_key"`
		PrivateKeyPath string   `yaml:"private_key"`
		BitSize        int      `yaml:"bit_size"`
		CookieSecret   string   `yaml:"cookie_secret"`
		AllowedDomains []string `yaml:"allowed_redirect_domains"`
	}

	Customization struct {
		AppName string `yaml:"app_name"`
		LogoURL string `yaml:"logo_url"`
		SiteURL string `yaml:"site_url"`
	}

	Config struct {
		Customization Customization `yaml:"customization"`
		Security      Security      `yaml:"security"`
		Providers     []Provider    `yaml:"providers"`
	}
)
