package esi_client

import (
	"flag"
	"fmt"
	"github.com/eve-tmarg/esiapi/helper"
	"github.com/theatrus/mediate"
	"time"
	"github.com/go-openapi/strfmt"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/eve-tmarg/esiapi/client"
)

func GenerateClient(esiSettings string, esiToken string) (generatedClient *client.EVESwaggerInterface, err error) {
	flag.Parse()

	settings, err := helper.LoadSettings(esiSettings)
	if err != nil {
		fmt.Printf("Can't load settings file %v\n", settings)
		return nil,err
	}

	transport, err := helper.InteractiveStartup(esiToken, settings)
	if err != nil {
		fmt.Printf("Can't do startup %v\n", err)
		return nil,err
	}

	fmt.Printf("Token refresh success - %v\n", esiToken)

	cliTransport := httptransport.New("esi.tech.ccp.is", "/latest", []string{"https"})
	cliTransport.Transport = mediate.RateLimit(10, 1*time.Second, transport)
	return client.New(cliTransport, strfmt.Default),nil
}
