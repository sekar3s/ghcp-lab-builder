package lab

import (
	"github.com/s-samadi/ghas-lab-builder/internal/config"
	"github.com/spf13/cobra"
)

var (
	usersFile      string
	labDate        string
	enterpriseSlug string
	orgPrefix      string
)

var LabCmd = &cobra.Command{
	Use:   "lab",
	Short: "Manage complete lab environments (orgs, repos, users)",
	Long:  "The 'lab' command lets you create, destroy, and inspect complete GitHub Advanced Security lab environments.",
}

func init() {
	LabCmd.PersistentFlags().StringVar(&labDate, "lab-date", "", "Date string to identify date of the lab (e.g., '2024-06-15')")
	LabCmd.MarkPersistentFlagRequired("lab-date")
	LabCmd.PersistentFlags().StringVar(&usersFile, "users-file", "", "Path to user file (txt) (required)")
	LabCmd.MarkPersistentFlagRequired("users-file")
	LabCmd.PersistentFlags().StringVar(&facilitators, "facilitators", "", "lab facilitators usernames, comma-separated")
	LabCmd.MarkPersistentFlagRequired("facilitators")
	LabCmd.PersistentFlags().StringVar(&enterpriseSlug, "enterprise-slug", "", "GitHub Enterprise slug")
	LabCmd.MarkPersistentFlagRequired("enterprise-slug")
	LabCmd.PersistentFlags().StringVar(&orgPrefix, "org-prefix", "", "Prefix for organization names (default: \""+config.DefaultOrgPrefix+"\")")

	LabCmd.AddCommand(CreateCmd)
	LabCmd.AddCommand(DeleteCmd)
}
