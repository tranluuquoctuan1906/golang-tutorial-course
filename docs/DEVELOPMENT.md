# Development Guide

Best practices and workflows for daily development.

---

## Daily Workflow

### 1. Start Development

```bash
# Pull latest changes
git checkout develop
git pull origin develop

# Create feature branch
git checkout -b feat/your-feature-name

# Start dev server
npm run dev
```

### 2. Develop Feature

```bash
# Generate feature (if needed)
npm run generate:form product
# or
npm run generate:table employee

# Make changes
# Test locally
# Run linter
npm run lint:fix
```

### 3. Commit Changes

```bash
# Check what changed
git status
git diff

# Stage changes
git add .

# Commit with conventional format
git commit -m "feat: add product creation form"
```

**Commit Format:**
- `feat:` - New feature
- `fix:` - Bug fix
- `refactor:` - Code refactoring
- `docs:` - Documentation
- `style:` - Formatting, missing semicolons
- `test:` - Adding tests
- `chore:` - Maintenance

### 4. Push & Create PR

```bash
# Push to remote
git push origin feat/your-feature-name

# Create PR on GitLab/GitHub
# Fill in PR template
# Request review
```

---

## Code Quality

### Linting

```bash
# Check for issues
npm run lint

# Auto-fix issues
npm run lint:fix
```

**ESLint will check:**
- Code style
- Best practices
- Unused variables
- TypeScript errors

### Formatting

Code is auto-formatted on save (if you have Prettier extension).

**Manual format:**
```bash
npx prettier --write .
```

### Type Checking

```bash
# Check TypeScript errors
npx tsc --noEmit
```

---

## Git Workflow

### Branch Strategy

```
main (production)
  └── develop (main development)
       ├── feat/feature-name
       ├── fix/bug-name
       └── refactor/refactor-name
```

### Branch Naming

- Features: `feat/add-product-filter`
- Bug fixes: `fix/validation-error`
- Refactoring: `refactor/api-structure`
- Documentation: `docs/update-readme`

### Commit Guidelines

**Good commits:**
```bash
feat: add product creation form
fix: resolve date picker validation error
refactor: improve API error handling
docs: update getting started guide
```

**Bad commits:**
```bash
update           # Too vague
fix bug          # Which bug?
WIP              # Work in progress (don't commit)
asdf             # Not descriptive
```

### Pull Request Process

1. **Create PR** with clear title and description
2. **Link issues** if applicable
3. **Request review** from team members
4. **Address feedback** promptly
5. **Squash and merge** when approved

**PR Template:**
```markdown
## What does this PR do?
Brief description of changes

## How to test?
1. Go to /admin/products/create
2. Fill in form
3. Submit

## Screenshots (if UI changes)
[Add screenshots]

## Checklist
- [ ] Code follows style guidelines
- [ ] Self-review completed
- [ ] Comments added for complex logic
- [ ] No console errors
- [ ] Tested on Chrome and Firefox
```

---

## Development Server

### Start Server

```bash
npm run dev
```

