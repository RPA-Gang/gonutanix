package gonut

type INutanixApi interface {
	Client() INutanixClient
	SetClient(client INutanixClient)
	WithClient(client INutanixClient)
}
