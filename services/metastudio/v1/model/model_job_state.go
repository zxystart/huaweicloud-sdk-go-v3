package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobState 任务状态。 * CREATING: 创建中 * SYSTEM_AUDITING: 系统审核中 * AUDITING: 人工审核中 * WAITING: 等待训练 * PROCESSING: 任务训练中 * RESULT_REVIEW: 审核结果 * AUDIT_FAILED: 审核失败,等待用户重传数据 * FAILED: 失败 * SUCCEED: 成功
type JobState struct {
	value string
}

type JobStateEnum struct {
	CREATING        JobState
	SYSTEM_AUDITING JobState
	AUDITING        JobState
	WAITING         JobState
	PROCESSING      JobState
	RESULT_REVIEW   JobState
	AUDIT_FAILED    JobState
	FAILED          JobState
	SUCCEED         JobState
}

func GetJobStateEnum() JobStateEnum {
	return JobStateEnum{
		CREATING: JobState{
			value: "CREATING",
		},
		SYSTEM_AUDITING: JobState{
			value: "SYSTEM_AUDITING",
		},
		AUDITING: JobState{
			value: "AUDITING",
		},
		WAITING: JobState{
			value: "WAITING",
		},
		PROCESSING: JobState{
			value: "PROCESSING",
		},
		RESULT_REVIEW: JobState{
			value: "RESULT_REVIEW",
		},
		AUDIT_FAILED: JobState{
			value: "AUDIT_FAILED",
		},
		FAILED: JobState{
			value: "FAILED",
		},
		SUCCEED: JobState{
			value: "SUCCEED",
		},
	}
}

func (c JobState) Value() string {
	return c.value
}

func (c JobState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobState) UnmarshalJSON(b []byte) error {
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
