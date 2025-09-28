// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CarbonTrustConnectCertificateRecord is an auto generated low-level Go binding around an user-defined struct.
type CarbonTrustConnectCertificateRecord struct {
	Timestamp         *big.Int
	Id                *big.Int
	CreditBatchId     *big.Int
	Quantity          *big.Int
	Status            uint8
	Owner             common.Address
	OwnerRecordIndex  *big.Int
	Editor            common.Address
	EditorRecordIndex *big.Int
	EditReason        string
}

// CarbonTrustConnectCreditRecord is an auto generated low-level Go binding around an user-defined struct.
type CarbonTrustConnectCreditRecord struct {
	Timestamp        *big.Int
	EventType        uint8
	ProjectId        *big.Int
	VintageYear      *big.Int
	Quantity         *big.Int
	CurrentOwner     common.Address
	OwnerRecordIndex *big.Int
	From             common.Address
	FromRecordIndex  *big.Int
	To               common.Address
	ToRecordIndex    *big.Int
	ParentBatchId    *big.Int
	ChildBatchId     *big.Int
	Reason           string
}

// CarbonTrustConnectProjectRecord is an auto generated low-level Go binding around an user-defined struct.
type CarbonTrustConnectProjectRecord struct {
	Timestamp         *big.Int
	Id                *big.Int
	Owner             common.Address
	OwnerRecordIndex  *big.Int
	Name              string
	Location          string
	Technologies      []uint8
	Description       string
	Reduction         *big.Int
	Status            uint8
	DocumentsHash     []string
	Editor            common.Address
	EditorRecordIndex *big.Int
	EditReason        string
}

// CarbonTrustConnectReviewRecord is an auto generated low-level Go binding around an user-defined struct.
type CarbonTrustConnectReviewRecord struct {
	Timestamp           *big.Int
	HandlingVerifier    common.Address
	VerifierRecordIndex *big.Int
	Approved            bool
	Comment             string
	IssuedBatchId       *big.Int
}

// CarbonTrustConnectUserRecord is an auto generated low-level Go binding around an user-defined struct.
type CarbonTrustConnectUserRecord struct {
	Timestamp         *big.Int
	Addr              common.Address
	Name              string
	Bio               string
	CredentialsHash   []string
	Roles             []uint8
	Editor            common.Address
	EditorRecordIndex *big.Int
	EditReason        string
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"certificateId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"recordIndex\",\"type\":\"uint256\"}],\"name\":\"CertificateHistoryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"recordIndex\",\"type\":\"uint256\"}],\"name\":\"CreditHistoryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"recordIndex\",\"type\":\"uint256\"}],\"name\":\"ProjectHistoryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"recordIndex\",\"type\":\"uint256\"}],\"name\":\"ReviewHistoryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"recordIndex\",\"type\":\"uint256\"}],\"name\":\"UserHistoryUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"acceptVerification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"enumCarbonTrustConnect.Role[]\",\"name\":\"_newRoles\",\"type\":\"uint8[]\"},{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"}],\"name\":\"adminSetUserRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"},{\"internalType\":\"enumCarbonTrustConnect.Technology[]\",\"name\":\"_technologies\",\"type\":\"uint8[]\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_reduction\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"_documentsHash\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"}],\"name\":\"adminUpdateProjectDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_bio\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_credentialsHash\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"}],\"name\":\"adminUpdateUserProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allUsers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"}],\"name\":\"banProject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"certificateHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditBatchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"enumCarbonTrustConnect.CertificateStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ownerRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"editor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"editorRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"editReason\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"creditHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumCarbonTrustConnect.CreditEventType\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vintageYear\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ownerRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"parentBatchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"childBatchId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isApproved\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"_comment\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_creditQuantityToIssue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_vintageYear\",\"type\":\"uint256\"}],\"name\":\"finalizeVerification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllCertificatesWithDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditBatchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"enumCarbonTrustConnect.CertificateStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ownerRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"editor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"editorRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"editReason\",\"type\":\"string\"}],\"internalType\":\"structCarbonTrustConnect.CertificateRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllCreditBatchesWithDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumCarbonTrustConnect.CreditEventType\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vintageYear\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ownerRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"parentBatchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"childBatchId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"internalType\":\"structCarbonTrustConnect.CreditRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllProjectsWithDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ownerRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"enumCarbonTrustConnect.Technology[]\",\"name\":\"technologies\",\"type\":\"uint8[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"reduction\",\"type\":\"uint256\"},{\"internalType\":\"enumCarbonTrustConnect.ProjectStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string[]\",\"name\":\"documentsHash\",\"type\":\"string[]\"},{\"internalType\":\"address\",\"name\":\"editor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"editorRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"editReason\",\"type\":\"string\"}],\"internalType\":\"structCarbonTrustConnect.ProjectRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllTrackedUsers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllUserRecords\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bio\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"credentialsHash\",\"type\":\"string[]\"},{\"internalType\":\"enumCarbonTrustConnect.Role[]\",\"name\":\"roles\",\"type\":\"uint8[]\"},{\"internalType\":\"address\",\"name\":\"editor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"editorRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"editReason\",\"type\":\"string\"}],\"internalType\":\"structCarbonTrustConnect.UserRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_certificateId\",\"type\":\"uint256\"}],\"name\":\"getCertificateHistoryRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditBatchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"enumCarbonTrustConnect.CertificateStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ownerRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"editor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"editorRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"editReason\",\"type\":\"string\"}],\"internalType\":\"structCarbonTrustConnect.CertificateRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCounters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"projectCounter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditBatchCounter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"certificateCounter\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchId\",\"type\":\"uint256\"}],\"name\":\"getCreditHistoryRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumCarbonTrustConnect.CreditEventType\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vintageYear\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ownerRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"parentBatchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"childBatchId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"internalType\":\"structCarbonTrustConnect.CreditRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPlatformSummary\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"getProjectHistoryRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ownerRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"enumCarbonTrustConnect.Technology[]\",\"name\":\"technologies\",\"type\":\"uint8[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"reduction\",\"type\":\"uint256\"},{\"internalType\":\"enumCarbonTrustConnect.ProjectStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string[]\",\"name\":\"documentsHash\",\"type\":\"string[]\"},{\"internalType\":\"address\",\"name\":\"editor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"editorRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"editReason\",\"type\":\"string\"}],\"internalType\":\"structCarbonTrustConnect.ProjectRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"getReviewHistoryRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"handlingVerifier\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"verifierRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"issuedBatchId\",\"type\":\"uint256\"}],\"internalType\":\"structCarbonTrustConnect.ReviewRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserHistoryRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bio\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"credentialsHash\",\"type\":\"string[]\"},{\"internalType\":\"enumCarbonTrustConnect.Role[]\",\"name\":\"roles\",\"type\":\"uint8[]\"},{\"internalType\":\"address\",\"name\":\"editor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"editorRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"editReason\",\"type\":\"string\"}],\"internalType\":\"structCarbonTrustConnect.UserRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_vintageYear\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"}],\"name\":\"mintCreditsForBuyer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"projectHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ownerRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"reduction\",\"type\":\"uint256\"},{\"internalType\":\"enumCarbonTrustConnect.ProjectStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"editor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"editorRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"editReason\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"projectReviewHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"handlingVerifier\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"verifierRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"issuedBatchId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"},{\"internalType\":\"enumCarbonTrustConnect.Technology[]\",\"name\":\"_technologies\",\"type\":\"uint8[]\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_reduction\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"_documentsHash\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"}],\"name\":\"registerProject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_retirementReason\",\"type\":\"string\"}],\"name\":\"retireCredits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_certificateId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"}],\"name\":\"revokeCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"}],\"name\":\"revokeCredits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"splitAndTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"}],\"name\":\"submitForVerification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"},{\"internalType\":\"enumCarbonTrustConnect.Technology[]\",\"name\":\"_technologies\",\"type\":\"uint8[]\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_reduction\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"_documentsHash\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"}],\"name\":\"updateProjectDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_bio\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_credentialsHash\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"}],\"name\":\"updateUserProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userRecordHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bio\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"editor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"editorRecordIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"editReason\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// AllUsers is a free data retrieval call binding the contract method 0xa2bdedf4.
//
// Solidity: function allUsers(uint256 ) view returns(address)
func (_Contracts *ContractsCaller) AllUsers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "allUsers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllUsers is a free data retrieval call binding the contract method 0xa2bdedf4.
//
// Solidity: function allUsers(uint256 ) view returns(address)
func (_Contracts *ContractsSession) AllUsers(arg0 *big.Int) (common.Address, error) {
	return _Contracts.Contract.AllUsers(&_Contracts.CallOpts, arg0)
}

// AllUsers is a free data retrieval call binding the contract method 0xa2bdedf4.
//
// Solidity: function allUsers(uint256 ) view returns(address)
func (_Contracts *ContractsCallerSession) AllUsers(arg0 *big.Int) (common.Address, error) {
	return _Contracts.Contract.AllUsers(&_Contracts.CallOpts, arg0)
}

