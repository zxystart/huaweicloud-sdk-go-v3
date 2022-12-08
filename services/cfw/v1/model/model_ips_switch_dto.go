package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// description
type IpsSwitchDto struct {

	// 防护对象id
	ObjectId string `json:"object_id"`

	// 补丁类型，1-基础补丁 2=虚拟补丁
	IpsType IpsSwitchDtoIpsType `json:"ips_type"`

	// ips特性开关状态
	Status int32 `json:"status"`
}

func (o IpsSwitchDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpsSwitchDto struct{}"
	}

	return strings.Join([]string{"IpsSwitchDto", string(data)}, " ")
}

type IpsSwitchDtoIpsType struct {
	value int32
}

type IpsSwitchDtoIpsTypeEnum struct {
	E_1 IpsSwitchDtoIpsType
	E_2 IpsSwitchDtoIpsType
}

func GetIpsSwitchDtoIpsTypeEnum() IpsSwitchDtoIpsTypeEnum {
	return IpsSwitchDtoIpsTypeEnum{
		E_1: IpsSwitchDtoIpsType{
			value: 1,
		}, E_2: IpsSwitchDtoIpsType{
			value: 2,
		},
	}
}

func (c IpsSwitchDtoIpsType) Value() int32 {
	return c.value
}

func (c IpsSwitchDtoIpsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IpsSwitchDtoIpsType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
