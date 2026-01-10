# Table Component - Hướng Dẫn Tích Hợp

> Hệ thống Table component hoàn chỉnh với Sorting, Filtering, Pagination và SWR Integration

---

## 📋 Mục Lục

1. [Tổng Quan](#-tổng-quan)
2. [Yêu Cầu](#-yêu-cầu)
3. [Cấu Trúc Thư Mục](#-cấu-trúc-thư-mục)
4. [Cài Đặt](#-cài-đặt)
5. [TypeScript Types](#-typescript-types)
6. [Hướng Dẫn Sử Dụng](#-hướng-dẫn-sử-dụng)
7. [API Reference](#-api-reference)
8. [Ví Dụ Chi Tiết](#-ví-dụ-chi-tiết)
9. [Best Practices](#-best-practices)

---

## 🎯 Tổng Quan

Table component là một hệ thống hoàn chỉnh cung cấp:

- ✅ **Server-side Pagination** - Phân trang từ server
- ✅ **Multi-column Sorting** - Sort nhiều cột cùng lúc
- ✅ **Advanced Filtering** - Filter linh hoạt với URL sync
- ✅ **SWR Integration** - Tích hợp SWR cho data fetching
- ✅ **Global Refresh** - Refresh tất cả tables cùng lúc
- ✅ **TypeScript Support** - Type-safe hoàn toàn
- ✅ **URL State Management** - State được sync với URL (bookmarkable)
- ✅ **Responsive Design** - Tương thích mobile
- ✅ **Loading & Error States** - Xử lý loading và error tốt
- ✅ **Sticky Header/Columns** - Header và columns có thể fixed

---

## 📦 Yêu Cầu

```json
{
  "dependencies": {
    "@mui/material": "^5.x.x",
    "@mui/icons-material": "^5.x.x",
    "next": "^14.x.x",
    "react": "^18.x.x",
    "swr": "^2.x.x",
    "axios": "^1.x.x",
    "lodash": "^4.x.x",
    "dayjs": "^1.x.x"
  }
}
```

---

## 📁 Cấu Trúc Thư Mục

Copy toàn bộ thư mục này vào dự án mới:

```
src/components/shared/Table/
├── components/
│   ├── Table.tsx              # Main table component
│   ├── TableHeader.tsx        # Header with sorting
│   ├── TableBody.tsx          # Body with data rendering
│   ├── TablePagination.tsx    # Pagination controls
│   ├── TableLoading.tsx       # Loading overlay
│   ├── ErrorPage.tsx          # Error display
│   ├── Filter.tsx             # Filter container
│   └── HeaderContent.tsx      # Page header with toolbar
├── hooks/
│   ├── useTableFilter.ts      # Filter management
│   ├── useTableSort.ts        # Sorting logic
│   ├── useTableWithSWR.ts     # SWR integration
│   └── useRefresh.ts          # Global refresh
├── store/
│   └── loader-inventory.ts    # Global refresh registry
├── utils/
│   └── helper.ts              # URL and params conversion
├── types.ts                   # TypeScript definitions
└── index.ts                   # Barrel export
```

---

## 🔧 Cài Đặt

### Bước 1: Copy Files

Copy toàn bộ thư mục `src/components/shared/Table/` vào dự án mới của bạn.

### Bước 2: Setup Axios Instance

Tạo file `src/libs/api/axios.ts`:

```typescript
import axios from 'axios'

export const axiosAdminTool = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL || 'http://localhost:3000/api',
  headers: {
    'Content-Type': 'application/json',
  },
})

// Add interceptors if needed
axiosAdminTool.interceptors.request.use((config) => {
  // Add auth token, etc.
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

axiosAdminTool.interceptors.response.use(
  (response) => response,
  (error) => {
    // Handle errors globally
    return Promise.reject(error)
  }
)
```

### Bước 3: Setup Loader Inventory

Tạo file `src/components/shared/Table/store/loader-inventory.ts`:

```typescript
type RefreshFunction = () => void

const loaderInventory: RefreshFunction[] = []

export const add = (refreshFn: RefreshFunction) => {
  loaderInventory.push(refreshFn)
}

export const remove = (refreshFn: RefreshFunction) => {
  const index = loaderInventory.indexOf(refreshFn)
  if (index > -1) {
    loaderInventory.splice(index, 1)
  }
}

export const refreshAll = () => {
  loaderInventory.forEach((fn) => fn())
}
```

---

## 📝 TypeScript Types

### Core Types

```typescript
// Data pagination structure
export interface DataPagination {
  currentPage: number
  perPage: number
  totalItems: number
  totalPages: number
}

// API response with pagination
export interface Pagination<R> {
  data: R[]
  pagination: DataPagination
}

// Column definition
export interface ColumnsProps {
  field: string                           // Field name in data object
  title: string | React.ReactElement      // Column header
  canSort?: boolean                       // Enable sorting
  fixed?: 'left' | 'right'                // Sticky column
  align?: 'left' | 'center' | 'right'     // Text alignment
  width?: number                          // Column width
}

// Sort values
export type SortValue = 'asc' | 'desc' | 'none'

// Fetch parameters
export interface FetchProps<F> {
  page: number
  size: number
  filter: F
  sort?: Record<string, SortValue>
}

// Loader pattern
export interface Loader<R, F> {
  fetch: (input: FetchProps<F>) => Promise<Pagination<R>>
  cancel?: () => void
  isLoading?: boolean
}

// Table props
export interface TableProps<R = ResultType, F = Record<string, unknown>> {
  loader: Loader<R, F>
  columns: ColumnsProps[]
  Wrapper?: React.FC<{ children: React.ReactNode }>
  render?: (data: R, column: ColumnsProps) => React.JSX.Element | React.ReactElement | null
  stickyHeader?: boolean
  maxHeight?: string | number
}

// Filter props
export interface FilterProps {
  lable?: string
  gridClassName?: string
  colClassName?: string
  gridSx?: SxProps<Theme>
  FilterComponents: React.ReactElement[]
  actions?: React.ReactElement
}
```

---

## 🚀 Hướng Dẫn Sử Dụng

### 1. Setup API Configuration

Tạo file config cho API endpoint:

```typescript
// src/libs/api/user.ts
import { SWRConfiguration } from 'swr'
import { Pagination } from '@/components/shared/Table'
import { UserResponse } from '@/types/user'

export const UserTableConfig = {
  url: '/api/users',
  options: {
    swrOptions: {
      revalidateOnFocus: false,
      revalidateOnReconnect: true,
    } as SWRConfiguration<Pagination<UserResponse>>,
  },
}
```

### 2. Define Data Types

```typescript
// src/types/user.ts
export interface UserResponse {
  id: number
  name: string
  userName: string
  email: string
  role: 'Admin' | 'Sale' | 'User'
  status: 'active' | 'inactive'
  createdAt: string
}

export interface UserFilter {
  name?: string
  role?: string
  status?: string
}
```

### 3. Define Columns

```typescript
// src/components/features/user-management/list/user.data.ts
import { ColumnsProps } from '@/components/shared/Table'

export const columns: ColumnsProps[] = [
  {
    field: 'id',
    title: 'No.',
    width: 20,
    align: 'left'
  },
  {
    field: 'name',
    canSort: true,
    title: 'User Name',
    width: 150
  },
  {
    field: 'role',
    title: 'Role',
    width: 80,
    align: 'center'
  },
  {
    field: 'status',
    title: 'Status',
    width: 80,
    align: 'center'
  },
  {
    field: 'createdAt',
    canSort: true,
    title: 'Created Date',
    width: 120
  },
  {
    field: 'actions',
    title: 'Actions',
    width: 100,
    align: 'center'
  },
]
```

### 4. Create Table Component

```typescript
'use client'

import AddIcon from '@mui/icons-material/Add'
import { Avatar, Box, Button, Chip, Typography } from '@mui/material'
import dayjs from 'dayjs'
import get from 'lodash/get'
import Link from 'next/link'
import React, { useMemo } from 'react'

import {
  createSWRLoader,
  HeaderContent,
  Table,
  ColumnsProps,
  useTableWithSWR,
} from '@/components/shared/Table'
import { UserTableConfig } from '@/libs/api/user'
import { UserFilter, UserResponse } from '@/types/user'
import { columns } from './user.data'

// Custom render function for each cell
const renderColumn = (data: UserResponse, column: ColumnsProps) => {
  switch (column.field) {
    case 'name':
      return (
        <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
          <Avatar>{data.name.charAt(0)}</Avatar>
          <Box>
            <Typography variant="body2" sx={{ fontWeight: 'bold' }}>
              {data.name}
            </Typography>
            <Typography variant="body2" sx={{ color: 'error.main' }}>
              {data.userName}
            </Typography>
            <Typography variant="inherit" sx={{ color: 'info.primary' }}>
              {data.email}
            </Typography>
          </Box>
        </Box>
      )

    case 'status':
      return (
        <Chip
          label={data.status}
          color={data.status === 'active' ? 'success' : 'error'}
          size="small"
          variant="outlined"
          sx={{
            fontWeight: 'bold',
            minWidth: 80,
            textTransform: 'capitalize',
            py: 1.5,
            px: 1,
          }}
        />
      )

    case 'role':
      const roleColor: 'error' | 'primary' | 'secondary' | 'default' =
        data.role === 'Admin'
          ? 'error'
          : data.role === 'Sale'
            ? 'primary'
            : 'default'
      return (
        <Chip
          label={data.role}
          color={roleColor}
          size="small"
          sx={{
            textTransform: 'uppercase',
            fontWeight: 'medium',
            minWidth: 100,
            py: 1.5,
            px: 1,
          }}
        />
      )

    case 'createdAt':
      return (
        <Typography variant="body2">
          {dayjs(data.createdAt).format('DD/MM/YYYY')}
        </Typography>
      )

    case 'actions':
      return (
        <Box sx={{ display: 'flex', gap: 1, justifyContent: 'center' }}>
          <Button size="small" variant="outlined">Edit</Button>
          <Button size="small" variant="outlined" color="error">Delete</Button>
        </Box>
      )

    default:
      // Default rendering using lodash get for nested properties
      const value = get(data, column.field)
      if (value === undefined || value === null) {
        return null
      }
      return <Typography variant="body2">{value}</Typography>
  }
}

export default function UserList() {
  // 1. Setup SWR data fetching
  const data = useTableWithSWR<UserResponse, UserFilter>(
    UserTableConfig.url,
    UserTableConfig.options
  )

  // 2. Create loader for Table component
  const loader = useMemo(() => createSWRLoader(data), [data])

  // 3. Create toolbar with action buttons
  const Toolbar = () => (
    <Link href="/users/create">
      <Button
        variant="outlined"
        color="primary"
        sx={{ mt: 2 }}
        startIcon={<AddIcon />}
      >
        Create User
      </Button>
    </Link>
  )

  return (
    <>
      <HeaderContent title="User Management" toolbar={<Toolbar />} />

      {/* Filter component (optional) */}
      {/* <UserFilter /> */}

      {/* Table component */}
      <Table
        loader={loader}
        columns={columns}
        render={renderColumn}
        stickyHeader={true}
        maxHeight="calc(100vh - 300px)"
      />
    </>
  )
}
```

### 5. Create Filter Component

```typescript
'use client'

import { Button, TextField } from '@mui/material'
import React from 'react'
import { Filter, useTableFilter, useFilter } from '@/components/shared/Table'

export default function UserFilter() {
  // Get individual filter values
  const [name, setName] = useTableFilter('name')
  const [role, setRole] = useTableFilter('role')
  const [status, setStatus] = useTableFilter('status')

  // Get bulk filter updater
  const updateFilters = useFilter()

  const handleClearAll = () => {
    updateFilters({
      name: undefined,
      role: undefined,
      status: undefined,
    })
  }

  const handleSearch = () => {
    // Filters are already applied via URL, this is optional
    console.log('Search triggered')
  }

  return (
    <Filter
      lable="Filter Users"
      FilterComponents={[
        <TextField
          key="name"
          label="Name"
          variant="outlined"
          size="small"
          value={name?.[0] || ''}
          onChange={(e) => setName(e.target.value ? [e.target.value] : undefined)}
          fullWidth
        />,
        <TextField
          key="role"
          label="Role"
          variant="outlined"
          size="small"
          select
          value={role?.[0] || ''}
          onChange={(e) => setRole(e.target.value ? [e.target.value] : undefined)}
          fullWidth
          SelectProps={{ native: true }}
        >
          <option value="">All Roles</option>
          <option value="Admin">Admin</option>
          <option value="Sale">Sale</option>
          <option value="User">User</option>
        </TextField>,
        <TextField
          key="status"
          label="Status"
          variant="outlined"
          size="small"
          select
          value={status?.[0] || ''}
          onChange={(e) => setStatus(e.target.value ? [e.target.value] : undefined)}
          fullWidth
          SelectProps={{ native: true }}
        >
          <option value="">All Status</option>
          <option value="active">Active</option>
          <option value="inactive">Inactive</option>
        </TextField>,
      ]}
      actions={
        <>
          <Button variant="outlined" onClick={handleClearAll}>
            Clear All
          </Button>
          <Button variant="contained" onClick={handleSearch}>
            Search
          </Button>
        </>
      }
    />
  )
}
```

---

## 📚 API Reference

### Components

#### `<Table />`

Main table component.

**Props:**
```typescript
{
  loader: Loader<R, F>           // Data loader
  columns: ColumnsProps[]        // Column definitions
  render?: (data, column) => JSX // Custom cell renderer
  Wrapper?: React.FC             // Wrapper component
  stickyHeader?: boolean         // Enable sticky header (default: true)
  maxHeight?: string | number    // Max height for scroll (default: calc(100vh - 300px))
}
```

#### `<Filter />`

Filter container component.

**Props:**
```typescript
{
  lable?: string                      // Filter label (default: 'Filter')
  FilterComponents: React.ReactElement[] // Array of filter components
  actions?: React.ReactElement        // Action buttons (Clear, Search, etc.)
  gridSx?: SxProps<Theme>             // Custom grid styles
  gridClassName?: string              // Grid CSS class
  colClassName?: string               // Column CSS class
}
```

#### `<HeaderContent />`

Page header with title and toolbar.

**Props:**
```typescript
{
  title: string | React.ReactElement  // Page title
  toolbar?: React.ReactNode           // Toolbar (action buttons)
}
```

### Hooks

#### `useTableWithSWR()`

Hook for SWR data fetching with automatic URL param handling.

```typescript
const data = useTableWithSWR<UserResponse, UserFilter>(
  apiEndpoint: string,
  options?: {
    fetcher?: (url: string) => Promise<Pagination<R>>
    swrOptions?: SWRConfiguration
  }
)

// Returns:
{
  data: Pagination<R> | undefined
  error: Error | undefined
  isLoading: boolean
  mutate: () => void
  queryParams: FetchProps<F>
}
```

#### `createSWRLoader()`

Creates a loader compatible with Table component from SWR data.

```typescript
const loader = createSWRLoader(swrData)
```

#### `useTableFilter()`

Manages a single filter field.

```typescript
const [value, setValue] = useTableFilter('fieldName')

// value: string[] | undefined
// setValue: (newValue: string[] | undefined) => void

// Example:
const [name, setName] = useTableFilter('name')
setName(['John'])      // Set filter
setName(undefined)     // Clear filter
```

#### `useFilter()`

Updates multiple filter fields at once.

```typescript
const updateFilters = useFilter()

// Usage:
updateFilters({
  name: ['John'],
  role: ['Admin'],
  status: undefined,  // Clear this filter
})
```

#### `useFilterParams()`

Gets all current filter params.

```typescript
const filters = useFilterParams()
// Returns: Record<string, string[]>
```

#### `useSort()`

Manages sorting for a specific field.

```typescript
const [sortValue, setSortValue] = useSort('', 'fieldName')

// sortValue: 'asc' | 'desc' | 'none'
// setSortValue: (newValue: SortValue) => void
```

#### `useRefresh()`

Triggers global refresh for all tables.

```typescript
const refresh = useRefresh()

// Usage:
refresh()  // Refreshes all tables using useTableWithSWR
```

### Utility Functions

#### `objectToSearchParams()`

Converts object to URL search params string.

```typescript
const str = objectToSearchParams({ page: 1, name: 'John' })
// Returns: "page=1&name=John"
```

#### `searchParamsToObject()`

Parses URL search params to object.

```typescript
const obj = searchParamsToObject(searchParams)
// Returns: { page: '1', name: 'John' }
```

#### `parsedSort()`

Parses sort string from URL to object.

```typescript
const sortObj = parsedSort('', 'name+asc,age+desc')
// Returns: { name: 'asc', age: 'desc' }
```

---

## 💡 Ví Dụ Chi Tiết

### Example 1: Basic Table

```typescript
'use client'

import { Table, useTableWithSWR, createSWRLoader, ColumnsProps } from '@/components/shared/Table'
import { useMemo } from 'react'

interface Product {
  id: number
  name: string
  price: number
}

const columns: ColumnsProps[] = [
  { field: 'id', title: 'ID', width: 50 },
  { field: 'name', title: 'Name', canSort: true },
  { field: 'price', title: 'Price', canSort: true, align: 'right' },
]

export default function ProductList() {
  const data = useTableWithSWR<Product>('/api/products')
  const loader = useMemo(() => createSWRLoader(data), [data])

  return <Table loader={loader} columns={columns} />
}
```

### Example 2: Custom Cell Rendering

```typescript
const renderColumn = (data: Product, column: ColumnsProps) => {
  switch (column.field) {
    case 'price':
      return (
        <Typography variant="body2" color="primary" fontWeight="bold">
          ${data.price.toFixed(2)}
        </Typography>
      )

    case 'name':
      return (
        <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
          <ShoppingCart />
          <Typography>{data.name}</Typography>
        </Box>
      )

    default:
      return <Typography>{get(data, column.field)}</Typography>
  }
}

// Use it:
<Table loader={loader} columns={columns} render={renderColumn} />
```

### Example 3: With Filters

```typescript
'use client'

import { Filter, useTableFilter, useFilter } from '@/components/shared/Table'
import { TextField, Button, MenuItem } from '@mui/material'

export default function ProductFilter() {
  const [category, setCategory] = useTableFilter('category')
  const [minPrice, setMinPrice] = useTableFilter('minPrice')
  const [maxPrice, setMaxPrice] = useTableFilter('maxPrice')
  const updateFilters = useFilter()

  const handleClearAll = () => {
    updateFilters({
      category: undefined,
      minPrice: undefined,
      maxPrice: undefined,
    })
  }

  return (
    <Filter
      FilterComponents={[
        <TextField
          select
          label="Category"
          value={category?.[0] || ''}
          onChange={(e) => setCategory(e.target.value ? [e.target.value] : undefined)}
          fullWidth
        >
          <MenuItem value="">All</MenuItem>
          <MenuItem value="electronics">Electronics</MenuItem>
          <MenuItem value="clothing">Clothing</MenuItem>
        </TextField>,

        <TextField
          type="number"
          label="Min Price"
          value={minPrice?.[0] || ''}
          onChange={(e) => setMinPrice(e.target.value ? [e.target.value] : undefined)}
          fullWidth
        />,

        <TextField
          type="number"
          label="Max Price"
          value={maxPrice?.[0] || ''}
          onChange={(e) => setMaxPrice(e.target.value ? [e.target.value] : undefined)}
          fullWidth
        />,
      ]}
      actions={
        <Button variant="outlined" onClick={handleClearAll}>
          Clear All
        </Button>
      }
    />
  )
}
```

### Example 4: Sticky Columns

```typescript
const columns: ColumnsProps[] = [
  { field: 'id', title: 'ID', fixed: 'left', width: 50 },
  { field: 'name', title: 'Name', width: 200 },
  { field: 'description', title: 'Description', width: 300 },
  { field: 'price', title: 'Price', width: 100 },
  { field: 'actions', title: 'Actions', fixed: 'right', width: 150 },
]

<Table
  loader={loader}
  columns={columns}
  stickyHeader={true}
  maxHeight={600}
/>
```

### Example 5: Custom Fetcher with Authentication

```typescript
import { axiosAdminTool } from '@/libs/api/axios'

const customFetcher = async (url: string) => {
  const token = localStorage.getItem('token')
  const res = await axiosAdminTool.get(url, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  })
  return res.data
}

const data = useTableWithSWR<User>('/api/users', {
  fetcher: customFetcher,
  swrOptions: {
    revalidateOnFocus: false,
    refreshInterval: 30000, // Refresh every 30s
  },
})
```

### Example 6: Manual Loader (Without SWR)

```typescript
import { Loader, FetchProps, Pagination } from '@/components/shared/Table'

const customLoader: Loader<User, UserFilter> = {
  fetch: async (params: FetchProps<UserFilter>): Promise<Pagination<User>> => {
    const response = await fetch('/api/users?' + new URLSearchParams({
      page: String(params.page),
      size: String(params.size),
      ...params.filter,
    }))

    const data = await response.json()

    return {
      data: data.items,
      pagination: {
        currentPage: data.page,
        perPage: data.size,
        totalItems: data.total,
        totalPages: Math.ceil(data.total / data.size),
      },
    }
  },
}

// Use it:
<Table loader={customLoader} columns={columns} />
```

### Example 7: Global Refresh

```typescript
'use client'

import { useRefresh } from '@/components/shared/Table'
import { Button } from '@mui/material'
import RefreshIcon from '@mui/icons-material/Refresh'

export default function RefreshButton() {
  const refresh = useRefresh()

  return (
    <Button
      variant="outlined"
      startIcon={<RefreshIcon />}
      onClick={refresh}
    >
      Refresh All Tables
    </Button>
  )
}
```

### Example 8: Nested Property Access

```typescript
interface User {
  id: number
  profile: {
    name: string
    avatar: string
  }
  company: {
    name: string
    address: {
      city: string
    }
  }
}

const columns: ColumnsProps[] = [
  { field: 'id', title: 'ID' },
  { field: 'profile.name', title: 'Name', canSort: true },
  { field: 'company.name', title: 'Company' },
  { field: 'company.address.city', title: 'City' },
]

// The table will automatically use lodash.get() to access nested properties
```

---

## ✨ Best Practices

### 1. Server Response Format

Backend API phải trả về format này:

```typescript
{
  "data": [
    { "id": 1, "name": "John" },
    { "id": 2, "name": "Jane" }
  ],
  "pagination": {
    "currentPage": 1,
    "perPage": 10,
    "totalItems": 100,
    "totalPages": 10
  }
}
```

### 2. API Query Parameters

Table tự động gửi các params này lên server:

```
GET /api/users?page=1&size=10&name=John&role=Admin&sort=name+asc,createdAt+desc
```

Server cần parse và xử lý:
- `page`: Trang hiện tại
- `size`: Số items per page
- `sort`: Sort format `field+direction,field2+direction`
- Các params khác: Filter fields

### 3. TypeScript Best Practices

```typescript
// ✅ GOOD: Define specific types
interface UserResponse {
  id: number
  name: string
  email: string
}

interface UserFilter {
  name?: string
  email?: string
}

const data = useTableWithSWR<UserResponse, UserFilter>('/api/users')

// ❌ BAD: Using any or unknown
const data = useTableWithSWR<any>('/api/users')
```

### 4. Memoization

```typescript
// ✅ GOOD: Memoize loader to prevent re-renders
const loader = useMemo(() => createSWRLoader(data), [data])

// ❌ BAD: Creating new loader on every render
const loader = createSWRLoader(data)
```

### 5. Filter State Management

```typescript
// ✅ GOOD: Use URL-based filters (shareable, bookmarkable)
const [name, setName] = useTableFilter('name')

// ❌ BAD: Local state (lost on refresh)
const [name, setName] = useState('')
```

### 6. Custom Rendering

```typescript
// ✅ GOOD: Use switch for multiple fields
const renderColumn = (data, column) => {
  switch (column.field) {
    case 'status':
      return <StatusChip status={data.status} />
    case 'actions':
      return <ActionButtons data={data} />
    default:
      return <Typography>{get(data, column.field)}</Typography>
  }
}

// ❌ BAD: Multiple if-else
const renderColumn = (data, column) => {
  if (column.field === 'status') return <StatusChip />
  if (column.field === 'actions') return <ActionButtons />
  // ...
}
```

### 7. Error Handling

```typescript
const data = useTableWithSWR<User>('/api/users', {
  swrOptions: {
    onError: (error) => {
      console.error('Failed to fetch users:', error)
      // Show toast notification
    },
    revalidateOnFocus: false,
    shouldRetryOnError: true,
    errorRetryCount: 3,
  },
})
```

### 8. Performance Optimization

```typescript
// ✅ GOOD: Disable unnecessary revalidation
const data = useTableWithSWR<User>('/api/users', {
  swrOptions: {
    revalidateOnFocus: false,      // Don't refetch on window focus
    revalidateOnReconnect: false,  // Don't refetch on reconnect
    dedupingInterval: 5000,        // Dedupe requests within 5s
  },
})
```

### 9. Responsive Design

```typescript
// Hide columns on mobile
const columns: ColumnsProps[] = [
  { field: 'id', title: 'ID', width: 50 },
  { field: 'name', title: 'Name' },
  // Only show on desktop
  { field: 'email', title: 'Email', /* add custom logic in render */ },
  { field: 'actions', title: 'Actions', fixed: 'right' },
]

const renderColumn = (data, column) => {
  // Hide email on mobile
  if (column.field === 'email') {
    return (
      <Box sx={{ display: { xs: 'none', md: 'block' } }}>
        {data.email}
      </Box>
    )
  }
  // ...
}
```

### 10. Loading States

```typescript
// Table automatically handles loading state from SWR
// But you can also add custom loading indicators

const data = useTableWithSWR<User>('/api/users')

if (data.isLoading && !data.data) {
  return <CustomLoadingSkeleton />
}

return <Table loader={loader} columns={columns} />
```

---

## 🔍 Troubleshooting

### Issue 1: Table không hiển thị data

**Nguyên nhân:** Backend response không đúng format

**Giải pháp:**
```typescript
// Backend phải trả về:
{
  "data": [...],
  "pagination": {
    "currentPage": 1,
    "perPage": 10,
    "totalItems": 100,
    "totalPages": 10
  }
}
```

### Issue 2: Filter không hoạt động

**Nguyên nhân:** Backend không nhận query params

**Giải pháp:** Kiểm tra backend có parse query params đúng không
```typescript
// URL: /api/users?name=John&role=Admin
// Backend phải đọc: req.query.name, req.query.role
```

### Issue 3: Sort không hoạt động

**Nguyên nhân:** Backend không parse sort param

**Giải pháp:**
```typescript
// URL: /api/users?sort=name+asc,age+desc
// Backend parse thành:
// sort = [{ field: 'name', direction: 'asc' }, { field: 'age', direction: 'desc' }]
```

### Issue 4: Pagination không chính xác

**Nguyên nhân:** Backend tính toán totalPages sai

**Giải pháp:**
```typescript
totalPages = Math.ceil(totalItems / perPage)
```

### Issue 5: Global refresh không hoạt động

**Nguyên nhân:** Chưa setup loader-inventory

**Giải pháp:** Đảm bảo đã tạo file `store/loader-inventory.ts` và import đúng

---

## 📖 Advanced Topics

### Custom Pagination Component

Nếu muốn tùy chỉnh pagination:

```typescript
import { TablePagination } from '@/components/shared/Table'

// Override styles
<TablePagination
  data={data}
  sx={{
    '& .MuiSelect-select': {
      color: 'primary.main',
    },
  }}
/>
```

### Server-Side Sorting Example (Backend)

```typescript
// Express.js example
app.get('/api/users', (req, res) => {
  const { page = 1, size = 10, sort } = req.query

  // Parse sort: "name+asc,age+desc"
  const sortFields = sort ? sort.split(',').map(s => {
    const [field, direction] = s.split('+')
    return { field, direction }
  }) : []

  // Build SQL query
  let query = 'SELECT * FROM users WHERE 1=1'

  // Add filters
  if (req.query.name) {
    query += ` AND name LIKE '%${req.query.name}%'`
  }

  // Add sorting
  if (sortFields.length > 0) {
    query += ' ORDER BY '
    query += sortFields.map(s => `${s.field} ${s.direction}`).join(', ')
  }

  // Add pagination
  const offset = (page - 1) * size
  query += ` LIMIT ${size} OFFSET ${offset}`

  // Execute query and return response
  const users = await db.query(query)
  const total = await db.query('SELECT COUNT(*) FROM users')

  res.json({
    data: users,
    pagination: {
      currentPage: parseInt(page),
      perPage: parseInt(size),
      totalItems: total,
      totalPages: Math.ceil(total / size),
    },
  })
})
```

---

## 📞 Support

Nếu gặp vấn đề, check lại:

1. ✅ Backend response đúng format
2. ✅ TypeScript types match với data
3. ✅ Axios instance được config đúng
4. ✅ URL params được backend xử lý đúng
5. ✅ Dependencies được cài đầy đủ

---

## 📝 Changelog

### Version 1.0.0
- Initial release
- Support sorting, filtering, pagination
- SWR integration
- Global refresh system
- TypeScript support
- URL state management

---

## 📄 License

MIT License - Free to use in your projects

---

**Happy Coding! 🚀**
