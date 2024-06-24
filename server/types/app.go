package types

import (
	"encoding/json"
	"io"

	ocabci "github.com/Finschia/ostracon/abci/types"
	"github.com/Finschia/ostracon/libs/log"
	octypes "github.com/Finschia/ostracon/types"
	"github.com/gogo/protobuf/grpc"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/Finschia/finschia-sdk/client"
	"github.com/Finschia/finschia-sdk/server/api"
	"github.com/Finschia/finschia-sdk/server/config"
	sdk "github.com/Finschia/finschia-sdk/types"
)

type (
	// AppOptions defines an interface that is passed into an application
	// constructor, typically used to set BaseApp options that are either supplied
	// via config file or through CLI arguments/flags. The underlying implementation
	// is defined by the server package and is typically implemented via a Viper
	// literal defined on the server Context. Note, casting Get calls may not yield
	// the expected types and could result in type assertion errors. It is recommend
	// to either use the cast package or perform manual conversion for safety.
	AppOptions interface {
		Get(string) interface{}
	}

	// Application defines an application interface that wraps abci.Application.
	// The interface defines the necessary contracts to be implemented in order
	// to fully bootstrap and start an application.
	Application interface {
		ocabci.Application

		RegisterAPIRoutes(*api.Server, config.APIConfig)

		// RegisterGRPCServer registers gRPC services directly with the gRPC
		// server.
		RegisterGRPCServer(grpc.Server)

		// RegisterTxService registers the gRPC Query service for tx (such as tx
		// simulation, fetching txs by hash...).
		RegisterTxService(client.Context)

		// RegisterTendermintService registers the gRPC Query service for ostracon queries.
		RegisterTendermintService(client.Context)

		// CommitMultiStore Returns the multistore instance
		CommitMultiStore() sdk.CommitMultiStore
	}

	// ApplicationQueryService defines an extension of the Application interface
	// that facilitates gRPC query Services.
	//
	// NOTE: This interfaces exists only in the v0.45.x line to ensure the existing
	// Application interface does not introduce API breaking changes.
	ApplicationQueryService interface {
		// RegisterNodeService registers the node gRPC Query service.
		RegisterNodeService(client.Context)
	}

	// AppCreator is a function that allows us to lazily initialize an
	// application using various configurations.
	AppCreator func(log.Logger, dbm.DB, io.Writer, AppOptions) Application

	// ModuleInitFlags takes a start command and adds modules specific init flags.
	ModuleInitFlags func(startCmd *cobra.Command)

	// ExportedApp represents an exported app state, along with
	// validators, consensus params and latest app height.
	ExportedApp struct {
		// AppState is the application state as JSON.
		AppState json.RawMessage
		// Validators is the exported validator set.
		Validators []octypes.GenesisValidator
		// Height is the app's latest block height.
		Height int64
		// ConsensusParams are the exported consensus params for ABCI.
		ConsensusParams *abci.ConsensusParams
	}

	// AppExporter is a function that dumps all app state to
	// JSON-serializable structure and returns the current validator set.
	AppExporter func(log.Logger, dbm.DB, io.Writer, int64, bool, []string, AppOptions) (ExportedApp, error)
)
