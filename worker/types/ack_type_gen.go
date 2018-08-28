package types

// Code generated by github.com/CovenantSQL/HashStablePack DO NOT EDIT.

import (
	hsp "github.com/CovenantSQL/HashStablePack/marshalhash"
)

// MarshalHash marshals for hash
func (z *Ack) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 2
	o = append(o, 0x82, 0x82)
	if oTemp, err := z.Header.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x82)
	if oTemp, err := z.Envelope.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Ack) Msgsize() (s int) {
	s = 1 + 7 + z.Header.Msgsize() + 9 + z.Envelope.Msgsize()
	return
}

// MarshalHash marshals for hash
func (z *AckHeader) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 3
	o = append(o, 0x83, 0x83)
	if oTemp, err := z.Response.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x83)
	if oTemp, err := z.NodeID.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x83)
	o = hsp.AppendTime(o, z.Timestamp)
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *AckHeader) Msgsize() (s int) {
	s = 1 + 9 + z.Response.Msgsize() + 7 + z.NodeID.Msgsize() + 10 + hsp.TimeSize
	return
}

// MarshalHash marshals for hash
func (z AckResponse) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 0
	o = append(o, 0x80)
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z AckResponse) Msgsize() (s int) {
	s = 1
	return
}

// MarshalHash marshals for hash
func (z *SignedAckHeader) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 4
	o = append(o, 0x84, 0x84)
	if z.Signee == nil {
		o = hsp.AppendNil(o)
	} else {
		if oTemp, err := z.Signee.MarshalHash(); err != nil {
			return nil, err
		} else {
			o = hsp.AppendBytes(o, oTemp)
		}
	}
	o = append(o, 0x84)
	if z.Signature == nil {
		o = hsp.AppendNil(o)
	} else {
		if oTemp, err := z.Signature.MarshalHash(); err != nil {
			return nil, err
		} else {
			o = hsp.AppendBytes(o, oTemp)
		}
	}
	// map header, size 3
	o = append(o, 0x84, 0x83, 0x83)
	if oTemp, err := z.AckHeader.Response.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x83)
	if oTemp, err := z.AckHeader.NodeID.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x83)
	o = hsp.AppendTime(o, z.AckHeader.Timestamp)
	o = append(o, 0x84)
	if oTemp, err := z.HeaderHash.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SignedAckHeader) Msgsize() (s int) {
	s = 1 + 7
	if z.Signee == nil {
		s += hsp.NilSize
	} else {
		s += z.Signee.Msgsize()
	}
	s += 10
	if z.Signature == nil {
		s += hsp.NilSize
	} else {
		s += z.Signature.Msgsize()
	}
	s += 10 + 1 + 9 + z.AckHeader.Response.Msgsize() + 7 + z.AckHeader.NodeID.Msgsize() + 10 + hsp.TimeSize + 11 + z.HeaderHash.Msgsize()
	return
}