# 🌱 Carbon Trust Connect Platform

> 一个基于区块链的碳信用交易和管理平台，提供完整的碳信用生命周期管理解决方案

[![Vue 3](https://img.shields.io/badge/Vue-3.x-4FC08D?style=flat-square&logo=vue.js)](https://vuejs.org/)
[![TypeScript](https://img.shields.io/badge/TypeScript-5.x-3178C6?style=flat-square&logo=typescript)](https://www.typescriptlang.org/)
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go)](https://golang.org/)
[![Solidity](https://img.shields.io/badge/Solidity-0.8+-363636?style=flat-square&logo=solidity)](https://soliditylang.org/)
[![Ant Design](https://img.shields.io/badge/Ant%20Design-4.x-0170FE?style=flat-square&logo=ant-design)](https://ant.design/)

## 📋 目录

- [项目简介](#-项目简介)
- [核心功能](#-核心功能)
- [技术栈](#-技术栈)
- [项目结构](#-项目结构)
- [快速开始](#-快速开始)
- [环境配置](#-环境配置)
- [开发指南](#-开发指南)
- [部署指南](#-部署指南)
- [API文档](#-api文档)
- [智能合约](#-智能合约)
- [贡献指南](#-贡献指南)
- [许可证](#-许可证)

## 🌟 项目简介

Carbon Trust Connect Platform 是一个基于区块链技术的碳信用交易和管理平台，旨在为碳信用市场提供透明、安全、高效的解决方案。平台支持碳信用的发行、验证、交易、转移和注销等完整生命周期管理。

### 🎯 核心价值

- **透明度**：基于区块链的不可篡改记录
- **安全性**：智能合约保障交易安全
- **效率**：自动化流程减少人工干预
- **合规性**：符合国际碳信用标准
- **可扩展性**：支持多网络部署

## 🚀 核心功能

### 👥 角色管理
- **监管方 (Regulator)**：平台监管和规则制定
- **项目方 (Project Owner)**：碳信用项目创建和管理
- **审核员 (Verifier)**：项目审核和验证
- **购买方 (Buyer)**：碳信用购买和交易

### 📊 项目管理
- 项目创建和编辑
- 项目状态跟踪
- 审核流程管理
- 项目文档管理

### 💰 碳信用交易
- 碳信用发行
- 市场交易
- 价格发现
- 交易历史

### 🏆 证书管理
- 数字证书生成
- 证书验证
- 证书转移
- 证书注销

### 🎨 用户体验
- 现代化UI设计
- 浅色/深色主题切换
- 响应式布局
- 多语言支持

## 🛠 技术栈

### 前端技术
- **Vue 3** - 渐进式JavaScript框架
- **TypeScript** - 类型安全的JavaScript
- **Ant Design Vue** - 企业级UI组件库
- **Vue Router** - 官方路由管理器
- **Pinia** - 状态管理库
- **Axios** - HTTP客户端
- **Web3.js** - 区块链交互库

### 后端技术
- **Go** - 高性能编程语言
- **Gin** - Web框架
- **GORM** - ORM库
- **JWT** - 身份验证
- **Swagger** - API文档

### 区块链技术
- **Solidity** - 智能合约语言
- **Hardhat** - 开发环境
- **MetaMask** - 钱包集成
- **Sepolia Testnet** - 测试网络

### 开发工具
- **Vite** - 构建工具
- **ESLint** - 代码检查
- **Prettier** - 代码格式化
- **Git** - 版本控制

## 📁 项目结构

```
carbon_credit_lifecycle/
├── 📁 frontend/                 # 前端应用
│   ├── 📁 public/              # 静态资源
│   ├── 📁 src/                 # 源代码
│   │   ├── 📁 components/      # 公共组件
│   │   ├── 📁 views/           # 页面组件
│   │   ├── 📁 layouts/         # 布局组件
│   │   ├── 📁 stores/          # 状态管理
│   │   ├── 📁 services/        # 服务层
│   │   ├── 📁 composables/     # 组合式函数
│   │   ├── 📁 router/          # 路由配置
│   │   └── 📁 utils/           # 工具函数
│   ├── 📄 package.json         # 依赖配置
│   └── 📄 vite.config.ts       # 构建配置
├── 📁 backend/                 # 后端应用
│   ├── 📁 cmd/                 # 应用入口
│   ├── 📁 internal/            # 内部包
│   │   ├── 📁 api/             # API处理
│   │   ├── 📁 config/          # 配置管理
│   │   ├── 📁 database/        # 数据库
│   │   ├── 📁 middleware/      # 中间件
│   │   └── 📁 service/         # 业务逻辑
│   ├── 📁 pkg/                 # 公共包
│   ├── 📄 go.mod               # Go模块
│   └── 📄 main.go              # 主程序
├── 📁 contract/                # 智能合约
│   ├── 📁 contracts/           # 合约源码
│   ├── 📁 scripts/             # 部署脚本
│   ├── 📁 test/                # 合约测试
│   └── 📄 hardhat.config.js    # 配置
├── 📁 documents/               # 项目文档
│   ├── 📄 frontend_design.md   # 前端设计文档
│   ├── 📄 AddressAndKey.md     # 测试账户信息
│   └── 📄 ErrorCodes.md        # 错误代码
├── 📄 .gitignore              # Git忽略文件
└── 📄 README.md               # 项目说明
```

## 🚀 快速开始

### 环境要求

- **Node.js** >= 18.0.0
- **Go** >= 1.21
- **Git** >= 2.0
- **MetaMask** 浏览器扩展

### 克隆项目

```bash
git clone https://github.com/WH1Cy/carbon_credit_lifecycle.git
cd carbon_credit_lifecycle
```

### 启动前端

```bash
cd frontend
npm install
npm run dev
```

访问：http://localhost:3000

### 启动后端

```bash
cd backend
go mod tidy
go run main.go
```

API服务：http://localhost:8080

### 部署智能合约

```bash
cd contract
npm install
npx hardhat compile
npx hardhat run scripts/deploy.js --network sepolia
```

## ⚙️ 环境配置

### 前端环境变量

创建 `frontend/.env.local`：

```env
VITE_API_BASE_URL=http://localhost:8080
VITE_CONTRACT_ADDRESS=0x71F9Cbd442304A86ba905968fEa9d250d9F08b04
VITE_NETWORK_ID=11155111
```

### 后端环境变量

创建 `backend/.env`：

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=carbon_trust
JWT_SECRET=your_jwt_secret
```

### MetaMask配置

1. 安装MetaMask浏览器扩展
2. 切换到Sepolia测试网络
3. 导入测试账户（见 `documents/AddressAndKey.md`）

## 🛠 开发指南

### 前端开发

```bash
# 安装依赖
npm install

# 启动开发服务器
npm run dev

# 构建生产版本
npm run build

# 代码检查
npm run lint

# 代码格式化
npm run format
```

### 后端开发

```bash
# 安装依赖
go mod tidy

# 运行测试
go test ./...

# 启动服务
go run main.go

# 构建二进制文件
go build -o carbon-trust-backend
```

### 智能合约开发

```bash
# 安装依赖
npm install

# 编译合约
npx hardhat compile

# 运行测试
npx hardhat test

# 部署到测试网
npx hardhat run scripts/deploy.js --network sepolia
```

## 🚀 部署指南

### 前端部署

```bash
# 构建生产版本
npm run build

# 部署到静态托管服务
# 推荐：Vercel, Netlify, GitHub Pages
```

### 后端部署

```bash
# 构建Docker镜像
docker build -t carbon-trust-backend .

# 运行容器
docker run -p 8080:8080 carbon-trust-backend
```

### 智能合约部署

```bash
# 部署到主网
npx hardhat run scripts/deploy.js --network mainnet
```

## 📚 API文档

### 认证接口

```http
POST /api/auth/login
POST /api/auth/register
POST /api/auth/logout
```

### 项目管理

```http
GET    /api/projects          # 获取项目列表
POST   /api/projects          # 创建项目
GET    /api/projects/:id      # 获取项目详情
PUT    /api/projects/:id      # 更新项目
DELETE /api/projects/:id      # 删除项目
```

### 碳信用交易

```http
GET  /api/credits/market      # 获取市场数据
POST /api/credits/buy         # 购买碳信用
POST /api/credits/sell        # 出售碳信用
GET  /api/credits/history     # 交易历史
```

### 证书管理

```http
GET  /api/certificates        # 获取证书列表
POST /api/certificates        # 创建证书
GET  /api/certificates/:id    # 获取证书详情
PUT  /api/certificates/:id    # 更新证书
```

## 🔗 智能合约

### 主要合约

- **CarbonTrustConnect.sol** - 主合约，管理碳信用生命周期
- **ProjectManager.sol** - 项目管理合约
- **CreditMarket.sol** - 碳信用交易合约
- **CertificateManager.sol** - 证书管理合约

### 合约地址

- **Sepolia测试网**：`0x71F9Cbd442304A86ba905968fEa9d250d9F08b04`

### 主要功能

```solidity
// 创建项目
function createProject(string memory name, string memory description) external

// 发行碳信用
function issueCredits(uint256 projectId, uint256 amount) external

// 交易碳信用
function tradeCredits(uint256 creditId, uint256 price) external

// 转移碳信用
function transferCredits(address to, uint256 creditId) external

// 注销碳信用
function retireCredits(uint256 creditId) external
```

## 🧪 测试

### 前端测试

```bash
npm run test
npm run test:coverage
```

### 后端测试

```bash
go test ./...
go test -cover ./...
```

### 智能合约测试

```bash
npx hardhat test
npx hardhat coverage
```

## 📊 性能优化

### 前端优化

- **代码分割**：使用动态导入
- **懒加载**：路由和组件懒加载
- **缓存策略**：API响应缓存
- **图片优化**：WebP格式和压缩

### 后端优化

- **数据库索引**：优化查询性能
- **连接池**：数据库连接管理
- **缓存机制**：Redis缓存
- **负载均衡**：多实例部署

## 🔒 安全考虑

- **输入验证**：前后端数据验证
- **SQL注入防护**：参数化查询
- **XSS防护**：内容安全策略
- **CSRF防护**：令牌验证
- **智能合约审计**：代码安全审查

## 🤝 贡献指南

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开 Pull Request

### 代码规范

- 使用 ESLint 和 Prettier 格式化代码
- 遵循 Vue 3 组合式API规范
- 编写单元测试
- 更新相关文档

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 📞 联系我们

- **项目维护者**：WH1Cy
- **GitHub Issues**：[提交问题](https://github.com/WH1Cy/carbon_credit_lifecycle/issues)
- **邮箱**：your-email@example.com

## 🙏 致谢

感谢以下开源项目的支持：

- [Vue.js](https://vuejs.org/)
- [Ant Design Vue](https://antdv.com/)
- [Go](https://golang.org/)
- [Hardhat](https://hardhat.org/)
- [MetaMask](https://metamask.io/)

---

⭐ 如果这个项目对您有帮助，请给我们一个星标！