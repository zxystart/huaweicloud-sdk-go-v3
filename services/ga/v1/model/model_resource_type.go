package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 资源类型，取值： - ga-accelerators：加速实例 - ga-listeners：监听器
type ResourceType struct {
	value string
}

type ResourceTypeEnum struct {
	GA_ACCELERATORS ResourceType
	GA_LISTENERS    ResourceType
}

func GetResourceTypeEnum() ResourceTypeEnum {
	return ResourceTypeEnum{
		GA_ACCELERATORS: ResourceType{
			value: "ga-accelerators",
		},
		GA_LISTENERS: ResourceType{
			value: "ga-listeners",
		},
	}
}

func (c ResourceType) Value() string {
	return c.value
}

func (c ResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