// CertificateHistory is a free data retrieval call binding the contract method 0x96d2a731.
//
// Solidity: function certificateHistory(uint256 , uint256 ) view returns(uint256 timestamp, uint256 id, uint256 creditBatchId, uint256 quantity, uint8 status, address owner, uint256 ownerRecordIndex, address editor, uint256 editorRecordIndex, string editReason)
func (_Contracts *ContractsCaller) CertificateHistory(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Timestamp         *big.Int
	Id                *big.Int
	CreditBatchId     *big.Int
	Quantity          *big.Int
	Status            uint8
	Owner             common.Address
	OwnerRecordIndex  *big.Int
	Editor            common.Address
	EditorRecordIndex *big.Int
	EditReason        string
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "certificateHistory", arg0, arg1)

	outstruct := new(struct {
		Timestamp         *big.Int
		Id                *big.Int
		CreditBatchId     *big.Int
		Quantity          *big.Int
		Status            uint8
		Owner             common.Address
		OwnerRecordIndex  *big.Int
		Editor            common.Address
		EditorRecordIndex *big.Int
		EditReason        string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Id = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CreditBatchId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Quantity = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.Owner = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.OwnerRecordIndex = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Editor = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.EditorRecordIndex = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.EditReason = *abi.ConvertType(out[9], new(string)).(*string)

	return *outstruct, err

}

// CertificateHistory is a free data retrieval call binding the contract method 0x96d2a731.
//
// Solidity: function certificateHistory(uint256 , uint256 ) view returns(uint256 timestamp, uint256 id, uint256 creditBatchId, uint256 quantity, uint8 status, address owner, uint256 ownerRecordIndex, address editor, uint256 editorRecordIndex, string editReason)
func (_Contracts *ContractsSession) CertificateHistory(arg0 *big.Int, arg1 *big.Int) (struct {
	Timestamp         *big.Int
	Id                *big.Int
	CreditBatchId     *big.Int
	Quantity          *big.Int
	Status            uint8
	Owner             common.Address
	OwnerRecordIndex  *big.Int
	Editor            common.Address
	EditorRecordIndex *big.Int
	EditReason        string
}, error) {
	return _Contracts.Contract.CertificateHistory(&_Contracts.CallOpts, arg0, arg1)
}

// CertificateHistory is a free data retrieval call binding the contract method 0x96d2a731.
//
// Solidity: function certificateHistory(uint256 , uint256 ) view returns(uint256 timestamp, uint256 id, uint256 creditBatchId, uint256 quantity, uint8 status, address owner, uint256 ownerRecordIndex, address editor, uint256 editorRecordIndex, string editReason)
func (_Contracts *ContractsCallerSession) CertificateHistory(arg0 *big.Int, arg1 *big.Int) (struct {
	Timestamp         *big.Int
	Id                *big.Int
	CreditBatchId     *big.Int
	Quantity          *big.Int
	Status            uint8
	Owner             common.Address
	OwnerRecordIndex  *big.Int
	Editor            common.Address
	EditorRecordIndex *big.Int
	EditReason        string
}, error) {
	return _Contracts.Contract.CertificateHistory(&_Contracts.CallOpts, arg0, arg1)
}

// CreditHistory is a free data retrieval call binding the contract method 0x64f298bf.
//
// Solidity: function creditHistory(uint256 , uint256 ) view returns(uint256 timestamp, uint8 eventType, uint256 projectId, uint256 vintageYear, uint256 quantity, address currentOwner, uint256 ownerRecordIndex, address from, uint256 fromRecordIndex, address to, uint256 toRecordIndex, uint256 parentBatchId, uint256 childBatchId, string reason)
func (_Contracts *ContractsCaller) CreditHistory(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Timestamp        *big.Int
	EventType        uint8
	ProjectId        *big.Int
	VintageYear      *big.Int
	Quantity         *big.Int
	CurrentOwner     common.Address
	OwnerRecordIndex *big.Int
	From             common.Address
	FromRecordIndex  *big.Int
	To               common.Address
	ToRecordIndex    *big.Int
	ParentBatchId    *big.Int
	ChildBatchId     *big.Int
	Reason           string
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "creditHistory", arg0, arg1)

	outstruct := new(struct {
		Timestamp        *big.Int
		EventType        uint8
		ProjectId        *big.Int
		VintageYear      *big.Int
		Quantity         *big.Int
		CurrentOwner     common.Address
		OwnerRecordIndex *big.Int
		From             common.Address
		FromRecordIndex  *big.Int
		To               common.Address
		ToRecordIndex    *big.Int
		ParentBatchId    *big.Int
		ChildBatchId     *big.Int
		Reason           string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EventType = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.ProjectId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.VintageYear = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Quantity = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.CurrentOwner = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.OwnerRecordIndex = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.From = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.FromRecordIndex = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.To = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)
	outstruct.ToRecordIndex = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.ParentBatchId = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.ChildBatchId = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)
	outstruct.Reason = *abi.ConvertType(out[13], new(string)).(*string)

	return *outstruct, err

}

// CreditHistory is a free data retrieval call binding the contract method 0x64f298bf.
//
// Solidity: function creditHistory(uint256 , uint256 ) view returns(uint256 timestamp, uint8 eventType, uint256 projectId, uint256 vintageYear, uint256 quantity, address currentOwner, uint256 ownerRecordIndex, address from, uint256 fromRecordIndex, address to, uint256 toRecordIndex, uint256 parentBatchId, uint256 childBatchId, string reason)
func (_Contracts *ContractsSession) CreditHistory(arg0 *big.Int, arg1 *big.Int) (struct {
	Timestamp        *big.Int
	EventType        uint8
	ProjectId        *big.Int
	VintageYear      *big.Int
	Quantity         *big.Int
	CurrentOwner     common.Address
	OwnerRecordIndex *big.Int
	From             common.Address
	FromRecordIndex  *big.Int
	To               common.Address
	ToRecordIndex    *big.Int
	ParentBatchId    *big.Int
	ChildBatchId     *big.Int
	Reason           string
}, error) {
	return _Contracts.Contract.CreditHistory(&_Contracts.CallOpts, arg0, arg1)
}

// CreditHistory is a free data retrieval call binding the contract method 0x64f298bf.
//
// Solidity: function creditHistory(uint256 , uint256 ) view returns(uint256 timestamp, uint8 eventType, uint256 projectId, uint256 vintageYear, uint256 quantity, address currentOwner, uint256 ownerRecordIndex, address from, uint256 fromRecordIndex, address to, uint256 toRecordIndex, uint256 parentBatchId, uint256 childBatchId, string reason)
func (_Contracts *ContractsCallerSession) CreditHistory(arg0 *big.Int, arg1 *big.Int) (struct {
	Timestamp        *big.Int
	EventType        uint8
	ProjectId        *big.Int
	VintageYear      *big.Int
	Quantity         *big.Int
	CurrentOwner     common.Address
	OwnerRecordIndex *big.Int
	From             common.Address
	FromRecordIndex  *big.Int
	To               common.Address
	ToRecordIndex    *big.Int
	ParentBatchId    *big.Int
	ChildBatchId     *big.Int
	Reason           string
}, error) {
	return _Contracts.Contract.CreditHistory(&_Contracts.CallOpts, arg0, arg1)
}

// GetAllCertificatesWithDetails is a free data retrieval call binding the contract method 0xf1c8e978.
//
// Solidity: function getAllCertificatesWithDetails() view returns((uint256,uint256,uint256,uint256,uint8,address,uint256,address,uint256,string)[])
func (_Contracts *ContractsCaller) GetAllCertificatesWithDetails(opts *bind.CallOpts) ([]CarbonTrustConnectCertificateRecord, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAllCertificatesWithDetails")

	if err != nil {
		return *new([]CarbonTrustConnectCertificateRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]CarbonTrustConnectCertificateRecord)).(*[]CarbonTrustConnectCertificateRecord)

	return out0, err

}

// GetAllCertificatesWithDetails is a free data retrieval call binding the contract method 0xf1c8e978.
//
// Solidity: function getAllCertificatesWithDetails() view returns((uint256,uint256,uint256,uint256,uint8,address,uint256,address,uint256,string)[])
func (_Contracts *ContractsSession) GetAllCertificatesWithDetails() ([]CarbonTrustConnectCertificateRecord, error) {
	return _Contracts.Contract.GetAllCertificatesWithDetails(&_Contracts.CallOpts)
}

// GetAllCertificatesWithDetails is a free data retrieval call binding the contract method 0xf1c8e978.
//
// Solidity: function getAllCertificatesWithDetails() view returns((uint256,uint256,uint256,uint256,uint8,address,uint256,address,uint256,string)[])
func (_Contracts *ContractsCallerSession) GetAllCertificatesWithDetails() ([]CarbonTrustConnectCertificateRecord, error) {
	return _Contracts.Contract.GetAllCertificatesWithDetails(&_Contracts.CallOpts)
}

// GetAllCreditBatchesWithDetails is a free data retrieval call binding the contract method 0x801316c7.
//
// Solidity: function getAllCreditBatchesWithDetails() view returns((uint256,uint8,uint256,uint256,uint256,address,uint256,address,uint256,address,uint256,uint256,uint256,string)[])
func (_Contracts *ContractsCaller) GetAllCreditBatchesWithDetails(opts *bind.CallOpts) ([]CarbonTrustConnectCreditRecord, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAllCreditBatchesWithDetails")

	if err != nil {
		return *new([]CarbonTrustConnectCreditRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]CarbonTrustConnectCreditRecord)).(*[]CarbonTrustConnectCreditRecord)

	return out0, err

}

// GetAllCreditBatchesWithDetails is a free data retrieval call binding the contract method 0x801316c7.
//
// Solidity: function getAllCreditBatchesWithDetails() view returns((uint256,uint8,uint256,uint256,uint256,address,uint256,address,uint256,address,uint256,uint256,uint256,string)[])
func (_Contracts *ContractsSession) GetAllCreditBatchesWithDetails() ([]CarbonTrustConnectCreditRecord, error) {
	return _Contracts.Contract.GetAllCreditBatchesWithDetails(&_Contracts.CallOpts)
}

