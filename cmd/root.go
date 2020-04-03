package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zaracooper/doc/cli"
	"github.com/zaracooper/doc/gdocs"
	"golang.org/x/oauth2/google"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	credsViper *viper.Viper
	tokenViper *viper.Viper
	tkFile     string
	rootCmd    = &cobra.Command{
		Use:   "doc",
		Short: "Generate your design docs/tech specs",
		Long:  `Generate your design docs/tech specs and upload them to Google Docs`,
		Run:   runCmd,
	}
)

// Execute TODO
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		exitWithError(err, "Failed to get config location: ")
	}

	tkFile = home + "/.docTK.json"

	tokenViper = viper.New()

	tokenViper.AddConfigPath(home)

	tokenViper.SetConfigName(".docTK")
	tokenViper.SetConfigType("json")

	tokenViper.AutomaticEnv()

	tokenViper.ReadInConfig()

	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		exitWithError(err, "Failed to read app config: ")
	}

	credsViper = viper.New()
	credsViper.SetDefault("creds", b)
}

func exitWithError(err error, msg string) {
	fmt.Println(msg, err.Error())
	os.Exit(1)
}

func runCmd(cmd *cobra.Command, args []string) {
	client := generateClient()

	docContents, err := cli.SelectContent()
	if err != nil {
		exitWithError(err, "Failed to get contents of document: ")
	}

	document, err := gdocs.CreateDocument(client, docContents.Title)
	if err != nil {
		exitWithError(err, "Failed to create google doc: ")
	}

	titlesInsertIndex := len(strings.Join(docContents.FrontMatter, "")) + (len(docContents.FrontMatter) * 18) + 4

	err = gdocs.AddText(
		client,
		document.DocumentId,
		append(
			gdocs.CreateFrontmatterRequests(docContents.FrontMatter, 5),
			gdocs.CreateTitleRequests(
				titlesInsertIndex,
				docContents.Introduction,
				docContents.Solutions,
				docContents.Considerations,
				docContents.SuccessEval,
				docContents.Work,
				docContents.Deliberation,
				docContents.EndMatter,
			)...),
	)
	if err != nil {
		exitWithError(err, "Failed to add text to google doc: ")
	}
}

func generateClient() *http.Client {
	creds, ok := credsViper.Get("creds").([]byte)
	if !ok {
		exitWithError(errors.New("Could not parse credentials"), "Failed to parse app credentials: ")
	}

	config, err := google.ConfigFromJSON(creds, "https://www.googleapis.com/auth/documents")
	if err != nil {
		exitWithError(err, "Unable to parse client secret file to config: ")
	}

	config.RedirectURL = "http://localhost:8000/code"

	token := tokenViper.Get("token")
	if token == nil {
		gdocs.GetTokenFromWeb(config, tokenViper, tkFile)
	}

	tk, err := gdocs.ConvertToToken(token)
	if err != nil {
		exitWithError(err, "Failed to get retrieve config: ")
	}

	return gdocs.GetClient(config, tk)
}
