package bootstrapper

import (
	"testing"

//	"github.com/topicai/candy"
//	"gopkg.in/yaml.v2"

//	"github.com/k8sp/auto-install/config"
	"github.com/stretchr/testify/assert"
  "github.com/k8sp/auto-install/bootstrapper"
)

const (
	tmplDnsmasqConf = `
# listen interface
interface={{ .Ipaddr }}
ip = {{ .Interface }}
bind-interfaces
domain={{ .Domain }}
dhcp-range={{ .DHCPRange }},{{ .NetMask }},{{ .Leases }}
#Gateway
dhcp-option=3,{{ .GateWay }}
#DNS
dhcp-option=6,{{ .DNSServer }}

no-hosts
expand-hosts
no-resolv
local=/ailabs.baifendian.com/
domain-needed
# PXE
dhcp-boot=pxelinux.0
# Broadcast Address
dhcp-option=28,{{ .BroadCastAddr }}

pxe-prompt="Press F8 for menu.", 60
pxe-service=x86PC, "Install CoreOS from network server 192.168.2.10", pxelinux
enable-tftp
tftp-root=/var/lib/tftpboot
`
)

func TestNginxConf(t *testing.T) {
//	c := &config.Cluster{}
//	candy.Must(yaml.Unmarshal([]byte(config.ExampleYAML), c))
	assert.Equal(t, tmplDnsmasqConf, GenDNSmasqConf())
}
