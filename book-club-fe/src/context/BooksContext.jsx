import { createContext, useContext, useReducer, useCallback } from 'react';
import * as booksApi from '../api/books';
import { DEMO_BOOKS } from '../api/constants';

const BooksContext = createContext(null);
const BooksDispatchContext = createContext(null);

function booksReducer(state, action) {
  switch (action.type) {
    case 'SET_LOADING':
      return { ...state, loading: action.payload };
    case 'SET_ERROR':
      return { ...state, error: action.payload, loading: false };
    case 'SET_BOOKS':
      return { ...state, books: action.payload, loading: false, error: null };
    case 'ADD_BOOK':
      return { ...state, books: [action.payload, ...state.books] };
    case 'UPDATE_BOOK':
      return {
        ...state,
        books: state.books.map((b) =>
          b.id === action.payload.id ? action.payload : b
        ),
      };
    case 'DELETE_BOOK':
      return {
        ...state,
        books: state.books.filter((b) => b.id !== action.payload),
      };
    default:
      return state;
  }
}

const initialState = {
  books: [],
  loading: false,
  error: null,
};

export function BooksProvider({ children }) {
  const [state, dispatch] = useReducer(booksReducer, initialState);

  return (
    <BooksContext.Provider value={state}>
      <BooksDispatchContext.Provider value={dispatch}>
        {children}
      </BooksDispatchContext.Provider>
    </BooksContext.Provider>
  );
}

export function useBooksState() {
  return useContext(BooksContext);
}

export function useBooksDispatch() {
  return useContext(BooksDispatchContext);
}

// Actions — these are used inside hooks or components
export function useBooksActions() {
  const dispatch = useBooksDispatch();

  const fetchBooks = useCallback(async () => {
    dispatch({ type: 'SET_LOADING', payload: true });
    try {
      const data = await booksApi.getBooks();
      dispatch({ type: 'SET_BOOKS', payload: data || [] });
    } catch {
      // Fall back to demo data when API is unreachable
      dispatch({ type: 'SET_BOOKS', payload: [...DEMO_BOOKS] });
    }
  }, [dispatch]);

  const addBook = useCallback(
    async (formData) => {
      const book = await booksApi.createBook(formData).catch(() => ({
        ...formData,
        id: crypto.randomUUID(),
      }));
      dispatch({ type: 'ADD_BOOK', payload: book });
      return book;
    },
    [dispatch]
  );

  const editBook = useCallback(
    async (id, formData) => {
      const updated = await booksApi.updateBook(id, formData).catch(() => ({
        ...formData,
        id,
      }));
      dispatch({ type: 'UPDATE_BOOK', payload: updated });
      return updated;
    },
    [dispatch]
  );

  const removeBook = useCallback(
    async (id) => {
      await booksApi.deleteBook(id).catch(() => {});
      dispatch({ type: 'DELETE_BOOK', payload: id });
    },
    [dispatch]
  );

  return { fetchBooks, addBook, editBook, removeBook };
}
