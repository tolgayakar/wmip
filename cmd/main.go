package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "wmip",
		Short: "Fetch and copy your public IP address to the clipboard",
		Run: func(cmd *cobra.Command, args []string) {
			// Fetch public IP
			publicIP, err := fetchPublicIP()
			if err != nil {
				fmt.Println("Error fetching public IP:", err)
				os.Exit(1)
			}

			// Copy to clipboard
			err = clipboard.WriteAll(publicIP)
			if err != nil {
				fmt.Println("Error copying to clipboard:", err)
				os.Exit(1)
			}

			fmt.Printf("Your public IP (%s) has been copied to the clipboard.\n", publicIP)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func fetchPublicIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
