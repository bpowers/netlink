//+build !linux

package netlink

// these are constant across all Linux archs
const (
	NLMSG_ALIGNTO = 0x4
	NLMSG_HDRLEN = 0x10
)

type NlMsghdr struct {
        Len   uint32
        Type  uint16
        Flags uint16
        Seq   uint32
        Pid   uint32
}

type NetlinkMessage struct {
	Header NlMsghdr
	Data   []byte
}