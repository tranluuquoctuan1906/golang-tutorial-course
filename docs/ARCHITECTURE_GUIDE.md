# Architecture Guide - Feature-Based Organization

> **Complete guide for new developers - Schemas, Types, API Services, Hooks**

## 📋 Table of Contents

- [Quick Start](#quick-start)
- [Project Structure](#project-structure)
- [Creating New Feature](#creating-new-feature)
- [Import Rules](#import-rules)
- [Best Practices](#best-practices)
- [Common Patterns](#common-patterns)
- [Troubleshooting](#troubleshooting)

---

## 🚀 Quick Start

### 5-Minute Setup

1. **Understand the 4 layers:**
   - **Types** → Data structure definitions
   - **Schemas** → Validation rules (Zod)
   - **API Services** → Backend communication
   - **Hooks** → React state management

2. **Use absolute imports:**
   ```typescript
   import { Candidate } from '@/types'
   import { candidateApi } from '@/services/api'
   import { createCandidateSchema } from '@/schemas'
   import { useMutation } from '@/hooks'
   ```

3. **Copy templates for new features:**
   ```bash
   cp -r src/types/api/features/_TEMPLATE src/types/api/features/candidate
   cp -r src/schemas/features/_TEMPLATE src/schemas/features/candidate
   cp -r src/services/api/features/_TEMPLATE src/services/api/features/candidate
   ```

---

## 📁 Project Structure

```
src/
├── types/api/                       # TypeScript Types
│   ├── common/                     # Shared types (PaginationState, SortState)
│   └── features/                   # Feature types
│       ├── requisition/
│       ├── candidate/
│       └── _TEMPLATE/             # ⭐ Copy này để tạo feature mới
│
├── schemas/                         # Validation Schemas (Zod)
│   ├── common/                     # 50+ reusable validators
│   │   ├── fields.schema.ts       # emailSchema, phoneSchema, etc.
│   │   ├── patterns.schema.ts     # Regex patterns
│   │   └── validators.schema.ts   # Custom validators
│   └── features/                   # Feature schemas
│       ├── requisition/
│       └── _TEMPLATE/             # ⭐ Copy này để tạo feature mới
│
├── services/api/                    # API Services
│   ├── common/
│   │   └── base.api.ts            # createBaseApi (generic CRUD)
│   └── features/                   # Feature API services
│       ├── requisition/
│       └── _TEMPLATE/             # ⭐ Copy này để tạo feature mới
│
└── hooks/                           # React Hooks
    ├── common/                     # Core hooks
    │   ├── useMutation.ts         # Generic mutation (CREATE/UPDATE/DELETE)
    │   ├── useQuery.ts            # Generic query (FETCH)
    │   └── useToast.ts            # Toast notifications
    └── features/                   # Feature hooks (optional)
        └── _TEMPLATE/             # Copy nếu cần custom hooks
```

---

## 🎯 Creating New Feature

### Example: "Candidate" Feature

#### Step 1: Types (Data Structure)

```bash
cp -r src/types/api/features/_TEMPLATE src/types/api/features/candidate
cd src/types/api/features/candidate
mv [feature].types.ts candidate.types.ts
```

**candidate.types.ts:**
```typescript
import { PaginationState, SortState } from '@/types/common'

// Base entity
export interface Candidate {
  id: string
  name: string
  email: string
  phone: string
  status: 'active' | 'archived'
  createdAt: string
  updatedAt: string
}

// Request types
export interface CreateCandidateRequest {
  name: string
  email: string
  phone: string
}

export interface UpdateCandidateRequest {
  name?: string
  email?: string
  phone?: string
  status?: 'active' | 'archived'
}

// Filter types
export interface CandidateFilters extends PaginationState, SortState {
  search?: string
  status?: 'active' | 'archived'
}

// Response types
export interface CandidateListResponse {
  items: Candidate[]
  total: number
  page: number
  limit: number
  totalPages: number
  hasMore: boolean
}
```

**Export in index.ts:**
```typescript
export * from '@/types/api/features/candidate/candidate.types'
```

**Export in main index:**
```typescript
// src/types/api/index.ts
export * from '@/types/api/features/candidate'
```

#### Step 2: Schemas (Validation)

```bash
cp -r src/schemas/features/_TEMPLATE src/schemas/features/candidate
cd src/schemas/features/candidate
mv [feature].schema.ts candidate.schema.ts
```

**candidate.schema.ts:**
```typescript
import { z } from 'zod'
import { emailSchema, phoneSchema, requiredString } from '@/schemas/common'

export const createCandidateSchema = z.object({
  name: requiredString,
  email: emailSchema,
  phone: phoneSchema,
})

export const updateCandidateSchema = z.object({
  name: z.string().optional(),
  email: emailSchema.optional(),
  phone: phoneSchema.optional(),
  status: z.enum(['active', 'archived']).optional(),
})

export type CreateCandidateInput = z.infer<typeof createCandidateSchema>
export type UpdateCandidateInput = z.infer<typeof updateCandidateSchema>
```

**Export in main index:**
```typescript
// src/schemas/index.ts
export * from '@/schemas/features/candidate'
```

#### Step 3: API Service (Backend Communication)

```bash
cp -r src/services/api/features/_TEMPLATE src/services/api/features/candidate
cd src/services/api/features/candidate
mv [feature].api.ts candidate.api.ts
```

**candidate.api.ts:**
```typescript
import { createBaseApi } from '@/services/api/common'
import { apiPost } from '@/services/fetcher'
import {
  Candidate,
  CreateCandidateRequest,
  UpdateCandidateRequest,
  CandidateFilters,
  CandidateListResponse,
} from '@/types'

const BASE_URL = '/api/candidates'

// Use createBaseApi for CRUD operations
const base = createBaseApi<Candidate, CreateCandidateRequest, UpdateCandidateRequest>({
  baseUrl: BASE_URL,
})

export const candidateApi = {
  ...base,  // Gets: create, update, delete, getById, getList

  // Add custom methods if needed
  archive: async (id: string) => {
    const response = await apiPost(`${BASE_URL}/${id}/archive`)
    return response.json()
  },
}
```

**Export in main index:**
```typescript
// src/services/api/index.ts
export * from '@/services/api/features/candidate'
```

#### Step 4: Use in Component

```typescript
'use client'

import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { createCandidateSchema } from '@/schemas'
import { useMutation } from '@/hooks'
import { candidateApi } from '@/services/api'
import { Candidate, CreateCandidateRequest } from '@/types'

export function CreateCandidateForm() {
  // 1. Form with validation
  const form = useForm<CreateCandidateRequest>({
    resolver: zodResolver(createCandidateSchema),
  })

  // 2. Mutation hook
  const { trigger: createCandidate, isMutating } = useMutation<
    Candidate,
    CreateCandidateRequest
  >('/api/candidates', candidateApi.create, {
    successMessage: 'Candidate created!',
    onSuccess: (data) => {
      console.log('Created:', data)
      router.push(`/candidates/${data.id}`)
    },
  })

  // 3. Submit handler
  const onSubmit = form.handleSubmit((data) => {
    createCandidate(data)
  })

  return (
    <form onSubmit={onSubmit}>
      <input {...form.register('name')} />
      {form.formState.errors.name && <span>Name is required</span>}

      <input {...form.register('email')} />
      {form.formState.errors.email && <span>Invalid email</span>}

      <input {...form.register('phone')} />
      {form.formState.errors.phone && <span>Invalid phone</span>}

      <button type="submit" disabled={isMutating}>
        {isMutating ? 'Creating...' : 'Create Candidate'}
      </button>
    </form>
  )
}
```

**Done! 🎉**

---

## 📏 Import Rules

### ✅ DO - Use Absolute Imports

```typescript
// Types
import { Candidate, CreateCandidateRequest } from '@/types'
import { PaginationState, SortState } from '@/types/common'

// Schemas
import { createCandidateSchema } from '@/schemas'
import { emailSchema, phoneSchema } from '@/schemas/common'

// API Services
import { candidateApi } from '@/services/api'
import { createBaseApi } from '@/services/api/common'

// Hooks
import { useMutation, useQuery } from '@/hooks'
```

### ❌ DON'T - Relative Imports (in feature folders)

```typescript
// ❌ NOT ALLOWED - ESLint will error
import { Candidate } from '../types'
import { candidateApi } from '../../services/api'
import { createCandidateSchema } from '../../../schemas'
```

**Why?**
- ✅ Easier to read
- ✅ Refactor-safe (move files without changing imports)
- ✅ Consistent across codebase
- ✅ Better IDE autocomplete

---

## 🎨 Best Practices

### 1. Reuse Common Schemas

```typescript
// ✅ GOOD - Reuse common validators
import { emailSchema, phoneSchema, requiredString } from '@/schemas/common'

export const createCandidateSchema = z.object({
  name: requiredString,
  email: emailSchema,
  phone: phoneSchema,
})

// ❌ BAD - Duplicate validation logic
export const createCandidateSchema = z.object({
  name: z.string().min(1),
  email: z.string().email(),
  phone: z.string().regex(/^[0-9+\-\s()]*$/),
})
```

### 2. Use Generic Hooks Directly

```typescript
// ✅ GOOD - Use generic hooks (recommended)
const { trigger: create } = useMutation('/api/candidates', candidateApi.create)

// ⚠️ Only create feature hooks if you need shared defaults
export function useCreateCandidate() {
  return useMutation('/api/candidates', candidateApi.create, {
    successMessage: 'Candidate created!',
    onSuccess: (data) => trackCandidateCreation(data), // Shared logic
  })
}
```

### 3. Use createBaseApi for CRUD

```typescript
// ✅ GOOD - Reuse base API
const base = createBaseApi<Candidate, CreateRequest, UpdateRequest>({
  baseUrl: '/api/candidates'
})

export const candidateApi = {
  ...base,  // Gets CRUD for free
  // Add custom methods only
}

// ❌ BAD - Reimplementing CRUD
export const candidateApi = {
  create: async (data) => { ... },
  update: async (id, data) => { ... },
  delete: async (id) => { ... },
  // Duplicating code!
}
```

### 4. Type Safety

```typescript
// ✅ Ensure types match schemas
export const createCandidateSchema = z.object({
  name: requiredString,
  email: emailSchema,
})

export type CreateCandidateInput = z.infer<typeof createCandidateSchema>

// CreateCandidateInput should match CreateCandidateRequest
export interface CreateCandidateRequest {
  name: string
  email: string
}
```

---

## 💡 Common Patterns

### CRUD Operations

```typescript
// CREATE
const { trigger: create, isMutating } = useMutation(
  '/api/candidates',
  candidateApi.create,
  { successMessage: 'Created!' }
)
await create({ name: 'John', email: 'john@example.com' })

// UPDATE
const { trigger: update } = useMutation(
  `/api/candidates/${id}`,
  (data) => candidateApi.update(id, data),
  { successMessage: 'Updated!' }
)
await update({ name: 'John Doe' })

// DELETE
const { trigger: remove } = useMutation(
  `/api/candidates/${id}`,
  () => candidateApi.delete(id),
  { successMessage: 'Deleted!' }
)
await remove(undefined)

// FETCH SINGLE
const { data: candidate, isLoading } = useQuery(
  `/api/candidates/${id}`,
  () => candidateApi.getById(id)
)

// FETCH LIST
const { data: candidates } = useQuery(
  '/api/candidates',
  () => candidateApi.getList({ status: 'active' })
)
```

### Custom API Methods

```typescript
export const candidateApi = {
  ...base,

  // Archive candidate
  archive: async (id: string) => {
    const response = await apiPost(`/api/candidates/${id}/archive`)
    return response.json()
  },

  // Bulk delete
  bulkDelete: async (ids: string[]) => {
    const response = await apiPost('/api/candidates/bulk-delete', { ids })
    return response.json()
  },

  // Search
  search: async (query: string) => {
    const response = await apiGet(`/api/candidates/search?q=${query}`)
    return response.json()
  },
}
```

### Form Validation

```typescript
const form = useForm<CreateCandidateRequest>({
  resolver: zodResolver(createCandidateSchema),
  defaultValues: {
    name: '',
    email: '',
    phone: '',
  },
})

const onSubmit = form.handleSubmit((data) => {
  createCandidate(data)
})
```

---

## 🆘 Troubleshooting

### ESLint Error: Relative imports not allowed

```
Error: Relative imports are not allowed in feature-based architecture...
  > import { Button } from '../components/Button'
```

**Fix:**
```typescript
// Change
import { Button } from '../components/Button'

// To
import { Button } from '@/components/Button'
```

### TypeScript Error: Cannot find module

```
Error: Cannot find module '@/types/candidate'
```

**Fix:** Import from main index
```typescript
// ❌ Wrong
import { Candidate } from '@/types/candidate'

// ✅ Correct
import { Candidate } from '@/types'
```

### Type Mismatch: Schema vs Interface

```typescript
// Make sure they match!
export const createCandidateSchema = z.object({
  name: z.string(),
  email: z.string().email(),
})

export type CreateCandidateInput = z.infer<typeof createCandidateSchema>

export interface CreateCandidateRequest {
  name: string  // Must match schema
  email: string // Must match schema
}
```

---

## 📚 Available Common Schemas

### Field Validators

```typescript
import {
  requiredString,        // Non-empty string
  emailSchema,          // Valid email
  phoneSchema,          // Phone number
  urlSchema,            // Valid URL
  numberInRange,        // Number with min/max
  positiveNumber,       // > 0
  nonNegativeNumber,    // >= 0
  imageFileSchema,      // Image file (jpg, png, etc.)
  videoFileSchema,      // Video file
  dateSchema,           // Valid date
  futureDate,           // Date in future
  pastDate,             // Date in past
  // ... 40+ more validators
} from '@/schemas/common'
```

### Usage Example

```typescript
export const createJobSchema = z.object({
  title: requiredString,
  description: requiredString,
  salary: numberInRange(1000, 1000000),
  email: emailSchema,
  website: urlSchema.optional(),
  startDate: futureDate,
  logo: imageFileSchema.optional(),
})
```

---

## 🔗 Data Flow

```
User Input
    ↓
React Hook Form + Zod Schema
    ↓
useMutation Hook
    ↓
API Service (candidateApi)
    ↓
Backend API
    ↓
Response (Candidate type)
    ↓
Update UI + Success Toast
```

---

## ✅ Checklist: New Feature

- [ ] Copy `_TEMPLATE` folders for types, schemas, API
- [ ] Rename files to match feature name
- [ ] Define types (Entity, Request, Response, Filters)
- [ ] Create validation schemas (reuse common validators)
- [ ] Create API service (use `createBaseApi`)
- [ ] Export everything in main index files
- [ ] Use `useMutation` and `useQuery` in component
- [ ] Test: validation, API calls, loading states, toasts

---

## 🎓 Summary

**4 Layers:**
1. **Types** → `@/types` - Data structures
2. **Schemas** → `@/schemas` - Validation rules
3. **API** → `@/services/api` - Backend calls
4. **Hooks** → `@/hooks` - React state

**Key Rules:**
- ✅ Always use absolute imports (`@/`)
- ✅ Copy templates for new features
- ✅ Reuse common utilities
- ✅ Use `createBaseApi` for CRUD
- ✅ Use generic hooks (`useMutation`, `useQuery`)

**Code Generation:**

Use the automated code generator instead of manual templates:

```bash
# Generate complete flows with TODO guides
npm run generate:form product      # CREATE form flow
npm run generate:table employee    # TABLE/LIST flow
```

See [Code Generation Guide](./CODE_GENERATION.md) for details.

---

**Welcome to the team! Happy coding! 🚀**
