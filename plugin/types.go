package types

import (
	pb "github.com/regen-network/keystone2/keystone"
)

// Options struct contains anything that is needed for configuring a
// plugin with one element initially, a path to a configuration, which
// can be interpreted differently by an individual plugin
type Options struct {
	ConfigPath string
}

// Plugin interface specifies the methods required for implementation
// by a plugin
type Plugin interface {
	NewKey(in *pb.KeySpec) (*pb.KeyRef, error)
	PubKey(in *pb.KeySpec) (*pb.PublicKey, error)
	Sign(in *pb.Msg) (*pb.Signed, error)
}
