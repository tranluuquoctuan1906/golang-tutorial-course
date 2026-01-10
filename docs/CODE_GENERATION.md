# Code Generation Guide

Generate complete feature flows in seconds with **Mock API** and **sample fields** ready to test.

---

## Quick Commands

```bash
# Create form flow (CREATE operation)
npm run generate:form product

# Create table flow (LIST operation)
npm run generate:table employee

# Interactive mode
npm run generate:flow
```

---

## What Gets Generated?

### Form Flow (CREATE)

```bash
npm run generate:form product
```

**Creates:**

- ✅ **Types** - TypeScript definitions with sample fields (price, category, dates, checkbox)
- ✅ **API Service** - CRUD methods
- ✅ **API Routes** - Real Next.js Route Handlers at `/api/products`:
  - `GET /api/products` - List with filters
  - `POST /api/products` - Create new
  - `GET /api/products/[id]` - Get by ID
  - `PUT /api/products/[id]` - Update
  - `DELETE /api/products/[id]` - Delete
- ✅ **Schema** - Zod validation with sample field schemas
- ✅ **Hook** - useCreateProduct with form logic
- ✅ **Form Component** - Ready-to-use form with all field types:
  - TextField (name)
  - TextArea (description)
  - Number input (price)
  - Select single (status)
  - Select multi with checkboxes (category)
  - CommonDatePicker (startDate, endDate)
  - Checkbox (isActive)
  - CheckboxGroup (interests - multiple selection)
  - RadioGroup (notification preference)
- ✅ **Page** - Admin route at `/admin/products/create`

**🎯 Works immediately** - Real API endpoints with mock data!

### Table Flow (LIST)

```bash
npm run generate:table employee
```

**Creates:**

- ✅ **Types** - With filters and pagination
- ✅ **API Service** - With search and bulk operations
- ✅ **Schema** - With filter validation
- ✅ **Hook** - useEmployeeList with table logic
- ✅ **Table Component** - With sorting, pagination, selection
- ✅ **Page** - Admin route at `/admin/employees`

---

## The TODO System

Every generated file includes:

### 1. Customization Checklist

At the top of each file:

```typescript
/**
 * 📋 CUSTOMIZATION CHECKLIST:
 * ├─ [ ] Step 1: Update main entity fields (line ~35)
 * ├─ [ ] Step 2: Update CreateRequest fields (line ~55)
 * ├─ [ ] Step 3: Update UpdateRequest fields (line ~70)
 * └─ [ ] Step 4: Export in src/types/api/index.ts
 */
```

### 2. Step-by-Step TODOs

At each customization point:

```typescript
// 📝 TODO [STEP 1]: Add your entity fields here
// Examples:
// price: number           // Product price
// category: string        // Product category
// inStock: boolean        // Availability
```

### 3. Code Examples

Ready-to-use examples you can uncomment:

```typescript
// 📝 TODO [STEP 2]: Add your custom fields here

/*
Example: Price field
<TextField
  control={form.control}
  name="price"
  label="Price"
  type="number"
  required
/>
*/
```

---

## Real API Routes - Test Immediately

Generated code includes **working Next.js API Route Handlers** so you can test with real HTTP endpoints immediately.

### File Structure

```
src/app/api/products/
├── route.ts              # GET (list), POST (create)
└── [id]/
    └── route.ts          # GET (detail), PUT (update), DELETE
```

### Available Endpoints

After running `npm run generate:form product`, you get:

| Method   | Endpoint          | Description                                     |
| -------- | ----------------- | ----------------------------------------------- |
| `GET`    | `/api/products`   | List products with filters (?search=, ?status=) |
| `POST`   | `/api/products`   | Create new product                              |
| `GET`    | `/api/products/1` | Get product by ID                               |
| `PUT`    | `/api/products/1` | Update product                                  |
| `DELETE` | `/api/products/1` | Delete product                                  |

### Test with curl or Postman

