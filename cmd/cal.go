/*
Copyright © 2020 Mike Girouard <mgirouard@gmail.com>
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice,
   this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors
   may be used to endorse or promote products derived from this software
   without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
POSSIBILITY OF SUCH DAMAGE.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// calCmd represents the cal command
var calCmd = &cobra.Command{
	Use:   "cal",
	Short: "display a calendar",
	//--------------------------------------------------------------------------------
	Long: `cal displays a simple calendar. If no arguments are specified, the current
month is displayed.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cal called")
	},
}

func init() {
	rootCmd.AddCommand(calCmd)
	calCmd.Flags().BoolP("one", "1", true, "Display single month output.")
	calCmd.Flags().BoolP("three", "3", false, "Display three months spanning the date.")
	calCmd.Flags().IntP("months", "n", 0, `Display n (number) of months, starting from the month
containing the date.`)
	calCmd.Flags().BoolP("sunday", "s", true, "Display Sunday as the first day of the week.")
	calCmd.Flags().BoolP("monday", "m", false, "Display Monday as the first day of the week.")
	calCmd.Flags().BoolP("vertical", "v", false, "Display using a vertical layout.")
	calCmd.Flags().BoolP("year", "y", false, "Display a calendar for the whole year.")
	calCmd.Flags().BoolP("twelve", "Y", false, "Display a calendar for the next twelve months.")
	calCmd.Flags().BoolP("week", "w", false, "Display a calendar for the next twelve months.")
}
