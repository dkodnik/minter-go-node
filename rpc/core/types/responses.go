package core_types

import (
	"encoding/json"
	"strings"
	"time"

	abci "github.com/tendermint/abci/types"
	crypto "github.com/tendermint/go-crypto"
	cmn "github.com/tendermint/tmlibs/common"

	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/state"
	types "minter/tmtypes"
)

// List of blocks
type ResultBlockchainInfo struct {
	LastHeight int64              `json:"last_height"`
	BlockMetas []*types.BlockMeta `json:"block_metas"`
}

// Genesis file
type ResultGenesis struct {
	Genesis *types.GenesisDoc `json:"genesis"`
}

// Single block (with meta)
type ResultBlock struct {
	BlockMeta *types.BlockMeta `json:"block_meta"`
	Block     *types.Block     `json:"block"`
}

// Commit and Header
type ResultCommit struct {
	// SignedHeader is header and commit, embedded so we only have
	// one level in the json output
	types.SignedHeader
	CanonicalCommit bool `json:"canonical"`
}

// ABCI results from a block
type ResultBlockResults struct {
	Height  int64                `json:"height"`
	Results *state.ABCIResponses `json:"results"`
}

// NewResultCommit is a helper to initialize the ResultCommit with
// the embedded struct
func NewResultCommit(header *types.Header, commit *types.Commit,
	canonical bool) *ResultCommit {

	return &ResultCommit{
		SignedHeader: types.SignedHeader{
			Header: header,
			Commit: commit,
		},
		CanonicalCommit: canonical,
	}
}

// Info about the node's syncing state
type SyncInfo struct {
	LatestBlockHash   cmn.HexBytes `json:"latest_block_hash"`
	LatestAppHash     cmn.HexBytes `json:"latest_app_hash"`
	LatestBlockHeight int64        `json:"latest_block_height"`
	LatestBlockTime   time.Time    `json:"latest_block_time"`
	Syncing           bool         `json:"syncing"`
}

// Info about the node's validator
type ValidatorInfo struct {
	Address     cmn.HexBytes  `json:"address"`
	PubKey      crypto.PubKey `json:"pub_key"`
	VotingPower int64         `json:"voting_power"`
}

// Node Status
type ResultStatus struct {
	NodeInfo      p2p.NodeInfo  `json:"node_info"`
	SyncInfo      SyncInfo      `json:"sync_info"`
	ValidatorInfo ValidatorInfo `json:"validator_info"`
}

// Is TxIndexing enabled
func (s *ResultStatus) TxIndexEnabled() bool {
	if s == nil {
		return false
	}
	for _, s := range s.NodeInfo.Other {
		info := strings.Split(s, "=")
		if len(info) == 2 && info[0] == "tx_index" {
			return info[1] == "on"
		}
	}
	return false
}

// Info about peer connections
type ResultNetInfo struct {
	Listening bool     `json:"listening"`
	Listeners []string `json:"listeners"`
	NPeers    int      `json:"n_peers"`
	Peers     []Peer   `json:"peers"`
}

// Log from dialing seeds
type ResultDialSeeds struct {
	Log string `json:"log"`
}

// Log from dialing peers
type ResultDialPeers struct {
	Log string `json:"log"`
}

// A peer
type Peer struct {
	p2p.NodeInfo     `json:"node_info"`
	IsOutbound       bool                 `json:"is_outbound"`
	ConnectionStatus p2p.ConnectionStatus `json:"connection_status"`
}

// Validators for a height
type ResultValidators struct {
	BlockHeight int64              `json:"block_height"`
	Validators  []*types.Validator `json:"validators"`
}

// Info about the consensus state.
// UNSTABLE
type ResultDumpConsensusState struct {
	RoundState json.RawMessage `json:"round_state"`
	Peers      []PeerStateInfo `json:"peers"`
}

// UNSTABLE
type PeerStateInfo struct {
	NodeAddress string          `json:"node_address"`
	PeerState   json.RawMessage `json:"peer_state"`
}

// UNSTABLE
type ResultConsensusState struct {
	RoundState json.RawMessage `json:"round_state"`
}

// CheckTx result
type ResultBroadcastTx struct {
	Code uint32       `json:"code"`
	Data cmn.HexBytes `json:"data"`
	Log  string       `json:"log"`

	Hash cmn.HexBytes `json:"hash"`
}

// CheckTx and DeliverTx results
type ResultBroadcastTxCommit struct {
	CheckTx   abci.ResponseCheckTx   `json:"check_tx"`
	DeliverTx abci.ResponseDeliverTx `json:"deliver_tx"`
	Hash      cmn.HexBytes           `json:"hash"`
	Height    int64                  `json:"height"`
}

// Result of querying for a tx
type ResultTx struct {
	Hash     cmn.HexBytes           `json:"hash"`
	Height   int64                  `json:"height"`
	Index    uint32                 `json:"index"`
	TxResult abci.ResponseDeliverTx `json:"tx_result"`
	Tx       types.Tx               `json:"tx"`
	Proof    types.TxProof          `json:"proof,omitempty"`
}

// List of mempool txs
type ResultUnconfirmedTxs struct {
	N   int        `json:"n_txs"`
	Txs []types.Tx `json:"txs"`
}

// Info abci msg
type ResultABCIInfo struct {
	Response abci.ResponseInfo `json:"response"`
}

// Query abci msg
type ResultABCIQuery struct {
	Response abci.ResponseQuery `json:"response"`
}

// empty results
type (
	ResultUnsafeFlushMempool struct{}
	ResultUnsafeProfile      struct{}
	ResultSubscribe          struct{}
	ResultUnsubscribe        struct{}
	ResultHealth             struct{}
)

// Event data from a subscription
type ResultEvent struct {
	Query string            `json:"query"`
	Data  types.TMEventData `json:"data"`
}