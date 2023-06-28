package cli

import (
	"fmt"
	"strconv"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"

	"github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
)

// GetCmdParams returns the command handler for the host submodule parameter querying.
func GetCmdParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "params",
		Short:   "Query the current interchain-accounts host submodule parameters",
		Long:    "Query the current interchain-accounts host submodule parameters",
		Args:    cobra.NoArgs,
		Example: fmt.Sprintf("%s query interchain-accounts host params", version.AppName),
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(cmd.Context(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res.Params)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdPacketEvents returns the command handler for the host packet events querying.
func GetCmdPacketEvents() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "packet-events [channel-id] [sequence]",
		Short:   "Query the interchain-accounts host submodule packet events",
		Long:    "Query the interchain-accounts host submodule packet events for a particular channel and sequence",
		Args:    cobra.ExactArgs(2),
		Example: fmt.Sprintf("%s query interchain-accounts host packet-events channel-0 100", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			channelID, portID := args[0], icatypes.HostPortID
			if err := host.ChannelIdentifierValidator(channelID); err != nil {
				return err
			}

			seq, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			searchEvents := []string{
				fmt.Sprintf("%s.%s='%s'", channeltypes.EventTypeRecvPacket, channeltypes.AttributeKeyDstChannel, channelID),
				fmt.Sprintf("%s.%s='%s'", channeltypes.EventTypeRecvPacket, channeltypes.AttributeKeyDstPort, portID),
				fmt.Sprintf("%s.%s='%d'", channeltypes.EventTypeRecvPacket, channeltypes.AttributeKeySequence, seq),
			}

			results, err := processEvents(searchEvents, clientCtx)

			var resEvents []abci.Event
			for _, r := range results {
				for _, tx := range r.Txs {
					resEvents = append(resEvents, tx.Events...)
				}
			}

			return clientCtx.PrintString(sdk.StringifyEvents(resEvents).String())
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func processEvents(events []string, clientCtx client.Context) ([]sdk.SearchTxsResult, error) {
	var results []sdk.SearchTxsResult
	for _, event := range events {
		result, err := tx.QueryTxsByEvents(clientCtx, 1, 1, event, "")
		if err != nil {
			return nil, err
		}
		results = append(results, *result)
	}
	return results, nil
}
