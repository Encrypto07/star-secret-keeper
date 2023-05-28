package secret

import (
	"errors"

	"github.com/Encrypto07/star-secret-keeper/pkg/decryption"
	"github.com/Encrypto07/star-secret-keeper/pkg/encryption"
	"github.com/spf13/cobra"
)

var key string

func Execute() error {
	rootCmd := &cobra.Command{
		Use:   "secret",
		Short: "A CLI Secrets Manager Tool",
	}

	setCmd := &cobra.Command{
		Use:   "set",
		Short: "Set a secret",
		RunE:  setSecret,
	}
	rootCmd.AddCommand(setCmd)

	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Get a secret",
		RunE:  getSecret,
	}
	rootCmd.AddCommand(getCmd)

	rootCmd.PersistentFlags().StringVar(&key, "key", "", "AES encryption key")

	return rootCmd.Execute()
}

func setSecret(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("invalid number of arguments")
	}

	value := args[0]

	return encryption.SetSecret(value, key)
}

func getSecret(cmd *cobra.Command, args []string) error {
	return decryption.GetSecret(key)
}