// GetAllCreditBatchesWithDetails is a free data retrieval call binding the contract method 0x801316c7.
//
// Solidity: function getAllCreditBatchesWithDetails() view returns((uint256,uint8,uint256,uint256,uint256,address,uint256,address,uint256,address,uint256,uint256,uint256,string)[])
func (_Contracts *ContractsCallerSession) GetAllCreditBatchesWithDetails() ([]CarbonTrustConnectCreditRecord, error) {
	return _Contracts.Contract.GetAllCreditBatchesWithDetails(&_Contracts.CallOpts)
}

// GetAllProjectsWithDetails is a free data retrieval call binding the contract method 0x84c66da9.
//
// Solidity: function getAllProjectsWithDetails() view returns((uint256,uint256,address,uint256,string,string,uint8[],string,uint256,uint8,string[],address,uint256,string)[])
func (_Contracts *ContractsCaller) GetAllProjectsWithDetails(opts *bind.CallOpts) ([]CarbonTrustConnectProjectRecord, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAllProjectsWithDetails")

	if err != nil {
		return *new([]CarbonTrustConnectProjectRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]CarbonTrustConnectProjectRecord)).(*[]CarbonTrustConnectProjectRecord)

	return out0, err

}

// GetAllProjectsWithDetails is a free data retrieval call binding the contract method 0x84c66da9.
//
// Solidity: function getAllProjectsWithDetails() view returns((uint256,uint256,address,uint256,string,string,uint8[],string,uint256,uint8,string[],address,uint256,string)[])
func (_Contracts *ContractsSession) GetAllProjectsWithDetails() ([]CarbonTrustConnectProjectRecord, error) {
	return _Contracts.Contract.GetAllProjectsWithDetails(&_Contracts.CallOpts)
}

// GetAllProjectsWithDetails is a free data retrieval call binding the contract method 0x84c66da9.
//
// Solidity: function getAllProjectsWithDetails() view returns((uint256,uint256,address,uint256,string,string,uint8[],string,uint256,uint8,string[],address,uint256,string)[])
func (_Contracts *ContractsCallerSession) GetAllProjectsWithDetails() ([]CarbonTrustConnectProjectRecord, error) {
	return _Contracts.Contract.GetAllProjectsWithDetails(&_Contracts.CallOpts)
}

// GetAllTrackedUsers is a free data retrieval call binding the contract method 0x7295cafa.
//
// Solidity: function getAllTrackedUsers() view returns(address[])
func (_Contracts *ContractsCaller) GetAllTrackedUsers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAllTrackedUsers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllTrackedUsers is a free data retrieval call binding the contract method 0x7295cafa.
//
// Solidity: function getAllTrackedUsers() view returns(address[])
func (_Contracts *ContractsSession) GetAllTrackedUsers() ([]common.Address, error) {
	return _Contracts.Contract.GetAllTrackedUsers(&_Contracts.CallOpts)
}

// GetAllTrackedUsers is a free data retrieval call binding the contract method 0x7295cafa.
//
// Solidity: function getAllTrackedUsers() view returns(address[])
func (_Contracts *ContractsCallerSession) GetAllTrackedUsers() ([]common.Address, error) {
	return _Contracts.Contract.GetAllTrackedUsers(&_Contracts.CallOpts)
}

// GetAllUserRecords is a free data retrieval call binding the contract method 0x5ed416d2.
//
// Solidity: function getAllUserRecords() view returns((uint256,address,string,string,string[],uint8[],address,uint256,string)[])
func (_Contracts *ContractsCaller) GetAllUserRecords(opts *bind.CallOpts) ([]CarbonTrustConnectUserRecord, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAllUserRecords")

	if err != nil {
		return *new([]CarbonTrustConnectUserRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]CarbonTrustConnectUserRecord)).(*[]CarbonTrustConnectUserRecord)

	return out0, err

}

// GetAllUserRecords is a free data retrieval call binding the contract method 0x5ed416d2.
//
// Solidity: function getAllUserRecords() view returns((uint256,address,string,string,string[],uint8[],address,uint256,string)[])
func (_Contracts *ContractsSession) GetAllUserRecords() ([]CarbonTrustConnectUserRecord, error) {
	return _Contracts.Contract.GetAllUserRecords(&_Contracts.CallOpts)
}

// GetAllUserRecords is a free data retrieval call binding the contract method 0x5ed416d2.
//
// Solidity: function getAllUserRecords() view returns((uint256,address,string,string,string[],uint8[],address,uint256,string)[])
func (_Contracts *ContractsCallerSession) GetAllUserRecords() ([]CarbonTrustConnectUserRecord, error) {
	return _Contracts.Contract.GetAllUserRecords(&_Contracts.CallOpts)
}

// GetCertificateHistoryRecord is a free data retrieval call binding the contract method 0x404e47ae.
//
// Solidity: function getCertificateHistoryRecord(uint256 _certificateId) view returns((uint256,uint256,uint256,uint256,uint8,address,uint256,address,uint256,string)[])
func (_Contracts *ContractsCaller) GetCertificateHistoryRecord(opts *bind.CallOpts, _certificateId *big.Int) ([]CarbonTrustConnectCertificateRecord, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getCertificateHistoryRecord", _certificateId)

	if err != nil {
		return *new([]CarbonTrustConnectCertificateRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]CarbonTrustConnectCertificateRecord)).(*[]CarbonTrustConnectCertificateRecord)

	return out0, err

}

// GetCertificateHistoryRecord is a free data retrieval call binding the contract method 0x404e47ae.
//
// Solidity: function getCertificateHistoryRecord(uint256 _certificateId) view returns((uint256,uint256,uint256,uint256,uint8,address,uint256,address,uint256,string)[])
func (_Contracts *ContractsSession) GetCertificateHistoryRecord(_certificateId *big.Int) ([]CarbonTrustConnectCertificateRecord, error) {
	return _Contracts.Contract.GetCertificateHistoryRecord(&_Contracts.CallOpts, _certificateId)
}

// GetCertificateHistoryRecord is a free data retrieval call binding the contract method 0x404e47ae.
//
// Solidity: function getCertificateHistoryRecord(uint256 _certificateId) view returns((uint256,uint256,uint256,uint256,uint8,address,uint256,address,uint256,string)[])
func (_Contracts *ContractsCallerSession) GetCertificateHistoryRecord(_certificateId *big.Int) ([]CarbonTrustConnectCertificateRecord, error) {
	return _Contracts.Contract.GetCertificateHistoryRecord(&_Contracts.CallOpts, _certificateId)
}

