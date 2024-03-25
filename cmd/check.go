package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkCommand)
}

var wg = &sync.WaitGroup{}

var checkCommand = &cobra.Command{
	Use:   "check",
	Short: "tracking the ip address",
	Long:  `tracking the ip address`,
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		for _, ip := range args {
			wg.Add(1)
			go ipTracker(ip)
		}
		wg.Wait()
		fmt.Println(time.Since(now))
	},
}

type Response struct {
	Status  string `json:"status"`
	Country string `json:"country"`
	City    string `json:"city"`
}

func ipTracker(ip string) {
	defer wg.Done()
	url := "http://ip-api.com/json/" + ip
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("this ip address not valid")
	} else {
		byteval, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			fmt.Println("failed to get body")
		} else {
			data := Response{}
			json.Unmarshal(byteval, &data)
			fmt.Printf("%+v\n", data)
		}
	}
}
