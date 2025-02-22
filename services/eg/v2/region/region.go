package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"sort"
	"strings"
)

var (
	CN_NORTH_4 = region.NewRegion("cn-north-4",
		"https://eg.cn-north-4.myhuaweicloud.com")
	CN_EAST_2 = region.NewRegion("cn-east-2",
		"https://eg.cn-east-2.myhuaweicloud.com")
	CN_EAST_3 = region.NewRegion("cn-east-3",
		"https://eg.cn-east-3.myhuaweicloud.com")
)

var staticFields = map[string]*region.Region{
	"cn-north-4": CN_NORTH_4,
	"cn-east-2":  CN_EAST_2,
	"cn-east-3":  CN_EAST_3,
}

var provider = region.DefaultProviderChain("EG")

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
	panic(fmt.Sprintf("region id '%s' is not in the following supported regions of service 'EG': [%s]", regionId, strings.Join(getRegionIds(), ", ")))
}