// GetCounters is a free data retrieval call binding the contract method 0x6abb871f.
//
// Solidity: function getCounters() view returns(uint256 projectCounter, uint256 creditBatchCounter, uint256 certificateCounter)
func (_Contracts *ContractsCaller) GetCounters(opts *bind.CallOpts) (struct {
	ProjectCounter     *big.Int
	CreditBatchCounter *big.Int
	CertificateCounter *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getCounters")

	outstruct := new(struct {
		ProjectCounter     *big.Int
		CreditBatchCounter *big.Int
		CertificateCounter *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProjectCounter = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CreditBatchCounter = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CertificateCounter = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCounters is a free data retrieval call binding the contract method 0x6abb871f.
//
// Solidity: function getCounters() view returns(uint256 projectCounter, uint256 creditBatchCounter, uint256 certificateCounter)
func (_Contracts *ContractsSession) GetCounters() (struct {
	ProjectCounter     *big.Int
	CreditBatchCounter *big.Int
	CertificateCounter *big.Int
}, error) {
	return _Contracts.Contract.GetCounters(&_Contracts.CallOpts)
}

// GetCounters is a free data retrieval call binding the contract method 0x6abb871f.
//
// Solidity: function getCounters() view returns(uint256 projectCounter, uint256 creditBatchCounter, uint256 certificateCounter)
func (_Contracts *ContractsCallerSession) GetCounters() (struct {
	ProjectCounter     *big.Int
	CreditBatchCounter *big.Int
	CertificateCounter *big.Int
}, error) {
	return _Contracts.Contract.GetCounters(&_Contracts.CallOpts)
}

// GetCreditHistoryRecord is a free data retrieval call binding the contract method 0xc3b5e9d7.
//
// Solidity: function getCreditHistoryRecord(uint256 _batchId) view returns((uint256,uint8,uint256,uint256,uint256,address,uint256,address,uint256,address,uint256,uint256,uint256,string)[])
func (_Contracts *ContractsCaller) GetCreditHistoryRecord(opts *bind.CallOpts, _batchId *big.Int) ([]CarbonTrustConnectCreditRecord, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getCreditHistoryRecord", _batchId)

	if err != nil {
		return *new([]CarbonTrustConnectCreditRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]CarbonTrustConnectCreditRecord)).(*[]CarbonTrustConnectCreditRecord)

	return out0, err

}

// GetCreditHistoryRecord is a free data retrieval call binding the contract method 0xc3b5e9d7.
//
// Solidity: function getCreditHistoryRecord(uint256 _batchId) view returns((uint256,uint8,uint256,uint256,uint256,address,uint256,address,uint256,address,uint256,uint256,uint256,string)[])
func (_Contracts *ContractsSession) GetCreditHistoryRecord(_batchId *big.Int) ([]CarbonTrustConnectCreditRecord, error) {
	return _Contracts.Contract.GetCreditHistoryRecord(&_Contracts.CallOpts, _batchId)
}

// GetCreditHistoryRecord is a free data retrieval call binding the contract method 0xc3b5e9d7.
//
// Solidity: function getCreditHistoryRecord(uint256 _batchId) view returns((uint256,uint8,uint256,uint256,uint256,address,uint256,address,uint256,address,uint256,uint256,uint256,string)[])
func (_Contracts *ContractsCallerSession) GetCreditHistoryRecord(_batchId *big.Int) ([]CarbonTrustConnectCreditRecord, error) {
	return _Contracts.Contract.GetCreditHistoryRecord(&_Contracts.CallOpts, _batchId)
}

// GetPlatformSummary is a free data retrieval call binding the contract method 0x0f546a1e.
//
// Solidity: function getPlatformSummary() view returns(uint256, uint256, uint256, uint256)
func (_Contracts *ContractsCaller) GetPlatformSummary(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getPlatformSummary")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetPlatformSummary is a free data retrieval call binding the contract method 0x0f546a1e.
//
// Solidity: function getPlatformSummary() view returns(uint256, uint256, uint256, uint256)
func (_Contracts *ContractsSession) GetPlatformSummary() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Contracts.Contract.GetPlatformSummary(&_Contracts.CallOpts)
}

// GetPlatformSummary is a free data retrieval call binding the contract method 0x0f546a1e.
//
// Solidity: function getPlatformSummary() view returns(uint256, uint256, uint256, uint256)
func (_Contracts *ContractsCallerSession) GetPlatformSummary() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Contracts.Contract.GetPlatformSummary(&_Contracts.CallOpts)
}

// GetProjectHistoryRecord is a free data retrieval call binding the contract method 0xb7c2a279.
//
// Solidity: function getProjectHistoryRecord(uint256 _projectId) view returns((uint256,uint256,address,uint256,string,string,uint8[],string,uint256,uint8,string[],address,uint256,string)[])
func (_Contracts *ContractsCaller) GetProjectHistoryRecord(opts *bind.CallOpts, _projectId *big.Int) ([]CarbonTrustConnectProjectRecord, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getProjectHistoryRecord", _projectId)

	if err != nil {
		return *new([]CarbonTrustConnectProjectRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]CarbonTrustConnectProjectRecord)).(*[]CarbonTrustConnectProjectRecord)

	return out0, err

}

// GetProjectHistoryRecord is a free data retrieval call binding the contract method 0xb7c2a279.
//
// Solidity: function getProjectHistoryRecord(uint256 _projectId) view returns((uint256,uint256,address,uint256,string,string,uint8[],string,uint256,uint8,string[],address,uint256,string)[])
func (_Contracts *ContractsSession) GetProjectHistoryRecord(_projectId *big.Int) ([]CarbonTrustConnectProjectRecord, error) {
	return _Contracts.Contract.GetProjectHistoryRecord(&_Contracts.CallOpts, _projectId)
}

// GetProjectHistoryRecord is a free data retrieval call binding the contract method 0xb7c2a279.
//
// Solidity: function getProjectHistoryRecord(uint256 _projectId) view returns((uint256,uint256,address,uint256,string,string,uint8[],string,uint256,uint8,string[],address,uint256,string)[])
func (_Contracts *ContractsCallerSession) GetProjectHistoryRecord(_projectId *big.Int) ([]CarbonTrustConnectProjectRecord, error) {
	return _Contracts.Contract.GetProjectHistoryRecord(&_Contracts.CallOpts, _projectId)
}

// GetReviewHistoryRecord is a free data retrieval call binding the contract method 0xdaf8724a.
//
// Solidity: function getReviewHistoryRecord(uint256 _projectId) view returns((uint256,address,uint256,bool,string,uint256)[])
func (_Contracts *ContractsCaller) GetReviewHistoryRecord(opts *bind.CallOpts, _projectId *big.Int) ([]CarbonTrustConnectReviewRecord, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getReviewHistoryRecord", _projectId)

	if err != nil {
		return *new([]CarbonTrustConnectReviewRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]CarbonTrustConnectReviewRecord)).(*[]CarbonTrustConnectReviewRecord)

	return out0, err

}

// GetReviewHistoryRecord is a free data retrieval call binding the contract method 0xdaf8724a.
//
// Solidity: function getReviewHistoryRecord(uint256 _projectId) view returns((uint256,address,uint256,bool,string,uint256)[])
func (_Contracts *ContractsSession) GetReviewHistoryRecord(_projectId *big.Int) ([]CarbonTrustConnectReviewRecord, error) {
	return _Contracts.Contract.GetReviewHistoryRecord(&_Contracts.CallOpts, _projectId)
}

// GetReviewHistoryRecord is a free data retrieval call binding the contract method 0xdaf8724a.
//
// Solidity: function getReviewHistoryRecord(uint256 _projectId) view returns((uint256,address,uint256,bool,string,uint256)[])
func (_Contracts *ContractsCallerSession) GetReviewHistoryRecord(_projectId *big.Int) ([]CarbonTrustConnectReviewRecord, error) {
	return _Contracts.Contract.GetReviewHistoryRecord(&_Contracts.CallOpts, _projectId)
}

// GetUserHistoryRecord is a free data retrieval call binding the contract method 0xa5fc92ec.
//
// Solidity: function getUserHistoryRecord(address _user) view returns((uint256,address,string,string,string[],uint8[],address,uint256,string)[])
func (_Contracts *ContractsCaller) GetUserHistoryRecord(opts *bind.CallOpts, _user common.Address) ([]CarbonTrustConnectUserRecord, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getUserHistoryRecord", _user)

	if err != nil {
		return *new([]CarbonTrustConnectUserRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]CarbonTrustConnectUserRecord)).(*[]CarbonTrustConnectUserRecord)

	return out0, err

}

// GetUserHistoryRecord is a free data retrieval call binding the contract method 0xa5fc92ec.
//
// Solidity: function getUserHistoryRecord(address _user) view returns((uint256,address,string,string,string[],uint8[],address,uint256,string)[])
func (_Contracts *ContractsSession) GetUserHistoryRecord(_user common.Address) ([]CarbonTrustConnectUserRecord, error) {
	return _Contracts.Contract.GetUserHistoryRecord(&_Contracts.CallOpts, _user)
}

// GetUserHistoryRecord is a free data retrieval call binding the contract method 0xa5fc92ec.
//
// Solidity: function getUserHistoryRecord(address _user) view returns((uint256,address,string,string,string[],uint8[],address,uint256,string)[])
func (_Contracts *ContractsCallerSession) GetUserHistoryRecord(_user common.Address) ([]CarbonTrustConnectUserRecord, error) {
	return _Contracts.Contract.GetUserHistoryRecord(&_Contracts.CallOpts, _user)
}

// ProjectHistory is a free data retrieval call binding the contract method 0xffee4f1b.
//
// Solidity: function projectHistory(uint256 , uint256 ) view returns(uint256 timestamp, uint256 id, address owner, uint256 ownerRecordIndex, string name, string location, string description, uint256 reduction, uint8 status, address editor, uint256 editorRecordIndex, string editReason)
func (_Contracts *ContractsCaller) ProjectHistory(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Timestamp         *big.Int
	Id                *big.Int
	Owner             common.Address
	OwnerRecordIndex  *big.Int
	Name              string
	Location          string
	Description       string
	Reduction         *big.Int
	Status            uint8
	Editor            common.Address
	EditorRecordIndex *big.Int
	EditReason        string
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "projectHistory", arg0, arg1)

	outstruct := new(struct {
		Timestamp         *big.Int
		Id                *big.Int
		Owner             common.Address
		OwnerRecordIndex  *big.Int
		Name              string
		Location          string
		Description       string
		Reduction         *big.Int
		Status            uint8
		Editor            common.Address
		EditorRecordIndex *big.Int
		EditReason        string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Id = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.OwnerRecordIndex = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Location = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.Reduction = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[8], new(uint8)).(*uint8)
	outstruct.Editor = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)
	outstruct.EditorRecordIndex = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.EditReason = *abi.ConvertType(out[11], new(string)).(*string)

	return *outstruct, err

}

// ProjectHistory is a free data retrieval call binding the contract method 0xffee4f1b.
//
// Solidity: function projectHistory(uint256 , uint256 ) view returns(uint256 timestamp, uint256 id, address owner, uint256 ownerRecordIndex, string name, string location, string description, uint256 reduction, uint8 status, address editor, uint256 editorRecordIndex, string editReason)
func (_Contracts *ContractsSession) ProjectHistory(arg0 *big.Int, arg1 *big.Int) (struct {
	Timestamp         *big.Int
	Id                *big.Int
	Owner             common.Address
	OwnerRecordIndex  *big.Int
	Name              string
	Location          string
	Description       string
	Reduction         *big.Int
	Status            uint8
	Editor            common.Address
	EditorRecordIndex *big.Int
	EditReason        string
}, error) {
	return _Contracts.Contract.ProjectHistory(&_Contracts.CallOpts, arg0, arg1)
}

// ProjectHistory is a free data retrieval call binding the contract method 0xffee4f1b.
//
// Solidity: function projectHistory(uint256 , uint256 ) view returns(uint256 timestamp, uint256 id, address owner, uint256 ownerRecordIndex, string name, string location, string description, uint256 reduction, uint8 status, address editor, uint256 editorRecordIndex, string editReason)
func (_Contracts *ContractsCallerSession) ProjectHistory(arg0 *big.Int, arg1 *big.Int) (struct {
	Timestamp         *big.Int
	Id                *big.Int
	Owner             common.Address
	OwnerRecordIndex  *big.Int
	Name              string
	Location          string
	Description       string
	Reduction         *big.Int
	Status            uint8
	Editor            common.Address
	EditorRecordIndex *big.Int
	EditReason        string
}, error) {
	return _Contracts.Contract.ProjectHistory(&_Contracts.CallOpts, arg0, arg1)
}

