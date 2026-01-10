# 🎨 Typography & Switch Migration Documentation

## Typography System

### Font Family: Cera-Pro

**Weights Available:**

- Regular (400) - Default body text
- Medium (500) - Subtitles, emphasized text
- Bold (600) - Headings, buttons, strong emphasis

### Font Files Location

```
/public/fonts/
├── Cera-Pro-Regular.ttf
├── Cera-Pro-Medium.ttf
└── Cera-Pro-Bold.ttf
```

### Usage

Typography is automatically applied to all MUI components. No additional configuration needed.

**Heading Examples:**

```tsx
<Typography variant="h1">Heading 1 - 40px / 600</Typography>
<Typography variant="h2">Heading 2 - 32px / 600</Typography>
<Typography variant="h3">Heading 3 - 28px / 600</Typography>
<Typography variant="h4">Heading 4 - 24px / 600</Typography>
<Typography variant="h5">Heading 5 - 20px / 600</Typography>
<Typography variant="h6">Heading 6 - 16px / 600</Typography>
```

**Body Text Examples:**

```tsx
<Typography variant="body1">Body text - 14px / 400</Typography>
<Typography variant="body2">Small text - 12px / 400</Typography>
<Typography variant="subtitle1">Subtitle - 16px / 500</Typography>
<Typography variant="subtitle2">Small subtitle - 14px / 500</Typography>
```

**Other Variants:**

```tsx
<Button>Button text - 16px / 600</Button>
<Typography variant="caption">Caption - 12px / 400</Typography>
<Typography variant="overline">OVERLINE - 12px / 400</Typography>
```

### Typography Configuration

#### Base Settings

- **Font-family:** `'Cera-Pro', sans-serif`
- **Base font-size:** `14px`
- **Line-height:** `1.5` (default body)

#### Font Weights

- `fontWeightLight`: 400
- `fontWeightRegular`: 400
- `fontWeightMedium`: 500
- `fontWeightBold`: 600

---
