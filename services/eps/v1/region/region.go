package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"sort"
	"strings"
)

var (
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://eps.myhuaweicloud.com")
	EU_WEST_101 = region.NewRegion("eu-west-101",
		"https://eps.eu-west-101.myhuaweicloud.eu")
)

var staticFields = map[string]*region.Region{
	"cn-north-4":  CN_NORTH_4,
	"eu-west-101": EU_WEST_101,
}

var provider = region.DefaultProviderChain("EPS")

func getRegionIds() []string {
	ids := make([]string, 0, len(staticFields))
	for key := range staticFields {
		ids = append(ids, key)
	}
	sort.Strings(ids)
	return ids
}

func ValueOf(regionId string) *region.Region {
	if regionId == "" {
		panic("unexpected empty parameter: regionId")
	}

	reg := provider.GetRegion(regionId)
	if reg != nil {
		return reg
	}

	if _, ok := staticFields[regionId]; ok {
		return staticFields[regionId]
	}
	panic(fmt.Sprintf("region id '%s' is not in the following supported regions of service 'EPS': [%s]", regionId, strings.Join(getRegionIds(), ", ")))
}
