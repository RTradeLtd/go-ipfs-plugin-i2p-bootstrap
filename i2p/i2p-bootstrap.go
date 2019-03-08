// +build !libp2p

package i2pbootstrap

import (
	"github.com/RTradeLtd/go-ipfs-plugin-i2p-gateway/config"
	"github.com/eyedeekay/sam-forwarder"
)

func (i *I2PBootstrapPlugin) ConnectBootstraps() error {
	for _, address := range i.connectBootstraps {
		err := i.ConnectBootstrap(address)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *I2PBootstrapPlugin) ConnectBootstrap(address string) error {
	i.Log("Creating an i2p destination for the Swarm Bootstrap Client")
	host := address
	i.Log("Swarm Bootstrap host", host)
	port := "4001"
	i.Log("Swarm Bootstrap port", port)
	pathRoot, err := i2pgateconfig.PathRoot()
	if err != nil {
		return err
	}
	i.Log("Keys Path", pathRoot+"ipfs-bootstrap.i2pkeys")
	GarlicForwarder, err := samforwarder.NewSAMClientForwarderFromOptions(
		//samforwarder.SetClientType("client"),
		samforwarder.SetClientSaveFile(true),
		samforwarder.SetClientFilePath(pathRoot),
		samforwarder.SetClientSAMHost(i.i2pconfig.HostSAM()),
		samforwarder.SetClientSAMPort(i.i2pconfig.PortSAM()),
		samforwarder.SetClientHost(host),
		samforwarder.SetClientPort(port),
		samforwarder.SetClientName("ipfs-gateway-swarm"),
		samforwarder.SetClientInLength(i.i2pconfig.InLength),
		samforwarder.SetClientOutLength(i.i2pconfig.OutLength),
		samforwarder.SetClientInVariance(i.i2pconfig.InVariance),
		samforwarder.SetClientOutVariance(i.i2pconfig.OutVariance),
		samforwarder.SetClientInQuantity(i.i2pconfig.InQuantity),
		samforwarder.SetClientOutQuantity(i.i2pconfig.OutQuantity),
		samforwarder.SetClientInBackups(i.i2pconfig.InBackupQuantity),
		samforwarder.SetClientOutBackups(i.i2pconfig.OutBackupQuantity),
		samforwarder.SetClientAllowZeroIn(i.i2pconfig.InAllowZeroHop),
		samforwarder.SetClientAllowZeroOut(i.i2pconfig.OutAllowZeroHop),
		samforwarder.SetClientCompress(i.i2pconfig.UseCompression),
		samforwarder.SetClientReduceIdle(i.i2pconfig.ReduceIdle),
		samforwarder.SetClientReduceIdleTimeMs(i.i2pconfig.ReduceIdleTime),
		samforwarder.SetClientReduceIdleQuantity(i.i2pconfig.ReduceIdleQuantity),
		samforwarder.SetClientAccessListType(i.i2pconfig.AccessListType),
		samforwarder.SetClientAccessList(i.i2pconfig.AccessList),
		samforwarder.SetClientEncrypt(i.i2pconfig.EncryptLeaseSet),
		samforwarder.SetClientLeaseSetKey(i.i2pconfig.EncryptedLeaseSetKey),
		samforwarder.SetClientLeaseSetPrivateKey(i.i2pconfig.EncryptedLeaseSetPrivateKey),
		samforwarder.SetClientLeaseSetPrivateSigningKey(i.i2pconfig.EncryptedLeaseSetPrivateSigningKey),
		samforwarder.SetClientMessageReliability(i.i2pconfig.MessageReliability),
	)
	if err != nil {
		return err

	}
	i.Log("SAM Generated Garlic Forwarder")
	go GarlicForwarder.Serve()
	for {
		if len(GarlicForwarder.Base32()) > 51 {
			i.Log("i2p base32(swarm): ", GarlicForwarder.Base32())
			break
		} else {
			i.Log("waiting for address")
		}
	}
	err = i2pgateconfig.ListenerBase32(GarlicForwarder.Base32(), i.i2pconfig)
	if err != nil {
		return err
	}
	err = i2pgateconfig.ListenerBase64(GarlicForwarder.Base64(), i.i2pconfig)
	if err != nil {
		return err
	}
	i.i2pconfig, err = i.i2pconfig.Save(i.configPath)
	if err != nil {
		return err
	}
	return nil
}
