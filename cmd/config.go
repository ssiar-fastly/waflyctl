// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

// TOMLConfig is the applications config file
type TOMLConfig struct {
	Logpath       string
	APIEndpoint   string
	Tags          []string
	Publisher     []string
	Action        string
	Rules         []int64
	DisabledRules []int64
	Owasp         owaspSettings
	Weblog        WeblogSettings
	Waflog        WaflogSettings
	Vclsnippet    VCLSnippetSettings
	Response      ResponseSettings
	Prefetch      PrefetchSettings
}

type owaspSettings struct {
	AllowedHTTPVersions           string
	AllowedMethods                string
	AllowedRequestContentType     string
	ArgLength                     int
	ArgNameLength                 int
	CombinedFileSizes             int
	CriticalAnomalyScore          int
	CRSValidateUTF8Encoding       bool
	ErrorAnomalyScore             int
	HTTPViolationScoreThreshold   int
	InboundAnomalyScoreThreshold  int
	LFIScoreThreshold             int
	MaxFileSize                   int
	MaxNumArgs                    int
	NoticeAnomalyScore            int
	ParanoiaLevel                 int
	PHPInjectionScoreThreshold    int
	RCEScoreThreshold             int
	RestrictedExtensions          string
	RestrictedHeaders             string
	RFIScoreThreshold             int
	SessionFixationScoreThreshold int
	SQLInjectionScoreThreshold    int
	XSSScoreThreshold             int
	TotalArgLength                int
	WarningAnomalyScore           int
}

// WeblogSettings parameters for logs in config file
type WeblogSettings struct {
	Name        string
	Address     string
	Port        uint
	Tlscacert   string
	Tlshostname string
	Format      string
}

// VCLSnippetSettings parameters for snippets in config file
type VCLSnippetSettings struct {
	Name     string
	Content  string
	Type     string
	Priority int
	Dynamic  int
}

// Version information from Fastly API
type Version struct {
	PublishKey string `json:"publish_key"`
	Name       string `json:"name"`
	Versions   []struct {
		Testing   bool        `json:"testing"`
		Locked    bool        `json:"locked"`
		Number    int         `json:"number"`
		Active    bool        `json:"active"`
		ServiceID string      `json:"service_id"`
		Staging   bool        `json:"staging"`
		CreatedAt time.Time   `json:"created_at"`
		DeletedAt interface{} `json:"deleted_at"`
		Comment   string      `json:"comment"`
		UpdatedAt time.Time   `json:"updated_at"`
		Deployed  bool        `json:"deployed"`
	} `json:"versions"`
	DeletedAt  interface{} `json:"deleted_at"`
	CreatedAt  time.Time   `json:"created_at"`
	Comment    string      `json:"comment"`
	CustomerID string      `json:"customer_id"`
	UpdatedAt  time.Time   `json:"updated_at"`
	ID         string      `json:"id"`
}

// WaflogSettings parameters from config
type WaflogSettings struct {
	Name        string
	Address     string
	Port        uint
	Tlscacert   string
	Tlshostname string
	Format      string
}

// ResponseSettings parameters from config
type ResponseSettings struct {
	Name           string
	HTTPStatusCode uint
	HTTPResponse   string
	ContentType    string
	Content        string
}

// PrefetchSettings parameters from config
type PrefetchSettings struct {
	Name      string
	Statement string
	Type      string
	Priority  int
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
