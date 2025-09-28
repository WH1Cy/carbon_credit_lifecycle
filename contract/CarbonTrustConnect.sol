// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title CarbonTrustConnect
 * @dev "碳信通"平台的核心智能合约。
 */
contract CarbonTrustConnect {
    // =================================================================================================
    // 状态变量与常量
    // =================================================================================================
    
    // --- 全实体历史追溯 ---
    mapping(address => UserRecord[]) public userRecordHistory;
    mapping(uint256 => ProjectRecord[]) public projectHistory;
    mapping(uint256 => ReviewRecord[]) public projectReviewHistory;
    mapping(uint256 => CreditRecord[]) public creditHistory;
    mapping(uint256 => CertificateRecord[]) public certificateHistory;
    
    // --- 最新记录下标索引 ---
    mapping(address => uint256) private _latestUserRecordIndex;
    mapping(uint256 => uint256) private _latestProjectRecordIndex;
    mapping(uint256 => uint256) private _latestReviewRecordIndex;
    mapping(uint256 => uint256) private _latestCreditRecordIndex;
    mapping(uint256 => uint256) private _latestCertificateRecordIndex;
    
    // --- 用户追踪 ---
    address[] public allUsers;
    mapping(address => bool) private _isUserTracked;

    // --- 计数器 ---
    uint256 private _projectCounter;
    uint256 private _creditBatchCounter;
    uint256 private _certificateCounter;

    // =================================================================================================
    // 数据结构
    // =================================================================================================
    
    enum Role { NONE, REGULATOR, PROJECT_OWNER, VERIFIER, BUYER }
    enum Technology { SOLAR, WIND, HYDRO, FORESTRY, CAPTURE, OTHER }
    enum ProjectStatus { EDITING, REVIEWING, APPROVED, REVOKED }
    enum CreditEventType { MINT, TRANSFER, RETIRE, REVOKE }
    enum CertificateStatus { ACTIVE, REVOKED }

    struct UserRecord {
        uint256 timestamp;
        address addr;
        string name; 
        string bio; 
        string[] credentialsHash; 
        Role[] roles;
        address editor;
        uint256 editorRecordIndex;
        string editReason;
    }

    struct ProjectRecord {
        uint256 timestamp; 
        uint256 id; 
        address owner;
        uint256 ownerRecordIndex;
        string name; 
        string location; 
        Technology[] technologies;
        string description;
        uint256 reduction; 
        ProjectStatus status; 
        string[] documentsHash;
        address editor;
        uint256 editorRecordIndex;
        string editReason;
    }
    
    struct ReviewRecord {
        uint256 timestamp;
        address handlingVerifier;
        uint256 verifierRecordIndex;
        bool approved; 
        string comment; 
        uint256 issuedBatchId;
    }

    struct CreditRecord {
        uint256 timestamp;
        CreditEventType eventType;
        uint256 projectId;
        uint256 vintageYear;
        uint256 quantity;
        address currentOwner;
        uint256 ownerRecordIndex;
        address from;
        uint256 fromRecordIndex;
        address to;
        uint256 toRecordIndex;
        uint256 parentBatchId;
        uint256 childBatchId;
        string reason;
    }

    struct CertificateRecord {
        uint256 timestamp;
        uint256 id;
        uint256 creditBatchId;
        uint256 quantity;
        CertificateStatus status;
        address owner;
        uint256 ownerRecordIndex;
        address editor;
        uint256 editorRecordIndex;
        string editReason;
    }

    // =================================================================================================
    // 事件
    // =================================================================================================
    event UserHistoryUpdated(address indexed userAddress, uint256 recordIndex);
    event ProjectHistoryUpdated(uint256 indexed projectId, uint256 recordIndex);
    event ReviewHistoryUpdated(uint256 indexed projectId, uint256 recordIndex);
    event CreditHistoryUpdated(uint256 indexed batchId, uint256 recordIndex);
    event CertificateHistoryUpdated(uint256 indexed certificateId, uint256 recordIndex);

    // =================================================================================================
    // 修饰符
    // =================================================================================================
    
    modifier hasRole(Role _role) {
        UserRecord[] storage history = userRecordHistory[msg.sender];
        require(history.length > 0, "1");

        Role[] memory currentRoles = history[_latestUserRecordIndex[msg.sender]].roles;
        for (uint i = 0; i < currentRoles.length; i++) {
            if (currentRoles[i] == _role) {
                _;
                return;
            }
        }
        revert("2");
    }

    modifier isAuthorized() {
        UserRecord[] storage history = userRecordHistory[msg.sender];
        require(history.length > 0, "1");
        
        Role[] memory currentRoles = history[_latestUserRecordIndex[msg.sender]].roles;
        require(currentRoles.length > 0, "2");
        _;
    }


    // =================================================================================================
    // 构造函数
    // =================================================================================================
    constructor() {
        Role[] memory initialRoles = new Role[](4);
        initialRoles[0] = Role.REGULATOR;
        initialRoles[1] = Role.PROJECT_OWNER;
        initialRoles[2] = Role.VERIFIER;
        initialRoles[3] = Role.BUYER;

        _trackUser(msg.sender);

        _addUserRecord(UserRecord({
            timestamp: block.timestamp, 
            addr: msg.sender,
            name: "Admin", 
            bio: "", 
            credentialsHash: new string[](0), 
            roles: initialRoles,
            editor: msg.sender,
            editorRecordIndex: 0,
            editReason: "Contract deployment"
        }));
    }

    // =================================================================================================
    // 内部辅助函数
    // =================================================================================================
    
    function _trackUser(address _user) private {
        if (!_isUserTracked[_user]) {
            _isUserTracked[_user] = true;
            allUsers.push(_user);
        }
    }

    function _addUserRecord(UserRecord memory _record) private {
        uint256 newIndex = userRecordHistory[_record.addr].length;
        userRecordHistory[_record.addr].push(_record);
        _latestUserRecordIndex[_record.addr] = newIndex;
        emit UserHistoryUpdated(_record.addr, newIndex);
    }

    function _addProjectRecord(ProjectRecord memory _record) private {
        uint256 newIndex = projectHistory[_record.id].length;
        projectHistory[_record.id].push(_record);
        _latestProjectRecordIndex[_record.id] = newIndex;
        emit ProjectHistoryUpdated(_record.id, newIndex);
    }

    function _addCreditRecord(uint256 _batchId, CreditRecord memory _record) private {
        uint256 newIndex = creditHistory[_batchId].length;
        creditHistory[_batchId].push(_record);
        _latestCreditRecordIndex[_batchId] = newIndex;
        emit CreditHistoryUpdated(_batchId, newIndex);
    }

    function _addCertificateRecord(CertificateRecord memory _record) private {
        uint256 newIndex = certificateHistory[_record.id].length;
        certificateHistory[_record.id].push(_record);
        _latestCertificateRecordIndex[_record.id] = newIndex;
        emit CertificateHistoryUpdated(_record.id, newIndex);
    }

    // --- 视图辅助函数 ---
    function _getCurrentProjectRecord(uint256 _projectId) private view returns (ProjectRecord storage) {
        require(_projectId > 0 && _projectId <= _projectCounter, "4");
        ProjectRecord[] storage history = projectHistory[_projectId];
        require(history.length > 0, "5");
        return history[_latestProjectRecordIndex[_projectId]];
    }

    function _getCurrentReviewRecord(uint256 _projectId) private view returns (ReviewRecord storage) {
        require(_projectId > 0 && _projectId <= _projectCounter, "4");
        ReviewRecord[] storage reviewHistory = projectReviewHistory[_projectId];
        require(reviewHistory.length > 0, "10");
        return reviewHistory[_latestReviewRecordIndex[_projectId]];
    }
    
    function _getCurrentCreditRecord(uint256 _batchId) private view returns (CreditRecord storage) {
        require(_batchId > 0 && _batchId <= _creditBatchCounter, "14");
        CreditRecord[] storage history = creditHistory[_batchId];
        require(history.length > 0, "15");
        return history[_latestCreditRecordIndex[_batchId]];
    }

    function _getCurrentCertificateRecord(uint256 _certificateId) private view returns (CertificateRecord storage) {
        require(_certificateId > 0 && _certificateId <= _certificateCounter, "22");
        CertificateRecord[] storage history = certificateHistory[_certificateId];
        require(history.length > 0, "23");
        return history[_latestCertificateRecordIndex[_certificateId]];
    }

    // =================================================================================================
    // 模块一: 用户与治理
    // =================================================================================================

    function updateUserProfile(
        string memory _name, 
        string memory _bio, 
        string[] memory _credentialsHash,
        string memory _reason
    ) external {
        _trackUser(msg.sender);
        UserRecord[] storage history = userRecordHistory[msg.sender];

        Role[] memory currentRoles = new Role[](0);
        uint256 editorIndex = 0;
        if (history.length > 0) { 
            uint256 lastIndex = _latestUserRecordIndex[msg.sender];
            currentRoles = history[lastIndex].roles;
            editorIndex = lastIndex;
        }

        _addUserRecord(UserRecord({
            timestamp: block.timestamp, 
            addr: msg.sender, 
            name: _name, 
            bio: _bio, 
            credentialsHash: _credentialsHash, 
            roles: currentRoles,
            editor: msg.sender, 
            editorRecordIndex: editorIndex, 
            editReason: _reason
        }));
    }

    function adminSetUserRoles(
        address _user, 
        Role[] calldata _newRoles, 
        string memory _reason
    ) external hasRole(Role.REGULATOR) {
        _trackUser(_user);
        UserRecord[] storage history = userRecordHistory[_user];

        UserRecord memory lastRecord;
        if (history.length > 0) { 
            lastRecord = history[_latestUserRecordIndex[_user]]; 
        }
        
        _addUserRecord(UserRecord({
            timestamp: block.timestamp, 
            addr: _user, 
            name: lastRecord.name, 
            bio: lastRecord.bio, 
            credentialsHash: lastRecord.credentialsHash, 
            roles: _newRoles,
            editor: msg.sender, 
            editorRecordIndex: _latestUserRecordIndex[msg.sender], 
            editReason: _reason
        }));
    }

    function adminUpdateUserProfile(
        address _user, 
        string memory _name, 
        string memory _bio, 
        string[] memory _credentialsHash, 
        string memory _reason
    ) external hasRole(Role.REGULATOR) {
        _trackUser(_user);
        UserRecord[] storage history = userRecordHistory[_user];

        Role[] memory currentRoles = new Role[](0);
        if (history.length > 0) { 
            currentRoles = history[_latestUserRecordIndex[_user]].roles; 
        }
        
        _addUserRecord(UserRecord({
            timestamp: block.timestamp, 
            addr: _user, 
            name: _name, 
            bio: _bio, 
            credentialsHash: _credentialsHash, 
            roles: currentRoles,
            editor: msg.sender, 
            editorRecordIndex: _latestUserRecordIndex[msg.sender], 
            editReason: _reason
        }));
    }

    // =================================================================================================
    // 模块二: 项目与治理
    // =================================================================================================

    function registerProject(
        string memory _name, 
        string memory _location, 
        Technology[] memory _technologies, 
        string memory _description, 
        uint256 _reduction, 
        string[] memory _documentsHash,
        string memory _reason
    ) external hasRole(Role.PROJECT_OWNER) {
        _projectCounter++;
        uint256 projectId = _projectCounter;
        uint256 userIndex = _latestUserRecordIndex[msg.sender];

        _addProjectRecord(ProjectRecord({
            timestamp: block.timestamp, 
            id: projectId, 
            owner: msg.sender, 
            ownerRecordIndex: userIndex, 
            name: _name, 
            location: _location, 
            technologies: _technologies, 
            description: _description, 
            reduction: _reduction, 
            status: ProjectStatus.EDITING, 
            documentsHash: _documentsHash,
            editor: msg.sender, 
            editorRecordIndex: userIndex, 
            editReason: _reason
        }));
    }

    function updateProjectDetails(
        uint256 _projectId, 
        string memory _name, 
        string memory _location, 
        Technology[] memory _technologies, 
        string memory _description, 
        uint256 _reduction, 
        string[] memory _documentsHash, 
        string memory _reason
    ) external {
        ProjectRecord storage lastRecord = _getCurrentProjectRecord(_projectId);
        require(lastRecord.owner == msg.sender, "6");
        require(lastRecord.status == ProjectStatus.EDITING, "7");

        uint256 userIndex = _latestUserRecordIndex[msg.sender];
        _addProjectRecord(ProjectRecord({
            timestamp: block.timestamp, 
            id: _projectId, 
            owner: msg.sender, 
            ownerRecordIndex: userIndex, 
            name: _name, 
            location: _location, 
            technologies: _technologies, 
            description: _description, 
            reduction: _reduction, 
            status: ProjectStatus.EDITING, 
            documentsHash: _documentsHash,
            editor: msg.sender, 
            editorRecordIndex: userIndex, 
            editReason: _reason
        }));
    }

    function adminUpdateProjectDetails(
        uint256 _projectId, 
        string memory _name, 
        string memory _location, 
        Technology[] memory _technologies, 
        string memory _description, 
        uint256 _reduction, 
        string[] memory _documentsHash, 
        string memory _reason
    ) external hasRole(Role.REGULATOR) {
        ProjectRecord storage lastRecord = _getCurrentProjectRecord(_projectId);

        _addProjectRecord(ProjectRecord({
            timestamp: block.timestamp, 
            id: _projectId, 
            owner: lastRecord.owner, 
            ownerRecordIndex: lastRecord.ownerRecordIndex, 
            name: _name, 
            location: _location, 
            technologies: _technologies, 
            description: _description, 
            reduction: _reduction, 
            status: lastRecord.status, 
            documentsHash: _documentsHash,
            editor: msg.sender, 
            editorRecordIndex: _latestUserRecordIndex[msg.sender], 
            editReason: _reason
        }));
    }

    function submitForVerification(uint256 _projectId) external {
        ProjectRecord storage lastRecord = _getCurrentProjectRecord(_projectId);
        require(lastRecord.owner == msg.sender, "6");
        require(lastRecord.status == ProjectStatus.EDITING, "8");
        
        ProjectRecord memory newRecord = lastRecord;
        newRecord.timestamp = block.timestamp;
        newRecord.status = ProjectStatus.REVIEWING;
        newRecord.editor = msg.sender;
        newRecord.editorRecordIndex = _latestUserRecordIndex[msg.sender];
        newRecord.editReason = "Submitted for verification";
        _addProjectRecord(newRecord);

        projectReviewHistory[_projectId].push(ReviewRecord({
            handlingVerifier: address(0), 
            verifierRecordIndex: 0, 
            timestamp: block.timestamp, 
            approved: false,
            comment: "New review cycle", 
            issuedBatchId: 0
        }));

        uint256 newReviewIndex = projectReviewHistory[_projectId].length - 1;
        _latestReviewRecordIndex[_projectId] = newReviewIndex;
        emit ReviewHistoryUpdated(_projectId, newReviewIndex);
    }
    
    function acceptVerification(uint256 _projectId) external hasRole(Role.VERIFIER) {
        ProjectRecord storage lastProjectRecord = _getCurrentProjectRecord(_projectId);
        require(lastProjectRecord.status == ProjectStatus.REVIEWING, "11");
        
        ReviewRecord storage lastReviewRecord = _getCurrentReviewRecord(_projectId);
        require(lastReviewRecord.handlingVerifier == address(0), "12");

        lastReviewRecord.handlingVerifier = msg.sender;
        lastReviewRecord.verifierRecordIndex = _latestUserRecordIndex[msg.sender];
        emit ReviewHistoryUpdated(_projectId, _latestReviewRecordIndex[_projectId]);
    }
    
    function finalizeVerification(
        uint256 _projectId, 
        bool _isApproved, 
        string memory _comment, 
        uint256 _creditQuantityToIssue, 
        uint256 _vintageYear
    ) external {
        ReviewRecord storage lastReview = _getCurrentReviewRecord(_projectId);
        require(lastReview.handlingVerifier == msg.sender, "13");

        ProjectRecord storage lastProjectRecord = _getCurrentProjectRecord(_projectId);
        require(lastProjectRecord.status == ProjectStatus.REVIEWING, "11");

        uint256 issuedBatchId = 0;
        ProjectRecord memory newProjectRecord = lastProjectRecord;
        newProjectRecord.timestamp = block.timestamp;
        newProjectRecord.editor = msg.sender;
        newProjectRecord.editorRecordIndex = _latestUserRecordIndex[msg.sender];
        
        if (_isApproved) {
            newProjectRecord.status = ProjectStatus.APPROVED;
            newProjectRecord.editReason = "Verification approved";
            _creditBatchCounter++;
            issuedBatchId = _creditBatchCounter;

            _addCreditRecord(issuedBatchId, CreditRecord({
                timestamp: block.timestamp, 
                eventType: CreditEventType.MINT, 
                projectId: _projectId, 
                vintageYear: _vintageYear, 
                quantity: _creditQuantityToIssue, 
                currentOwner: newProjectRecord.owner, 
                ownerRecordIndex: newProjectRecord.ownerRecordIndex, 
                from: address(0), 
                fromRecordIndex: 0, 
                to: address(0), 
                toRecordIndex: 0, 
                parentBatchId: 0, 
                childBatchId: 0, 
                reason: ""
            }));
        } else {
            newProjectRecord.status = ProjectStatus.EDITING;
            newProjectRecord.editReason = "Verification rejected, returned to edit";
        }

        _addProjectRecord(newProjectRecord);
        
        lastReview.approved = _isApproved;
        lastReview.comment = _comment;
        lastReview.issuedBatchId = issuedBatchId;
        emit ReviewHistoryUpdated(_projectId, _latestReviewRecordIndex[_projectId]);
    }
    
    function banProject(uint256 _projectId, string memory _reason) external hasRole(Role.REGULATOR) {
        ProjectRecord storage lastRecord = _getCurrentProjectRecord(_projectId);
        require(lastRecord.status != ProjectStatus.REVOKED, "9");

        ProjectRecord memory newRecord = lastRecord;
        newRecord.timestamp = block.timestamp;
        newRecord.status = ProjectStatus.REVOKED;
        newRecord.editor = msg.sender;
        newRecord.editorRecordIndex = _latestUserRecordIndex[msg.sender];
        newRecord.editReason = _reason;
        _addProjectRecord(newRecord);
        
        // Cascade revoke credits and certificates
        for (uint256 i = 1; i <= _creditBatchCounter; i++) {
            if (creditHistory[i].length > 0 && creditHistory[i][0].projectId == _projectId) {
                CreditRecord storage lastCreditRecord = _getCurrentCreditRecord(i);
                if (lastCreditRecord.eventType != CreditEventType.RETIRE && lastCreditRecord.eventType != CreditEventType.REVOKE) {
                    _addCreditRecord(i, CreditRecord({
                        timestamp: block.timestamp, 
                        eventType: CreditEventType.REVOKE, 
                        projectId: _projectId, 
                        vintageYear: lastCreditRecord.vintageYear, 
                        quantity: lastCreditRecord.quantity, 
                        currentOwner: lastCreditRecord.currentOwner, 
                        ownerRecordIndex: lastCreditRecord.ownerRecordIndex,
                        from: address(0), 
                        fromRecordIndex: 0, 
                        to: address(0), 
                        toRecordIndex: 0, 
                        parentBatchId: 0, 
                        childBatchId: 0, 
                        reason: _reason
                    }));
                }
            }
        }

        for (uint256 i = 1; i <= _certificateCounter; i++) {
            CertificateRecord storage lastCertRecord = _getCurrentCertificateRecord(i);
            if(lastCertRecord.status == CertificateStatus.ACTIVE) {
                 CreditRecord storage creditOrigin = creditHistory[lastCertRecord.creditBatchId][0];
                 if(creditOrigin.projectId == _projectId) {
                    _revokeCertificateInternal(i, _reason, false); // Do not restore credit
                 }
            }
        }
    }

    // =================================================================================================
    // 模块三: 碳信用与证书生命周期管理
    // =================================================================================================
    
    function mintCreditsForBuyer(
        address _buyer, 
        uint256 _quantity, 
        uint256 _vintageYear, 
        string memory _reason
    ) external hasRole(Role.REGULATOR) {
        require(_buyer != address(0), "25");
        require(_quantity > 0, "26");
        _trackUser(_buyer);

        _creditBatchCounter++;
        uint256 issuedBatchId = _creditBatchCounter;
        
        _addCreditRecord(issuedBatchId, CreditRecord({
            timestamp: block.timestamp, 
            eventType: CreditEventType.MINT, 
            projectId: 0,
            vintageYear: _vintageYear, 
            quantity: _quantity, 
            currentOwner: _buyer, 
            ownerRecordIndex: _latestUserRecordIndex[_buyer],
            from: address(0), 
            fromRecordIndex: 0, 
            to: address(0), 
            toRecordIndex: 0, 
            parentBatchId: 0, 
            childBatchId: 0, 
            reason: _reason
        }));
    }

    function splitAndTransfer(
        uint256 _batchId, 
        uint256 _quantity, 
        address _to
    ) external isAuthorized {
        CreditRecord storage lastRecord = _getCurrentCreditRecord(_batchId);
        require(lastRecord.currentOwner == msg.sender, "16");
        require(lastRecord.eventType != CreditEventType.RETIRE && lastRecord.eventType != CreditEventType.REVOKE, "17");
        require(_quantity > 0 && _quantity <= lastRecord.quantity, "18");
        
        address from = msg.sender;
        uint256 remainingQuantity = lastRecord.quantity - _quantity;
        
        _creditBatchCounter++;
        uint256 newBatchId = _creditBatchCounter;
        
        uint256 fromIndex = _latestUserRecordIndex[from];
        uint256 toIndex = _latestUserRecordIndex[_to];

        // Update old batch record for the remaining amount
        _addCreditRecord(_batchId, CreditRecord({ 
            timestamp: block.timestamp, 
            eventType: CreditEventType.TRANSFER, 
            projectId: lastRecord.projectId, 
            vintageYear: lastRecord.vintageYear, 
            quantity: remainingQuantity, 
            currentOwner: from, 
            ownerRecordIndex: fromIndex, 
            from: from, 
            fromRecordIndex: fromIndex, 
            to: _to, 
            toRecordIndex: toIndex, 
            parentBatchId: 0, 
            childBatchId: newBatchId, 
            reason: "" 
        }));
        
        // Create new batch record for the transferred amount
        _addCreditRecord(newBatchId, CreditRecord({ 
            timestamp: block.timestamp, 
            eventType: CreditEventType.MINT, 
            projectId: lastRecord.projectId, 
            vintageYear: lastRecord.vintageYear, 
            quantity: _quantity, 
            currentOwner: _to, 
            ownerRecordIndex: toIndex, 
            from: address(0), 
            fromRecordIndex: 0, 
            to: address(0), 
            toRecordIndex: 0, 
            parentBatchId: _batchId, 
            childBatchId: 0, 
            reason: "" 
        }));
    }

    function retireCredits(
        uint256 _batchId, 
        uint256 _quantity, 
        string memory _retirementReason
    ) external isAuthorized {
        CreditRecord storage lastRecord = _getCurrentCreditRecord(_batchId);
        require(lastRecord.currentOwner == msg.sender, "16");
        require(lastRecord.eventType != CreditEventType.RETIRE && lastRecord.eventType != CreditEventType.REVOKE, "19");
        require(_quantity > 0 && _quantity <= lastRecord.quantity, "20");

        address retiree = msg.sender;
        uint256 retireeIndex = _latestUserRecordIndex[retiree];
        uint256 remainingQuantity = lastRecord.quantity - _quantity;
        uint256 newBatchId = 0;

        if (remainingQuantity > 0) {
            _creditBatchCounter++;
            newBatchId = _creditBatchCounter;

            _addCreditRecord(newBatchId, CreditRecord({
                timestamp: block.timestamp,
                eventType: CreditEventType.MINT,
                projectId: lastRecord.projectId,
                vintageYear: lastRecord.vintageYear,
                quantity: remainingQuantity,
                currentOwner: retiree,
                ownerRecordIndex: retireeIndex,
                from: address(0), fromRecordIndex: 0, to: address(0), toRecordIndex: 0,
                parentBatchId: _batchId,
                childBatchId: 0,
                reason: "Partial retirement remainder"
            }));
        }

        // Mark the original (or part of it) as retired
        _addCreditRecord(_batchId, CreditRecord({
            timestamp: block.timestamp,
            eventType: CreditEventType.RETIRE,
            projectId: lastRecord.projectId,
            vintageYear: lastRecord.vintageYear,
            quantity: 0, 
            currentOwner: retiree,
            ownerRecordIndex: retireeIndex,
            from: address(0), fromRecordIndex: 0, to: address(0), toRecordIndex: 0,
            parentBatchId: 0,
            childBatchId: newBatchId, 
            reason: _retirementReason
        }));

        // Issue a new certificate for the retired amount
        _certificateCounter++;
        uint256 certificateId = _certificateCounter;

        _addCertificateRecord(CertificateRecord({
            timestamp: block.timestamp,
            id: certificateId,
            creditBatchId: _batchId,
            quantity: _quantity, 
            status: CertificateStatus.ACTIVE,
            owner: retiree,
            ownerRecordIndex: retireeIndex,
            editor: retiree,
            editorRecordIndex: retireeIndex,
            editReason: "Certificate issuance"
        }));
    }
    
    function revokeCredits(uint256 _batchId, string memory _reason) external hasRole(Role.REGULATOR) {
        CreditRecord storage lastRecord = _getCurrentCreditRecord(_batchId);
        require(lastRecord.eventType != CreditEventType.RETIRE && lastRecord.eventType != CreditEventType.REVOKE, "19");

        _addCreditRecord(_batchId, CreditRecord({ 
            timestamp: block.timestamp, 
            eventType: CreditEventType.REVOKE, 
            projectId: lastRecord.projectId, 
            vintageYear: lastRecord.vintageYear, 
            quantity: lastRecord.quantity, 
            currentOwner: lastRecord.currentOwner, 
            ownerRecordIndex: lastRecord.ownerRecordIndex,
            from: address(0), fromRecordIndex: 0, to: address(0), toRecordIndex: 0, 
            parentBatchId: 0, 
            childBatchId: 0, 
            reason: _reason 
        }));
    }

    function revokeCertificate(uint256 _certificateId, string memory _reason) external hasRole(Role.REGULATOR) {
        _revokeCertificateInternal(_certificateId, _reason, true); // Restore credit
    }

    function _revokeCertificateInternal(uint256 _certificateId, string memory _reason, bool _shouldRestoreCredit) private {
        CertificateRecord storage lastCertRecord = _getCurrentCertificateRecord(_certificateId);
        require(lastCertRecord.status == CertificateStatus.ACTIVE, "24");

        if (_shouldRestoreCredit) {
            CreditRecord storage lastCreditRecord = _getCurrentCreditRecord(lastCertRecord.creditBatchId);
            require(lastCreditRecord.eventType == CreditEventType.RETIRE, "21");

            // Restore the credit by creating a new 'MINT' record
            _addCreditRecord(lastCertRecord.creditBatchId, CreditRecord({
                timestamp: block.timestamp,
                eventType: CreditEventType.MINT, // Effectively un-retiring
                projectId: lastCreditRecord.projectId,
                vintageYear: lastCreditRecord.vintageYear,
                quantity: lastCreditRecord.quantity,
                currentOwner: lastCreditRecord.currentOwner,
                ownerRecordIndex: lastCreditRecord.ownerRecordIndex,
                from: address(0), fromRecordIndex: 0, to: address(0), toRecordIndex: 0,
                parentBatchId: 0, childBatchId: 0,
                reason: "Restored due to certificate revocation"
            }));
        }

        // Mark the certificate as revoked
        _addCertificateRecord(CertificateRecord({
            timestamp: block.timestamp,
            id: _certificateId,
            creditBatchId: lastCertRecord.creditBatchId,
            quantity: lastCertRecord.quantity,
            status: CertificateStatus.REVOKED,
            owner: lastCertRecord.owner,
            ownerRecordIndex: lastCertRecord.ownerRecordIndex,
            editor: msg.sender,
            editorRecordIndex: _latestUserRecordIndex[msg.sender],
            editReason: _reason
        }));
    }
    
    // =================================================================================================
    // 模块四: 视图函数
    // =================================================================================================

    function getAllUserRecords() external view returns (UserRecord[] memory) {
        UserRecord[] memory users = new UserRecord[](allUsers.length);
        for (uint i = 0; i < allUsers.length; i++) {
            address userAddr = allUsers[i];
            UserRecord[] storage history = userRecordHistory[userAddr];
            if (history.length > 0) {
                users[i] = history[_latestUserRecordIndex[userAddr]];
            }
        }
        return users;
    }

    function getAllProjectsWithDetails() external view returns (ProjectRecord[] memory) {
        ProjectRecord[] memory allProjectRecords = new ProjectRecord[](_projectCounter);
        for (uint i = 1; i <= _projectCounter; i++) {
            allProjectRecords[i-1] = _getCurrentProjectRecord(i);
        }
        return allProjectRecords;
    }

    function getAllCreditBatchesWithDetails() external view returns (CreditRecord[] memory) {
        CreditRecord[] memory allCreditRecords = new CreditRecord[](_creditBatchCounter);
        for (uint i = 1; i <= _creditBatchCounter; i++) {
            allCreditRecords[i-1] = _getCurrentCreditRecord(i);
        }
        return allCreditRecords;
    }

    function getAllCertificatesWithDetails() external view returns (CertificateRecord[] memory) {
        CertificateRecord[] memory allCertificateRecords = new CertificateRecord[](_certificateCounter);
        for (uint i = 1; i <= _certificateCounter; i++) {
            allCertificateRecords[i-1] = _getCurrentCertificateRecord(i);
        }
        return allCertificateRecords;
    }

    function getPlatformSummary() external view returns (uint256, uint256, uint256, uint256) { 
        return (allUsers.length, _projectCounter, _creditBatchCounter, _certificateCounter); 
    }

    // =================================================================================================
    // 模块五: 视图函数（获取完整历史记录与计数器）
    // =================================================================================================

    function getProjectHistoryRecord(uint256 _projectId) external view returns (ProjectRecord[] memory) {
        require(_projectId > 0 && _projectId <= _projectCounter, "4");
        return projectHistory[_projectId];
    }

    function getReviewHistoryRecord(uint256 _projectId) external view returns (ReviewRecord[] memory) {
        require(_projectId > 0 && _projectId <= _projectCounter, "4");
        return projectReviewHistory[_projectId];
    }

    function getUserHistoryRecord(address _user) external view returns (UserRecord[] memory) {
        require(userRecordHistory[_user].length > 0, "3");
        return userRecordHistory[_user];
    }

    function getCreditHistoryRecord(uint256 _batchId) external view returns (CreditRecord[] memory) {
        require(_batchId > 0 && _batchId <= _creditBatchCounter, "14");
        return creditHistory[_batchId];
    }

    function getCertificateHistoryRecord(uint256 _certificateId) external view returns (CertificateRecord[] memory) {
        require(_certificateId > 0 && _certificateId <= _certificateCounter, "22");
        return certificateHistory[_certificateId];
    }

    function getAllTrackedUsers() external view returns (address[] memory) {
        return allUsers;
    }

    function getCounters() external view returns (uint256 projectCounter, uint256 creditBatchCounter, uint256 certificateCounter) {
        return (_projectCounter, _creditBatchCounter, _certificateCounter);
    }
}