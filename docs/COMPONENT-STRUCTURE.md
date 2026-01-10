# React Function Component Structure Guidelines

## Table of Contents

- [Recommended Order of Declarations](#recommended-order-of-declarations)
- [Detailed Guidelines](#detailed-guidelines)
- [Project-Specific Patterns](#project-specific-patterns)
- [Common Mistakes to Avoid](#common-mistakes-to-avoid)
- [Code Review Checklist](#code-review-checklist)

---

## Recommended Order of Declarations

Follow this order for consistent and readable React function components:

```tsx
function MyComponent(props: Props) {
  // 1. DESTRUCTURING & CONSTANTS
  const { propA, propB } = props
  const CONSTANT_VALUE = 100

  // 2. HOOKS - STATE
  const [state, setState] = useState(initialState)
  const [count, setCount] = useState(0)

  // 3. HOOKS - CONTEXT & REFS
  const context = useContext(MyContext)
  const ref = useRef<HTMLDivElement>(null)

  // 4. CUSTOM HOOKS
  const { data, loading } = useCustomHook()
  const translation = useTranslation()

  // 5. COMPUTED VALUES - useMemo
  const expensiveValue = useMemo(() => {
    return computeExpensiveValue(state)
  }, [state])

  // 6. REGULAR FUNCTIONS & HANDLERS
  const handleClick = () => {
    // handler logic
  }

  const processData = (data: Data) => {
    // processing logic
  }

  // 7. MEMOIZED CALLBACKS - useCallback
  const memoizedCallback = useCallback(
    () => {
      handleClick()
    },
    [
      /* deps */
    ],
  )

  // 8. SIDE EFFECTS - useEffect
  useEffect(
    () => {
      // effect logic
      return () => {
        // cleanup
      }
    },
    [
      /* deps */
    ],
  )

  // 9. EARLY RETURNS (if any)
  if (loading) return <Loading />
  if (!data) return null

  // 10. RENDER HELPERS (optional)
  const renderHeader = () => <Header />

  // 11. JSX RETURN
  return <div>{/* component JSX */}</div>
}
```

## Key Principles

1. **Top to Bottom Flow**: Read from data declaration → processing → effects → render
2. **Hook Order**: useState → useContext/useRef → custom hooks → useMemo → regular functions → useCallback → useEffect
3. **Group Related Logic**: Keep related state and effects together when needed
4. **Early Returns**: Place conditional returns before the main render
5. **Consistency**: Always follow the same order across all components
6. **Simplicity**: Don't over-optimize. Use useMemo/useCallback only when needed
7. **Readability**: Code should be easy to understand and maintain

---

## Why This Order?

- **useState first**: Declare state before using it
- **Custom hooks early**: Get data/logic before computing
- **useMemo before functions**: Computed values available for handlers
- **Functions before useCallback**: Define logic before memoizing
- **useCallback before useEffect**: Memoized callbacks ready for effects
- **useEffect last**: Side effects run after all logic is defined
- **Return at end**: Render uses all declared logic above

---

## ESLint Support

Currently, ESLint cannot automatically enforce function component internal ordering.
This is a convention-based guideline that should be followed during code review.

### Enforced Rules

Our ESLint configuration enforces:

- Hook rules (`eslint-plugin-react-hooks`)
- Import order (`eslint-plugin-import`)
- TypeScript types (`@typescript-eslint`)
- Next.js best practices (`@next/eslint-plugin-next`)

### Manual Review Required

- Component internal ordering (this guide)
- Atomic Design classification
- Performance optimization decisions
- Business logic organization

---

## References

- [React Hooks Documentation](https://react.dev/reference/react)
- [React Hook Form](https://react-hook-form.com/)
- [SWR Documentation](https://swr.vercel.app/)
- [Material-UI Components](https://mui.com/)
- [Next.js Best Practices](https://nextjs.org/docs)

---

**Last Updated**: January 2026
**Maintained By**: Development Team
