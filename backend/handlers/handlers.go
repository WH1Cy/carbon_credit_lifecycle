package handlers

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"strconv"

	"go-carbon-backend/blockchain"
	"go-carbon-backend/contracts"
	"go-carbon-backend/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
)

// AppHandler 持有对区块链连接器的引用
type AppHandler struct {
	Connector *blockchain.BlockchainConnector
}

// 自定义的角色数组类型，确保JSON序列化为数字数组而不是Base64
// Go语言默认会将[]uint8序列化为Base64字符串，但前端期望数字数组
type RoleArray []uint8

// MarshalJSON 自定义JSON序列化，将[]uint8转换为[]int以避免Base64编码
func (r RoleArray) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}

	// 转换为int数组进行序列化，这样前端会收到 [1,2,3,4] 而不是 "AQIDBA=="
	intArray := make([]int, len(r))
	for i, role := range r {
		intArray[i] = int(role)
	}

	return json.Marshal(intArray)
}

// 用户记录结构体
type UserRecordResponse struct {
	Timestamp         int64     `json:"timestamp"`
	Addr              string    `json:"addr"`
	Name              string    `json:"name"`
	Bio               string    `json:"bio"`
	CredentialsHash   []string  `json:"credentialsHash"`
	Roles             RoleArray `json:"roles"`
	Editor            string    `json:"editor"`
	EditorRecordIndex int64     `json:"editorRecordIndex"`
	EditReason        string    `json:"editReason"`
}

// 自定义的技术数组类型，确保JSON序列化为数字数组而不是Base64
type TechnologyArray []uint8

// MarshalJSON 自定义JSON序列化
func (t TechnologyArray) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte("null"), nil
	}

	// 转换为int数组进行序列化
	intArray := make([]int, len(t))
	for i, tech := range t {
		intArray[i] = int(tech)
	}

	return json.Marshal(intArray)
}

// 项目记录结构体
type ProjectRecordResponse struct {
	Timestamp         int64           `json:"timestamp"`
	Id                int64           `json:"id"`
	Owner             string          `json:"owner"`
	OwnerRecordIndex  int64           `json:"ownerRecordIndex"`
	Name              string          `json:"name"`
	Location          string          `json:"location"`
	Technologies      TechnologyArray `json:"technologies"`
	Description       string          `json:"description"`
	Reduction         int64           `json:"reduction"`
	Status            uint8           `json:"status"`
	DocumentsHash     []string        `json:"documentsHash"`
	Editor            string          `json:"editor"`
	EditorRecordIndex int64           `json:"editorRecordIndex"`
	EditReason        string          `json:"editReason"`
}

// 碳信用记录结构体
type CreditRecordResponse struct {
	Timestamp        int64  `json:"timestamp"`
	EventType        uint8  `json:"eventType"`
	ProjectId        int64  `json:"projectId"`
	VintageYear      int64  `json:"vintageYear"`
	Quantity         int64  `json:"quantity"`
	CurrentOwner     string `json:"currentOwner"`
	OwnerRecordIndex int64  `json:"ownerRecordIndex"`
	From             string `json:"from"`
	FromRecordIndex  int64  `json:"fromRecordIndex"`
	To               string `json:"to"`
	ToRecordIndex    int64  `json:"toRecordIndex"`
	ParentBatchId    int64  `json:"parentBatchId"`
	ChildBatchId     int64  `json:"childBatchId"`
	Reason           string `json:"reason"`
}

// 证书记录结构体
type CertificateRecordResponse struct {
	Timestamp         int64  `json:"timestamp"`
	Id                int64  `json:"id"`
	CreditBatchId     int64  `json:"creditBatchId"`
	Quantity          int64  `json:"quantity"`
	Status            uint8  `json:"status"`
	Owner             string `json:"owner"`
	OwnerRecordIndex  int64  `json:"ownerRecordIndex"`
	Editor            string `json:"editor"`
	EditorRecordIndex int64  `json:"editorRecordIndex"`
	EditReason        string `json:"editReason"`
}

// 转换合约UserRecord到前端友好格式
func convertUserRecord(contractUser contracts.CarbonTrustConnectUserRecord) UserRecordResponse {
	return UserRecordResponse{
		Timestamp:         contractUser.Timestamp.Int64(),
		Addr:              contractUser.Addr.Hex(),
		Name:              contractUser.Name,
		Bio:               contractUser.Bio,
		CredentialsHash:   contractUser.CredentialsHash,
		Roles:             RoleArray(contractUser.Roles),
		Editor:            contractUser.Editor.Hex(),
		EditorRecordIndex: contractUser.EditorRecordIndex.Int64(),
		EditReason:        contractUser.EditReason,
	}
}

