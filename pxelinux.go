package pxelinux

import (
	"fmt"
	"io"

	"github.com/k8sp/auto-install/bootstrapper/cmd"
	"github.com/k8sp/auto-install/config"
)

func Install(){
	const (
		centos = "centos"
		ubuntu = "ubuntu"
	)
	
	dist := config.LinuxDistro()
	if dist != centos && dist != ubuntu {
		log.Panicf("Unsupported OS: %s", dist)
	}

	switch dist {
	case centos:
		cmd.Run("yum", "-y", "install", "syslinux")
		io.Copy("/var/lib/tftpboot/", "/usr/share/syslinux/pxelinux.0")
	case ubuntu:
		cmd.Run("apt-get","update")
		cmd.Run("apt-get", "-y", "install", "pxelinux", "syslinux-common")
		io.Copy("/var/lib/tftpboot/", "/usr/lib/PXELINUX/pxelinux.0")
		io.Copy("/var/lib/tftpboot/", "/usr/lib/syslinux/modules/bios/ldlinux.c32")
	}

}

//add the Copy function to cmd Package
/*func Copy(dst string, src string){
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}*/
