package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	"sort"
	"strings"
)

var (
	AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1",
		"https://sms.ap-southeast-1.myhuaweicloud.com")
	AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3",
		"https://sms.ap-southeast-3.myhuaweicloud.com")
	EU_WEST_101 = region.NewRegion("eu-west-101",
		"https://sms.eu-west-101.myhuaweicloud.eu")
)

var staticFields = map[string]*region.Region{
	"ap-southeast-1": AP_SOUTHEAST_1,
	"ap-southeast-3": AP_SOUTHEAST_3,
	"eu-west-101":    EU_WEST_101,
}

var provider = region.DefaultProviderChain("SMS")

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
	panic(fmt.Sprintf("region id '%s' is not in the following supported regions of service 'SMS': [%s]", regionId, strings.Join(getRegionIds(), ", ")))
}
