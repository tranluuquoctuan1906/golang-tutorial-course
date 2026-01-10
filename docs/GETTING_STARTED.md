# Getting Started Guide

Welcome to the team! This guide will help you set up and create your first feature.

---

## Day 1: Setup & First Feature

### Step 1: Setup Environment (10 minutes)

#### 1.1 Clone and Install

```bash
# Clone repository
git clone <repository-url>
cd ts-employer-fe

# Install dependencies (this may take 2-3 minutes)
npm install
```

#### 1.2 Environment Configuration

```bash
# Copy environment template
cp .env.example .env.local
```

Edit `.env.local`:
```env
NEXT_PUBLIC_API_URL=your-api-url-here
NEXT_PUBLIC_APP_URL=http://localhost:3000
```

Ask your team lead for the correct API URL.

#### 1.3 Start Development Server

```bash
npm run dev
```

Open [http://localhost:3000](http://localhost:3000)

**You should see the login page.** If not, check [Troubleshooting](#troubleshooting) below.

---

### Step 2: Understand Project Structure (5 minutes)

Our project uses **feature-based architecture**:

```
src/
├── features/           # Each feature is self-contained
│   ├── candidates/    # Example: Candidate management
│   │   ├── components/
│   │   ├── hooks/
│   │   └── types/
│   └── dashboard/     # Example: Dashboard widgets
│
├── components/         # Shared UI components (Atomic Design)
│   ├── atoms/         # Button, Input, Badge
│   ├── molecules/     # FormField, DatePicker
│   └── organisms/     # Table, Form, Layout
│
├── services/          # API calls
└── app/               # Next.js pages
```

**Key Concept**: Everything related to one feature lives in its folder.

---

### Step 3: Create Your First Feature (15 minutes)

Let's create a simple "Product" feature with a create form.

#### 3.1 Generate Feature

```bash
npm run generate:form product
```

This creates everything inside `src/features/product/`:
```
src/features/product/
├── types/product.types.ts           - TypeScript types
├── services/product.api.ts          - API service
├── schemas/product.schema.ts        - Form validation
├── constants/product.keys.ts        - SWR cache keys
├── hooks/useCreateProduct.ts        - Custom hook
├── components/CreateProductForm.tsx - Form component
└── index.tsx                        - Feature export
```
Plus: `src/app/[locale]/admin/(private)/products/create/page.tsx` - Page

#### 3.2 Find What to Customize

Open VSCode and press `Ctrl+Shift+F` (Windows) or `Cmd+Shift+F` (Mac).

Search for: `📝 TODO`

You'll see all the places you need to customize:

```
product.types.ts
  Line 35: 📝 TODO [STEP 1]: Add your entity fields here
  Line 55: 📝 TODO [STEP 2]: Add create request fields here

product.api.ts
  Line 30: 📝 TODO [STEP 1]: Update endpoint to match backend

CreateProductForm.tsx
  Line 85: 📝 TODO [STEP 1]: Add your custom fields here
```

#### 3.3 Add Your Fields

**Example: Adding price and category fields**

1. **Update Types** (`product.types.ts`):

```typescript
export interface Product {
  id: string
  name: string
  description?: string

  // Your custom fields
  price: number
  category: string

  createdAt: string
  updatedAt: string
}
```

2. **Update Schema** (`product.schema.ts`):

```typescript
export const createProductSchema = () =>
  z.object({
    name: productNameSchema,
    description: optionalString('Description'),

    // Your custom fields
    price: z.number().min(0, 'Price must be positive'),
    category: z.string().min(1, 'Category is required'),
  })
```

3. **Update Form** (`CreateProductForm.tsx`):

```typescript
{/* Add these fields after the description field */}
<FormTextField
  control={form.control}
  name="price"
  label="Price"
  type="number"
  required
/>

<FormTextField
  control={form.control}
  name="category"
  label="Category"
  required
/>
```

#### 3.4 Test Your Feature

```bash
# Make sure dev server is running
npm run dev
```

Navigate to: [http://localhost:3000/en/admin/products/create](http://localhost:3000/en/admin/products/create)

**Test cases:**
1. Submit empty form → Should show validation errors
2. Enter invalid data → Should show specific error messages
3. Enter valid data → Should submit successfully

---

### Step 4: Common Tasks

#### Add a Table/List View

```bash
npm run generate:table product
```

This creates a complete table with:
- Pagination
- Sorting
- Filters
- Bulk actions

#### Update API Endpoint

Edit `src/features/product/services/product.api.ts`:

```typescript
const BASE_URL = '/api/products'  // Change this to match your backend
```

#### Add More Form Fields

See the TODO comments in your form component. We provide examples for:
- Text inputs
- Dropdowns
- Date pickers
- File uploads
- Rich text editors

---

## Troubleshooting

### Port 3000 already in use

```bash
# Windows
netstat -ano | findstr :3000
taskkill /F /PID <process-id>

# Mac/Linux
lsof -ti:3000 | xargs kill -9
```

### Module not found errors

```bash
# Clear and reinstall
rm -rf node_modules package-lock.json
npm install
```

### TypeScript errors after generation

Make sure you export new types in index files:

```typescript
// src/types/api/index.ts
export * from './features/product/product.types'

// src/services/api/index.ts
export * from './features/product/product.api'
```

### Environment variables not working

- Restart dev server after changing `.env.local`
- Variables must start with `NEXT_PUBLIC_` to be available in browser

---

## Next Steps

Now that you've created your first feature:

1. **Learn Code Generation** - [Code Generation Guide](./CODE_GENERATION.md)
2. **Understand Architecture** - [Architecture Guide](./ARCHITECTURE_GUIDE.md)
3. **Review Tech Stack** - [Tech Stack Guide](./TECH_STACK.md)
4. **Development Workflow** - [Development Guide](./DEVELOPMENT.md)

---

## Quick Reference

```bash
# Start development
npm run dev

# Generate features
npm run generate:form <name>
npm run generate:table <name>

# Find TODOs
# VSCode: Ctrl+Shift+F → search "📝 TODO"

# Check code quality
npm run lint
npm run lint:fix

# Build for production
npm run build
```

---

## Need Help?

- Search for `📝 TODO` to find customization points
- Check examples in generated files
- Review [Architecture Guide](./ARCHITECTURE_GUIDE.md)
- Ask your team lead

**Welcome aboard! Happy coding!**
