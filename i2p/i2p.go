package i2pbootstrap

import (
	//"log"
	//"os"

	config "gx/ipfs/QmTbcMKv6GU3fxhnNcbzYChdox9Fdd7VpucM3PQ7UWjX3D/go-ipfs-config"

	"github.com/RTradeLtd/go-ipfs-plugin-i2p-gateway/config"
	plugin "github.com/ipfs/go-ipfs/plugin"
	//fsrepo "github.com/ipfs/go-ipfs/repo/fsrepo"
	coreiface "gx/ipfs/QmNmqKNivNTN11HrKWJYt29n6Z2fuzkeDheQV62dbxNuLb/interface-go-ipfs-core"
)


// I2PBootstrapPlugin is a structure containing information which is used for
// setting up an i2p tunnel that connects an IPFS bootstrap to a tunnel over i2p.
type I2PBootstrapPlugin struct {
	configPath    string
	config        *config.Config
	i2pconfigPath string
	i2pconfig     *i2pgateconfig.Config

	connectBootstraps []string
}

// I2PType will be used to identify this as the i2p bootstrap plugin to things
// that use it.
var I2PType = "i2pbootstrap"

var _ plugin.PluginDaemon = (*I2PBootstrapPlugin)(nil)

// Name returns the plugin's name, satisfying the plugin.Plugin interface.
func (*I2PBootstrapPlugin) Name() string {
	return "fwd-i2pbootstrap"
}

// Version returns the plugin's version, satisfying the plugin.Plugin interface.
func (*I2PBootstrapPlugin) Version() string {
	return "0.0.0"
}

// Init initializes plugin, satisfying the plugin.Plugin interface. Put any
// initialization logic here.
func (i *I2PBootstrapPlugin) Init() error {
	/*i := Setup()
	    if err != nil {
			return nil, err
		}*/
	return nil
}

// Setup creates an I2PBootstrapPlugin and config file, but it doesn't start
// any tunnels.
func Setup() (*I2PBootstrapPlugin, error) {
	return nil, nil
}

func (i *I2PBootstrapPlugin) swarmString() string {
	swarmaddressbytes := ""
	for _, v := range i.config.Addresses.Swarm {
		swarmaddressbytes += v
	}
	return i2pgateconfig.Unquote(string(swarmaddressbytes))
}

func (i *I2PBootstrapPlugin) idString() string {
	idbytes := i.config.Identity.PeerID
	return i2pgateconfig.Unquote(string(idbytes))
}

// I2PTypeName returns I2PType
func (*I2PBootstrapPlugin) I2PTypeName() string {
	return I2PType
}

func (i *I2PBootstrapPlugin) falseStart() error {
	i2p, err := Setup()
	if err != nil {
		return err
	}
    i2p.ConnectBootstraps()

	return nil
}

// Start starts the tunnels and also satisfies the Daemon plugin interface
func (i *I2PBootstrapPlugin) Start(coreiface.CoreAPI) error {
	i2p, err := Setup()
	if err != nil {
		return err
	}
    i2p.ConnectBootstraps()
	return nil
}

// Close satisfies the Daemon plugin interface
func (*I2PBootstrapPlugin) Close() error {
	return nil
}
