# Hướng Dẫn Sử Dụng Storybook

Tài liệu hướng dẫn cài đặt và sử dụng Storybook cho dự án **Employer Management System**.

---

## 📖 Mục Lục

- [Giới Thiệu](#giới-thiệu)
- [Cài Đặt](#cài-đặt)
- [Khởi Chạy Storybook](#khởi-chạy-storybook)
- [Cấu Trúc Thư Mục](#cấu-trúc-thư-mục)
- [Cách Sử Dụng](#cách-sử-dụng)
- [Tạo Story Mới](#tạo-story-mới)
- [Best Practices](#best-practices)
- [Troubleshooting](#troubleshooting)

---

## Giới Thiệu

**Storybook** là công cụ phát triển UI component một cách độc lập, giúp:

- 🎨 **Xem trước components** mà không cần chạy toàn bộ ứng dụng
- 📚 **Tài liệu hóa** các components cho team
- 🧪 **Test các trường hợp** khác nhau (states, props, variants)
- 🎯 **Single source of truth** cho UI (thay thế Figma)

### Công Nghệ

- **Storybook**: v10.1.11
- **Framework**: Next.js (App Router)
- **UI Library**: Material-UI v7
- **Addons**:
  - `@storybook/addon-docs` - Auto-generate documentation
  - `@storybook/addon-a11y` - Accessibility testing
  - `@storybook/addon-links` - Navigate between stories
  - `@storybook/addon-onboarding` - Getting started guide

---

## Cài Đặt

### Yêu Cầu Tiên Quyết

- Node.js 20+
- npm hoặc yarn

### Bước 1: Clone Project

```bash
git clone <repository-url>
cd ts-employer-fe
```

### Bước 2: Cài Đặt Dependencies

```bash
npm install
# hoặc
yarn install
```

> **Lưu ý**: Storybook và các addons đã được cài sẵn trong `devDependencies`, bạn không cần cài thêm gì.

### Bước 3: Kiểm Tra Cấu Hình

Storybook đã được cấu hình sẵn tại `.storybook/`:

- [`.storybook/main.ts`](file:///.storybook/main.ts) - Cấu hình chính
- [`.storybook/preview.tsx`](file:///.storybook/preview.tsx) - Theme và global decorators
- [`.storybook/Introduction.mdx`](file:///.storybook/Introduction.mdx) - Trang giới thiệu
- [`.storybook/fonts.css`](file:///.storybook/fonts.css) - Font imports

---

## Khởi Chạy Storybook

### Development Mode

```bash
npm run storybook
# hoặc
yarn storybook
```

Storybook sẽ chạy tại: **http://localhost:6006**

### Build Static Storybook

Để build static files cho deploy:

```bash
npm run build-storybook
# hoặc
yarn build-storybook
```

Output sẽ được tạo tại `storybook-static/`

---

## Cấu Trúc Thư Mục

```
.storybook/
├── main.ts              # Cấu hình chính
├── preview.tsx          # Theme, decorators
├── Introduction.mdx     # Trang chủ Storybook
└── fonts.css           # Font imports

src/components/
├── atoms/              # Components cơ bản
│   ├── Button/
│   │   └── CommonButton.stories.tsx
│   ├── Input/
│   │   └── TextField.stories.tsx
│   ├── Select/
│   │   └── stories/
│   │       ├── Select.stories.tsx
│   │       └── SelectUser.stories.tsx
│   ├── Checkbox/
│   │   └── Checkbox.stories.tsx
│   └── ...
├── molecules/          # Components phức hợp
│   └── DatePicker/
│       └── CommonDatePicker/
│           └── CommonDatePicker.stories.tsx
└── storybook/         # Theme showcases
    ├── Colors.stories.tsx
    ├── Typography.stories.tsx
    └── ThemeOverrides.stories.tsx
```

---

## Cách Sử Dụng

### 1. Khởi Động Storybook

```bash
npm run storybook
```

### 2. Browse Components

- Mở **http://localhost:6006**
- Click vào sidebar để xem các components:
  - **Introduction** - Hướng dẫn tổng quan
  - **Storybook/** - Theme showcases (Colors, Typography, Theme Overrides)
  - **Atoms/** - Button, Input, Select, Checkbox, Radio, TextArea
  - **Molecules/** - DatePicker, Dialog, Upload, FormField

### 3. Interactive Controls

Mỗi story có **Controls panel** ở dưới, cho phép bạn:

- Thay đổi props real-time
- Test các states (disabled, error, loading)
- Copy code examples

![Storybook Controls](https://storybook.js.org/images/addon-controls.gif)

### 4. Xem Documentation

Click tab **Docs** để xem:

- Auto-generated prop tables
- Component usage examples
- Implementation notes

---

## Tạo Story Mới

### Bước 1: Tạo File Story

Tạo file `ComponentName.stories.tsx` cùng folder với component:

```tsx
// src/components/atoms/Button/MyButton.stories.tsx
import type { Meta, StoryObj } from '@storybook/react'
import { MyButton } from './MyButton'

const meta: Meta<typeof MyButton> = {
  title: 'Atoms/Button/MyButton',
  component: MyButton,
  tags: ['autodocs'],
  argTypes: {
    variant: {
      control: 'select',
      options: ['contained', 'outlined', 'text'],
    },
    color: {
      control: 'select',
      options: ['primary', 'secondary', 'error'],
    },
    disabled: {
      control: 'boolean',
    },
  },
}

export default meta
type Story = StoryObj<typeof MyButton>

// Default story
export const Default: Story = {
  args: {
    children: 'Click Me',
    variant: 'contained',
    color: 'primary',
  },
}

// Variants
export const Outlined: Story = {
  args: {
    children: 'Outlined Button',
    variant: 'outlined',
  },
}

export const Disabled: Story = {
  args: {
    children: 'Disabled',
    disabled: true,
  },
}
```

### Bước 2: Tổ Chức Stories

#### Naming Convention

```
title: '[Category]/[Subcategory]/[ComponentName]'

Examples:
- 'Atoms/Button/CommonButton'
- 'Atoms/Input/TextField'
- 'Molecules/DatePicker/CommonDatePicker'
- 'Storybook/Colors'
```

#### File Structure

```
src/components/
└── atoms/
    └── Button/
        ├── CommonButton.tsx          # Component
        ├── CommonButton.stories.tsx  # Stories
        ├── ButtonStyles.tsx          # Styles (if any)
        └── index.tsx                 # Export
```

### Bước 3: Story với React Hook Form

Để showcase components dùng `react-hook-form`:

```tsx
import { useForm } from 'react-hook-form'
import { TextField } from './TextField'

export const WithValidation: Story = {
  render: () => {
    const { control } = useForm({
      defaultValues: {
        email: '',
      },
    })

    return (
      <form>
        <TextField
          name="email"
          control={control}
          label="Email"
          rules={{
            required: 'Email is required',
            pattern: {
              value: /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i,
              message: 'Invalid email address',
            },
          }}
        />
      </form>
    )
  },
}
```

### Bước 4: Add Documentation

Sử dụng JSDoc hoặc MDX:

````tsx
/**
 * CommonButton is a customized MUI Button component.
 *
 * ## Usage
 * ```tsx
 * import { CommonButton } from '@/components/atoms'
 *
 * <CommonButton variant="contained" color="primary">
 *   Click Me
 * </CommonButton>
 * ```
 */
export const CommonButton = (props: CommonButtonProps) => {
  // ...
}
````

---

## Best Practices

### 1. Story Naming

```tsx
// ✅ Good - Descriptive names
export const Default: Story = { ... }
export const WithError: Story = { ... }
export const Disabled: Story = { ... }
export const Loading: Story = { ... }

// ❌ Avoid - Generic names
export const Story1: Story = { ... }
export const Example: Story = { ... }
```

### 2. Use ArgTypes for Controls

```tsx
const meta: Meta<typeof Select> = {
  argTypes: {
    options: {
      control: 'object',
      description: 'Array of select options',
    },
    disabled: {
      control: 'boolean',
      description: 'Disable the select',
    },
    multiple: {
      control: 'boolean',
      description: 'Allow multiple selection',
    },
  },
}
```

### 3. Showcase All States

Mỗi component nên có stories cho:

- ✅ Default state
- ✅ Different variants
- ✅ Error state
- ✅ Disabled state
- ✅ Loading state (if applicable)
- ✅ With validation (for form components)

### 4. Organize by Atomic Design

```
Storybook/
├── Introduction
├── Storybook/               # Theme showcases
│   ├── Colors
│   ├── Typography
│   └── Theme Overrides
├── Atoms/                   # Basic components
│   ├── Button/
│   ├── Input/
│   ├── Select/
│   └── ...
├── Molecules/               # Composite components
│   ├── DatePicker/
│   └── ...
└── Organisms/               # Complex components
    ├── Table/
    └── ...
```

### 5. Keep Stories Simple

```tsx
// ✅ Good - One concern per story
export const Default: Story = {
  args: { label: 'Username' },
}

export const WithError: Story = {
  args: {
    label: 'Username',
    error: true,
    helperText: 'Username is required',
  },
}

// ❌ Avoid - Too many variants in one story
export const AllVariants: Story = {
  render: () => (
    <>
      <TextField label="Default" />
      <TextField label="Error" error />
      <TextField label="Disabled" disabled />
    </>
  ),
}
```

---

## Troubleshooting

### Storybook không khởi động

```bash
# Xóa cache và reinstall
rm -rf node_modules package-lock.json
npm install
npm run storybook
```

### Stories không hiển thị

Kiểm tra pattern trong `.storybook/main.ts`:

```ts
stories: ['../src/**/*.mdx', '../src/**/*.stories.@(js|jsx|mjs|ts|tsx)']
```

Đảm bảo file story của bạn có đuôi `.stories.tsx`

### Theme không apply

Kiểm tra `.storybook/preview.tsx` đã wrap với `ThemeProvider`:

```tsx
decorators: [
  (Story) => (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Story />
    </ThemeProvider>
  ),
]
```

### Fonts không load

Kiểm tra:

1. `.storybook/fonts.css` có import đúng fonts
2. `preview.tsx` có import fonts.css
3. Public fonts folder đúng path

### Port 6006 đã được sử dụng

Chạy với port khác:

```bash
npm run storybook -- -p 6007
```

---

## Tài Nguyên Bổ Sung

### Official Documentation

- [Storybook for Next.js](https://storybook.js.org/docs/get-started/nextjs)
- [Storybook Addons](https://storybook.js.org/addons)
- [Writing Stories](https://storybook.js.org/docs/writing-stories)

### Project Documentation

- [Getting Started Guide](./GETTING_STARTED.md)
- [Architecture Guide](./ARCHITECTURE_GUIDE.md)
- [Form Integration Guide](./FORM_INTEGRATION_GUIDE.md)

### Component Documentation

Trong Storybook UI:

- **Introduction** - Component architecture overview
- **Storybook/Colors** - Color palette showcase
- **Storybook/Typography** - Typography system
- **Storybook/Theme Overrides** - MUI overrides

---

## FAQ

### Khi nào nên tạo story?

- ✅ Khi tạo component mới (atoms, molecules)
- ✅ Khi update component có thay đổi UI đáng kể
- ✅ Khi muốn showcase theme overrides
- ⚠️ Không bắt buộc cho organisms phức tạp (forms, tables)

### Làm sao để share Storybook với team?

**Option 1: Local development**

```bash
npm run storybook
# Share http://localhost:6006 trên local network
```

**Option 2: Build và deploy**

```bash
npm run build-storybook
# Deploy folder storybook-static/ lên hosting (Vercel, Netlify, etc.)
```

### Storybook có thay thế Figma không?

Trong project này: **CÓ** ✅

- Chúng ta không có Figma designs
- Storybook là single source of truth cho UI
- Designer/PM review UI qua Storybook

---

## 🎯 Quick Commands

```bash
# Start Storybook
npm run storybook

# Build static Storybook
npm run build-storybook

# Run on different port
npm run storybook -- -p 6007

# Clear cache and start
rm -rf node_modules/.cache/storybook
npm run storybook
```

---

**Happy Documenting! 📚✨**

Nếu có câu hỏi, liên hệ team lead hoặc tạo issue trong repository.
