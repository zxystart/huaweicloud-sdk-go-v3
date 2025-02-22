package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// CreateFunctionTriggerResponse Response Object
type CreateFunctionTriggerResponse struct {

	// 触发器ID。
	TriggerId *string `json:"trigger_id,omitempty"`

	// 触发器类型。  - TIMER: \"定时触发器。\" - APIG: \"APIG触发器。\" - CTS: \"云审计服务触发器。\" - DDS: \"文档数据库服务触发器。\" - DMS: \"分布式服务触发器。\" - DIS: \"数据接入服务触发器。\" - LTS: \"云日志服务触发器。\" - OBS: \"对象存储触发器。\" - SMN: \"消息通知服务触发器。\" - KAFKA: \"专享版消息通知服务触发器。\"
	TriggerTypeCode *CreateFunctionTriggerResponseTriggerTypeCode `json:"trigger_type_code,omitempty"`

	// \"触发器状态\"  - ACTIVE: 启用状态。 - DISABLED: 禁用状态。
	TriggerStatus *CreateFunctionTriggerResponseTriggerStatus `json:"trigger_status,omitempty"`

	EventData *TriggerEventDataRequestBody `json:"event_data,omitempty"`

	// 最后更新时间。
	LastUpdatedTime *sdktime.SdkTime `json:"last_updated_time,omitempty"`

	// 触发器创建时间。
	CreatedTime    *sdktime.SdkTime `json:"created_time,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateFunctionTriggerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFunctionTriggerResponse struct{}"
	}

	return strings.Join([]string{"CreateFunctionTriggerResponse", string(data)}, " ")
}

type CreateFunctionTriggerResponseTriggerTypeCode struct {
	value string
}

type CreateFunctionTriggerResponseTriggerTypeCodeEnum struct {
	TIMER            CreateFunctionTriggerResponseTriggerTypeCode
	APIG             CreateFunctionTriggerResponseTriggerTypeCode
	CTS              CreateFunctionTriggerResponseTriggerTypeCode
	DDS              CreateFunctionTriggerResponseTriggerTypeCode
	DMS              CreateFunctionTriggerResponseTriggerTypeCode
	DIS              CreateFunctionTriggerResponseTriggerTypeCode
	LTS              CreateFunctionTriggerResponseTriggerTypeCode
	OBS              CreateFunctionTriggerResponseTriggerTypeCode
	SMN              CreateFunctionTriggerResponseTriggerTypeCode
	KAFKA            CreateFunctionTriggerResponseTriggerTypeCode
	RABBITMQ         CreateFunctionTriggerResponseTriggerTypeCode
	DEDICATEDGATEWAY CreateFunctionTriggerResponseTriggerTypeCode
	OPENSOURCEKAFKA  CreateFunctionTriggerResponseTriggerTypeCode
	APIC             CreateFunctionTriggerResponseTriggerTypeCode
	GAUSSMONGO       CreateFunctionTriggerResponseTriggerTypeCode
	EVENTGRID        CreateFunctionTriggerResponseTriggerTypeCode
}

func GetCreateFunctionTriggerResponseTriggerTypeCodeEnum() CreateFunctionTriggerResponseTriggerTypeCodeEnum {
	return CreateFunctionTriggerResponseTriggerTypeCodeEnum{
		TIMER: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "TIMER",
		},
		APIG: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "APIG",
		},
		CTS: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "CTS",
		},
		DDS: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "DDS",
		},
		DMS: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "DMS",
		},
		DIS: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "DIS",
		},
		LTS: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "LTS",
		},
		OBS: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "OBS",
		},
		SMN: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "SMN",
		},
		KAFKA: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "KAFKA",
		},
		RABBITMQ: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "RABBITMQ",
		},
		DEDICATEDGATEWAY: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "DEDICATEDGATEWAY",
		},
		OPENSOURCEKAFKA: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "OPENSOURCEKAFKA",
		},
		APIC: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "APIC",
		},
		GAUSSMONGO: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "GAUSSMONGO",
		},
		EVENTGRID: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "EVENTGRID",
		},
	}
}

func (c CreateFunctionTriggerResponseTriggerTypeCode) Value() string {
	return c.value
}

func (c CreateFunctionTriggerResponseTriggerTypeCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFunctionTriggerResponseTriggerTypeCode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateFunctionTriggerResponseTriggerStatus struct {
	value string
}

type CreateFunctionTriggerResponseTriggerStatusEnum struct {
	ACTIVE   CreateFunctionTriggerResponseTriggerStatus
	DISABLED CreateFunctionTriggerResponseTriggerStatus
}

func GetCreateFunctionTriggerResponseTriggerStatusEnum() CreateFunctionTriggerResponseTriggerStatusEnum {
	return CreateFunctionTriggerResponseTriggerStatusEnum{
		ACTIVE: CreateFunctionTriggerResponseTriggerStatus{
			value: "ACTIVE",
		},
		DISABLED: CreateFunctionTriggerResponseTriggerStatus{
			value: "DISABLED",
		},
	}
}

func (c CreateFunctionTriggerResponseTriggerStatus) Value() string {
	return c.value
}

func (c CreateFunctionTriggerResponseTriggerStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFunctionTriggerResponseTriggerStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
