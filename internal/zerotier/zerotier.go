package zerotier

//go:generate go run ../../tools/zerotier-oapi-refresher/ -output-spec-file=../../api/zerotier-openapi.yaml

const (
	AuthFile            = "/var/lib/zerotier-one/authtoken.secret"
	AuthTokenHeader     = "X-ZT1-Auth"
	PublicID            = "/var/lib/zerotier-one/identity.public"
	NetworkPrivate      = 1
	NetworkBridging     = 0
	NetworkV4AssignMode = "zt"
	NetworkV6AssignMode = "zt"
)
