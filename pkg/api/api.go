package api

import (
	"github.com/RPA-Gang/gonutanix/pkg/client"
)

type INutanixApi interface {
	Client() client.INutanixClient
	SetClient(client client.INutanixClient)
	WithClient(client client.INutanixClient)
}