// 转换合约ProjectRecord到前端友好格式
func convertProjectRecord(contractProject contracts.CarbonTrustConnectProjectRecord) ProjectRecordResponse {
	return ProjectRecordResponse{
		Timestamp:         contractProject.Timestamp.Int64(),
		Id:                contractProject.Id.Int64(),
		Owner:             contractProject.Owner.Hex(),
		OwnerRecordIndex:  contractProject.OwnerRecordIndex.Int64(),
		Name:              contractProject.Name,
		Location:          contractProject.Location,
		Technologies:      TechnologyArray(contractProject.Technologies),
		Description:       contractProject.Description,
		Reduction:         contractProject.Reduction.Int64(),
		Status:            contractProject.Status,
		DocumentsHash:     contractProject.DocumentsHash,
		Editor:            contractProject.Editor.Hex(),
		EditorRecordIndex: contractProject.EditorRecordIndex.Int64(),
		EditReason:        contractProject.EditReason,
	}
}

// 转换合约CreditRecord到前端友好格式
func convertCreditRecord(contractCredit contracts.CarbonTrustConnectCreditRecord) CreditRecordResponse {
	return CreditRecordResponse{
		Timestamp:        contractCredit.Timestamp.Int64(),
		EventType:        contractCredit.EventType,
		ProjectId:        contractCredit.ProjectId.Int64(),
		VintageYear:      contractCredit.VintageYear.Int64(),
		Quantity:         contractCredit.Quantity.Int64(),
		CurrentOwner:     contractCredit.CurrentOwner.Hex(),
		OwnerRecordIndex: contractCredit.OwnerRecordIndex.Int64(),
		From:             contractCredit.From.Hex(),
		FromRecordIndex:  contractCredit.FromRecordIndex.Int64(),
		To:               contractCredit.To.Hex(),
		ToRecordIndex:    contractCredit.ToRecordIndex.Int64(),
		ParentBatchId:    contractCredit.ParentBatchId.Int64(),
		ChildBatchId:     contractCredit.ChildBatchId.Int64(),
		Reason:           contractCredit.Reason,
	}
}

// 转换合约CertificateRecord到前端友好格式
func convertCertificateRecord(contractCert contracts.CarbonTrustConnectCertificateRecord) CertificateRecordResponse {
	return CertificateRecordResponse{
		Timestamp:         contractCert.Timestamp.Int64(),
		Id:                contractCert.Id.Int64(),
		CreditBatchId:     contractCert.CreditBatchId.Int64(),
		Quantity:          contractCert.Quantity.Int64(),
		Status:            contractCert.Status,
		Owner:             contractCert.Owner.Hex(),
		OwnerRecordIndex:  contractCert.OwnerRecordIndex.Int64(),
		Editor:            contractCert.Editor.Hex(),
		EditorRecordIndex: contractCert.EditorRecordIndex.Int64(),
		EditReason:        contractCert.EditReason,
	}
}

// =================================================================================
// 读操作处理器
// =================================================================================

func (h *AppHandler) GetSummaryHandler(w http.ResponseWriter, r *http.Request) {
	totalUsers, totalProjects, totalCreditBatches, totalCertificates, err := h.Connector.Contract.GetPlatformSummary(&bind.CallOpts{})
	if err != nil {
		handleReadError(w, "GetPlatformSummary", err)
		return
	}
	response := map[string]interface{}{
		"totalUsers":         totalUsers,
		"totalProjects":      totalProjects,
		"totalCreditBatches": totalCreditBatches,
		"totalCertificates":  totalCertificates,
	}
	respondWithJSON(w, http.StatusOK, response)
}

func (h *AppHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.Connector.Contract.GetAllUserRecords(&bind.CallOpts{})
	if err != nil {
		handleReadError(w, "GetAllUserRecords", err)
		return
	}

	// 转换为前端友好格式
	var convertedUsers []UserRecordResponse
	for _, user := range users {
		convertedUsers = append(convertedUsers, convertUserRecord(user))
	}

	respondWithJSON(w, http.StatusOK, convertedUsers)
}

