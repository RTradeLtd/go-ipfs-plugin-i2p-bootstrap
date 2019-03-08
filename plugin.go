package main

import (
	bootstrapi2p "github.com/RTradeLtd/go-ipfs-plugin-i2p-bootstrap/i2p"
	i2p "github.com/RTradeLtd/go-ipfs-plugin-i2p-gateway/i2p"
	plugin "github.com/ipfs/go-ipfs/plugin"
)

// Plugins is an exported list of plugins that will be loaded by go-ipfs.
var Plugins = []plugin.Plugin{
	&bootstrapi2p.I2PBootstrapPlugin{},
	&i2p.I2PGatePlugin{},
}
