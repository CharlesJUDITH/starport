package starportcmd

import (
	"github.com/spf13/cobra"
)

// NewNetworkCampaign creates a new campaign command that holds other
// subcommands related to launching a network for a campaign.
func NewNetworkCampaign() *cobra.Command {
	c := &cobra.Command{
		Use:   "campaign",
		Short: "Handle campaigns",
	}

	c.AddCommand(
		NewNetworkCampaignList(),
		NewNetworkCampaignShow(),
	)

	return c
}
