package network

type StatusMessage struct {
	ID 									string // Id of the server
	CurrentHeight 			uint32
	Version 						uint32
}

type GetStatusMessage struct {}
