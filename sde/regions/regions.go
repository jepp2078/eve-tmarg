package regions

import (
	"strconv"
	"strings"
)

const regionsData = `
10000035,Deklein
10000036,Devoid
10000037,Everyshore
10000038,The Bleak Lands
10000039,Esoteria
10000040,Oasa
10000041,Syndicate
10000042,Metropolis
10000043,Domain
10000044,Solitude
10000045,Tenal
10000046,Fade
10000047,Providence
10000048,Placid
10000049,Khanid
10000050,Querious
10000051,Cloud Ring
10000052,Kador
10000053,Cobalt Edge
10000054,Aridia
10000055,Branch
10000001,Derelik
10000002,The Forge
10000003,Vale of the Silent
10000004,UUA-F4
10000005,Detorid
10000006,Wicked Creek
10000007,Cache
10000008,Scalding Pass
10000009,Insmother
10000010,Tribute
10000011,Great Wildlands
10000012,Curse
10000013,Malpais
10000014,Catch
10000015,Venal
10000016,Lonetrek
10000017,J7HZ-F
10000018,The Spire
10000019,A821-A
10000020,Tash-Murkon
10000021,Outer Passage
10000022,Stain
10000023,Pure Blind
10000025,Immensea
10000027,Etherium Reach
10000028,Molden Heath
10000029,Geminate
10000030,Heimatar
10000031,Impass
10000032,Sinq Laison
10000033,The Citadel
10000034,The Kalevala Expanse
10000056,Feythabolis
10000057,Outer Ring
10000058,Fountain
10000059,Paragon Soul
10000060,Delve
10000061,Tenerifis
10000062,Omist
10000063,Period Basis
10000064,Essence
10000065,Kor-Azor
10000066,Perrigen Falls
10000067,Genesis
10000068,Verge Vendor
10000069,Black Rise
11000001,A-R00001
11000017,D-R00017
11000018,D-R00018
11000019,D-R00019
11000020,D-R00020
11000021,D-R00021
11000002,A-R00002
11000003,A-R00003
11000004,B-R00004
11000005,B-R00005
11000006,B-R00006
11000007,B-R00007
11000008,B-R00008
11000009,C-R00009
11000010,C-R00010
11000011,C-R00011
11000012,C-R00012
11000013,C-R00013
11000014,C-R00014
11000015,C-R00015
11000016,D-R00016
11000022,D-R00022
11000023,D-R00023
11000024,E-R00024
11000025,E-R00025
11000026,E-R00026
11000027,E-R00027
11000028,E-R00028
11000029,E-R00029
11000030,F-R00030
11000031,G-R00031
11000032,H-R00032
11000033,K-R00033
`

// An EVE-Online Region
type Region struct {
	Name string
	Id   int32
}

// A set of Regions, with pre-filtered views
type Regions struct {
	AllRegions []*Region
	byId       map[int]*Region
	byName     map[string]*Region
}

var regions *Regions

func GenerateRegions() *Regions {
	regions = &Regions{AllRegions: make([]*Region, 0)}
	for _, line := range strings.Split(regionsData, "\n") {
		parts := strings.Split(line, ",")
		if len(parts) < 2 {
			continue
		}

		id, _ := strconv.ParseInt(parts[0], 10, 64)
		name := parts[1]

		regions.AllRegions = append(regions.AllRegions, &Region{Name: name, Id: int32(id)})
	}

	regions.byName = make(map[string]*Region)
	for _, region := range regions.AllRegions {
		regions.byName[region.Name] = region
	}

	regions.byId = make(map[int]*Region)
	for _, region := range regions.AllRegions {
		regions.byId[int(region.Id)] = region
	}

	return regions
}

func (r *Regions) ByName(name string) *Region {
	return r.byName[name]
}

func (r *Regions) ById(id int) *Region {
	return r.byId[id]
}
