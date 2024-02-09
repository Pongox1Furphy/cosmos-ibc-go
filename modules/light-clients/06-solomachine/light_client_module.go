package solomachine

import (
	errorsmod "cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v8/modules/core/exported"
)

var _ exported.LightClientModule = (*LightClientModule)(nil)

// LightClientModule implements the core IBC api.LightClientModule interface?
type LightClientModule struct {
	cdc           codec.BinaryCodec
	storeProvider exported.ClientStoreProvider
}

// NewLightClientModule creates and returns a new 06-solomachine LightClientModule.
func NewLightClientModule(cdc codec.BinaryCodec) LightClientModule {
	return LightClientModule{
		cdc: cdc,
	}
}

func (lcm *LightClientModule) RegisterStoreProvider(storeProvider exported.ClientStoreProvider) {
	lcm.storeProvider = storeProvider
}

// Initialize is called upon client creation, it allows the client to perform validation on the initial consensus state and set the
// client state, consensus state and any client-specific metadata necessary for correct light client operation in the provided client store.
func (lcm LightClientModule) Initialize(ctx sdk.Context, clientID string, clientStateBz, consensusStateBz []byte) error {
	if err := validateClientID(clientID); err != nil {
		return err
	}

	var clientState ClientState
	if err := lcm.cdc.Unmarshal(clientStateBz, &clientState); err != nil {
		return err
	}

	if err := clientState.Validate(); err != nil {
		return err
	}

	var consensusState ConsensusState
	if err := lcm.cdc.Unmarshal(consensusStateBz, &consensusState); err != nil {
		return err
	}

	if err := consensusState.ValidateBasic(); err != nil {
		return err
	}

	clientStore := lcm.storeProvider.ClientStore(ctx, clientID)
	return clientState.Initialize(ctx, lcm.cdc, clientStore, &consensusState)
}

// VerifyClientMessage must verify a ClientMessage. A ClientMessage could be a Header, Misbehaviour, or batch update.
// It must handle each type of ClientMessage appropriately. Calls to CheckForMisbehaviour, UpdateState, and UpdateStateOnMisbehaviour
// will assume that the content of the ClientMessage has been verified and can be trusted. An error should be returned
// if the ClientMessage fails to verify.
func (lcm LightClientModule) VerifyClientMessage(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) error {
	if err := validateClientID(clientID); err != nil {
		return err
	}

	clientStore := lcm.storeProvider.ClientStore(ctx, clientID)
	clientState, found := getClientState(clientStore, lcm.cdc)
	if !found {
		return errorsmod.Wrap(clienttypes.ErrClientNotFound, clientID)
	}

	return clientState.VerifyClientMessage(ctx, lcm.cdc, clientStore, clientMsg)
}

// Checks for evidence of a misbehaviour in Header or Misbehaviour type. It assumes the ClientMessage
// has already been verified.
func (lcm LightClientModule) CheckForMisbehaviour(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) bool {
	if err := validateClientID(clientID); err != nil {
		panic(err)
	}

	clientStore := lcm.storeProvider.ClientStore(ctx, clientID)
	clientState, found := getClientState(clientStore, lcm.cdc)
	if !found {
		panic(errorsmod.Wrap(clienttypes.ErrClientNotFound, clientID))
	}

	return clientState.CheckForMisbehaviour(ctx, lcm.cdc, clientStore, clientMsg)
}

// UpdateStateOnMisbehaviour should perform appropriate state changes on a client state given that misbehaviour has been detected and verified
func (lcm LightClientModule) UpdateStateOnMisbehaviour(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) {
	if err := validateClientID(clientID); err != nil {
		panic(err)
	}

	clientStore := lcm.storeProvider.ClientStore(ctx, clientID)
	clientState, found := getClientState(clientStore, lcm.cdc)
	if !found {
		panic(errorsmod.Wrap(clienttypes.ErrClientNotFound, clientID))
	}

	clientState.UpdateStateOnMisbehaviour(ctx, lcm.cdc, clientStore, clientMsg)
}

// UpdateState updates and stores as necessary any associated information for an IBC client, such as the ClientState and corresponding ConsensusState.
// Upon successful update, a list of consensus heights is returned. It assumes the ClientMessage has already been verified.
func (lcm LightClientModule) UpdateState(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) []exported.Height {
	if err := validateClientID(clientID); err != nil {
		panic(err)
	}

	clientStore := lcm.storeProvider.ClientStore(ctx, clientID)
	clientState, found := getClientState(clientStore, lcm.cdc)
	if !found {
		panic(errorsmod.Wrap(clienttypes.ErrClientNotFound, clientID))
	}

	return clientState.UpdateState(ctx, lcm.cdc, clientStore, clientMsg)
}