// ProjectReviewHistory is a free data retrieval call binding the contract method 0x0c8eb864.
//
// Solidity: function projectReviewHistory(uint256 , uint256 ) view returns(uint256 timestamp, address handlingVerifier, uint256 verifierRecordIndex, bool approved, string comment, uint256 issuedBatchId)
func (_Contracts *ContractsCaller) ProjectReviewHistory(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Timestamp           *big.Int
	HandlingVerifier    common.Address
	VerifierRecordIndex *big.Int
	Approved            bool
	Comment             string
	IssuedBatchId       *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "projectReviewHistory", arg0, arg1)

	outstruct := new(struct {
		Timestamp           *big.Int
		HandlingVerifier    common.Address
		VerifierRecordIndex *big.Int
		Approved            bool
		Comment             string
		IssuedBatchId       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.HandlingVerifier = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.VerifierRecordIndex = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Approved = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.Comment = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.IssuedBatchId = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProjectReviewHistory is a free data retrieval call binding the contract method 0x0c8eb864.
//
// Solidity: function projectReviewHistory(uint256 , uint256 ) view returns(uint256 timestamp, address handlingVerifier, uint256 verifierRecordIndex, bool approved, string comment, uint256 issuedBatchId)
func (_Contracts *ContractsSession) ProjectReviewHistory(arg0 *big.Int, arg1 *big.Int) (struct {
	Timestamp           *big.Int
	HandlingVerifier    common.Address
	VerifierRecordIndex *big.Int
	Approved            bool
	Comment             string
	IssuedBatchId       *big.Int
}, error) {
	return _Contracts.Contract.ProjectReviewHistory(&_Contracts.CallOpts, arg0, arg1)
}

// ProjectReviewHistory is a free data retrieval call binding the contract method 0x0c8eb864.
//
// Solidity: function projectReviewHistory(uint256 , uint256 ) view returns(uint256 timestamp, address handlingVerifier, uint256 verifierRecordIndex, bool approved, string comment, uint256 issuedBatchId)
func (_Contracts *ContractsCallerSession) ProjectReviewHistory(arg0 *big.Int, arg1 *big.Int) (struct {
	Timestamp           *big.Int
	HandlingVerifier    common.Address
	VerifierRecordIndex *big.Int
	Approved            bool
	Comment             string
	IssuedBatchId       *big.Int
}, error) {
	return _Contracts.Contract.ProjectReviewHistory(&_Contracts.CallOpts, arg0, arg1)
}

// UserRecordHistory is a free data retrieval call binding the contract method 0x5e97b343.
//
// Solidity: function userRecordHistory(address , uint256 ) view returns(uint256 timestamp, address addr, string name, string bio, address editor, uint256 editorRecordIndex, string editReason)
func (_Contracts *ContractsCaller) UserRecordHistory(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Timestamp         *big.Int
	Addr              common.Address
	Name              string
	Bio               string
	Editor            common.Address
	EditorRecordIndex *big.Int
	EditReason        string
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "userRecordHistory", arg0, arg1)

	outstruct := new(struct {
		Timestamp         *big.Int
		Addr              common.Address
		Name              string
		Bio               string
		Editor            common.Address
		EditorRecordIndex *big.Int
		EditReason        string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Addr = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Bio = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Editor = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.EditorRecordIndex = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.EditReason = *abi.ConvertType(out[6], new(string)).(*string)

	return *outstruct, err

}

// UserRecordHistory is a free data retrieval call binding the contract method 0x5e97b343.
//
// Solidity: function userRecordHistory(address , uint256 ) view returns(uint256 timestamp, address addr, string name, string bio, address editor, uint256 editorRecordIndex, string editReason)
func (_Contracts *ContractsSession) UserRecordHistory(arg0 common.Address, arg1 *big.Int) (struct {
	Timestamp         *big.Int
	Addr              common.Address
	Name              string
	Bio               string
	Editor            common.Address
	EditorRecordIndex *big.Int
	EditReason        string
}, error) {
	return _Contracts.Contract.UserRecordHistory(&_Contracts.CallOpts, arg0, arg1)
}

// UserRecordHistory is a free data retrieval call binding the contract method 0x5e97b343.
//
// Solidity: function userRecordHistory(address , uint256 ) view returns(uint256 timestamp, address addr, string name, string bio, address editor, uint256 editorRecordIndex, string editReason)
func (_Contracts *ContractsCallerSession) UserRecordHistory(arg0 common.Address, arg1 *big.Int) (struct {
	Timestamp         *big.Int
	Addr              common.Address
	Name              string
	Bio               string
	Editor            common.Address
	EditorRecordIndex *big.Int
	EditReason        string
}, error) {
	return _Contracts.Contract.UserRecordHistory(&_Contracts.CallOpts, arg0, arg1)
}

// AcceptVerification is a paid mutator transaction binding the contract method 0xc46b0c68.
//
// Solidity: function acceptVerification(uint256 _projectId) returns()
func (_Contracts *ContractsTransactor) AcceptVerification(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "acceptVerification", _projectId)
}

// AcceptVerification is a paid mutator transaction binding the contract method 0xc46b0c68.
//
// Solidity: function acceptVerification(uint256 _projectId) returns()
func (_Contracts *ContractsSession) AcceptVerification(_projectId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AcceptVerification(&_Contracts.TransactOpts, _projectId)
}

// AcceptVerification is a paid mutator transaction binding the contract method 0xc46b0c68.
//
// Solidity: function acceptVerification(uint256 _projectId) returns()
func (_Contracts *ContractsTransactorSession) AcceptVerification(_projectId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AcceptVerification(&_Contracts.TransactOpts, _projectId)
}

// AdminSetUserRoles is a paid mutator transaction binding the contract method 0x13eccbfe.
//
// Solidity: function adminSetUserRoles(address _user, uint8[] _newRoles, string _reason) returns()
func (_Contracts *ContractsTransactor) AdminSetUserRoles(opts *bind.TransactOpts, _user common.Address, _newRoles []uint8, _reason string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "adminSetUserRoles", _user, _newRoles, _reason)
}

// AdminSetUserRoles is a paid mutator transaction binding the contract method 0x13eccbfe.
//
// Solidity: function adminSetUserRoles(address _user, uint8[] _newRoles, string _reason) returns()
func (_Contracts *ContractsSession) AdminSetUserRoles(_user common.Address, _newRoles []uint8, _reason string) (*types.Transaction, error) {
	return _Contracts.Contract.AdminSetUserRoles(&_Contracts.TransactOpts, _user, _newRoles, _reason)
}

// AdminSetUserRoles is a paid mutator transaction binding the contract method 0x13eccbfe.
//
// Solidity: function adminSetUserRoles(address _user, uint8[] _newRoles, string _reason) returns()
func (_Contracts *ContractsTransactorSession) AdminSetUserRoles(_user common.Address, _newRoles []uint8, _reason string) (*types.Transaction, error) {
	return _Contracts.Contract.AdminSetUserRoles(&_Contracts.TransactOpts, _user, _newRoles, _reason)
}

// AdminUpdateProjectDetails is a paid mutator transaction binding the contract method 0x16a61cf3.
//
// Solidity: function adminUpdateProjectDetails(uint256 _projectId, string _name, string _location, uint8[] _technologies, string _description, uint256 _reduction, string[] _documentsHash, string _reason) returns()
func (_Contracts *ContractsTransactor) AdminUpdateProjectDetails(opts *bind.TransactOpts, _projectId *big.Int, _name string, _location string, _technologies []uint8, _description string, _reduction *big.Int, _documentsHash []string, _reason string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "adminUpdateProjectDetails", _projectId, _name, _location, _technologies, _description, _reduction, _documentsHash, _reason)
}

// AdminUpdateProjectDetails is a paid mutator transaction binding the contract method 0x16a61cf3.
//
// Solidity: function adminUpdateProjectDetails(uint256 _projectId, string _name, string _location, uint8[] _technologies, string _description, uint256 _reduction, string[] _documentsHash, string _reason) returns()
func (_Contracts *ContractsSession) AdminUpdateProjectDetails(_projectId *big.Int, _name string, _location string, _technologies []uint8, _description string, _reduction *big.Int, _documentsHash []string, _reason string) (*types.Transaction, error) {
	return _Contracts.Contract.AdminUpdateProjectDetails(&_Contracts.TransactOpts, _projectId, _name, _location, _technologies, _description, _reduction, _documentsHash, _reason)
}

// AdminUpdateProjectDetails is a paid mutator transaction binding the contract method 0x16a61cf3.
//
// Solidity: function adminUpdateProjectDetails(uint256 _projectId, string _name, string _location, uint8[] _technologies, string _description, uint256 _reduction, string[] _documentsHash, string _reason) returns()
func (_Contracts *ContractsTransactorSession) AdminUpdateProjectDetails(_projectId *big.Int, _name string, _location string, _technologies []uint8, _description string, _reduction *big.Int, _documentsHash []string, _reason string) (*types.Transaction, error) {
	return _Contracts.Contract.AdminUpdateProjectDetails(&_Contracts.TransactOpts, _projectId, _name, _location, _technologies, _description, _reduction, _documentsHash, _reason)
}

// AdminUpdateUserProfile is a paid mutator transaction binding the contract method 0xf536be68.
//
// Solidity: function adminUpdateUserProfile(address _user, string _name, string _bio, string[] _credentialsHash, string _reason) returns()
func (_Contracts *ContractsTransactor) AdminUpdateUserProfile(opts *bind.TransactOpts, _user common.Address, _name string, _bio string, _credentialsHash []string, _reason string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "adminUpdateUserProfile", _user, _name, _bio, _credentialsHash, _reason)
}

