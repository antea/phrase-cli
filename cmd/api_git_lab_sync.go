package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	helpers "github.com/phrase/phrase-cli/helpers"
	api "github.com/phrase/phrase-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initGitlabSyncDelete()
	initGitlabSyncExport()
	initGitlabSyncHistory()
	initGitlabSyncImport()
	initGitlabSyncList()
	initGitlabSyncShow()
	initGitlabSyncUpdate()

	rootCmd.AddCommand(gitLabSyncApiCmd)
}

var gitLabSyncApiCmd = &cobra.Command{
	// this weird approach is due to mustache template limitations
	Use:   strings.TrimSuffix("gitlabsyncapi", "api"),
	Short: strings.Join([]string{strings.TrimSuffix("GitLabSyncApi", "Api"), "API"}, " "),
}


func initGitlabSyncDelete() {
	params := viper.New()
	var gitlabSyncDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GitlabSyncDelete", strings.TrimSuffix("GitLabSyncApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GitLabSyncApi", "Api"), "s"))),
		Short: "Delete single Sync Setting",
		Long:  `Deletes a single GitLab Sync Setting.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GitlabSyncDeleteOpts{}

			
			id := params.GetString("id")

			

			data, api_response, err := client.GitLabSyncApi.GitlabSyncDelete(auth, id, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	gitLabSyncApiCmd.AddCommand(gitlabSyncDelete)

	
	AddFlag(gitlabSyncDelete, "string", "id", "", "ID", true)
	

	params.BindPFlags(gitlabSyncDelete.Flags())
}

func initGitlabSyncExport() {
	params := viper.New()
	var gitlabSyncExport = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GitlabSyncExport", strings.TrimSuffix("GitLabSyncApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GitLabSyncApi", "Api"), "s"))),
		Short: "Export from Phrase to GitLab",
		Long:  `Export translations from Phrase to GitLab according to the .phraseapp.yml file within the GitLab repository.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GitlabSyncExportOpts{}

			
			gitlabSyncId := params.GetString("gitlabSyncId")

			

			gitlabSyncExportParameters := api.GitlabSyncExportParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &gitlabSyncExportParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", gitlabSyncExportParameters)
			}
			

			data, api_response, err := client.GitLabSyncApi.GitlabSyncExport(auth, gitlabSyncId, gitlabSyncExportParameters, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	gitLabSyncApiCmd.AddCommand(gitlabSyncExport)

	
	AddFlag(gitlabSyncExport, "string", "gitlabSyncId", "", "Gitlab Sync ID", true)
	
	AddFlag(gitlabSyncExport, "string", "data", "d", "payload in JSON format", true)
	// gitlabSyncExportParameters := api.GitlabSyncExportParameters{}
	

	params.BindPFlags(gitlabSyncExport.Flags())
}

func initGitlabSyncHistory() {
	params := viper.New()
	var gitlabSyncHistory = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GitlabSyncHistory", strings.TrimSuffix("GitLabSyncApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GitLabSyncApi", "Api"), "s"))),
		Short: "History of single Sync Setting",
		Long:  `List history for a single Sync Setting.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GitlabSyncHistoryOpts{}

			
			gitlabSyncId := params.GetString("gitlabSyncId")

			

			data, api_response, err := client.GitLabSyncApi.GitlabSyncHistory(auth, gitlabSyncId, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	gitLabSyncApiCmd.AddCommand(gitlabSyncHistory)

	
	AddFlag(gitlabSyncHistory, "string", "gitlabSyncId", "", "Gitlab Sync ID", true)
	

	params.BindPFlags(gitlabSyncHistory.Flags())
}

func initGitlabSyncImport() {
	params := viper.New()
	var gitlabSyncImport = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GitlabSyncImport", strings.TrimSuffix("GitLabSyncApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GitLabSyncApi", "Api"), "s"))),
		Short: "Import from GitLab to Phrase",
		Long:  `Import translations from GitLab to Phrase according to the .phraseapp.yml file within the GitLab repository.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GitlabSyncImportOpts{}

			
			gitlabSyncId := params.GetString("gitlabSyncId")

			

			gitlabSyncImportParameters := api.GitlabSyncImportParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &gitlabSyncImportParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", gitlabSyncImportParameters)
			}
			

			data, api_response, err := client.GitLabSyncApi.GitlabSyncImport(auth, gitlabSyncId, gitlabSyncImportParameters, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	gitLabSyncApiCmd.AddCommand(gitlabSyncImport)

	
	AddFlag(gitlabSyncImport, "string", "gitlabSyncId", "", "Gitlab Sync ID", true)
	
	AddFlag(gitlabSyncImport, "string", "data", "d", "payload in JSON format", true)
	// gitlabSyncImportParameters := api.GitlabSyncImportParameters{}
	

	params.BindPFlags(gitlabSyncImport.Flags())
}

func initGitlabSyncList() {
	params := viper.New()
	var gitlabSyncList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GitlabSyncList", strings.TrimSuffix("GitLabSyncApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GitLabSyncApi", "Api"), "s"))),
		Short: "List GitLab syncs",
		Long:  `List all GitLab Sync Settings for which synchronisation with Phrase and GitLab is activated.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GitlabSyncListOpts{}

			

			data, api_response, err := client.GitLabSyncApi.GitlabSyncList(auth, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	gitLabSyncApiCmd.AddCommand(gitlabSyncList)

	

	params.BindPFlags(gitlabSyncList.Flags())
}

func initGitlabSyncShow() {
	params := viper.New()
	var gitlabSyncShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GitlabSyncShow", strings.TrimSuffix("GitLabSyncApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GitLabSyncApi", "Api"), "s"))),
		Short: "Get single Sync Setting",
		Long:  `Shows a single GitLab Sync Setting.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GitlabSyncShowOpts{}

			
			id := params.GetString("id")

			

			data, api_response, err := client.GitLabSyncApi.GitlabSyncShow(auth, id, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	gitLabSyncApiCmd.AddCommand(gitlabSyncShow)

	
	AddFlag(gitlabSyncShow, "string", "id", "", "ID", true)
	

	params.BindPFlags(gitlabSyncShow.Flags())
}

func initGitlabSyncUpdate() {
	params := viper.New()
	var gitlabSyncUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("GitlabSyncUpdate", strings.TrimSuffix("GitLabSyncApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("GitLabSyncApi", "Api"), "s"))),
		Short: "Update single Sync Setting",
		Long:  `Updates a single GitLab Sync Setting.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.GitlabSyncUpdateOpts{}

			
			id := params.GetString("id")

			

			data, api_response, err := client.GitLabSyncApi.GitlabSyncUpdate(auth, id, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	gitLabSyncApiCmd.AddCommand(gitlabSyncUpdate)

	
	AddFlag(gitlabSyncUpdate, "string", "id", "", "ID", true)
	

	params.BindPFlags(gitlabSyncUpdate.Flags())
}