// VerifyMembership is a generic proof verification method which verifies a proof of the existence of a value at a given CommitmentPath at the specified height.
// The caller is expected to construct the full CommitmentPath from a CommitmentPrefix and a standardized path (as defined in ICS 24).
func (lcm LightClientModule) VerifyMembership(
	ctx sdk.Context,
	clientID string,
	height exported.Height, // TODO: change to concrete type
	delayTimePeriod uint64,
	delayBlockPeriod uint64,
	proof []byte,
	path exported.Path, // TODO: change to conrete type
	value []byte,
) error {
	if err := validateClientID(clientID); err != nil {
		return err
	}

	clientStore := lcm.storeProvider.ClientStore(ctx, clientID)
	clientState, found := getClientState(clientStore, lcm.cdc)
	if !found {
		return errorsmod.Wrap(clienttypes.ErrClientNotFound, clientID)
	}

	return clientState.VerifyMembership(ctx, clientStore, lcm.cdc, height, delayTimePeriod, delayBlockPeriod, proof, path, value)
}

// VerifyNonMembership is a generic proof verification method which verifies the absence of a given CommitmentPath at a specified height.
// The caller is expected to construct the full CommitmentPath from a CommitmentPrefix and a standardized path (as defined in ICS 24).
func (lcm LightClientModule) VerifyNonMembership(
	ctx sdk.Context,
	clientID string,
	height exported.Height, // TODO: change to concrete type
	delayTimePeriod uint64,
	delayBlockPeriod uint64,
	proof []byte,
	path exported.Path, // TODO: change to conrete type
) error {
	if err := validateClientID(clientID); err != nil {
		return err
	}

	clientStore := lcm.storeProvider.ClientStore(ctx, clientID)
	clientState, found := getClientState(clientStore, lcm.cdc)
	if !found {
		return errorsmod.Wrap(clienttypes.ErrClientNotFound, clientID)
	}

	return clientState.VerifyNonMembership(ctx, clientStore, lcm.cdc, height, delayTimePeriod, delayBlockPeriod, proof, path)
}

// Status must return the status of the client. Only Active clients are allowed to process packets.
func (lcm LightClientModule) Status(ctx sdk.Context, clientID string) exported.Status {
	if err := validateClientID(clientID); err != nil {
		return exported.Unknown // TODO: or panic
	}

	clientStore := lcm.storeProvider.ClientStore(ctx, clientID)
	clientState, found := getClientState(clientStore, lcm.cdc)
	if !found {
		return exported.Unknown
	}

	return clientState.Status(ctx, clientStore, lcm.cdc)
}

// TimestampAtHeight must return the timestamp for the consensus state associated with the provided height.
func (lcm LightClientModule) TimestampAtHeight(ctx sdk.Context, clientID string, height exported.Height) (uint64, error) {
	if err := validateClientID(clientID); err != nil {
		return 0, err
	}

	clientStore := lcm.storeProvider.ClientStore(ctx, clientID)
	clientState, found := getClientState(clientStore, lcm.cdc)
	if !found {
		return 0, errorsmod.Wrap(clienttypes.ErrClientNotFound, clientID)
	}

	return clientState.GetTimestampAtHeight(ctx, clientStore, lcm.cdc, height)
}

func validateClientID(clientID string) error {
	clientType, _, err := clienttypes.ParseClientIdentifier(clientID)
	if err != nil {
		return err
	}

	if clientType != exported.Solomachine {
		return errorsmod.Wrapf(clienttypes.ErrInvalidClientType, "expected: %s, got: %s", exported.Solomachine, clientType)
	}

	return nil
}

// // CheckSubstituteAndUpdateState must verify that the provided substitute may be used to update the subject client.
// // The light client must set the updated client and consensus states within the clientStore for the subject client.
// // Deprecated: will be removed as performs internal functionality
func (LightClientModule) RecoverClient(ctx sdk.Context, clientID, substituteClientID string) error {
	return nil
}

// VerifyUpgradeAndUpdateState returns an error since solomachine client does not support upgrades
func (LightClientModule) VerifyUpgradeAndUpdateState(
	ctx sdk.Context,
	clientID string,
	newClient []byte,
	newConsState []byte,
	upgradeClientProof,
	upgradeConsensusStateProof []byte,
) error {
	return errorsmod.Wrap(clienttypes.ErrInvalidUpgradeClient, "cannot upgrade solomachine client")
}