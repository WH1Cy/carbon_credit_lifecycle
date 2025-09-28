# 碳信通错误代码映射表

## 错误代码说明
错误代码格式：纯数字 `1` - `26`
- 使用最简洁的数字字符串表示错误
- 极致优化合约大小和gas消耗

## 错误代码映射表

| 错误代码 | 错误信息 | 模块 |
|---------|---------|------|
| 1 | User record does not exist for caller | 用户权限 |
| 2 | Caller is not authorized to perform this action | 用户权限 |
| 3 | User does not exist | 用户权限 |
| 4 | Project ID is out of bounds | 项目管理 |
| 5 | Project history is empty, cannot get current record | 项目管理 |
| 6 | Caller must be the project owner | 项目管理 |
| 7 | Project must be in EDITING status to be updated | 项目管理 |
| 8 | Project must be in EDITING status to be submitted | 项目管理 |
| 9 | Project is already revoked | 项目管理 |
| 10 | Project has no review history | 项目管理 |
| 11 | Project is not in REVIEWING status | 审核流程 |
| 12 | Review task has already been accepted by another verifier | 审核流程 |
| 13 | Caller is not the designated verifier for this project | 审核流程 |
| 14 | Credit batch ID is out of bounds | 碳信用 |
| 15 | Credit batch history is empty, cannot get current record | 碳信用 |
| 16 | Caller is not the current owner of this credit batch | 碳信用 |
| 17 | Credit batch is finalized and cannot be transferred | 碳信用 |
| 18 | Transfer quantity exceeds available balance | 碳信用 |
| 19 | Credit batch is already finalized | 碳信用 |
| 20 | Retirement quantity must be valid and within available balance | 碳信用 |
| 21 | Source credit is not in retired state | 碳信用 |
| 22 | Certificate ID is out of bounds | 证书管理 |
| 23 | Certificate history is empty, cannot get current record | 证书管理 |
| 24 | Certificate is already revoked | 证书管理 |
| 25 | Buyer address cannot be zero | 数据验证 |
| 26 | Quantity must be greater than zero | 数据验证 |

## 使用示例
```solidity
// 原来的错误信息
require(history.length > 0, "CC: User record does not exist for caller");

// 极致优化的数字错误代码
require(history.length > 0, "1");
```