func (h *AppHandler) GetUserHistoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := common.HexToAddress(vars["address"])

	history, err := h.Connector.Contract.GetUserHistoryRecord(&bind.CallOpts{}, address)
	if err != nil {
		handleReadError(w, "GetUserHistoryRecord", err)
		return
	}

	// 转换为前端友好格式
	var convertedHistory []UserRecordResponse
	for _, record := range history {
		convertedHistory = append(convertedHistory, convertUserRecord(record))
	}

	respondWithJSON(w, http.StatusOK, convertedHistory)
}

func (h *AppHandler) GetAllProjectsHandler(w http.ResponseWriter, r *http.Request) {
	projects, err := h.Connector.Contract.GetAllProjectsWithDetails(&bind.CallOpts{})
	if err != nil {
		handleReadError(w, "GetAllProjectsWithDetails", err)
		return
	}

	// 转换为前端友好格式
	var convertedProjects []ProjectRecordResponse
	for _, project := range projects {
		convertedProjects = append(convertedProjects, convertProjectRecord(project))
	}

	respondWithJSON(w, http.StatusOK, convertedProjects)
}

func (h *AppHandler) GetProjectHistoryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	history, err := h.Connector.Contract.GetProjectHistoryRecord(&bind.CallOpts{}, id)
	if err != nil {
		handleReadError(w, "GetProjectHistoryRecord", err)
		return
	}

	// 转换为前端友好格式
	var convertedHistory []ProjectRecordResponse
	for _, record := range history {
		convertedHistory = append(convertedHistory, convertProjectRecord(record))
	}

	respondWithJSON(w, http.StatusOK, convertedHistory)
}

func (h *AppHandler) GetProjectReviewHistoryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	reviews, err := h.Connector.Contract.GetReviewHistoryRecord(&bind.CallOpts{}, id)
	if err != nil {
		handleReadError(w, "GetReviewHistoryRecord", err)
		return
	}
	respondWithJSON(w, http.StatusOK, reviews)
}

func (h *AppHandler) GetAllCreditsHandler(w http.ResponseWriter, r *http.Request) {
	credits, err := h.Connector.Contract.GetAllCreditBatchesWithDetails(&bind.CallOpts{})
	if err != nil {
		handleReadError(w, "GetAllCreditBatchesWithDetails", err)
		return
	}

	// 转换为前端友好格式
	var convertedCredits []CreditRecordResponse
	for _, credit := range credits {
		convertedCredits = append(convertedCredits, convertCreditRecord(credit))
	}

	respondWithJSON(w, http.StatusOK, convertedCredits)
}

func (h *AppHandler) GetCreditHistoryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	history, err := h.Connector.Contract.GetCreditHistoryRecord(&bind.CallOpts{}, id)
	if err != nil {
		handleReadError(w, "GetCreditHistoryRecord", err)
		return
	}

	// 转换为前端友好格式
	var convertedHistory []CreditRecordResponse
	for _, record := range history {
		convertedHistory = append(convertedHistory, convertCreditRecord(record))
	}

	respondWithJSON(w, http.StatusOK, convertedHistory)
}

func (h *AppHandler) GetAllCertificatesHandler(w http.ResponseWriter, r *http.Request) {
	certs, err := h.Connector.Contract.GetAllCertificatesWithDetails(&bind.CallOpts{})
	if err != nil {
		handleReadError(w, "GetAllCertificatesWithDetails", err)
		return
	}

	// 转换为前端友好格式
	var convertedCerts []CertificateRecordResponse
	for _, cert := range certs {
		convertedCerts = append(convertedCerts, convertCertificateRecord(cert))
	}

	respondWithJSON(w, http.StatusOK, convertedCerts)
}

func (h *AppHandler) GetCertificateHistoryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	history, err := h.Connector.Contract.GetCertificateHistoryRecord(&bind.CallOpts{}, id)
	if err != nil {
		handleReadError(w, "GetCertificateHistoryRecord", err)
		return
	}

	// 转换为前端友好格式
	var convertedHistory []CertificateRecordResponse
	for _, record := range history {
		convertedHistory = append(convertedHistory, convertCertificateRecord(record))
	}

	respondWithJSON(w, http.StatusOK, convertedHistory)
}

// =================================================================================
// 辅助函数
// =================================================================================

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func handleReadError(w http.ResponseWriter, funcName string, err error) {
	log.Printf("Error calling %s: %v", funcName, err)
	http.Error(w, utils.MapContractError(err), http.StatusInternalServerError)
}

func getIDFromRequest(r *http.Request, key string) (*big.Int, error) {
	vars := mux.Vars(r)
	idStr := vars[key]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, err
	}
	return big.NewInt(int64(id)), nil
}
