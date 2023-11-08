package keys

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/cosmos/go-bip39"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/require"

	"github.com/Finschia/ostracon/libs/cli"

	"github.com/Finschia/finschia-sdk/client"
	"github.com/Finschia/finschia-sdk/client/flags"
	"github.com/Finschia/finschia-sdk/crypto/hd"
	"github.com/Finschia/finschia-sdk/crypto/keyring"
	"github.com/Finschia/finschia-sdk/simapp"
	"github.com/Finschia/finschia-sdk/testutil"
	"github.com/Finschia/finschia-sdk/testutil/testdata"
	sdk "github.com/Finschia/finschia-sdk/types"
)

func Test_runAddCmdBasic(t *testing.T) {
	cmd := AddKeyCommand()
	cmd.Flags().AddFlagSet(Commands("home").PersistentFlags())

	mockIn := testutil.ApplyMockIODiscardOutErr(cmd)
	kbHome := t.TempDir()

	kb, err := keyring.New(sdk.KeyringServiceName(), keyring.BackendTest, kbHome, mockIn)
	require.NoError(t, err)

	clientCtx := client.Context{}.WithKeyringDir(kbHome).WithInput(mockIn).WithKeyring(kb)
	clientCtxPtr := &clientCtx
	ctx := context.WithValue(context.Background(), client.ClientContextKey, clientCtxPtr)

	t.Cleanup(func() {
		_ = kb.Delete("keyname1")
		_ = kb.Delete("keyname2")
	})

	cmd.SetArgs([]string{
		"keyname1",
		fmt.Sprintf("--%s=%s", flags.FlagHome, kbHome),
		fmt.Sprintf("--%s=%s", cli.OutputFlag, OutputFormatText),
		fmt.Sprintf("--%s=%s", flags.FlagKeyAlgorithm, string(hd.Secp256k1Type)),
		fmt.Sprintf("--%s=%s", flags.FlagKeyringBackend, keyring.BackendTest),
	})
	mockIn.Reset("y\n")
	require.NoError(t, cmd.ExecuteContext(ctx))

	mockIn.Reset("N\n")
	require.Error(t, cmd.ExecuteContext(ctx))

	cmd.SetArgs([]string{
		"keyname2",
		fmt.Sprintf("--%s=%s", flags.FlagHome, kbHome),
		fmt.Sprintf("--%s=%s", cli.OutputFlag, OutputFormatText),
		fmt.Sprintf("--%s=%s", flags.FlagKeyAlgorithm, string(hd.Secp256k1Type)),
		fmt.Sprintf("--%s=%s", flags.FlagKeyringBackend, keyring.BackendTest),
	})

	require.NoError(t, cmd.ExecuteContext(ctx))

	mockIn.Reset("y\n")
	require.NoError(t, cmd.ExecuteContext(ctx))

	cmd.SetArgs([]string{
		"keyname4",
		fmt.Sprintf("--%s=%s", flags.FlagHome, kbHome),
		fmt.Sprintf("--%s=%s", cli.OutputFlag, OutputFormatText),
		fmt.Sprintf("--%s=%s", flags.FlagKeyAlgorithm, string(hd.Secp256k1Type)),
		fmt.Sprintf("--%s=%s", flags.FlagKeyringBackend, keyring.BackendTest),
	})

	require.NoError(t, cmd.ExecuteContext(ctx))

	// In Multisig
	tcs := []struct {
		args []string
		err  string
	}{
		{
			[]string{
				"keyname1",
				fmt.Sprintf("--%s=%s", flags.FlagHome, kbHome),
				fmt.Sprintf("--%s=%s", cli.OutputFlag, OutputFormatText),
				fmt.Sprintf("--%s=%s", flags.FlagKeyAlgorithm, string(hd.Secp256k1Type)),
				fmt.Sprintf("--%s=%s", flags.FlagKeyringBackend, keyring.BackendTest),
				fmt.Sprintf("--%s=%s", flagMultisig, "keyname1,keyname2"),
			},
			"you cannot specify a new key as one of the names of the keys that make up a multisig",
		},
		{
			[]string{
				"keyname-multi",
				fmt.Sprintf("--%s=%s", flags.FlagHome, kbHome),
				fmt.Sprintf("--%s=%s", cli.OutputFlag, OutputFormatText),
				fmt.Sprintf("--%s=%s", flags.FlagKeyAlgorithm, string(hd.Secp256k1Type)),
				fmt.Sprintf("--%s=%s", flags.FlagKeyringBackend, keyring.BackendTest),
				fmt.Sprintf("--%s=%s", flagMultisig, "keyname1,keyname11"),
			},
			"part of the multisig target key does not exist",
		},
		{
			[]string{
				"keyname-multi",
				fmt.Sprintf("--%s=%s", flags.FlagHome, kbHome),
				fmt.Sprintf("--%s=%s", cli.OutputFlag, OutputFormatText),
				fmt.Sprintf("--%s=%s", flags.FlagKeyAlgorithm, string(hd.Secp256k1Type)),
				fmt.Sprintf("--%s=%s", flags.FlagKeyringBackend, keyring.BackendTest),
				fmt.Sprintf("--%s=%s", flagMultisig, "keyname1,keyname2"),
				fmt.Sprintf("--%s=%d", flagMultiSigThreshold, 3),
			},
			"threshold k of n multisignature",
		},
		{
			[]string{
				"keyname-multi",
				fmt.Sprintf("--%s=%s", flags.FlagHome, kbHome),
				fmt.Sprintf("--%s=%s", cli.OutputFlag, OutputFormatText),
				fmt.Sprintf("--%s=%s", flags.FlagKeyAlgorithm, string(hd.Secp256k1Type)),
				fmt.Sprintf("--%s=%s", flags.FlagKeyringBackend, keyring.BackendTest),
				fmt.Sprintf("--%s=%s", flagMultisig, "keyname1,keyname2"),
				fmt.Sprintf("--%s=%d", flagMultiSigThreshold, -1),
			},
			"threshold must be a positive integer",
		},
		{
			[]string{
				"keyname-multi",
				fmt.Sprintf("--%s=%s", flags.FlagHome, kbHome),
				fmt.Sprintf("--%s=%s", cli.OutputFlag, OutputFormatText),
				fmt.Sprintf("--%s=%s", flags.FlagKeyAlgorithm, string(hd.Secp256k1Type)),
				fmt.Sprintf("--%s=%s", flags.FlagKeyringBackend, keyring.BackendTest),
				fmt.Sprintf("--%s=%s", flagMultisig, "keyname1,keyname2"),
				fmt.Sprintf("--%s=%d", flagMultiSigThreshold, 2),
			},
			"",
		},
	}

	for _, tc := range tcs {
		cmd.SetArgs(tc.args)
		if tc.err != "" {
			require.Contains(t, cmd.ExecuteContext(ctx).Error(), tc.err)
		} else {
			require.NoError(t, cmd.ExecuteContext(ctx))
		}

		cmd.Flags().Visit(func(f *pflag.Flag) {
			if f.Name == flagMultisig {
				err = f.Value.(pflag.SliceValue).Replace([]string{})
				require.NoError(t, err)
			}
		})
	}

	cmd.SetArgs([]string{
		"keyname5",
		fmt.Sprintf("--%s=%s", flags.FlagHome, kbHome),
		fmt.Sprintf("--%s=true", flags.FlagDryRun),
		fmt.Sprintf("--%s=%s", cli.OutputFlag, OutputFormatText),
		fmt.Sprintf("--%s=%s", flags.FlagKeyAlgorithm, string(hd.Secp256k1Type)),
	})
	mockIn.Reset("\n")
	require.NoError(t, cmd.ExecuteContext(ctx))

	// In recovery mode
	cmd.SetArgs([]string{
		"keyname6",
		fmt.Sprintf("--%s=true", flagRecover),
	})

	// use valid mnemonic and complete recovery key generation successfully
	mockIn.Reset("decide praise business actor peasant farm drastic weather extend front hurt later song give verb rhythm worry fun pond reform school tumble august one\n")
	require.NoError(t, cmd.ExecuteContext(ctx))

	// use invalid mnemonic and fail recovery key generation
	mockIn.Reset("invalid mnemonic\n")
	require.Error(t, cmd.ExecuteContext(ctx))

	// In interactive mode
	cmd.SetArgs([]string{
		"keyname7",
		"-i",
		fmt.Sprintf("--%s=false", flagRecover),
	})

	const password = "password1!"

	// set password and complete interactive key generation successfully
	mockIn.Reset("\n" + password + "\n" + password + "\n")
	require.NoError(t, cmd.ExecuteContext(ctx))

	// passwords don't match and fail interactive key generation
	mockIn.Reset("\n" + password + "\n" + "fail" + "\n")
	require.Error(t, cmd.ExecuteContext(ctx))
}

