# 碳信通前端应用

基于 Vue 3 + Ant Design Vue + TypeScript 开发的碳信用生命周期管理平台前端应用。

## 技术栈

- **框架**: Vue 3 + TypeScript
- **UI组件库**: Ant Design Vue 4.x
- **构建工具**: Vite
- **状态管理**: Pinia
- **路由**: Vue Router 4
- **Web3集成**: Web3.js + MetaMask
- **样式**: CSS + Ant Design Vue主题

## 环境要求

- Node.js >= 16.0.0
- npm >= 8.0.0 或 yarn >= 1.22.0
- MetaMask浏览器插件

## 安装和运行

### 1. 安装依赖

```bash
npm install
# 或
yarn install
```

### 2. 启动开发服务器

```bash
npm run dev
# 或
yarn dev
```

应用将在 `http://localhost:3000` 启动

### 3. 构建生产版本

```bash
npm run build
# 或
yarn build
```

### 4. 预览构建结果

```bash
npm run preview
# 或
yarn preview
```

## 项目结构

```
src/
├── components/          # 公共组件
├── layouts/            # 布局组件
│   └── BasicLayout.vue # 主布局
├── router/             # 路由配置
│   └── index.ts
├── services/           # 服务层
│   ├── api.ts          # API接口
│   └── web3.ts         # Web3服务
├── stores/             # 状态管理
│   └── user.ts         # 用户状态
├── types/              # TypeScript类型定义
│   └── index.ts
├── views/              # 页面组件
│   ├── dashboard/      # 概览页面
│   ├── users/          # 用户管理
│   ├── projects/       # 项目管理
│   ├── credits/        # 碳信用管理
│   └── certificates/   # 证书管理
├── App.vue             # 根组件
└── main.ts             # 应用入口
```