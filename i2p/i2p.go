package i2pbootstrap

import (
	"log"
	"os"

	config "gx/ipfs/QmTbcMKv6GU3fxhnNcbzYChdox9Fdd7VpucM3PQ7UWjX3D/go-ipfs-config"

	"github.com/RTradeLtd/go-ipfs-plugin-i2p-gateway/config"
	coreiface "github.com/ipfs/go-ipfs/core/coreapi/interface"
	plugin "github.com/ipfs/go-ipfs/plugin"
	fsrepo "github.com/ipfs/go-ipfs/repo/fsrepo"
)

// I2PGatePlugin is a structure containing information which is used for
// setting up an i2p tunnel that connects an IPFS bootstrap to a tunnel over i2p.
type I2PGatePlugin struct {
	configPath    string
	config        *config.Config
	i2pconfigPath string
	i2pconfig     *i2pgateconfig.Config

	forwardHTTP  string
	forwardRPC   string
	forwardSwarm string
}

// I2PType will be used to identify this as the i2p bootstrap plugin to things
// that use it.
var I2PType = "i2pbootstrap"

var _ plugin.PluginDaemon = (*I2PGatePlugin)(nil)

// Name returns the plugin's name, satisfying the plugin.Plugin interface.
func (*I2PGatePlugin) Name() string {
	return "fwd-i2pbootstrap"
}

// Version returns the plugin's version, satisfying the plugin.Plugin interface.
func (*I2PGatePlugin) Version() string {
	return "0.0.0"
}

// Init initializes plugin, satisfying the plugin.Plugin interface. Put any
// initialization logic here.
func (i *I2PGatePlugin) Init() error {
	/*i := Setup()
	    if err != nil {
			return nil, err
		}*/
	return nil
}

// Setup creates an I2PGatePlugin and config file, but it doesn't start
// any tunnels.
func Setup() (*I2PGatePlugin, error) {
	return nil, nil
}

func (i *I2PGatePlugin) swarmString() string {
	swarmaddressbytes := ""
	for _, v := range i.config.Addresses.Swarm {
		swarmaddressbytes += v
	}
	return i2pgateconfig.Unquote(string(swarmaddressbytes))
}

func (i *I2PGatePlugin) idString() string {
	idbytes := i.config.Identity.PeerID
	return i2pgateconfig.Unquote(string(idbytes))
}

// I2PTypeName returns I2PType
func (*I2PGatePlugin) I2PTypeName() string {
	return I2PType
}

func (i *I2PGatePlugin) falseStart() error {
	i2p, err := Setup()
	if err != nil {
		return err
	}

	return nil
}

// Start starts the tunnels and also satisfies the Daemon plugin interface
func (i *I2PGatePlugin) Start(coreiface.CoreAPI) error {
	i2p, err := Setup()
	if err != nil {
		return err
	}

	return nil
}

// Close satisfies the Daemon plugin interface
func (*I2PGatePlugin) Close() error {
	return nil
}
