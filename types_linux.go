package netlink

import (
	"syscall"
)

const (
	NLMSG_ALIGNTO = syscall.NLMSG_ALIGNTO
	NLMSG_HDRLEN = syscall.NLMSG_HDRLEN
)

type NlMsghdr = syscall.NlMsghdr
type NetlinkMessage = syscall.NetlinkMessage