func TestAddRecoverFileBackend(t *testing.T) {
	cmd := AddKeyCommand()
	cmd.Flags().AddFlagSet(Commands("home").PersistentFlags())

	mockIn := testutil.ApplyMockIODiscardOutErr(cmd)
	kbHome := t.TempDir()

	clientCtx := client.Context{}.WithKeyringDir(kbHome).WithInput(mockIn)
	ctx := context.WithValue(context.Background(), client.ClientContextKey, &clientCtx)

	cmd.SetArgs([]string{
		"keyname1",
		fmt.Sprintf("--%s=%s", flags.FlagHome, kbHome),
		fmt.Sprintf("--%s=%s", cli.OutputFlag, OutputFormatText),
		fmt.Sprintf("--%s=%s", flags.FlagKeyAlgorithm, string(hd.Secp256k1Type)),
		fmt.Sprintf("--%s=%s", flags.FlagKeyringBackend, keyring.BackendFile),
		fmt.Sprintf("--%s", flagRecover),
	})

	keyringPassword := "12345678"

	entropySeed, err := bip39.NewEntropy(mnemonicEntropySize)
	require.NoError(t, err)

	mnemonic, err := bip39.NewMnemonic(entropySeed)
	require.NoError(t, err)

	mockIn.Reset(fmt.Sprintf("%s\n%s\n%s\n", mnemonic, keyringPassword, keyringPassword))
	require.NoError(t, cmd.ExecuteContext(ctx))

	kb, err := keyring.New(sdk.KeyringServiceName(), keyring.BackendFile, kbHome, mockIn)
	require.NoError(t, err)

	t.Cleanup(func() {
		mockIn.Reset(fmt.Sprintf("%s\n%s\n", keyringPassword, keyringPassword))
		_ = kb.Delete("keyname1")
	})

	mockIn.Reset(fmt.Sprintf("%s\n%s\n", keyringPassword, keyringPassword))
	info, err := kb.Key("keyname1")
	require.NoError(t, err)
	require.Equal(t, "keyname1", info.GetName())
}

