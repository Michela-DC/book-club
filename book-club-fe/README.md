# React + Vite

This template provides a minimal setup to get React working in Vite with HMR and some ESLint rules.

Currently, two official plugins are available:

- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react) uses [Oxc](https://oxc.rs)
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react-swc) uses [SWC](https://swc.rs/)

## React Compiler

The React Compiler is not enabled on this template because of its impact on dev & build performances. To add it, see [this documentation](https://react.dev/learn/react-compiler/installation).

## Expanding the ESLint configuration

If you are developing a production application, we recommend using TypeScript with type-aware lint rules enabled. Check out the [TS template](https://github.com/vitejs/vite/tree/main/packages/create-vite/template-react-ts) for information on how to integrate TypeScript and [`typescript-eslint`](https://typescript-eslint.io) in your project.

Project structure:
```
book-club/
├── .env.example              <- set VITE_API_URL here
└── src/
    ├── api/
    │   ├── books.js          <- all fetch calls (GET/PUT/PATCH/DELETE)
    │   └── constants.js      <- status labels, filter tabs, demo data
    ├── context/
    │   └── BooksContext.jsx  <- useReducer global state + action hooks
    ├── hooks/
    │   ├── useFilter.js      <- tab filtering logic
    │   └── useToast.js       <- toast notification hook
    ├── components/
    │   ├── ui/
    │   │   ├── Button        <- primary / secondary / danger / ghost variants
    │   │   ├── Modal         <- accessible overlay, Esc to close
    │   │   ├── Header        <- sticky top bar
    │   │   ├── StatusBadge   <- color-coded status pill
    │   │   └── Toast         <- bottom-right notification
    │   └── books/
    │       ├── BookCard      <- spine-style card with hover tilt
    │       ├── BookForm      <- controlled form with validation
    │       ├── BookGrid      <- grid + loading + empty states
    │       └── FilterTabs    <- status tabs with live counts
    ├── pages/
    │   └── BookShelf.jsx     <- main page, wires everything together
    ├── styles/
    │   └── globals.css       <- CSS variables, resets, Google Fonts
    └── App.jsx               <- BooksProvider + BrowserRouter
```