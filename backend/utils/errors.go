package utils

import (
	"regexp"
)

// errorCodeMap 存储了从智能合约错误代码到可读信息的映射
var errorCodeMap = map[string]string{
	"1":  "User record does not exist for caller",
	"2":  "Caller is not authorized to perform this action",
	"3":  "User does not exist",
	"4":  "Project ID is out of bounds",
	"5":  "Project history is empty, cannot get current record",
	"6":  "Caller must be the project owner",
	"7":  "Project must be in EDITING status to be updated",
	"8":  "Project must be in EDITING status to be submitted",
	"9":  "Project is already revoked",
	"10": "Project has no review history",
	"11": "Project is not in REVIEWING status",
	"12": "Review task has already been accepted by another verifier",
	"13": "Caller is not the designated verifier for this project",
	"14": "Credit batch ID is out of bounds",
	"15": "Credit batch history is empty, cannot get current record",
	"16": "Caller is not the current owner of this credit batch",
	"17": "Credit batch is finalized and cannot be transferred",
	"18": "Transfer quantity exceeds available balance",
	"19": "Credit batch is already finalized",
	"20": "Retirement quantity must be valid and within available balance",
	"21": "Source credit is not in retired state",
	"22": "Certificate ID is out of bounds",
	"23": "Certificate history is empty, cannot get current record",
	"24": "Certificate is already revoked",
	"25": "Buyer address cannot be zero",
	"26": "Quantity must be greater than zero",
}

// MapContractError 尝试从 go-ethereum 的错误中解析出合约的 revert 原因
func MapContractError(err error) string {
	if err == nil {
		return ""
	}

	// 尝试匹配 "execution reverted: 1" 这种模式
	re := regexp.MustCompile(`execution reverted: (\d+)`)
	matches := re.FindStringSubmatch(err.Error())

	if len(matches) > 1 {
		code := matches[1]
		if msg, ok := errorCodeMap[code]; ok {
			return "Contract Error: " + msg // 返回映射后的错误信息
		}
		return "Unknown Contract Error Code: " + code
	}

	// 如果不是 revert 错误，返回原始错误信息
	return err.Error()
}
