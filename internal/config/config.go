package config

type Config struct {
	Active      string                `toml:"active"`
	Controllers map[string]Controller `toml:"controllers"`
}

type Controller struct {
	Address   string `toml:"address"`
	AuthToken string `toml:"auth-token"`
}

func (c *Config) ActiveController() Controller {
	return c.Controllers[c.Active]
}