```bash
# Get list
curl http://localhost:3000/api/products

# Get with filters
curl "http://localhost:3000/api/products?search=test&status=active"

# Create new
curl -X POST http://localhost:3000/api/products \
  -H "Content-Type: application/json" \
  -d '{"name":"New Product","price":100,"category":"electronics"}'

# Get by ID
curl http://localhost:3000/api/products/1

# Update
curl -X PUT http://localhost:3000/api/products/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated Product","price":200}'

# Delete
curl -X DELETE http://localhost:3000/api/products/1
```

---

## Mock Data in Route Handlers

Route handlers include an in-memory mock database that persists during development:

### How It Works

**In `product.api.ts`:**

```typescript
// ============================================
// 🔧 MOCK API (For Development)
// ============================================
const ENABLE_MOCK = true // Set to false when backend is ready

// Mock data generator
const generateMockProduct = (id: string): Product => {
  const categories = ['electronics', 'clothing', 'food', 'books']
  const today = new Date()

  return {
    id,
    name: `Product ${id}`,
    description: `Sample product description`,
    status: 'active',
    price: Math.floor(Math.random() * 1000) + 100,
    category: categories[Math.floor(Math.random() * categories.length)],
    startDate: today.toISOString().split('T')[0],
    endDate: /* 30 days later */,
    isActive: Math.random() > 0.5,
    createdAt: new Date().toISOString(),
    updatedAt: new Date().toISOString(),
  }
}

let mockDatabase: Product[] = [
  generateMockProduct('1'),
  generateMockProduct('2'),
  generateMockProduct('3'),
]

// Mock delay to simulate network
const mockDelay = () => new Promise(resolve => setTimeout(resolve, 500))
```

### Features

✅ **CREATE** - Adds to mock database
✅ **READ** - Returns from mock database
✅ **UPDATE** - Updates mock database
✅ **DELETE** - Removes from mock database
✅ **FILTER** - Searches mock data
✅ **Network delay** - Simulates real API (500ms)

### When to Disable

When your backend is ready:

```typescript
// Change this line in product.api.ts
const ENABLE_MOCK = false // Disable mock, use real API
```

---

## Sample Fields - Ready to Use

Generated forms include working examples of all field types:

### Text Fields

- **Name** (required text)
- **Description** (textarea)

### Number Fields

- **Price** (number input with validation)

### Select/Dropdown

- **Status** (single select: draft, active, inactive)
- **Category** (multi-select with checkboxes: electronics, clothing, food, books)

### CheckboxGroup & RadioGroup

- **Interests** (multiple checkbox selection: technology, design, marketing, sales, finance)
- **Notification Preference** (radio selection: email, SMS, both, none)

### Date Fields

- **Start Date** (date picker)
- **End Date** (date picker with validation)

### Checkbox

- **Is Active** (boolean checkbox)

### Schema Validation

All fields have Zod validation:

```typescript
// In product.schema.ts
export const productPriceSchema = z
  .number({ required_error: 'Price is required' })
  .min(0, 'Price must be positive')
  .max(1000000, 'Price too high')
  .optional()

export const createProductSchema = () =>
  z
    .object({
      name: productNameSchema,
      price: productPriceSchema,
      startDate: productStartDateSchema,
      // ...
    })
    .refine(
      (data) => {
        if (data.startDate && data.endDate) {
          return new Date(data.startDate) <= new Date(data.endDate)
        }
        return true
      },
      {
        message: 'End date must be after start date',
        path: ['endDate'],
      },
    )
```

---

## Workflow

### Step 1: Generate

```bash
npm run generate:form product
```

Output:

```
✨ Generated files:
  ✅ src/types/api/features/product/product.types.ts
  ✅ src/services/api/features/product/product.api.ts
  ✅ src/schemas/features/product/product.schema.ts
  ✅ src/features/product/hooks/useCreateProduct.ts
  ✅ src/features/product/components/CreateProductForm.tsx
  ✅ src/app/[locale]/admin/(private)/products/create/page.tsx

📋 NEXT STEPS:
  1. Search "📝 TODO" in VSCode
  2. Follow step-by-step instructions
  3. Test at: http://localhost:3000/en/admin/products/create
```

