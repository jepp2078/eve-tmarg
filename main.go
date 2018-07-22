package main

import (
	"github.com/eve-tmarg/esiapi/cli/esi_client"
	"flag"
	"fmt"
	"github.com/eve-tmarg/sde/tools/sdetype"
	"github.com/eve-tmarg/esiapi/client/market"
	"github.com/eve-tmarg/sde/regions"
	"sync"
)

var sdePath string
var buildDB bool
var settingsFile string
var tokenFile string

func init() {
	flag.StringVar(&sdePath,"sde.path", "sdePath", "Path to the SDE root")
	flag.BoolVar(&buildDB, "sde.buildDB", false, "Convert to BoltDB")
	flag.StringVar(&settingsFile, "esi.settings", "settings.json", "Default settings file")
	flag.StringVar(&tokenFile, "esi.token", "", "Default token file")
}

func main() {
	flag.Parse()

	region := regions.GenerateRegions()

	sde, err := sdetype.GenerateSDE(sdePath, buildDB)
	if err != nil{
		fmt.Printf("Cannot generate SDE %v\n", err)
		return
	}
	client, err := esi_client.GenerateClient(settingsFile, tokenFile)
	if err != nil{
		fmt.Printf("Cannot generate ESI client %v\n", err)
		return
	}

	item, found := sde.GetTypeByExactName("Heavy Ion Blaster I")
	if found != true{
		fmt.Printf("Item not found %v\n", "Heavy Ion Blaster I")
		return
	}
	fmt.Print(item)

	params := market.NewGetMarketsRegionIDOrdersParams()
	params.SetRegionID(region.ByName("The Forge").Id)
	params.SetTypeID(&item.TypeId)
	params.SetOrderType("sell")
	result, err := client.Market.GetMarketsRegionIDOrders(params)

	wg := sync.WaitGroup{}
	wg.Add(len(result.Payload))
	for _, orders := range result.Payload {
		fmt.Printf("Item: %v Price: %v\n", *orders.TypeID, *orders.Price)
	}
	wg.Wait()


}
