package bootstrapper

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "text/template"
    "bytes"
    "github.com/topicai/candy"
    "fmt"
    "io"
)

type Config struct {
    Ipaddr string
    Interface string
    Domain string
    Leases string
    DHCPRange string
    GateWay string
    DNSServer string
    NetMask string
    BroadCastAddr string
}


func Conf(tf string) {
    var config Config
    source, err := ioutil.ReadFile("./clusterDesc.yaml")
    if err != nil {
        panic(err)
    }
    err = yaml.Unmarshal(source, &config)
    if err != nil {
        panic(err)
    }
    ec := Config {
      Ipaddr: config.Ipaddr,
      Interface: config.Interface,
      Domain: config.Domain,
      Leases: config.Leases,
      DHCPRange: config.DHCPRange,
      GateWay: config.GateWay,
      DNSServer: config.DNSServer,
      NetMask: config.NetMask,
      BroadCastAddr: config.BroadCastAddr,
  }
    tmpl := template.New("")
    if len(tf) > 0 {
      tmpl = template.Must(tmpl.Parse(tf))
    } else {
      tmpl = template.Must(tmpl.Parse(tmplDnsmasqConf))
    }

    var buf bytes.Buffer
    tmpl.Execute(&buf, ec)
    candy.WithCreated("./dnsmasq.conf", func(w io.Writer) {
      _, e := fmt.Fprint(w, buf.String())
      candy.Must(e)
    })

}

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