// AdminUpdateUserProfile is a paid mutator transaction binding the contract method 0xf536be68.
//
// Solidity: function adminUpdateUserProfile(address _user, string _name, string _bio, string[] _credentialsHash, string _reason) returns()
func (_Contracts *ContractsSession) AdminUpdateUserProfile(_user common.Address, _name string, _bio string, _credentialsHash []string, _reason string) (*types.Transaction, error) {
	return _Contracts.Contract.AdminUpdateUserProfile(&_Contracts.TransactOpts, _user, _name, _bio, _credentialsHash, _reason)
}

// AdminUpdateUserProfile is a paid mutator transaction binding the contract method 0xf536be68.
//
// Solidity: function adminUpdateUserProfile(address _user, string _name, string _bio, string[] _credentialsHash, string _reason) returns()
func (_Contracts *ContractsTransactorSession) AdminUpdateUserProfile(_user common.Address, _name string, _bio string, _credentialsHash []string, _reason string) (*types.Transaction, error) {
	return _Contracts.Contract.AdminUpdateUserProfile(&_Contracts.TransactOpts, _user, _name, _bio, _credentialsHash, _reason)
}

// BanProject is a paid mutator transaction binding the contract method 0x9c0bb2e9.
//
// Solidity: function banProject(uint256 _projectId, string _reason) returns()
func (_Contracts *ContractsTransactor) BanProject(opts *bind.TransactOpts, _projectId *big.Int, _reason string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "banProject", _projectId, _reason)
}

// BanProject is a paid mutator transaction binding the contract method 0x9c0bb2e9.
//
// Solidity: function banProject(uint256 _projectId, string _reason) returns()
func (_Contracts *ContractsSession) BanProject(_projectId *big.Int, _reason string) (*types.Transaction, error) {
	return _Contracts.Contract.BanProject(&_Contracts.TransactOpts, _projectId, _reason)
}

// BanProject is a paid mutator transaction binding the contract method 0x9c0bb2e9.
//
// Solidity: function banProject(uint256 _projectId, string _reason) returns()
func (_Contracts *ContractsTransactorSession) BanProject(_projectId *big.Int, _reason string) (*types.Transaction, error) {
	return _Contracts.Contract.BanProject(&_Contracts.TransactOpts, _projectId, _reason)
}

// FinalizeVerification is a paid mutator transaction binding the contract method 0x0ec1eeea.
//
// Solidity: function finalizeVerification(uint256 _projectId, bool _isApproved, string _comment, uint256 _creditQuantityToIssue, uint256 _vintageYear) returns()
func (_Contracts *ContractsTransactor) FinalizeVerification(opts *bind.TransactOpts, _projectId *big.Int, _isApproved bool, _comment string, _creditQuantityToIssue *big.Int, _vintageYear *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "finalizeVerification", _projectId, _isApproved, _comment, _creditQuantityToIssue, _vintageYear)
}

// FinalizeVerification is a paid mutator transaction binding the contract method 0x0ec1eeea.
//
// Solidity: function finalizeVerification(uint256 _projectId, bool _isApproved, string _comment, uint256 _creditQuantityToIssue, uint256 _vintageYear) returns()
func (_Contracts *ContractsSession) FinalizeVerification(_projectId *big.Int, _isApproved bool, _comment string, _creditQuantityToIssue *big.Int, _vintageYear *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.FinalizeVerification(&_Contracts.TransactOpts, _projectId, _isApproved, _comment, _creditQuantityToIssue, _vintageYear)
}

// FinalizeVerification is a paid mutator transaction binding the contract method 0x0ec1eeea.
//
// Solidity: function finalizeVerification(uint256 _projectId, bool _isApproved, string _comment, uint256 _creditQuantityToIssue, uint256 _vintageYear) returns()
func (_Contracts *ContractsTransactorSession) FinalizeVerification(_projectId *big.Int, _isApproved bool, _comment string, _creditQuantityToIssue *big.Int, _vintageYear *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.FinalizeVerification(&_Contracts.TransactOpts, _projectId, _isApproved, _comment, _creditQuantityToIssue, _vintageYear)
}

// MintCreditsForBuyer is a paid mutator transaction binding the contract method 0x88d4c2d8.
//
// Solidity: function mintCreditsForBuyer(address _buyer, uint256 _quantity, uint256 _vintageYear, string _reason) returns()
func (_Contracts *ContractsTransactor) MintCreditsForBuyer(opts *bind.TransactOpts, _buyer common.Address, _quantity *big.Int, _vintageYear *big.Int, _reason string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "mintCreditsForBuyer", _buyer, _quantity, _vintageYear, _reason)
}

// MintCreditsForBuyer is a paid mutator transaction binding the contract method 0x88d4c2d8.
//
// Solidity: function mintCreditsForBuyer(address _buyer, uint256 _quantity, uint256 _vintageYear, string _reason) returns()
func (_Contracts *ContractsSession) MintCreditsForBuyer(_buyer common.Address, _quantity *big.Int, _vintageYear *big.Int, _reason string) (*types.Transaction, error) {
	return _Contracts.Contract.MintCreditsForBuyer(&_Contracts.TransactOpts, _buyer, _quantity, _vintageYear, _reason)
}

// MintCreditsForBuyer is a paid mutator transaction binding the contract method 0x88d4c2d8.
//
// Solidity: function mintCreditsForBuyer(address _buyer, uint256 _quantity, uint256 _vintageYear, string _reason) returns()
func (_Contracts *ContractsTransactorSession) MintCreditsForBuyer(_buyer common.Address, _quantity *big.Int, _vintageYear *big.Int, _reason string) (*types.Transaction, error) {
	return _Contracts.Contract.MintCreditsForBuyer(&_Contracts.TransactOpts, _buyer, _quantity, _vintageYear, _reason)
}

// RegisterProject is a paid mutator transaction binding the contract method 0xfa840f7e.
//
// Solidity: function registerProject(string _name, string _location, uint8[] _technologies, string _description, uint256 _reduction, string[] _documentsHash, string _reason) returns()
func (_Contracts *ContractsTransactor) RegisterProject(opts *bind.TransactOpts, _name string, _location string, _technologies []uint8, _description string, _reduction *big.Int, _documentsHash []string, _reason string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerProject", _name, _location, _technologies, _description, _reduction, _documentsHash, _reason)
}

// RegisterProject is a paid mutator transaction binding the contract method 0xfa840f7e.
//
// Solidity: function registerProject(string _name, string _location, uint8[] _technologies, string _description, uint256 _reduction, string[] _documentsHash, string _reason) returns()
func (_Contracts *ContractsSession) RegisterProject(_name string, _location string, _technologies []uint8, _description string, _reduction *big.Int, _documentsHash []string, _reason string) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterProject(&_Contracts.TransactOpts, _name, _location, _technologies, _description, _reduction, _documentsHash, _reason)
}

// RegisterProject is a paid mutator transaction binding the contract method 0xfa840f7e.
//
// Solidity: function registerProject(string _name, string _location, uint8[] _technologies, string _description, uint256 _reduction, string[] _documentsHash, string _reason) returns()
func (_Contracts *ContractsTransactorSession) RegisterProject(_name string, _location string, _technologies []uint8, _description string, _reduction *big.Int, _documentsHash []string, _reason string) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterProject(&_Contracts.TransactOpts, _name, _location, _technologies, _description, _reduction, _documentsHash, _reason)
}

// RetireCredits is a paid mutator transaction binding the contract method 0xc081ec96.
//
// Solidity: function retireCredits(uint256 _batchId, uint256 _quantity, string _retirementReason) returns()
func (_Contracts *ContractsTransactor) RetireCredits(opts *bind.TransactOpts, _batchId *big.Int, _quantity *big.Int, _retirementReason string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "retireCredits", _batchId, _quantity, _retirementReason)
}

// RetireCredits is a paid mutator transaction binding the contract method 0xc081ec96.
//
// Solidity: function retireCredits(uint256 _batchId, uint256 _quantity, string _retirementReason) returns()
func (_Contracts *ContractsSession) RetireCredits(_batchId *big.Int, _quantity *big.Int, _retirementReason string) (*types.Transaction, error) {
	return _Contracts.Contract.RetireCredits(&_Contracts.TransactOpts, _batchId, _quantity, _retirementReason)
}

// RetireCredits is a paid mutator transaction binding the contract method 0xc081ec96.
//
// Solidity: function retireCredits(uint256 _batchId, uint256 _quantity, string _retirementReason) returns()
func (_Contracts *ContractsTransactorSession) RetireCredits(_batchId *big.Int, _quantity *big.Int, _retirementReason string) (*types.Transaction, error) {
	return _Contracts.Contract.RetireCredits(&_Contracts.TransactOpts, _batchId, _quantity, _retirementReason)
}

// RevokeCertificate is a paid mutator transaction binding the contract method 0xb1466d1d.
//
// Solidity: function revokeCertificate(uint256 _certificateId, string _reason) returns()
func (_Contracts *ContractsTransactor) RevokeCertificate(opts *bind.TransactOpts, _certificateId *big.Int, _reason string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "revokeCertificate", _certificateId, _reason)
}

// RevokeCertificate is a paid mutator transaction binding the contract method 0xb1466d1d.
//
// Solidity: function revokeCertificate(uint256 _certificateId, string _reason) returns()
func (_Contracts *ContractsSession) RevokeCertificate(_certificateId *big.Int, _reason string) (*types.Transaction, error) {
	return _Contracts.Contract.RevokeCertificate(&_Contracts.TransactOpts, _certificateId, _reason)
}

// RevokeCertificate is a paid mutator transaction binding the contract method 0xb1466d1d.
//
// Solidity: function revokeCertificate(uint256 _certificateId, string _reason) returns()
func (_Contracts *ContractsTransactorSession) RevokeCertificate(_certificateId *big.Int, _reason string) (*types.Transaction, error) {
	return _Contracts.Contract.RevokeCertificate(&_Contracts.TransactOpts, _certificateId, _reason)
}