### Step 2: Find TODOs

**In VSCode:**

1. Press `Ctrl+Shift+F` (Windows) or `Cmd+Shift+F` (Mac)
2. Search: `📝 TODO`
3. See all customization points

**Result:**

```
product.types.ts (3 matches)
  Line 35: 📝 TODO [STEP 1]: Add your entity fields here
  Line 55: 📝 TODO [STEP 2]: Add create request fields

product.api.ts (2 matches)
  Line 30: 📝 TODO [STEP 1]: Update endpoint

CreateProductForm.tsx (2 matches)
  Line 85: 📝 TODO [STEP 1]: Add your custom fields
  Line 140: 📝 TODO [STEP 4]: Test the form
```

### Step 3: Customize

Go through each TODO in order:

#### Example: Adding Custom Fields

**In `product.types.ts`:**

```typescript
export interface Product {
  id: string
  name: string
  description?: string

  // 📝 TODO [STEP 1]: Add your entity fields here
  // Uncomment and modify these examples:
  price: number
  category: string
  inStock: boolean

  createdAt: string
  updatedAt: string
}
```

**In `product.schema.ts`:**

```typescript
export const createProductSchema = () =>
  z.object({
    name: productNameSchema,
    description: optionalString('Description'),

    // 📝 TODO [STEP 2]: Add validation for custom fields
    price: z.number().min(0, 'Price must be positive'),
    category: z.string().min(1, 'Category is required'),
    inStock: z.boolean().optional().default(true),
  })
```

**In `CreateProductForm.tsx`:**

```typescript
{/* 📝 TODO [STEP 1]: Add your custom fields here */}

<TextField
  control={form.control}
  name="price"
  label="Price"
  type="number"
  required
/>

<Select
  control={form.control}
  name="category"
  label="Category"
  options={[
    { value: 'electronics', label: 'Electronics' },
    { value: 'clothing', label: 'Clothing' },
  ]}
/>
```

### Step 4: Export

Don't forget to export your new types and services:

**`src/types/api/index.ts`:**

```typescript
export * from './features/product/product.types'
```

**`src/services/api/index.ts`:**

```typescript
export * from './features/product/product.api'
```

### Step 5: Test

```bash
npm run dev
```

Navigate to: `http://localhost:3000/en/admin/products/create`

**Test checklist** (provided in generated files):

- ✅ Submit empty form → Check validation
- ✅ Submit invalid data → Check error messages
- ✅ Submit valid data → Check success flow
- ✅ Click cancel → Check navigation

---

## Advanced Examples

### Custom API Methods

**In `product.api.ts`:**

```typescript
// 📝 TODO [STEP 2]: Add custom methods here

// Already have getList, create, update, delete from base

// Add your custom endpoints:
export const productApi = {
  ...base,

  // Get products by category
  getByCategory: async (category: string): Promise<ApiResponse<Product[]>> => {
    const response = await apiGet(`${BASE_URL}/category/${category}`)
    return response.json()
  },

  // Bulk update prices
  bulkUpdatePrices: async (updates: Array<{ id: string; price: number }>) => {
    const response = await apiPost(`${BASE_URL}/bulk-update-prices`, {
      updates,
    })
    return response.json()
  },
}
```

### Custom Validation

**In `product.schema.ts`:**

```typescript
// Custom slug validator
export const productSlugSchema = z
  .string()
  .min(2, 'Slug too short')
  .max(100, 'Slug too long')
  .regex(/^[a-z0-9-]+$/, 'Slug must be lowercase with hyphens only')

// Custom price validator with range
export const productPriceSchema = z
  .number()
  .min(0, 'Price cannot be negative')
  .max(1000000, 'Price too high')
  .refine((val) => val % 0.01 === 0, 'Price must have max 2 decimal places')
```

### Nested Forms

**In `CreateProductForm.tsx`:**

