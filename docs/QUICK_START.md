# 🚀 Quick Start Guide - Complete Flow Generator

## 📋 Commands

### Generate CREATE FORM
```bash
npm run generate:form <feature-name>
```

### Generate TABLE/LIST  
```bash
npm run generate:table <feature-name>
```

### Interactive Mode
```bash
npm run generate:flow
```

---

## 🎯 Examples

### Example 1: Product Management Form
```bash
npm run generate:form product
```

**Result:**
- ✅ 8 files created
- ✅ Each with TODO comments
- ✅ Ready to customize

### Example 2: Employee List Table
```bash
npm run generate:table employee
```

**Result:**
- ✅ 8 files created  
- ✅ Table with pagination, sorting, selection
- ✅ Ready to customize

---

## 🔍 Find Customization Points

1. **Open VSCode**
2. **Press:** `Ctrl+Shift+F` (Windows) or `Cmd+Shift+F` (Mac)
3. **Search:** `📝 TODO`
4. **Result:** All customization points highlighted!

---

## 📝 Customization Workflow

### Step 1: Types
```typescript
// src/types/api/features/product/product.types.ts

export interface Product {
  id: string
  name: string
  
  // 📝 TODO [STEP 1]: Add your fields
  // Uncomment and customize:
  // price: number
  // sku: string
  // category: string
}
```

### Step 2: Schema
```typescript
// src/schemas/features/product/product.schema.ts

export const createProductSchema = () =>
  z.object({
    name: requiredString('Name'),
    
    // 📝 TODO [STEP 2]: Add validation
    // price: z.number().min(0),
    // sku: z.string().min(3),
  })
```

### Step 3: Form
```typescript
// src/features/product/components/CreateProductForm.tsx

<FormTextField name="name" label="Name" required />

{/* 📝 TODO [STEP 1]: Add fields */}
{/* Uncomment:
<FormTextField 
  name="price" 
  label="Price" 
  type="number" 
/>
*/}
```

### Step 4: Export
```typescript
// src/services/api/index.ts
export * from './features/product/product.api'

// src/types/api/index.ts
export * from './features/product/product.types'
```

### Step 5: Test
```bash
npm run dev
# Navigate to: http://localhost:3000/en/admin/products/create
```

---

## ✅ Checklist After Generation

- [ ] Search for `📝 TODO` in all files
- [ ] Uncomment and customize examples
- [ ] Update API endpoint in `.api.ts`
- [ ] Add validation rules in `.schema.ts`
- [ ] Add form fields in component
- [ ] Export in `index.ts` files
- [ ] Test the feature
- [ ] Remove unused TODO comments

---

## 📚 Documentation

- **Examples:** [DEMO_GENERATED_OUTPUT.md](../DEMO_GENERATED_OUTPUT.md)
- **Detailed Examples:** [GENERATED_FILE_EXAMPLE.md](./GENERATED_FILE_EXAMPLE.md)
- **Architecture:** [../ARCHITECTURE_GUIDE.md](../ARCHITECTURE_GUIDE.md)

---

## 🎉 Tips

### Tip 1: Use Examples
All TODO comments include copy-paste ready examples!

### Tip 2: Follow Line Numbers
Each TODO shows approximate line number: `(line ~35)`

### Tip 3: Step Numbers
Follow `[STEP 1]`, `[STEP 2]`, etc. in order

### Tip 4: Test Early
Test after each step, don't wait until end

### Tip 5: Remove TODOs
Delete TODO comments after completing each step

---

**Happy Coding! 🚀**
