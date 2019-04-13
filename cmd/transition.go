// Copyright © 2019 Robert Sotomski <sotomski@gmail.com>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"github.com/sotomskir/goops/features/jira"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// pipelineJiraTransitionCmd represents the pipelineJiraTransition command
var transitionCmd = &cobra.Command{
	Use:     "transition STATE",
	Aliases: []string{"t"},
	Short:   "Transition all issues to desired state",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		j := jira.New()
		j.JiraTransition(viper.GetString("GOOPS_ISSUES"), args[0])
	},
}

func init() {
	rootCmd.AddCommand(transitionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pipelineJiraTransitionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pipelineJiraTransitionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
