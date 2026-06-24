import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { BooksProvider } from './context/BooksContext';
import BookShelf from './pages/BookShelf';
import './styles/globals.css';

export default function App() {
  return (
    <BooksProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<BookShelf />} />
        </Routes>
      </BrowserRouter>
    </BooksProvider>
  );
}
