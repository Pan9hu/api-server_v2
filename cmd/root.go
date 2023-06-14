package cmd

import (
	"github.com/Pan9Hu/api-server_v2/conf"
	"github.com/Pan9Hu/api-server_v2/core"
	"github.com/lithammer/dedent"
	"github.com/spf13/cobra"
	"log"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "api-server",
	Short: "API Server is a stateless service",
	Long: dedent.Dedent(`
			┌───────────────────────────────────────────────────────────────────────────────┐
			│ API Server                                                                    │
			│ API Server is a modular design API framework based on Gin,which services REST │
			│ operations and provides the HTTP methods CRUD to the Melo CMDB.               │
			│                                                                               │
			│                                                                               │
			│ Please give us feedback at:                                                   │
			│ https://github.com/Pan9hu/api-server_v2/issues                                │
			└───────────────────────────────────────────────────────────────────────────────┘
			(\__/) ||               
			(•ㅅ•) ||               
			/ 　 づv

		Example usage:

			You can edit the configuration file(api.properties) and use it to start the service.

			┌──────────────────────────────────────────────────────────┐
			│ Start the API Server from a configuration file           │
			├──────────────────────────────────────────────────────────┤
			│ api-server -c/--config 'file'                            │
			└──────────────────────────────────────────────────────────┘

			You can use api-server -h or api-server --help get more support.`),
	//Run: func(cmd *cobra.Command, args []string) {
	//	// Do something
	//},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("[Error] start melo api server error, because: ", err.Error())
		return
	}
	defer func() {
		conf.MeloVP = core.Configure(cfgFile)
	}()
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Start the API Server from a configuration file")
}