func Test_runAddCmdDryRun(t *testing.T) {
	pubkey1 := `{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AtObiFVE4s+9+RX5SP8TN9r2mxpoaT4eGj9CJfK7VRzN"}`
	pubkey2 := `{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A/se1vkqgdQ7VJQCM4mxN+L+ciGhnnJ4XYsQCRBMrdRi"}`

	testData := []struct {
		name  string
		args  []string
		added bool
	}{
		{
			name: "account is added",
			args: []string{
				"testkey",
				fmt.Sprintf("--%s=%s", flags.FlagDryRun, "false"),
			},
			added: true,
		},
		{
			name: "account is not added with dry run",
			args: []string{
				"testkey",
				fmt.Sprintf("--%s=%s", flags.FlagDryRun, "true"),
			},
			added: false,
		},
		{
			name: "multisig account is added",
			args: []string{
				"testkey",
				fmt.Sprintf("--%s=%s", flags.FlagDryRun, "false"),
				fmt.Sprintf("--%s=%s", flagMultisig, "subkey"),
			},
			added: true,
		},
		{
			name: "multisig account is not added with dry run",
			args: []string{
				"testkey",
				fmt.Sprintf("--%s=%s", flags.FlagDryRun, "true"),
				fmt.Sprintf("--%s=%s", flagMultisig, "subkey"),
			},
			added: false,
		},
		{
			name: "pubkey account is added",
			args: []string{
				"testkey",
				fmt.Sprintf("--%s=%s", flags.FlagDryRun, "false"),
				fmt.Sprintf("--%s=%s", FlagPublicKey, pubkey1),
			},
			added: true,
		},
		{
			name: "pubkey account is not added with dry run",
			args: []string{
				"testkey",
				fmt.Sprintf("--%s=%s", flags.FlagDryRun, "true"),
				fmt.Sprintf("--%s=%s", FlagPublicKey, pubkey2),
			},
			added: false,
		},
	}
	for _, tt := range testData {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			cmd := AddKeyCommand()
			cmd.Flags().AddFlagSet(Commands("home").PersistentFlags())

			kbHome := t.TempDir()
			mockIn := testutil.ApplyMockIODiscardOutErr(cmd)
			kb, err := keyring.New(sdk.KeyringServiceName(), keyring.BackendTest, kbHome, mockIn)
			require.NoError(t, err)

			appCodec := simapp.MakeTestEncodingConfig().Marshaler
			clientCtx := client.Context{}.
				WithCodec(appCodec).
				WithKeyringDir(kbHome).
				WithKeyring(kb)
			ctx := context.WithValue(context.Background(), client.ClientContextKey, &clientCtx)

			path := sdk.GetConfig().GetFullBIP44Path()
			_, err = kb.NewAccount("subkey", testdata.TestMnemonic, "", path, hd.Secp256k1)
			require.NoError(t, err)

			t.Cleanup(func() {
				_ = kb.Delete("subkey")
			})

			b := bytes.NewBufferString("")
			cmd.SetOut(b)

			cmd.SetArgs(tt.args)
			require.NoError(t, cmd.ExecuteContext(ctx))

			if tt.added {
				_, err = kb.Key("testkey")
				require.NoError(t, err)

				out, err := io.ReadAll(b)
				require.NoError(t, err)
				require.Contains(t, string(out), "name: testkey")
			} else {
				_, err = kb.Key("testkey")
				require.Error(t, err)
				require.Equal(t, "testkey.info: key not found", err.Error())
			}
		})
	}
}