// RevokeCredits is a paid mutator transaction binding the contract method 0xaa036112.
//
// Solidity: function revokeCredits(uint256 _batchId, string _reason) returns()
func (_Contracts *ContractsTransactor) RevokeCredits(opts *bind.TransactOpts, _batchId *big.Int, _reason string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "revokeCredits", _batchId, _reason)
}

// RevokeCredits is a paid mutator transaction binding the contract method 0xaa036112.
//
// Solidity: function revokeCredits(uint256 _batchId, string _reason) returns()
func (_Contracts *ContractsSession) RevokeCredits(_batchId *big.Int, _reason string) (*types.Transaction, error) {
	return _Contracts.Contract.RevokeCredits(&_Contracts.TransactOpts, _batchId, _reason)
}

// RevokeCredits is a paid mutator transaction binding the contract method 0xaa036112.
//
// Solidity: function revokeCredits(uint256 _batchId, string _reason) returns()
func (_Contracts *ContractsTransactorSession) RevokeCredits(_batchId *big.Int, _reason string) (*types.Transaction, error) {
	return _Contracts.Contract.RevokeCredits(&_Contracts.TransactOpts, _batchId, _reason)
}

// SplitAndTransfer is a paid mutator transaction binding the contract method 0x84707273.
//
// Solidity: function splitAndTransfer(uint256 _batchId, uint256 _quantity, address _to) returns()
func (_Contracts *ContractsTransactor) SplitAndTransfer(opts *bind.TransactOpts, _batchId *big.Int, _quantity *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "splitAndTransfer", _batchId, _quantity, _to)
}

// SplitAndTransfer is a paid mutator transaction binding the contract method 0x84707273.
//
// Solidity: function splitAndTransfer(uint256 _batchId, uint256 _quantity, address _to) returns()
func (_Contracts *ContractsSession) SplitAndTransfer(_batchId *big.Int, _quantity *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SplitAndTransfer(&_Contracts.TransactOpts, _batchId, _quantity, _to)
}

// SplitAndTransfer is a paid mutator transaction binding the contract method 0x84707273.
//
// Solidity: function splitAndTransfer(uint256 _batchId, uint256 _quantity, address _to) returns()
func (_Contracts *ContractsTransactorSession) SplitAndTransfer(_batchId *big.Int, _quantity *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SplitAndTransfer(&_Contracts.TransactOpts, _batchId, _quantity, _to)
}

// SubmitForVerification is a paid mutator transaction binding the contract method 0x4eb907b8.
//
// Solidity: function submitForVerification(uint256 _projectId) returns()
func (_Contracts *ContractsTransactor) SubmitForVerification(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "submitForVerification", _projectId)
}

// SubmitForVerification is a paid mutator transaction binding the contract method 0x4eb907b8.
//
// Solidity: function submitForVerification(uint256 _projectId) returns()
func (_Contracts *ContractsSession) SubmitForVerification(_projectId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SubmitForVerification(&_Contracts.TransactOpts, _projectId)
}

// SubmitForVerification is a paid mutator transaction binding the contract method 0x4eb907b8.
//
// Solidity: function submitForVerification(uint256 _projectId) returns()
func (_Contracts *ContractsTransactorSession) SubmitForVerification(_projectId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SubmitForVerification(&_Contracts.TransactOpts, _projectId)
}

// UpdateProjectDetails is a paid mutator transaction binding the contract method 0xecfd37d0.
//
// Solidity: function updateProjectDetails(uint256 _projectId, string _name, string _location, uint8[] _technologies, string _description, uint256 _reduction, string[] _documentsHash, string _reason) returns()
func (_Contracts *ContractsTransactor) UpdateProjectDetails(opts *bind.TransactOpts, _projectId *big.Int, _name string, _location string, _technologies []uint8, _description string, _reduction *big.Int, _documentsHash []string, _reason string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateProjectDetails", _projectId, _name, _location, _technologies, _description, _reduction, _documentsHash, _reason)
}

// UpdateProjectDetails is a paid mutator transaction binding the contract method 0xecfd37d0.
//
// Solidity: function updateProjectDetails(uint256 _projectId, string _name, string _location, uint8[] _technologies, string _description, uint256 _reduction, string[] _documentsHash, string _reason) returns()
func (_Contracts *ContractsSession) UpdateProjectDetails(_projectId *big.Int, _name string, _location string, _technologies []uint8, _description string, _reduction *big.Int, _documentsHash []string, _reason string) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateProjectDetails(&_Contracts.TransactOpts, _projectId, _name, _location, _technologies, _description, _reduction, _documentsHash, _reason)
}

// UpdateProjectDetails is a paid mutator transaction binding the contract method 0xecfd37d0.
//
// Solidity: function updateProjectDetails(uint256 _projectId, string _name, string _location, uint8[] _technologies, string _description, uint256 _reduction, string[] _documentsHash, string _reason) returns()
func (_Contracts *ContractsTransactorSession) UpdateProjectDetails(_projectId *big.Int, _name string, _location string, _technologies []uint8, _description string, _reduction *big.Int, _documentsHash []string, _reason string) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateProjectDetails(&_Contracts.TransactOpts, _projectId, _name, _location, _technologies, _description, _reduction, _documentsHash, _reason)
}

// UpdateUserProfile is a paid mutator transaction binding the contract method 0xc90b6fa2.
//
// Solidity: function updateUserProfile(string _name, string _bio, string[] _credentialsHash, string _reason) returns()
func (_Contracts *ContractsTransactor) UpdateUserProfile(opts *bind.TransactOpts, _name string, _bio string, _credentialsHash []string, _reason string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateUserProfile", _name, _bio, _credentialsHash, _reason)
}

// UpdateUserProfile is a paid mutator transaction binding the contract method 0xc90b6fa2.
//
// Solidity: function updateUserProfile(string _name, string _bio, string[] _credentialsHash, string _reason) returns()
func (_Contracts *ContractsSession) UpdateUserProfile(_name string, _bio string, _credentialsHash []string, _reason string) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateUserProfile(&_Contracts.TransactOpts, _name, _bio, _credentialsHash, _reason)
}

// UpdateUserProfile is a paid mutator transaction binding the contract method 0xc90b6fa2.
//
// Solidity: function updateUserProfile(string _name, string _bio, string[] _credentialsHash, string _reason) returns()
func (_Contracts *ContractsTransactorSession) UpdateUserProfile(_name string, _bio string, _credentialsHash []string, _reason string) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateUserProfile(&_Contracts.TransactOpts, _name, _bio, _credentialsHash, _reason)
}

