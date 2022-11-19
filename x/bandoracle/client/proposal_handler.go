package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/redactedfury/petri/x/bandoracle/client/cli"
	"github.com/redactedfury/petri/x/bandoracle/client/rest"
)

var AddFetchPriceHandler = govclient.NewProposalHandler(cli.NewCmdSubmitFetchPriceProposal, rest.SubmitFetchPriceProposalRESTHandler)
