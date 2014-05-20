package bitcoind

// An Info represent a response to getmininginfo
type Info struct {
	// The server version
	Version uint32 `json:"version"`

	// The protocol version
	Protocolversion uint32 `json:"protocolversion"`

	// The wallet version
	Walletversion uint32 `json:"walletversion"`

	// The total bitcoin balance of the wallet
	Balance float64 `json:"balance"`

	// The current number of blocks processed in the server
	Blocks uint32 `json:"blocks"`

	// The time offset
	Timeoffset uint32 `json:"timeoffset"`

	// The number of connections
	Connections uint32 `json:"connections"`

	// Tthe proxy used by the server
	Proxy string `json:"proxy,omitempty"`

	// Tthe current difficulty
	Difficulty float64 `json:"difficulty"`

	// If the server is using testnet or not
	Testnet bool `json:"testnet"`

	// The timestamp (seconds since GMT epoch) of the oldest pre-generated key in the key pool
	Keypoololdest uint64 `json:"keypoololdest"`

	// How many new keys are pre-generated
	KeypoolSize uint32 `json:"keypoolsize,omitempty"`

	// The timestamp in seconds since epoch (midnight Jan 1 1970 GMT) that the wallet is unlocked for transfers, or 0 if the wallet is locked
	UnlockedUntil int64 `json:"unlocked_until,omitempty"`

	// the transaction fee set in btc/kb
	Paytxfee float64 `json:"paytxfee"`

	// Minimum relay fee for non-free transactions in btc/kb
	Relayfee float64 `json:"relayfee"`

	//  Any error messages
	Errors string `json:"errors"`
}
