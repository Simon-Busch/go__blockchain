package network

type StatusMessage struct {
	ID 									string // Id of the server
	CurrentHeight 			uint32
	Version 						uint32
}

type GetStatusMessage struct {}

type GetBlocksMessage struct {
	From 								uint32 // from this height to that height
	To 									uint32 // If to == 0 the maximum blocks will be returned
}