// ContractsCertificateHistoryUpdatedIterator is returned from FilterCertificateHistoryUpdated and is used to iterate over the raw logs and unpacked data for CertificateHistoryUpdated events raised by the Contracts contract.
type ContractsCertificateHistoryUpdatedIterator struct {
	Event *ContractsCertificateHistoryUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsCertificateHistoryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsCertificateHistoryUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsCertificateHistoryUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsCertificateHistoryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsCertificateHistoryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsCertificateHistoryUpdated represents a CertificateHistoryUpdated event raised by the Contracts contract.
type ContractsCertificateHistoryUpdated struct {
	CertificateId *big.Int
	RecordIndex   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCertificateHistoryUpdated is a free log retrieval operation binding the contract event 0xb772c5a2e62c27f4742bb413977965eec9e15628e35f83ddf3360f66a01f94a0.
//
// Solidity: event CertificateHistoryUpdated(uint256 indexed certificateId, uint256 recordIndex)
func (_Contracts *ContractsFilterer) FilterCertificateHistoryUpdated(opts *bind.FilterOpts, certificateId []*big.Int) (*ContractsCertificateHistoryUpdatedIterator, error) {

	var certificateIdRule []interface{}
	for _, certificateIdItem := range certificateId {
		certificateIdRule = append(certificateIdRule, certificateIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "CertificateHistoryUpdated", certificateIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsCertificateHistoryUpdatedIterator{contract: _Contracts.contract, event: "CertificateHistoryUpdated", logs: logs, sub: sub}, nil
}

// WatchCertificateHistoryUpdated is a free log subscription operation binding the contract event 0xb772c5a2e62c27f4742bb413977965eec9e15628e35f83ddf3360f66a01f94a0.
//
// Solidity: event CertificateHistoryUpdated(uint256 indexed certificateId, uint256 recordIndex)
func (_Contracts *ContractsFilterer) WatchCertificateHistoryUpdated(opts *bind.WatchOpts, sink chan<- *ContractsCertificateHistoryUpdated, certificateId []*big.Int) (event.Subscription, error) {

	var certificateIdRule []interface{}
	for _, certificateIdItem := range certificateId {
		certificateIdRule = append(certificateIdRule, certificateIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "CertificateHistoryUpdated", certificateIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsCertificateHistoryUpdated)
				if err := _Contracts.contract.UnpackLog(event, "CertificateHistoryUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCertificateHistoryUpdated is a log parse operation binding the contract event 0xb772c5a2e62c27f4742bb413977965eec9e15628e35f83ddf3360f66a01f94a0.
//
// Solidity: event CertificateHistoryUpdated(uint256 indexed certificateId, uint256 recordIndex)
func (_Contracts *ContractsFilterer) ParseCertificateHistoryUpdated(log types.Log) (*ContractsCertificateHistoryUpdated, error) {
	event := new(ContractsCertificateHistoryUpdated)
	if err := _Contracts.contract.UnpackLog(event, "CertificateHistoryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsCreditHistoryUpdatedIterator is returned from FilterCreditHistoryUpdated and is used to iterate over the raw logs and unpacked data for CreditHistoryUpdated events raised by the Contracts contract.
type ContractsCreditHistoryUpdatedIterator struct {
	Event *ContractsCreditHistoryUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsCreditHistoryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsCreditHistoryUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsCreditHistoryUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsCreditHistoryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsCreditHistoryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsCreditHistoryUpdated represents a CreditHistoryUpdated event raised by the Contracts contract.
type ContractsCreditHistoryUpdated struct {
	BatchId     *big.Int
	RecordIndex *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCreditHistoryUpdated is a free log retrieval operation binding the contract event 0x359d841244ed343d506ef5eb242db90640c03a50084c1ca70aa2dc2640d04999.
//
// Solidity: event CreditHistoryUpdated(uint256 indexed batchId, uint256 recordIndex)
func (_Contracts *ContractsFilterer) FilterCreditHistoryUpdated(opts *bind.FilterOpts, batchId []*big.Int) (*ContractsCreditHistoryUpdatedIterator, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "CreditHistoryUpdated", batchIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsCreditHistoryUpdatedIterator{contract: _Contracts.contract, event: "CreditHistoryUpdated", logs: logs, sub: sub}, nil
}

// WatchCreditHistoryUpdated is a free log subscription operation binding the contract event 0x359d841244ed343d506ef5eb242db90640c03a50084c1ca70aa2dc2640d04999.
//
// Solidity: event CreditHistoryUpdated(uint256 indexed batchId, uint256 recordIndex)
func (_Contracts *ContractsFilterer) WatchCreditHistoryUpdated(opts *bind.WatchOpts, sink chan<- *ContractsCreditHistoryUpdated, batchId []*big.Int) (event.Subscription, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "CreditHistoryUpdated", batchIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsCreditHistoryUpdated)
				if err := _Contracts.contract.UnpackLog(event, "CreditHistoryUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreditHistoryUpdated is a log parse operation binding the contract event 0x359d841244ed343d506ef5eb242db90640c03a50084c1ca70aa2dc2640d04999.
//
// Solidity: event CreditHistoryUpdated(uint256 indexed batchId, uint256 recordIndex)
func (_Contracts *ContractsFilterer) ParseCreditHistoryUpdated(log types.Log) (*ContractsCreditHistoryUpdated, error) {
	event := new(ContractsCreditHistoryUpdated)
	if err := _Contracts.contract.UnpackLog(event, "CreditHistoryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsProjectHistoryUpdatedIterator is returned from FilterProjectHistoryUpdated and is used to iterate over the raw logs and unpacked data for ProjectHistoryUpdated events raised by the Contracts contract.
type ContractsProjectHistoryUpdatedIterator struct {
	Event *ContractsProjectHistoryUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsProjectHistoryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsProjectHistoryUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsProjectHistoryUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsProjectHistoryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsProjectHistoryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsProjectHistoryUpdated represents a ProjectHistoryUpdated event raised by the Contracts contract.
type ContractsProjectHistoryUpdated struct {
	ProjectId   *big.Int
	RecordIndex *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProjectHistoryUpdated is a free log retrieval operation binding the contract event 0x23490d2931b87c11bea1b8498c31cb08054ed2417fff125008f0f1b8e508ce4b.
//
// Solidity: event ProjectHistoryUpdated(uint256 indexed projectId, uint256 recordIndex)
func (_Contracts *ContractsFilterer) FilterProjectHistoryUpdated(opts *bind.FilterOpts, projectId []*big.Int) (*ContractsProjectHistoryUpdatedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ProjectHistoryUpdated", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsProjectHistoryUpdatedIterator{contract: _Contracts.contract, event: "ProjectHistoryUpdated", logs: logs, sub: sub}, nil
}

// WatchProjectHistoryUpdated is a free log subscription operation binding the contract event 0x23490d2931b87c11bea1b8498c31cb08054ed2417fff125008f0f1b8e508ce4b.
//
// Solidity: event ProjectHistoryUpdated(uint256 indexed projectId, uint256 recordIndex)
func (_Contracts *ContractsFilterer) WatchProjectHistoryUpdated(opts *bind.WatchOpts, sink chan<- *ContractsProjectHistoryUpdated, projectId []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ProjectHistoryUpdated", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsProjectHistoryUpdated)
				if err := _Contracts.contract.UnpackLog(event, "ProjectHistoryUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProjectHistoryUpdated is a log parse operation binding the contract event 0x23490d2931b87c11bea1b8498c31cb08054ed2417fff125008f0f1b8e508ce4b.
//
// Solidity: event ProjectHistoryUpdated(uint256 indexed projectId, uint256 recordIndex)
func (_Contracts *ContractsFilterer) ParseProjectHistoryUpdated(log types.Log) (*ContractsProjectHistoryUpdated, error) {
	event := new(ContractsProjectHistoryUpdated)
	if err := _Contracts.contract.UnpackLog(event, "ProjectHistoryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsReviewHistoryUpdatedIterator is returned from FilterReviewHistoryUpdated and is used to iterate over the raw logs and unpacked data for ReviewHistoryUpdated events raised by the Contracts contract.
type ContractsReviewHistoryUpdatedIterator struct {
	Event *ContractsReviewHistoryUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsReviewHistoryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsReviewHistoryUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsReviewHistoryUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsReviewHistoryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsReviewHistoryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsReviewHistoryUpdated represents a ReviewHistoryUpdated event raised by the Contracts contract.
type ContractsReviewHistoryUpdated struct {
	ProjectId   *big.Int
	RecordIndex *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReviewHistoryUpdated is a free log retrieval operation binding the contract event 0x451237cdf091f70f43a5ba6831df266d7b74adde2ece4355a1d3c5e036b3f608.
//
// Solidity: event ReviewHistoryUpdated(uint256 indexed projectId, uint256 recordIndex)
func (_Contracts *ContractsFilterer) FilterReviewHistoryUpdated(opts *bind.FilterOpts, projectId []*big.Int) (*ContractsReviewHistoryUpdatedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ReviewHistoryUpdated", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsReviewHistoryUpdatedIterator{contract: _Contracts.contract, event: "ReviewHistoryUpdated", logs: logs, sub: sub}, nil
}

// WatchReviewHistoryUpdated is a free log subscription operation binding the contract event 0x451237cdf091f70f43a5ba6831df266d7b74adde2ece4355a1d3c5e036b3f608.
//
// Solidity: event ReviewHistoryUpdated(uint256 indexed projectId, uint256 recordIndex)
func (_Contracts *ContractsFilterer) WatchReviewHistoryUpdated(opts *bind.WatchOpts, sink chan<- *ContractsReviewHistoryUpdated, projectId []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ReviewHistoryUpdated", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsReviewHistoryUpdated)
				if err := _Contracts.contract.UnpackLog(event, "ReviewHistoryUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReviewHistoryUpdated is a log parse operation binding the contract event 0x451237cdf091f70f43a5ba6831df266d7b74adde2ece4355a1d3c5e036b3f608.
//
// Solidity: event ReviewHistoryUpdated(uint256 indexed projectId, uint256 recordIndex)
func (_Contracts *ContractsFilterer) ParseReviewHistoryUpdated(log types.Log) (*ContractsReviewHistoryUpdated, error) {
	event := new(ContractsReviewHistoryUpdated)
	if err := _Contracts.contract.UnpackLog(event, "ReviewHistoryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsUserHistoryUpdatedIterator is returned from FilterUserHistoryUpdated and is used to iterate over the raw logs and unpacked data for UserHistoryUpdated events raised by the Contracts contract.
type ContractsUserHistoryUpdatedIterator struct {
	Event *ContractsUserHistoryUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsUserHistoryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsUserHistoryUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsUserHistoryUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsUserHistoryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsUserHistoryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsUserHistoryUpdated represents a UserHistoryUpdated event raised by the Contracts contract.
type ContractsUserHistoryUpdated struct {
	UserAddress common.Address
	RecordIndex *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUserHistoryUpdated is a free log retrieval operation binding the contract event 0x370e862073d2323f8f407e5f42332b6e48f34f880e38cc54f2b7aa06b156ebb5.
//
// Solidity: event UserHistoryUpdated(address indexed userAddress, uint256 recordIndex)
func (_Contracts *ContractsFilterer) FilterUserHistoryUpdated(opts *bind.FilterOpts, userAddress []common.Address) (*ContractsUserHistoryUpdatedIterator, error) {

	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "UserHistoryUpdated", userAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractsUserHistoryUpdatedIterator{contract: _Contracts.contract, event: "UserHistoryUpdated", logs: logs, sub: sub}, nil
}

// WatchUserHistoryUpdated is a free log subscription operation binding the contract event 0x370e862073d2323f8f407e5f42332b6e48f34f880e38cc54f2b7aa06b156ebb5.
//
// Solidity: event UserHistoryUpdated(address indexed userAddress, uint256 recordIndex)
func (_Contracts *ContractsFilterer) WatchUserHistoryUpdated(opts *bind.WatchOpts, sink chan<- *ContractsUserHistoryUpdated, userAddress []common.Address) (event.Subscription, error) {

	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "UserHistoryUpdated", userAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsUserHistoryUpdated)
				if err := _Contracts.contract.UnpackLog(event, "UserHistoryUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserHistoryUpdated is a log parse operation binding the contract event 0x370e862073d2323f8f407e5f42332b6e48f34f880e38cc54f2b7aa06b156ebb5.
//
// Solidity: event UserHistoryUpdated(address indexed userAddress, uint256 recordIndex)
func (_Contracts *ContractsFilterer) ParseUserHistoryUpdated(log types.Log) (*ContractsUserHistoryUpdated, error) {
	event := new(ContractsUserHistoryUpdated)
	if err := _Contracts.contract.UnpackLog(event, "UserHistoryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
