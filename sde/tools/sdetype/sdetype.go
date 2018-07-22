package sdetype

import (
	"flag"
	"fmt"

	"github.com/eve-tmarg/sde"
)

// Build a BoltDB of all of the relevant SDE items
func GenerateSDE(sdePath string, buildDB bool) (sdeClient sdetools.SDE,err error) {
	flag.Parse()

	sde := sdetools.SDE{
		BaseDir: sdePath,
	}
	sde.Init()
	if buildDB {
		sde.BuildBoltDB()
	}

	_, ok := sde.GetSystemNameById(30000142)
	if ok != true {
		return sde, fmt.Errorf("can't load types %v\n", ok)
	}
	fmt.Println("Names OK!")
	_, ok = sde.GetGroupById(8)
	if ok != true {
		return sde, fmt.Errorf("can't load groups\n")
	}
	fmt.Println("Groups OK!")

	_, ok = sde.GetTypeById(34)
	if ok != true {
		return sde, fmt.Errorf("can't load typeIds\n")
	}
	fmt.Println("MarketTypeByID OK!")
	_, ok = sde.GetTypeByExactName("Zealot")
	if ok != true {
		return sde, fmt.Errorf("can't load typeNameIds\n")
	}
	fmt.Println("MarketTypeByName OK!")

	return sde, nil
}
