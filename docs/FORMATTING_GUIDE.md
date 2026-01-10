# Code Formatting & Linting Guide

## Overview

Dự án sử dụng **ESLint** và **Prettier** để đảm bảo code quality và consistent formatting. VSCode đã được cấu hình sẵn để tự động format khi save file.

## VSCode Extensions (Auto-Recommended)

Khi mở project lần đầu, **VSCode sẽ tự động hiện popup** đề xuất cài đặt các extensions cần thiết.

### Required Extensions

Các extensions được định nghĩa trong `.vscode/extensions.json`:

1. **ESLint** (`dbaeumer.vscode-eslint`) - Code quality linting
2. **Prettier** (`esbenp.prettier-vscode`) - Code formatting

### Recommended Extensions

3. **TypeScript** (`ms-vscode.vscode-typescript-next`) - Enhanced TS support
4. **React Snippets** (`dsznajder.es7-react-js-snippets`) - React code snippets
5. **GitLens** (`eamodio.gitlens`) - Git integration
6. **Path Intellisense** (`christian-kohler.path-intellisense`) - Path autocomplete

### Cách cài đặt

**Option 1: Tự động (Recommended)**

- VSCode sẽ tự động hiện popup khi mở project
- Click **"Install All"** hoặc **"Show Recommended Extensions"**

**Option 2: Manual**

- Mở Extensions panel (`Cmd+Shift+X`)
- Search `@recommended`
- Click **"Install Workspace Recommended Extensions"**

**Option 3: Command Palette**

```
Cmd+Shift+P → "Extensions: Show Recommended Extensions"
```

## Auto Format on Save

VSCode đã được config sẵn trong `.vscode/settings.json`:

- **Format on Save**: Tự động chạy Prettier khi save
- **ESLint Auto-fix**: Tự động fix các lỗi ESLint có thể fix được

## Commands

### Formatting

```bash
# Format toàn bộ dự án
yarn format

# Kiểm tra xem code đã đúng format chưa
yarn format:check
```

### Linting

```bash
# Check ESLint
yarn lint

# Auto-fix ESLint issues
yarn lint:fix

# Fix cả Prettier + ESLint
yarn format:all
```

### Setup cho Dev mới

```bash
# Option 1: Dùng script setup (recommended)
./scripts/setup-dev.sh

# Option 2: Dùng yarn command
yarn setup
```

## ESLint Rules Highlights

### 1. Import Ordering

Imports phải được sắp xếp theo thứ tự:

1. React, Next.js, MUI
2. Internal types & constants
3. Components (atoms → molecules → organisms)
4. Hooks & utils
5. Styles (cuối cùng)

### 2. Absolute Imports (Warning)

- Prefer: `import { Button } from '@/components/atoms/Button'`
- Avoid: `import { Button } from './Button'` hoặc `import { Button } from '../atoms/Button'`

### 3. TypeScript Member Ordering

Interface/Type members sẽ được sắp xếp alphabetically theo thứ tự:

1. Signature
2. Fields
3. Constructor
4. Methods

### 4. React Props Ordering

JSX props được sắp xếp theo:

1. Reserved props (key, ref)
2. Shorthand props first
3. Callbacks last
4. Multiline props last

## Prettier Configuration

File `.prettierrc`:

```json
{
  "semi": false,
  "singleQuote": true,
  "tabWidth": 2,
  "useTabs": false,
  "trailingComma": "all",
  "printWidth": 80,
  "bracketSpacing": true,
  "arrowParens": "always"
}
```

## Pre-commit Hooks

Husky + lint-staged đã được setup để tự động:

- Format code với Prettier
- Fix ESLint issues
- Chỉ áp dụng cho files đã staged

## Troubleshooting

### Format on save không hoạt động

1. Kiểm tra VSCode extensions đã cài đặt chưa
2. Reload VSCode window (`Cmd+Shift+P` → "Reload Window")
3. Kiểm tra `.vscode/settings.json` có đúng config không

### ESLint báo lỗi quá nhiều

```bash
# Auto-fix toàn bộ
yarn format:all

# Hoặc fix từng loại
yarn format    # Fix Prettier
yarn lint:fix  # Fix ESLint
```

### Conflict giữa ESLint và Prettier

Dự án đã cài `eslint-config-prettier` để tránh conflict. Nếu vẫn gặp vấn đề:

1. `yarn install` lại
2. Restart VSCode
