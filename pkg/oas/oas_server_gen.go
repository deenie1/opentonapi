// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// DnsBackResolve implements dnsBackResolve operation.
	//
	// Get domains for wallet account.
	//
	// GET /v2/accounts/{account_id}/dns/backresolve
	DnsBackResolve(ctx context.Context, params DnsBackResolveParams) (DnsBackResolveRes, error)
	// DnsResolve implements dnsResolve operation.
	//
	// DNS resolve for domain name.
	//
	// GET /v2/dns/{domain_name}/resolve
	DnsResolve(ctx context.Context, params DnsResolveParams) (DnsResolveRes, error)
	// EmulateMessage implements emulateMessage operation.
	//
	// Emulate sending message to blockchain.
	//
	// POST /v2/blockchain/message/emulate
	EmulateMessage(ctx context.Context, req OptEmulateMessageReq) (EmulateMessageRes, error)
	// ExecGetMethod implements execGetMethod operation.
	//
	// Execute get method for account.
	//
	// POST /v2/blockchain/accounts/{account_id}/methods/{method_name}
	ExecGetMethod(ctx context.Context, req *ExecGetMethodReq, params ExecGetMethodParams) (ExecGetMethodRes, error)
	// GetAccount implements getAccount operation.
	//
	// Get human-friendly information about an account without low-level details.
	//
	// GET /v2/accounts/{account_id}
	GetAccount(ctx context.Context, params GetAccountParams) (GetAccountRes, error)
	// GetAccountTransactions implements getAccountTransactions operation.
	//
	// Get account transactions.
	//
	// GET /v2/blockchain/accounts/{account_id}/transactions
	GetAccountTransactions(ctx context.Context, params GetAccountTransactionsParams) (GetAccountTransactionsRes, error)
	// GetAllAuctions implements getAllAuctions operation.
	//
	// Get all auctions.
	//
	// GET /v2/dns/auctions
	GetAllAuctions(ctx context.Context, params GetAllAuctionsParams) (GetAllAuctionsRes, error)
	// GetBlock implements getBlock operation.
	//
	// Get block data.
	//
	// GET /v2/blockchain/blocks/{block_id}
	GetBlock(ctx context.Context, params GetBlockParams) (GetBlockRes, error)
	// GetBlockTransactions implements getBlockTransactions operation.
	//
	// Get transactions from block.
	//
	// GET /v2/blockchain/blocks/{block_id}/transactions
	GetBlockTransactions(ctx context.Context, params GetBlockTransactionsParams) (GetBlockTransactionsRes, error)
	// GetConfig implements getConfig operation.
	//
	// Get blockchain config.
	//
	// GET /v2/blockchain/config
	GetConfig(ctx context.Context) (GetConfigRes, error)
	// GetDomainBids implements getDomainBids operation.
	//
	// Get domain bids.
	//
	// GET /v2/dns/{domain_name}/bids
	GetDomainBids(ctx context.Context, params GetDomainBidsParams) (GetDomainBidsRes, error)
	// GetEvent implements getEvent operation.
	//
	// Get the event by event ID or hash of any transaction in trace.
	//
	// GET /v2/events/{event_id}
	GetEvent(ctx context.Context, params GetEventParams) (GetEventRes, error)
	// GetEventsByAccount implements getEventsByAccount operation.
	//
	// Get events for account.
	//
	// GET /v2/accounts/{account_id}/events
	GetEventsByAccount(ctx context.Context, params GetEventsByAccountParams) (GetEventsByAccountRes, error)
	// GetJettonInfo implements getJettonInfo operation.
	//
	// Get jetton metadata by jetton master address.
	//
	// GET /v2/jettons/{account_id}
	GetJettonInfo(ctx context.Context, params GetJettonInfoParams) (GetJettonInfoRes, error)
	// GetJettonsBalances implements getJettonsBalances operation.
	//
	// Get all Jettons balances by owner address.
	//
	// GET /v2/accounts/{account_id}/jettons
	GetJettonsBalances(ctx context.Context, params GetJettonsBalancesParams) (GetJettonsBalancesRes, error)
	// GetMasterchainHead implements getMasterchainHead operation.
	//
	// Get last known masterchain block.
	//
	// GET /v2/blockchain/masterchain-head
	GetMasterchainHead(ctx context.Context) (GetMasterchainHeadRes, error)
	// GetNftCollection implements getNftCollection operation.
	//
	// Get NFT collection by collection address.
	//
	// GET /v2/nfts/collections/{account_id}
	GetNftCollection(ctx context.Context, params GetNftCollectionParams) (GetNftCollectionRes, error)
	// GetNftCollections implements getNftCollections operation.
	//
	// Get NFT collections.
	//
	// GET /v2/nfts/collections
	GetNftCollections(ctx context.Context, params GetNftCollectionsParams) (GetNftCollectionsRes, error)
	// GetNftItemsByAddresses implements getNftItemsByAddresses operation.
	//
	// Get NFT items by its address.
	//
	// GET /v2/nfts/{account_ids}
	GetNftItemsByAddresses(ctx context.Context, params GetNftItemsByAddressesParams) (GetNftItemsByAddressesRes, error)
	// GetNftItemsByOwner implements getNftItemsByOwner operation.
	//
	// Get all NFT items by owner address.
	//
	// GET /v2/accounts/{account_id}/ntfs
	GetNftItemsByOwner(ctx context.Context, params GetNftItemsByOwnerParams) (GetNftItemsByOwnerRes, error)
	// GetRawAccount implements getRawAccount operation.
	//
	// Get low-level information about an account taken directly from the blockchain.
	//
	// GET /v2/blockchain/accounts/{account_id}
	GetRawAccount(ctx context.Context, params GetRawAccountParams) (GetRawAccountRes, error)
	// GetSubscriptionsByAccount implements getSubscriptionsByAccount operation.
	//
	// Get all subscriptions by wallet address.
	//
	// GET /v2/accounts/{account_id}/subscriptions
	GetSubscriptionsByAccount(ctx context.Context, params GetSubscriptionsByAccountParams) (GetSubscriptionsByAccountRes, error)
	// GetTrace implements getTrace operation.
	//
	// Get the trace by trace ID or hash of any transaction in trace.
	//
	// GET /v2/traces/{trace_id}
	GetTrace(ctx context.Context, params GetTraceParams) (GetTraceRes, error)
	// GetTracesByAccount implements getTracesByAccount operation.
	//
	// Get traces for account.
	//
	// GET /v2/accounts/{account_id}/traces
	GetTracesByAccount(ctx context.Context, params GetTracesByAccountParams) (GetTracesByAccountRes, error)
	// GetTransaction implements getTransaction operation.
	//
	// Get transaction data.
	//
	// GET /v2/blockchain/transactions/{transaction_id}
	GetTransaction(ctx context.Context, params GetTransactionParams) (GetTransactionRes, error)
	// GetValidators implements getValidators operation.
	//
	// Get validators.
	//
	// GET /v2/blockchain/validators
	GetValidators(ctx context.Context) (GetValidatorsRes, error)
	// PoolsByNominators implements poolsByNominators operation.
	//
	// All pools where account participates.
	//
	// GET /v2/stacking/nominator/{account_id}/pools
	PoolsByNominators(ctx context.Context, params PoolsByNominatorsParams) (PoolsByNominatorsRes, error)
	// SendMessage implements sendMessage operation.
	//
	// Send message to blockchain.
	//
	// POST /v2/blockchain/message
	SendMessage(ctx context.Context, req OptSendMessageReq) (SendMessageRes, error)
	// StackingPoolInfo implements stackingPoolInfo operation.
	//
	// Pool info.
	//
	// GET /v2/stacking/pool/{account_id}
	StackingPoolInfo(ctx context.Context, params StackingPoolInfoParams) (StackingPoolInfoRes, error)
	// StackingPools implements stackingPools operation.
	//
	// All pools available in network.
	//
	// GET /v2/stacking/pools
	StackingPools(ctx context.Context, params StackingPoolsParams) (StackingPoolsRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