Server runs at: [http://localhost:3000](http://localhost:3000)

### Hot Reload

Changes auto-reload in browser. If stuck:

```bash
# Stop server (Ctrl+C)
# Clear Next.js cache
rm -rf .next
# Restart
npm run dev
```

### Environment Variables

Edit `.env.local`:

```env
NEXT_PUBLIC_API_URL=http://localhost:8000/api
NEXT_PUBLIC_APP_URL=http://localhost:3000
```

**Note:** Restart dev server after changing `.env.local`

---

## Building for Production

### Build

```bash
npm run build
```

**Checks performed:**
- TypeScript compilation
- ESLint checks
- Next.js optimization
- Bundle size analysis

### Test Production Build

```bash
npm run build
npm start
```

Server runs at: [http://localhost:3000](http://localhost:3000)

### Build Errors

**Common issues:**

1. **TypeScript errors**
   - Fix type errors shown in output
   - Run `npx tsc --noEmit` to see all errors

2. **ESLint errors**
   - Run `npm run lint:fix`
   - Fix remaining errors manually

3. **Missing environment variables**
   - Check `.env.local` is set up
   - Ensure all `NEXT_PUBLIC_` vars are defined

---

## Common Tasks

### Add New Feature

```bash
# Generate feature
npm run generate:form product

# Customize generated files
# Follow TODO comments

# Export types and services
# Add to src/types/api/index.ts
# Add to src/services/api/index.ts

# Test feature
npm run dev

# Commit
git add .
git commit -m "feat: add product management"
```

### Add API Endpoint

1. **Update API service** (`src/services/api/features/product/product.api.ts`):

```typescript
export const productApi = {
  ...base,

  getByCategory: async (category: string) => {
    const response = await apiGet(`${BASE_URL}/category/${category}`)
    return response.json()
  },
}
```

2. **Add SWR key** (`src/services/api/features/product/product.keys.ts`):

```typescript
export const PRODUCT_KEYS = {
  byCategory: (category: string) => ['products', 'category', category] as const,
}
```

3. **Use in component**:

```typescript
const { data } = useSWR(
  PRODUCT_KEYS.byCategory('electronics'),
  () => productApi.getByCategory('electronics')
)
```

### Update Form Validation

Edit schema file (`src/schemas/features/product/product.schema.ts`):

```typescript
export const createProductSchema = () =>
  z.object({
    name: requiredString('Product Name', { minLength: 2, maxLength: 100 }),
    price: z.number().min(0, 'Price must be positive'),
    category: z.enum(['electronics', 'clothing', 'food']),
  })
```

### Add Translation

1. **Add to English** (`messages/en.json`):

```json
{
  "product": {
    "title": "Products",
    "createButton": "Create Product",
    "nameLabel": "Product Name"
  }
}
```

2. **Add to Vietnamese** (`messages/vi.json`):

```json
{
  "product": {
    "title": "Sản phẩm",
    "createButton": "Tạo sản phẩm",
    "nameLabel": "Tên sản phẩm"
  }
}
```

3. **Use in component**:

```typescript
import { useTranslations } from 'next-intl'

const t = useTranslations('product')

return <h1>{t('title')}</h1>
```

---

## Debugging

### Console Logs

```typescript
console.log('Debug:', data)
console.error('Error:', error)
console.table(arrayData)
```

**Remove before committing!**

### React DevTools

Install browser extension:
- Chrome: [React DevTools](https://chrome.google.com/webstore/detail/react-developer-tools/fmkadmapgofadopljbjfkapdkoienihi)
- Firefox: [React DevTools](https://addons.mozilla.org/en-US/firefox/addon/react-devtools/)

### Network Tab

Use browser DevTools → Network tab to:
- Check API calls
- Inspect request/response
- Debug CORS issues

### SWR DevTools

View cache state:

```typescript
import { useSWRConfig } from 'swr'

const { cache } = useSWRConfig()
console.log(cache)
```

---

## Testing

### Manual Testing Checklist

Before committing:

- [ ] Feature works as expected
- [ ] Form validation works
- [ ] Error handling works
- [ ] No console errors
- [ ] Responsive design (mobile, tablet, desktop)
- [ ] Works in Chrome and Firefox
- [ ] i18n works (switch EN/VI)

### Test URLs

```
# Development
http://localhost:3000

# Common routes
http://localhost:3000/en/admin/products
http://localhost:3000/en/admin/products/create
http://localhost:3000/vi/admin/products
```

### Test Data

Use realistic test data:

```typescript
// Good
name: "MacBook Pro 14-inch"
price: 1999.99

// Bad
name: "test"
price: 123
```

---

## Performance

### Check Bundle Size

```bash
npm run build
```

Look for:
```
Route (app)                              Size     First Load JS
┌ ○ /                                    1.2 kB         150 kB
├ ○ /admin/products                      5.3 kB         155 kB
└ ○ /admin/products/create               8.1 kB         158 kB
```

### Optimize Heavy Components

```typescript
// Lazy load heavy components
import dynamic from 'next/dynamic'

const Chart = dynamic(() => import('./Chart'), {
  ssr: false,
  loading: () => <Skeleton />,
})
```

### Optimize Images

```typescript
import Image from 'next/image'

<Image
  src="/image.jpg"
  width={500}
  height={300}
  alt="Description"
  loading="lazy"  // Lazy load
/>
```

---

## Troubleshooting

### Port Already in Use

```bash
# Windows
netstat -ano | findstr :3000
taskkill /F /PID <PID>

# Mac/Linux
lsof -ti:3000 | xargs kill -9
```

### Module Not Found

```bash
rm -rf node_modules package-lock.json
npm install
```

### TypeScript Errors After Generation

1. Export types in `src/types/api/index.ts`
2. Export services in `src/services/api/index.ts`
3. Restart TypeScript server in VS Code: `Ctrl+Shift+P` → "Restart TS Server"

### Git Hooks Failing

```bash
# Skip hooks (only if really needed!)
git commit --no-verify -m "message"

# Better: Fix the issues
npm run lint:fix
git add .
git commit -m "message"
```

### Environment Variables Not Working

1. Check `.env.local` exists
2. Check variables start with `NEXT_PUBLIC_`
3. Restart dev server
4. Clear browser cache

---

## VS Code Setup

### Recommended Extensions

```json
{
  "recommendations": [
    "dbaeumer.vscode-eslint",
    "esbenp.prettier-vscode",
    "bradlc.vscode-tailwindcss",
    "ms-vscode.vscode-typescript-next"
  ]
}
```

### Settings

Create `.vscode/settings.json`:

```json
{
  "editor.formatOnSave": true,
  "editor.defaultFormatter": "esbenp.prettier-vscode",
  "editor.codeActionsOnSave": {
    "source.fixAll.eslint": true
  },
  "typescript.preferences.importModuleSpecifier": "non-relative"
}
```

---

## Best Practices

### Code Organization

```typescript
// ✅ Good: Organized imports
import React from 'react'
import { Box, Button } from '@mui/material'
import { useForm } from 'react-hook-form'

import { FormTextField } from '@/components/organisms/Form'
import { productApi } from '@/services/api'

// ❌ Bad: Mixed imports
import { Box, Button } from '@mui/material'
import { productApi } from '@/services/api'
import React from 'react'
import { useForm } from 'react-hook-form'
```

### Component Structure

```typescript
// ✅ Good: Clear structure
export default function ProductForm() {
  // Hooks
  const form = useForm()
  const { data } = useSWR(key, fetcher)

  // Handlers
  const handleSubmit = async () => {
    // ...
  }

  // Render
  return (
    <Box>
      {/* ... */}
    </Box>
  )
}
```

### Type Safety

```typescript
// ✅ Good: Type-safe
interface Product {
  id: string
  name: string
  price: number
}

const product: Product = { id: '1', name: 'Test', price: 100 }

// ❌ Bad: Using any
const product: any = { id: '1', name: 'Test', price: 100 }
```

### Error Handling

```typescript
// ✅ Good: Handle errors
try {
  const result = await productApi.create(data)
  toast.success('Product created!')
} catch (error) {
  console.error('Failed to create product:', error)
  toast.error('Failed to create product')
}

// ❌ Bad: No error handling
const result = await productApi.create(data)
toast.success('Product created!')
```

---

## Quick Reference

```bash
# Development
npm run dev                    # Start dev server
npm run build                  # Build for production
npm start                      # Start production server

# Code Quality
npm run lint                   # Check code
npm run lint:fix               # Fix issues

# Code Generation
npm run generate:form <name>   # Generate form
npm run generate:table <name>  # Generate table

# Git
git checkout -b feat/name      # Create branch
git commit -m "feat: message"  # Commit
git push origin feat/name      # Push
```

---

**Next:** Explore the [Architecture Guide](./ARCHITECTURE_GUIDE.md) or [Form Integration Guide](./FORM_INTEGRATION_GUIDE.md)
