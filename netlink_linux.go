// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Netlink sockets and messages

package netlink

import (
	"syscall"
	"unsafe"
)

// Round the length of a netlink message up to align it properly.
func nlmAlignOf(msglen int) int {
	return (msglen + syscall.NLMSG_ALIGNTO - 1) & ^(syscall.NLMSG_ALIGNTO - 1)
}

// ParseNetlinkMessage parses b as an array of netlink messages and
// returns the slice containing the NetlinkMessage structures.
func ParseNetlinkMessage(b []byte) ([]syscall.NetlinkMessage, error) {
	var msgs []syscall.NetlinkMessage
	for len(b) >= syscall.NLMSG_HDRLEN {
		h, dbuf, dlen, err := netlinkMessageHeaderAndData(b)
		if err != nil {
			return nil, err
		}
		m := syscall.NetlinkMessage{
			Header: *h,
			Data:   dbuf[:int(h.Len)-syscall.NLMSG_HDRLEN],
		}
		msgs = append(msgs, m)
		b = b[dlen:]
	}
	return msgs, nil
}

func netlinkMessageHeaderAndData(b []byte) (*syscall.NlMsghdr, []byte, int, error) {
	h := (*syscall.NlMsghdr)(unsafe.Pointer(&b[0]))
	l := nlmAlignOf(int(h.Len))
	if int(h.Len) < syscall.NLMSG_HDRLEN || l > len(b) {
		return nil, nil, 0, syscall.EINVAL
	}
	return h, b[syscall.NLMSG_HDRLEN:], l, nil
}
