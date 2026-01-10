# Form Component - Hướng Dẫn Tích Hợp

> Hệ thống Form component hoàn chỉnh với DynamicForm, Validation (Zod), và React Hook Form

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
9. [Validation với Zod](#-validation-với-zod)
10. [Best Practices](#-best-practices)

---

## 🎯 Tổng Quan

Form component là một hệ thống hoàn chỉnh cung cấp:

- ✅ **DynamicForm** - Tạo form từ config array (declarative)
- ✅ **10+ Field Types** - Text, Email, Password, Textarea, Select, Autocomplete, Checkbox, Radio, Date, File
- ✅ **Zod Validation** - Type-safe validation với error messages
- ✅ **React Hook Form Integration** - Form state management hiệu quả
- ✅ **Async Options Loading** - Load options từ API cho select/autocomplete
- ✅ **File Upload** - Drag-drop, preview, size validation
- ✅ **Responsive Layout** - Grid system linh hoạt
- ✅ **Loading States** - Submit button với loading indicator
- ✅ **TypeScript Support** - Fully typed
- ✅ **MUI Integration** - Material-UI components

---

## 📦 Yêu Cầu

```json
{
  "dependencies": {
    "@mui/material": "^5.x.x",
    "@mui/icons-material": "^5.x.x",
    "@mui/x-date-pickers": "^6.x.x",
    "react": "^18.x.x",
    "react-hook-form": "^7.x.x",
    "@hookform/resolvers": "^3.x.x",
    "zod": "^3.x.x",
    "dayjs": "^1.x.x",
    "react-hot-toast": "^2.x.x"
  }
}
```

---

## 📁 Cấu Trúc Thư Mục

Copy toàn bộ thư mục này vào dự án mới:

```
src/components/shared/form/
├── DynamicForm.tsx           # Main dynamic form component
├── FormTextField.tsx         # Text input field
├── FormTextareaField.tsx     # Textarea field
├── FormSelectField.tsx       # Select/Autocomplete field
├── FormCheckboxField.tsx     # Checkbox field
├── FormRadioField.tsx        # Radio button field
├── FormFileField.tsx         # File upload field
├── types.ts                  # TypeScript type definitions
└── index.ts                  # Barrel export

src/libs/zod/
└── helper.ts                 # Zod validation helpers

src/components/shared/style-share/
└── index.tsx                 # Styled components for forms
```

---

## 🔧 Cài Đặt

### Bước 1: Copy Files

Copy toàn bộ thư mục `src/components/shared/form/` vào dự án mới.

### Bước 2: Setup Dependencies

```bash
npm install @mui/material @mui/icons-material @mui/x-date-pickers
npm install react-hook-form @hookform/resolvers zod
npm install dayjs react-hot-toast
```

### Bước 3: Setup Styled Components

Tạo file `src/components/shared/style-share/index.tsx`:

```typescript
import { FormControl, Typography, styled } from '@mui/material'

export const FormControlStyled = styled(FormControl, {
  shouldForwardProp: (prop) => prop !== 'labelWidth' && prop !== 'hiddenLabel',
})<{ labelWidth?: number; hiddenLabel?: boolean }>(
  ({ labelWidth, hiddenLabel = false }) => ({
    display: 'grid',
    gridTemplateColumns: hiddenLabel ? '1fr' : `${labelWidth ?? 180}px 1fr`,
    alignItems: 'start',
    paddingBottom: '20px',
  })
)

export const DescriptionText = styled(Typography)`
  color: #b0b0b0;
  font-size: 0.75rem;
  margin-top: 4px;
`
```

### Bước 4: Setup Error Helper

Tạo file `src/libs/utils/index.tsx`:

```typescript
import { FieldError, FieldErrors, FieldErrorsImpl, FieldValues, Merge } from 'react-hook-form'

export function getErrorName(
  obj: FieldErrors<FieldValues>,
  path: string
): FieldError | Merge<FieldError, FieldErrorsImpl<any>> | undefined {
  return path.split('.').reduce((o: any, k) => (o ? o[k] : undefined), obj) as
    | FieldError
    | Merge<FieldError, FieldErrorsImpl<any>>
    | undefined
}
```

### Bước 5: Setup Zod Validation Helpers (Optional)

Tạo file `src/libs/zod/helper.ts`:

```typescript
import dayjs, { Dayjs } from 'dayjs'
import isSameOrAfter from 'dayjs/plugin/isSameOrAfter'
import isSameOrBefore from 'dayjs/plugin/isSameOrBefore'
import { z, ZodType } from 'zod'

dayjs.extend(isSameOrAfter)
dayjs.extend(isSameOrBefore)

const NORMAL_INPUT_MAX_LENGTH = 255
const DATE_PICKER_MIN_DATE = '1900-01-01'
const DATE_PICKER_MAX_DATE = '2100-12-31'

const VALIDATION_ERROR_MESSAGES = {
  required: 'This field is required',
  dateInvalid: 'Invalid date',
  numberMaxLength: (max: number) => `Number cannot exceed ${max} digits`,
  invalidMinDateTime: (min: string) => `Date must be after ${min}`,
}

// Checks if a given Dayjs object represents a valid date
export function isValidDate(date?: Dayjs | null): boolean {
  if (!date) return false
  if (!dayjs.isDayjs(date) || !date.isValid()) return false

  const min = dayjs(DATE_PICKER_MIN_DATE)
  const max = dayjs(DATE_PICKER_MAX_DATE)

  return date.isSameOrAfter(min) && date.isSameOrBefore(max)
}

// Creates a Zod schema to validate text input with optional length constraints
export function zodValidateInputText(options?: {
  required?: boolean
  length?: number
  enableTrim?: boolean
  isNumberValidate?: boolean
}): z.ZodString {
  const length = options?.length || NORMAL_INPUT_MAX_LENGTH
  const required = options?.required ?? false
  const enableTrim = options?.enableTrim ?? true
  const isNumberValidate = options?.isNumberValidate ?? false

  let baseValidation = z
    .string({
      required_error: VALIDATION_ERROR_MESSAGES.required,
      invalid_type_error: VALIDATION_ERROR_MESSAGES.required,
    })
    .max(length, {
      message: isNumberValidate
        ? VALIDATION_ERROR_MESSAGES.numberMaxLength(length)
        : `The text must not exceed ${length} characters.`,
    })

  if (enableTrim) {
    baseValidation = baseValidation.trim()
  }

  return required
    ? baseValidation.min(1, { message: VALIDATION_ERROR_MESSAGES.required })
    : baseValidation
}

// Validates Dayjs date with constraints
export function zodValidateDayjs(options?: {
  required?: boolean
  notAllowSelectFutureDate?: boolean
  notAllowSelectPastDate?: boolean
}) {
  const required = Boolean(options?.required)
  const notAllowSelectFutureDate = Boolean(options?.notAllowSelectFutureDate)
  const notAllowSelectPastDate = Boolean(options?.notAllowSelectPastDate)

  const maxDate = notAllowSelectFutureDate
    ? dayjs()
    : dayjs(DATE_PICKER_MAX_DATE)
  const minDate = notAllowSelectPastDate ? dayjs() : dayjs(DATE_PICKER_MIN_DATE)

  const requiredDateZod = z
    .custom<Dayjs>()
    .refine((value) => !!value, VALIDATION_ERROR_MESSAGES.required)
    .refine(
      (value) =>
        dayjs.isDayjs(value) &&
        value.isValid() &&
        !value.isBefore(minDate, 'date') &&
        !value.isAfter(maxDate, 'date'),
      VALIDATION_ERROR_MESSAGES.dateInvalid
    )

  const nullishDateZod = z.custom<Dayjs>().refine((value) => {
    if (!value) return true

    return (
      dayjs.isDayjs(value) &&
      value.isValid() &&
      !value.isBefore(minDate, 'date') &&
      !value.isAfter(maxDate, 'date')
    )
  }, VALIDATION_ERROR_MESSAGES.dateInvalid)

  return required ? requiredDateZod : nullishDateZod
}

// Validates DateTime with min constraint
export function zodValidateDateTime(options: {
  required?: boolean
  minDateTime?: Dayjs
}) {
  const { required = false, minDateTime } = options

  let schema: ZodType<Dayjs | null> = z
    .any()
    .refine(
      (value) => {
        if (!required) return true
        return !!value
      },
      {
        message: VALIDATION_ERROR_MESSAGES.required,
      }
    )
    .refine(
      (value) => {
        if (!value) return true
        return isValidDate(value)
      },
      {
        message: VALIDATION_ERROR_MESSAGES.dateInvalid,
      }
    )

  if (minDateTime) {
    const min = minDateTime.startOf('minute')
    schema = schema.refine(
      (val) => {
        if (!val || !isValidDate(val)) return true
        return val.isSameOrAfter(min, 'minute')
      },
      {
        message: VALIDATION_ERROR_MESSAGES.invalidMinDateTime(
          min.toISOString()
        ),
      }
    )
  }

  return schema
}
```

---

## 📝 TypeScript Types

### Core Types

```typescript
import { Dayjs } from 'dayjs'

// Base field config
interface BaseFieldConfig {
  name: string
  label?: string
  required?: boolean
  disabled?: boolean
  placeholder?: string
  xs?: number // Grid column size on mobile (default: 12)
  sm?: number // Grid column size on desktop (default: 6)
}

// Text field
interface TextFieldConfig extends BaseFieldConfig {
  type: 'text' | 'email' | 'password'
  maxLength?: number
}

// Textarea field
interface TextareaFieldConfig extends BaseFieldConfig {
  type: 'textarea'
  rows?: number
  maxLength?: number
}

// Select option type
export type SelectOption = {
  value: string | number
  label: string
}

// Select variants
interface SelectFieldConfig extends BaseFieldConfig {
  type: 'select'
  options?: SelectOption[]
  loadOptions?: () => Promise<SelectOption[]>
  multiple?: boolean
  autocomplete?: boolean
}

// Checkbox field
interface CheckboxFieldConfig extends BaseFieldConfig {
  type: 'checkbox'
}

// Radio field
export type RadioOption = {
  value: string | number
  label: string
}

interface RadioFieldConfig extends BaseFieldConfig {
  type: 'radio'
  options: RadioOption[]
  row?: boolean
}

// Date fields
interface DateFieldConfig extends BaseFieldConfig {
  type: 'date' | 'datetime' | 'daterange'
  minDate?: Dayjs
  maxDate?: Dayjs
}

// File upload field
interface FileFieldConfig extends BaseFieldConfig {
  type: 'file'
  multiple?: boolean
  accept?: string // e.g., "image/*", ".pdf,.doc", "image/png,image/jpeg"
  maxSize?: number // in bytes
  maxFiles?: number // maximum number of files (only for multiple)
  showPreview?: boolean // show image preview (default: true)
}

// Union type for all field configs
export type FormFieldConfig =
  | TextFieldConfig
  | TextareaFieldConfig
  | SelectFieldConfig
  | CheckboxFieldConfig
  | RadioFieldConfig
  | DateFieldConfig
  | FileFieldConfig

// DynamicForm props
export interface DynamicFormProps<T extends Record<string, unknown>> {
  fields: FormFieldConfig[]
  columns?: number // Number of columns per row (default: 2)
  onSubmit: (data: T) => void | Promise<void>
  defaultValues?: Partial<T>
  schema?: unknown // Zod schema
  submitButtonText?: string
  showSubmitButton?: boolean
  clearButtonText?: string
  showClearButton?: boolean
  cancelButtonText?: string
  showCancelButton?: boolean
  onCancel?: () => void
  isSubmitting?: boolean // Loading state for submit button
}
```

---

## 🚀 Hướng Dẫn Sử Dụng

### 1. Basic Form với DynamicForm

#### Step 1: Define Schema và Form Data

```typescript
// AddUser.data.ts
import { z } from 'zod'
import { FormFieldConfig } from '@/components/shared/form'

// Zod validation schema
export const userSchema = z.object({
  username: z.string().min(1, 'Username is required'),
  fullname: z.string().min(1, 'Fullname is required'),
  email: z.string().email('Invalid email address'),
  password: z.string().min(6, 'Password must be at least 6 characters'),
  department: z.string().min(1, 'Department is required'),
  userType: z.string().min(1, 'User type is required'),
  saleGroup: z.string().min(1, 'Sale group is required'),
  status: z.enum(['active', 'inactive']),
  date: z.string().optional(),
})

// TypeScript type from Zod schema
export type UserFormData = z.infer<typeof userSchema>

// Default values
export const defaultValues: UserFormData = {
  username: '',
  fullname: '',
  email: '',
  password: '',
  department: '',
  userType: '',
  saleGroup: '',
  status: 'active',
  date: undefined,
}
```

#### Step 2: Define Field Configuration

```typescript
// AddUser.data.ts (continued)
import { SaleGroupApi } from '@/libs/api/sale-group'

// Async option loader
export const loadSaleGroups = async () => {
  try {
    const data = await SaleGroupApi.getListApi()
    return data
  } catch (error) {
    console.error('Failed to load sale groups:', error)
    return []
  }
}

// Field configuration array
export const formFields: FormFieldConfig[] = [
  { type: 'text', name: 'username', label: 'Username', required: true },
  { type: 'text', name: 'fullname', label: 'Fullname', required: true },
  { type: 'email', name: 'email', label: 'Email', required: true },
  { type: 'password', name: 'password', label: 'Password', required: true },
  { type: 'text', name: 'department', label: 'Department', required: true },
  { type: 'date', name: 'date', label: 'Date' },
  {
    type: 'select',
    name: 'userType',
    label: 'User type',
    required: true,
    options: [
      { value: '', label: 'All' },
      { value: 'admin', label: 'Admin' },
      { value: 'user', label: 'User' },
      { value: 'implementer', label: 'Implementer' },
    ],
  },
  {
    type: 'select',
    name: 'saleGroup',
    label: 'Sale group',
    required: true,
    autocomplete: true,
    loadOptions: loadSaleGroups,  // Async option loading
  },
  {
    type: 'file',
    name: 'avatar',
    label: 'Ảnh đại diện',
    accept: 'image/*',
    maxSize: 5 * 1024 * 1024,  // 5MB
    showPreview: true,
    sm: 3,  // Column width: 3/12 (25%)
  },
  {
    type: 'radio',
    name: 'status',
    label: 'Status',
    sm: 12,  // Full width
    options: [
      { value: 'active', label: 'Active' },
      { value: 'inactive', label: 'Inactive' },
    ],
  },
]
```

#### Step 3: Create Submit Handler

```typescript
// useAdd.ts
import { useRouter } from 'next/navigation'
import { useState } from 'react'
import toast from 'react-hot-toast'
import { UserFormData } from './AddUser.data'
import { UserApi } from '@/libs/api/user'

export default function useAdd() {
  const router = useRouter()
  const [isSubmitting, setIsSubmitting] = useState(false)

  const handleSubmit = async (data: UserFormData) => {
    setIsSubmitting(true)
    try {
      const payload = {
        userName: data.username,
        name: data.fullname,
        email: data.email,
        password: data.password,
        role: data.userType,
        status: data.status,
        saleGroup: data.saleGroup,
      }

      await UserApi.createApi(payload)
      toast.success('User created successfully!')
      router.push('/users')
    } catch (error) {
      toast.error('Failed to create user')
      console.error('Error creating user:', error)
    } finally {
      setIsSubmitting(false)
    }
  }

  return { isSubmitting, handleSubmit }
}
```

#### Step 4: Render DynamicForm

```typescript
// AddUser.tsx
'use client'

import { DynamicForm } from '@/components/shared/form'
import { formFields, userSchema, defaultValues } from './AddUser.data'
import useAdd from './useAdd'

export default function AddUser() {
  const { isSubmitting, handleSubmit } = useAdd()

  return (
    <DynamicForm
      fields={formFields}
      columns={2}
      schema={userSchema}
      onSubmit={handleSubmit}
      defaultValues={defaultValues}
      submitButtonText="Create User"
      showSubmitButton
      showClearButton
      cancelButtonText="Cancel"
      showCancelButton
      onCancel={() => router.back()}
      isSubmitting={isSubmitting}
    />
  )
}
```

---

## 📚 API Reference

### Components

#### `<DynamicForm />`

Main dynamic form component that generates form fields from configuration.

**Props:**
```typescript
{
  fields: FormFieldConfig[]           // Array of field configurations
  columns?: number                    // Grid columns (default: 2)
  onSubmit: (data: T) => void | Promise<void>
  defaultValues?: Partial<T>          // Initial form values
  schema?: unknown                    // Zod validation schema
  submitButtonText?: string           // Default: "Submit"
  showSubmitButton?: boolean          // Default: true
  clearButtonText?: string            // Default: "Clear"
  showClearButton?: boolean           // Default: false
  cancelButtonText?: string           // Default: "Cancel"
  showCancelButton?: boolean          // Default: false
  onCancel?: () => void               // Cancel handler
  isSubmitting?: boolean              // Loading state
}
```

#### `<FormTextField />`

Standalone text field component.

**Props:**
```typescript
{
  name: string
  label?: string
  required?: boolean
  disabled?: boolean
  maxLength?: number
  placeholder?: string
  labelWidth?: number
  description?: string
  labelDescription?: string
}
```

#### `<FormSelectField />`

Standalone select/autocomplete field.

**Props:**
```typescript
{
  name: string
  label?: string
  options?: SelectOption[]
  loadOptions?: () => Promise<SelectOption[]>
  required?: boolean
  disabled?: boolean
  multiple?: boolean
  autocomplete?: boolean
  placeholder?: string
}
```

#### `<FormFileField />`

Standalone file upload field with drag-drop.

**Props:**
```typescript
{
  name: string
  label?: string
  required?: boolean
  disabled?: boolean
  placeholder?: string
  multiple?: boolean
  accept?: string
  maxSize?: number
  maxFiles?: number
  showPreview?: boolean
}
```

---

## 💡 Ví Dụ Chi Tiết

### Example 1: Simple Contact Form

```typescript
import { z } from 'zod'
import { DynamicForm, FormFieldConfig } from '@/components/shared/form'

const contactSchema = z.object({
  name: z.string().min(1, 'Name is required'),
  email: z.string().email('Invalid email'),
  message: z.string().min(10, 'Message must be at least 10 characters'),
})

type ContactFormData = z.infer<typeof contactSchema>

const fields: FormFieldConfig[] = [
  { type: 'text', name: 'name', label: 'Name', required: true },
  { type: 'email', name: 'email', label: 'Email', required: true },
  {
    type: 'textarea',
    name: 'message',
    label: 'Message',
    required: true,
    rows: 5,
    sm: 12, // Full width
  },
]

export default function ContactForm() {
  const handleSubmit = async (data: ContactFormData) => {
    console.log('Form data:', data)
    // Submit to API
  }

  return (
    <DynamicForm
      fields={fields}
      columns={1}
      schema={contactSchema}
      onSubmit={handleSubmit}
      submitButtonText="Send Message"
    />
  )
}
```

### Example 2: Multi-Select with Async Options

```typescript
const fields: FormFieldConfig[] = [
  {
    type: 'select',
    name: 'countries',
    label: 'Select Countries',
    multiple: true,
    autocomplete: true,
    loadOptions: async () => {
      const response = await fetch('/api/countries')
      const data = await response.json()
      return data.map(country => ({
        value: country.code,
        label: country.name,
      }))
    },
    sm: 12,
  },
]
```

### Example 3: File Upload with Preview

```typescript
const fields: FormFieldConfig[] = [
  {
    type: 'file',
    name: 'documents',
    label: 'Upload Documents',
    multiple: true,
    accept: '.pdf,.doc,.docx',
    maxSize: 10 * 1024 * 1024, // 10MB
    maxFiles: 5,
    showPreview: false,
    sm: 12,
  },
  {
    type: 'file',
    name: 'avatar',
    label: 'Profile Picture',
    accept: 'image/*',
    maxSize: 2 * 1024 * 1024, // 2MB
    showPreview: true,
    sm: 6,
  },
]
```

### Example 4: Radio Buttons Layout

```typescript
const fields: FormFieldConfig[] = [
  {
    type: 'radio',
    name: 'gender',
    label: 'Gender',
    row: true, // Horizontal layout
    options: [
      { value: 'male', label: 'Male' },
      { value: 'female', label: 'Female' },
      { value: 'other', label: 'Other' },
    ],
  },
  {
    type: 'radio',
    name: 'subscription',
    label: 'Subscription Plan',
    row: false, // Vertical layout
    options: [
      { value: 'free', label: 'Free' },
      { value: 'basic', label: 'Basic - $9.99/mo' },
      { value: 'premium', label: 'Premium - $19.99/mo' },
    ],
  },
]
```

### Example 5: Date Range with Constraints

```typescript
import dayjs from 'dayjs'

const fields: FormFieldConfig[] = [
  {
    type: 'date',
    name: 'startDate',
    label: 'Start Date',
    required: true,
    minDate: dayjs(), // Cannot select past dates
    sm: 6,
  },
  {
    type: 'date',
    name: 'endDate',
    label: 'End Date',
    required: true,
    minDate: dayjs(), // Cannot select past dates
    sm: 6,
  },
]
```

### Example 6: Using Individual Form Components

```typescript
'use client'

import { FormProvider, useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { Button, Box } from '@mui/material'
import { FormTextField, FormSelectField } from '@/components/shared/form'
import { userSchema, UserFormData } from './schema'

export default function CustomForm() {
  const methods = useForm<UserFormData>({
    resolver: zodResolver(userSchema),
    defaultValues: {
      name: '',
      role: '',
    },
  })

  const onSubmit = (data: UserFormData) => {
    console.log(data)
  }

  return (
    <FormProvider {...methods}>
      <Box component="form" onSubmit={methods.handleSubmit(onSubmit)}>
        <FormTextField
          name="name"
          label="Full Name"
          required
          maxLength={100}
        />

        <FormSelectField
          name="role"
          label="Role"
          required
          options={[
            { value: 'admin', label: 'Admin' },
            { value: 'user', label: 'User' },
          ]}
        />

        <Button type="submit" variant="contained">
          Submit
        </Button>
      </Box>
    </FormProvider>
  )
}
```

### Example 7: Conditional Fields

```typescript
'use client'

import { useState } from 'react'
import { DynamicForm, FormFieldConfig } from '@/components/shared/form'

export default function ConditionalForm() {
  const [showAddress, setShowAddress] = useState(false)

  const baseFields: FormFieldConfig[] = [
    { type: 'text', name: 'name', label: 'Name', required: true },
    {
      type: 'checkbox',
      name: 'hasAddress',
      label: 'I want to add address',
    },
  ]

  const addressFields: FormFieldConfig[] = [
    { type: 'text', name: 'street', label: 'Street' },
    { type: 'text', name: 'city', label: 'City' },
    { type: 'text', name: 'zipcode', label: 'Zip Code' },
  ]

  const fields = showAddress ? [...baseFields, ...addressFields] : baseFields

  const handleSubmit = (data: any) => {
    console.log(data)
    // Check if hasAddress is checked
    if (data.hasAddress) {
      setShowAddress(true)
    }
  }

  return <DynamicForm fields={fields} onSubmit={handleSubmit} />
}
```

### Example 8: Custom Grid Layout

```typescript
const fields: FormFieldConfig[] = [
  // Full width header
  { type: 'text', name: 'title', label: 'Title', sm: 12 },

  // 2 columns
  { type: 'text', name: 'firstName', label: 'First Name', sm: 6 },
  { type: 'text', name: 'lastName', label: 'Last Name', sm: 6 },

  // 3 columns
  { type: 'text', name: 'city', label: 'City', sm: 4 },
  { type: 'text', name: 'state', label: 'State', sm: 4 },
  { type: 'text', name: 'zip', label: 'ZIP', sm: 4 },

  // Full width textarea
  { type: 'textarea', name: 'notes', label: 'Notes', rows: 5, sm: 12 },
]

// Use with columns prop
<DynamicForm fields={fields} columns={3} onSubmit={handleSubmit} />
```

---

## 🔒 Validation với Zod

### Basic Validation

```typescript
import { z } from 'zod'

const schema = z.object({
  // Required string
  name: z.string().min(1, 'Name is required'),

  // Email validation
  email: z.string().email('Invalid email address'),

  // Min/Max length
  password: z
    .string()
    .min(8, 'Password must be at least 8 characters')
    .max(50, 'Password cannot exceed 50 characters'),

  // Number validation
  age: z
    .number()
    .min(18, 'Must be at least 18 years old')
    .max(120, 'Invalid age'),

  // Optional field
  phone: z.string().optional(),

  // Enum validation
  status: z.enum(['active', 'inactive', 'pending']),

  // Boolean
  agreeToTerms: z.boolean().refine(val => val === true, {
    message: 'You must agree to the terms',
  }),
})
```

### Advanced Validation

```typescript
import { z } from 'zod'

const schema = z.object({
  // Custom regex
  username: z
    .string()
    .min(3)
    .regex(/^[a-zA-Z0-9_]+$/, 'Username can only contain letters, numbers, and underscores'),

  // URL validation
  website: z.string().url('Invalid URL'),

  // Date validation
  birthdate: z.string().refine(
    (val) => {
      const date = new Date(val)
      const age = (new Date().getTime() - date.getTime()) / (365.25 * 24 * 60 * 60 * 1000)
      return age >= 18
    },
    { message: 'Must be at least 18 years old' }
  ),

  // Conditional validation
  country: z.string(),
  state: z.string().optional(),
}).refine(
  (data) => {
    if (data.country === 'US') {
      return !!data.state
    }
    return true
  },
  {
    message: 'State is required for US addresses',
    path: ['state'],
  }
)

// Password confirmation
const passwordSchema = z
  .object({
    password: z.string().min(8),
    confirmPassword: z.string(),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "Passwords don't match",
    path: ['confirmPassword'],
  })
```

### Using Zod Helpers

```typescript
import { zodValidateInputText, zodValidateDayjs } from '@/libs/zod/helper'

const schema = z.object({
  // Text with max length
  name: zodValidateInputText({ required: true, length: 100 }),

  // Optional text
  nickname: zodValidateInputText({ required: false, length: 50 }),

  // Number validation
  code: zodValidateInputText({
    required: true,
    length: 6,
    isNumberValidate: true,
  }),

  // Date validation
  birthDate: zodValidateDayjs({
    required: true,
    notAllowSelectFutureDate: true, // Cannot select future dates
  }),

  // Date validation (no past dates)
  appointmentDate: zodValidateDayjs({
    required: true,
    notAllowSelectPastDate: true, // Cannot select past dates
  }),
})
```

---

## ✨ Best Practices

### 1. Form Organization

```typescript
// ✅ GOOD: Separate concerns
// schema.ts
export const userSchema = z.object({ ... })
export type UserFormData = z.infer<typeof userSchema>

// fields.ts
export const userFields: FormFieldConfig[] = [ ... ]

// hooks/useUserForm.ts
export function useUserForm() {
  const handleSubmit = async (data: UserFormData) => { ... }
  return { handleSubmit }
}

// UserForm.tsx
import { DynamicForm } from '@/components/shared/form'
import { userSchema, userFields } from './schema'
import { useUserForm } from './hooks/useUserForm'

export default function UserForm() {
  const { handleSubmit } = useUserForm()
  return <DynamicForm schema={userSchema} fields={userFields} onSubmit={handleSubmit} />
}

// ❌ BAD: Everything in one file
export default function UserForm() {
  const schema = z.object({ ... })
  const fields = [ ... ]
  const handleSubmit = () => { ... }
  return <DynamicForm ... />
}
```

### 2. Type Safety

```typescript
// ✅ GOOD: Use z.infer for type safety
const schema = z.object({
  name: z.string(),
  age: z.number(),
})

type FormData = z.infer<typeof schema>

const handleSubmit = (data: FormData) => {
  // TypeScript knows data.name is string and data.age is number
  console.log(data.name.toUpperCase())
  console.log(data.age + 1)
}

// ❌ BAD: Manual type definition (can get out of sync)
interface FormData {
  name: string
  age: number
}
```

### 3. Async Options Loading

```typescript
// ✅ GOOD: Error handling and loading state
export const loadCountries = async (): Promise<SelectOption[]> => {
  try {
    const response = await fetch('/api/countries')
    if (!response.ok) throw new Error('Failed to load countries')
    const data = await response.json()
    return data.map(country => ({
      value: country.code,
      label: country.name,
    }))
  } catch (error) {
    console.error('Error loading countries:', error)
    toast.error('Failed to load countries')
    return []
  }
}

// ❌ BAD: No error handling
export const loadCountries = async () => {
  const response = await fetch('/api/countries')
  const data = await response.json()
  return data
}
```

### 4. Default Values

```typescript
// ✅ GOOD: Provide default values matching schema
const defaultValues: UserFormData = {
  name: '',
  email: '',
  role: 'user',
  status: 'active',
  agreeToTerms: false,
}

<DynamicForm defaultValues={defaultValues} ... />

// ❌ BAD: Partial or missing defaults
const defaultValues = {
  name: '',
  // Missing other required fields
}
```

### 5. File Upload Validation

```typescript
// ✅ GOOD: Clear constraints
{
  type: 'file',
  name: 'avatar',
  label: 'Profile Picture',
  accept: 'image/png,image/jpeg,image/jpg',
  maxSize: 2 * 1024 * 1024, // 2MB
  showPreview: true,
  required: true,
}

// With Zod validation
const schema = z.object({
  avatar: z
    .instanceof(File)
    .refine(file => file.size <= 2 * 1024 * 1024, 'File size must be less than 2MB')
    .refine(
      file => ['image/png', 'image/jpeg', 'image/jpg'].includes(file.type),
      'Only PNG, JPEG, JPG files are allowed'
    ),
})
```

### 6. Grid Layout

```typescript
// ✅ GOOD: Responsive grid with clear structure
const fields: FormFieldConfig[] = [
  // Mobile: full width, Desktop: 6/12 (50%)
  { type: 'text', name: 'firstName', label: 'First Name', sm: 6 },
  { type: 'text', name: 'lastName', label: 'Last Name', sm: 6 },

  // Full width
  { type: 'email', name: 'email', label: 'Email', sm: 12 },

  // 3 columns on desktop
  { type: 'text', name: 'city', label: 'City', sm: 4 },
  { type: 'text', name: 'state', label: 'State', sm: 4 },
  { type: 'text', name: 'zip', label: 'ZIP', sm: 4 },
]
```

### 7. Submit Handler

```typescript
// ✅ GOOD: Proper error handling and loading state
const [isSubmitting, setIsSubmitting] = useState(false)

const handleSubmit = async (data: FormData) => {
  setIsSubmitting(true)
  try {
    await api.createUser(data)
    toast.success('User created successfully!')
    router.push('/users')
  } catch (error) {
    if (error instanceof ApiError) {
      toast.error(error.message)
    } else {
      toast.error('An unexpected error occurred')
    }
    console.error('Error:', error)
  } finally {
    setIsSubmitting(false)
  }
}

// ❌ BAD: No error handling
const handleSubmit = async (data: FormData) => {
  await api.createUser(data)
  router.push('/users')
}
```

### 8. Form Reset

```typescript
// ✅ GOOD: Provide clear button with reset
<DynamicForm
  fields={fields}
  onSubmit={handleSubmit}
  showClearButton
  clearButtonText="Reset Form"
  defaultValues={defaultValues}
/>

// Or manual reset
const methods = useForm({ defaultValues })

const handleReset = () => {
  methods.reset(defaultValues)
}
```

### 9. Accessibility

```typescript
// ✅ GOOD: Proper labels and required indicators
const fields: FormFieldConfig[] = [
  {
    type: 'text',
    name: 'email',
    label: 'Email Address', // Clear label
    required: true, // Shows asterisk
    placeholder: 'john@example.com', // Helpful placeholder
  },
]
```

### 10. Performance

```typescript
// ✅ GOOD: Memoize async loaders
import { useCallback } from 'react'

const loadOptions = useCallback(async () => {
  const data = await fetchOptions()
  return data
}, [])

// ✅ GOOD: Debounce for autocomplete
import { useMemo } from 'react'
import debounce from 'lodash/debounce'

const debouncedSearch = useMemo(
  () => debounce(async (query: string) => {
    const results = await api.search(query)
    setOptions(results)
  }, 300),
  []
)
```

---

## 🔍 Troubleshooting

### Issue 1: Validation errors not showing

**Nguyên nhân:** Thiếu `schema` prop hoặc field `name` không match với schema

**Giải pháp:**
```typescript
// Ensure schema is passed
<DynamicForm schema={userSchema} ... />

// Ensure field names match schema keys
const schema = z.object({
  email: z.string(), // Key is 'email'
})

const fields = [
  { type: 'email', name: 'email' } // Name must match
]
```

### Issue 2: Async options not loading

**Nguyên nhân:** `loadOptions` function không return `Promise<SelectOption[]>`

**Giải pháp:**
```typescript
// ✅ CORRECT
const loadOptions = async (): Promise<SelectOption[]> => {
  const data = await fetch('/api/options')
  return data.map(item => ({ value: item.id, label: item.name }))
}

// ❌ WRONG
const loadOptions = async () => {
  const data = await fetch('/api/options')
  return data // Wrong format
}
```

### Issue 3: File upload not working

**Nguyên nhân:** Backend không nhận multipart/form-data

**Giải pháp:**
```typescript
const handleSubmit = async (data: FormData) => {
  const formData = new FormData()

  // Add text fields
  formData.append('name', data.name)

  // Add file
  if (data.avatar instanceof File) {
    formData.append('avatar', data.avatar)
  }

  // Send with correct headers
  await fetch('/api/upload', {
    method: 'POST',
    body: formData,
    // Don't set Content-Type header - browser will set it automatically
  })
}
```

### Issue 4: Grid layout not responsive

**Nguyên nhân:** Không sử dụng `sm` prop

**Giải pháp:**
```typescript
// ✅ CORRECT
{ type: 'text', name: 'name', sm: 6 } // 50% on desktop

// ❌ WRONG - Will use default (50% with columns=2)
{ type: 'text', name: 'name' }
```

---

## 📖 Advanced Topics

### Custom Field Component

Nếu muốn tạo custom field type:

```typescript
// CustomColorPicker.tsx
import { Controller, useFormContext } from 'react-hook-form'

export function CustomColorPicker({ name, label }: { name: string, label: string }) {
  const { control } = useFormContext()

  return (
    <Controller
      name={name}
      control={control}
      render={({ field }) => (
        <div>
          <label>{label}</label>
          <input type="color" {...field} />
        </div>
      )}
    />
  )
}

// Usage
import { FormProvider, useForm } from 'react-hook-form'
import { CustomColorPicker } from './CustomColorPicker'

export default function MyForm() {
  const methods = useForm()

  return (
    <FormProvider {...methods}>
      <form>
        <CustomColorPicker name="favoriteColor" label="Favorite Color" />
      </form>
    </FormProvider>
  )
}
```

### Form with Multiple Steps

```typescript
// MultiStepForm.tsx
import { useState } from 'react'
import { DynamicForm } from '@/components/shared/form'

const step1Fields = [
  { type: 'text', name: 'name', label: 'Name', required: true },
  { type: 'email', name: 'email', label: 'Email', required: true },
]

const step2Fields = [
  { type: 'text', name: 'address', label: 'Address' },
  { type: 'text', name: 'city', label: 'City' },
]

export default function MultiStepForm() {
  const [step, setStep] = useState(1)
  const [formData, setFormData] = useState({})

  const handleStep1Submit = (data: any) => {
    setFormData({ ...formData, ...data })
    setStep(2)
  }

  const handleStep2Submit = (data: any) => {
    const finalData = { ...formData, ...data }
    console.log('Final data:', finalData)
    // Submit to API
  }

  return (
    <>
      {step === 1 && (
        <DynamicForm
          fields={step1Fields}
          onSubmit={handleStep1Submit}
          submitButtonText="Next"
        />
      )}
      {step === 2 && (
        <DynamicForm
          fields={step2Fields}
          onSubmit={handleStep2Submit}
          defaultValues={formData}
          submitButtonText="Finish"
          showCancelButton
          cancelButtonText="Back"
          onCancel={() => setStep(1)}
        />
      )}
    </>
  )
}
```

---

## 📞 Support

Nếu gặp vấn đề, check lại:

1. ✅ Đã cài đặt đầy đủ dependencies
2. ✅ Schema và field names match nhau
3. ✅ `loadOptions` return đúng format `SelectOption[]`
4. ✅ File upload: backend support multipart/form-data
5. ✅ Import đúng components từ `@/components/shared/form`

---

## 📄 License

MIT License - Free to use in your projects

---

**Happy Coding! 🚀**
