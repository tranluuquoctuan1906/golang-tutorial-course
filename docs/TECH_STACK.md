# Tech Stack

Technologies and libraries used in this project.

---

## Core Framework

### Next.js 15 (App Router)
- **Purpose:** React framework with server-side rendering
- **Why:** Best performance, SEO, and developer experience
- **Docs:** [nextjs.org](https://nextjs.org/)

**Key Features We Use:**
- App Router (new routing system)
- Server Components
- API Routes
- Image Optimization
- Internationalization (i18n)

### TypeScript 5.0
- **Purpose:** Type-safe JavaScript
- **Why:** Catch errors early, better IDE support, safer refactoring
- **Docs:** [typescriptlang.org](https://www.typescriptlang.org/)

**Configuration:**
- Strict mode enabled
- Path aliases (`@/components`, `@/services`, etc.)
- Type-safe API layer

---

## UI & Styling

### Material-UI v7
- **Purpose:** Component library
- **Why:** Professional components, theming, accessibility
- **Docs:** [mui.com](https://mui.com/)

**Components We Use:**
- DataGrid (tables)
- Form components (TextField, Select, etc.)
- Layout (Box, Stack, Grid)
- Navigation (Drawer, AppBar)
- Feedback (Dialog, Snackbar)

### Emotion
- **Purpose:** CSS-in-JS
- **Why:** Comes with MUI, dynamic styling, type-safe
- **Docs:** [emotion.sh](https://emotion.sh/)

**Usage:**
```typescript
import { styled } from '@mui/material/styles'

const StyledBox = styled(Box)(({ theme }) => ({
  padding: theme.spacing(2),
  backgroundColor: theme.palette.background.paper,
}))
```

---

## State Management

### SWR (Stale-While-Revalidate)
- **Purpose:** Server state management
- **Why:** Auto-revalidation, caching, optimistic updates
- **Docs:** [swr.vercel.app](https://swr.vercel.app/)

**Usage:**
```typescript
import useSWR from 'swr'

const { data, error, isLoading } = useSWR(
  PRODUCT_KEYS.list(),
  productApi.getList
)
```

**Features:**
- Automatic caching
- Revalidation on focus
- Revalidation on network recovery
- Optimistic UI updates
- Polling

### React Hook Form
- **Purpose:** Form state management
- **Why:** Minimal re-renders, easy validation, great DX
- **Docs:** [react-hook-form.com](https://react-hook-form.com/)

**Usage:**
```typescript
const form = useForm({
  resolver: zodResolver(createProductSchema()),
})
```

**Features:**
- Minimal re-renders
- Built-in validation
- Error handling
- Field arrays
- Watch values

---

## Validation

### Zod
- **Purpose:** Schema validation
- **Why:** Type-safe, composable, great error messages
- **Docs:** [zod.dev](https://zod.dev/)

**Usage:**
```typescript
const productSchema = z.object({
  name: z.string().min(2).max(100),
  price: z.number().min(0),
  category: z.enum(['electronics', 'clothing']),
})

type Product = z.infer<typeof productSchema>
```

**Integration:**
- Works with React Hook Form via `zodResolver`
- Reusable schemas in `src/schemas/`
- Automatic TypeScript types

---

## Data Handling

### Day.js
- **Purpose:** Date manipulation
- **Why:** Lightweight, moment.js replacement, i18n support
- **Docs:** [day.js.org](https://day.js.org/)

**Usage:**
```typescript
import dayjs from 'dayjs'

const formatted = dayjs(date).format('DD/MM/YYYY')
const isAfter = dayjs(date1).isAfter(date2)
```

### Chart.js + react-chartjs-2
- **Purpose:** Charts and graphs
- **Why:** Flexible, responsive, well-documented
- **Docs:** [chartjs.org](https://www.chartjs.org/)

**Charts We Use:**
- Line charts (trends)
- Bar charts (comparisons)
- Pie charts (distributions)
- Doughnut charts (percentages)

---

## Rich Content

### TipTap
- **Purpose:** Rich text editor
- **Why:** Headless, extensible, modern
- **Docs:** [tiptap.dev](https://tiptap.dev/)

**Features:**
- WYSIWYG editing
- Markdown support
- Collaborative editing ready
- Custom extensions

### dnd-kit
- **Purpose:** Drag and drop
- **Why:** Accessible, mobile-friendly, modular
- **Docs:** [dndkit.com](https://dndkit.com/)

**Use Cases:**
- Reorderable lists
- Kanban boards
- File uploads
- Dashboard widgets

---

## Internationalization

### next-intl
- **Purpose:** i18n for Next.js
- **Why:** App Router support, type-safe, server components
- **Docs:** [next-intl-docs.vercel.app](https://next-intl-docs.vercel.app/)

**Languages:**
- English (en)
- Vietnamese (vi)

**Usage:**
```typescript
import { useTranslations } from 'next-intl'

const t = useTranslations('common')
return <h1>{t('welcome')}</h1>
```

**Files:**
- `messages/en.json` - English translations
- `messages/vi.json` - Vietnamese translations

---

## Development Tools

### ESLint
- **Purpose:** Code linting
- **Why:** Enforce code quality and consistency
- **Config:** `.eslintrc.json`

**Rules:**
- Next.js recommended
- TypeScript recommended
- React hooks rules
- Custom rules for our project

### Prettier
- **Purpose:** Code formatting
- **Why:** Consistent code style, auto-format on save
- **Config:** `.prettierrc`

**Settings:**
- Single quotes
- 2 spaces indentation
- Trailing commas
- Semicolons

### Husky
- **Purpose:** Git hooks
- **Why:** Run checks before commit/push
- **Config:** `.husky/`

**Hooks:**
- `pre-commit` - Run linter and formatter
- Commit message validation

### Commitlint
- **Purpose:** Enforce commit message format
- **Why:** Consistent git history, automatic changelog
- **Config:** `.commitlintrc.json`

**Format:**
```
feat: add product creation
fix: resolve validation error
docs: update README
refactor: improve API structure
```

---

## Architecture Patterns

### Atomic Design
Components organized by complexity:

```
atoms/        → Button, Input, Badge
molecules/    → FormField, DatePicker, Card
organisms/    → Table, Form, Header, Sidebar
```

**Why:**
- Reusable components
- Easy to maintain
- Clear component hierarchy

### Feature-Based Structure
Each feature is self-contained:

```
features/product/
├── components/    → UI specific to products
├── hooks/         → Product-specific logic
├── types/         → Product types
└── utils/         → Product utilities
```

**Why:**
- Easy to find related code
- Can be extracted to separate package
- Clear boundaries

### API Layer
Centralized API handling:

```
services/api/
├── common/           → Base API utilities
└── features/         → Feature-specific APIs
    └── product/
        └── product.api.ts
```

**Why:**
- Single source of truth for API calls
- Type-safe
- Easy to mock for testing

---

## Package Management

### npm
- **Purpose:** Package manager
- **Why:** Standard, reliable, good performance

**Key Commands:**
```bash
npm install          # Install dependencies
npm run dev          # Start dev server
npm run build        # Build for production
npm run lint         # Run linter
```

**Lock File:** `package-lock.json` (commit this!)

---

## Browser Support

- **Chrome:** Latest 2 versions
- **Firefox:** Latest 2 versions
- **Safari:** Latest 2 versions
- **Edge:** Latest 2 versions

**Mobile:**
- iOS Safari 12+
- Chrome Android Latest

---

## Performance Features

### Code Splitting
Automatic with Next.js App Router:
```typescript
// Lazy load heavy components
const ChartComponent = dynamic(() => import('./Chart'), {
  ssr: false,
  loading: () => <Skeleton />,
})
```

### Image Optimization
Using Next.js Image:
```typescript
import Image from 'next/image'

<Image
  src="/logo.png"
  width={200}
  height={100}
  alt="Logo"
  priority // For above-fold images
/>
```

### Font Optimization
Using Next.js Font:
```typescript
import { Inter } from 'next/font/google'

const inter = Inter({ subsets: ['latin'] })
```

---

## Environment Variables

### Available Variables

```env
# API Configuration
NEXT_PUBLIC_API_URL=          # Backend API URL

# App Configuration
NEXT_PUBLIC_APP_URL=          # Frontend URL

# Feature Flags (optional)
NEXT_PUBLIC_ENABLE_ANALYTICS= # Enable analytics
```

**Note:** Variables with `NEXT_PUBLIC_` prefix are available in browser.

---

## Key Dependencies

```json
{
  "next": "^15.5.0",
  "react": "^19.0.0",
  "typescript": "^5.0.0",
  "@mui/material": "^7.0.0",
  "swr": "^2.2.0",
  "react-hook-form": "^7.0.0",
  "zod": "^3.22.0",
  "dayjs": "^1.11.0",
  "next-intl": "^3.0.0"
}
```

See `package.json` for complete list.

---

## Why These Choices?

### Performance
- Next.js SSR for faster initial load
- SWR for optimistic UI updates
- Code splitting for smaller bundles

### Developer Experience
- TypeScript for type safety
- Hot reload for fast development
- Code generation for productivity

### User Experience
- Material-UI for professional UI
- i18n for multi-language support
- Responsive design for all devices

### Maintainability
- Feature-based architecture
- Atomic Design pattern
- Comprehensive testing tools

---

## Learning Resources

### Official Docs
- [Next.js Learn](https://nextjs.org/learn)
- [React Docs](https://react.dev/)
- [TypeScript Handbook](https://www.typescriptlang.org/docs/)
- [MUI Documentation](https://mui.com/getting-started/)

### Our Guides
- [Architecture Guide](./ARCHITECTURE_GUIDE.md) - How it all fits together
- [Form Integration](./FORM_INTEGRATION_GUIDE.md) - Building forms
- [Table Integration](./TABLE_INTEGRATION_GUIDE.md) - Building tables

---

## Version Updates

We regularly update dependencies for:
- Security patches
- New features
- Performance improvements
- Bug fixes

**Update Strategy:**
- Patch versions: Auto-update
- Minor versions: Monthly review
- Major versions: Planned migrations

---

**Next:** Learn the [Development Workflow](./DEVELOPMENT.md) or explore the [Architecture](./ARCHITECTURE_GUIDE.md)