```typescript
// Array of variants
{/* Variants */}
<FormArrayField
  control={form.control}
  name="variants"
  label="Product Variants"
  renderItem={(index) => (
    <Stack spacing={2}>
      <FormTextField
        control={form.control}
        name={`variants.${index}.size`}
        label="Size"
      />
      <FormTextField
        control={form.control}
        name={`variants.${index}.color`}
        label="Color"
      />
    </Stack>
  )}
/>
```

---

## Available Components in Generated Forms

All form components require `control={form.control}` prop. Here's what gets generated:

### TextField

```typescript
<TextField
  control={form.control}
  name="fieldName"
  label="Label"
  placeholder="Enter value"
  required
/>
```

### TextArea

```typescript
<TextArea
  control={form.control}
  name="description"
  label="Description"
  rows={4}
/>
```

### Number Input

```typescript
<TextField
  control={form.control}
  name="price"
  label="Price"
  type="number"
  required
/>
```

### Select (Single)

```typescript
<Select
  control={form.control}
  name="status"
  label="Status"
  options={[
    { value: 'option1', label: 'Option 1' },
    { value: 'option2', label: 'Option 2' },
  ]}
/>
```

### Select (Multi with Checkboxes)

```typescript
<Select
  control={form.control}
  isMulti
  showCheckboxes
  name="category"
  label="Category"
  options={[
    { value: 'electronics', label: 'Electronics' },
    { value: 'clothing', label: 'Clothing' },
  ]}
/>
```

### CommonDatePicker

```typescript
<CommonDatePicker
  name="startDate"
  label="Start Date"
  value={null}
  onChange={() => {}}
/>
```

### Checkbox (Single)

```typescript
<Checkbox
  control={form.control}
  name="isActive"
  label="Active"
  sx={{ p: 0 }}
/>
```

### CheckboxGroup (Multiple Selection)

```typescript
<CheckboxGroup
  control={form.control}
  name="interests"
  label="Interests (Select multiple)"
  options={[
    { value: 'technology', label: 'Technology' },
    { value: 'design', label: 'Design' },
    { value: 'marketing', label: 'Marketing' },
  ]}
/>
```

### RadioGroup (Single Selection)

```typescript
<RadioGroup
  control={form.control}
  name="notificationPreference"
  label="Notification Preference"
  options={[
    { value: 'email', label: 'Email Only' },
    { value: 'sms', label: 'SMS Only' },
    { value: 'both', label: 'Email & SMS' },
  ]}
/>
```

---

## Tips

### 1. Start Simple

Generate first, customize later. Don't try to add everything at once.

### 2. Follow the Order

TODOs are numbered (STEP 1, STEP 2, etc.) for a reason. Follow them in order.

### 3. Use Examples

Every TODO has examples. Uncomment and modify them.

### 4. Check Generated Tests

At the bottom of each file is a testing guide. Follow it.

### 5. Don't Skip Exports

Always export new types and services in index files.

---

## Troubleshooting

### TODO not showing up in search

Make sure VSCode search settings:

- "Use Regular Expression" is OFF
- "Match Case" is OFF
- Search in: "files"

### Generated files have errors

This is normal! The generated code uses generic types. Fix them by:

1. Following the TODO comments
2. Updating types to match your needs
3. Running `npm run lint:fix`

### Form validation not working

Check that:

1. Schema is properly defined
2. Field names match between schema and form
3. You imported the schema in your hook

---

## More Examples

See [QUICK_START.md](./QUICK_START.md) for detailed command reference and examples.

---

## Quick Reference

```bash
# Commands
npm run generate:form <name>      # Form + Create flow
npm run generate:table <name>     # Table + List flow
npm run generate:flow              # Interactive mode

# Find TODOs
Ctrl+Shift+F → search "📝 TODO"

# Test
npm run dev → http://localhost:3000/en/admin/<name>/create
```

---

**Next:** Learn about the [Tech Stack](./TECH_STACK.md) or [Development Workflow](./DEVELOPMENT.md)
