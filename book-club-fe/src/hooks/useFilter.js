import { useState, useMemo } from 'react';
import { useBooksState } from '../context/BooksContext';

export function useFilter() {
  const { books } = useBooksState();
  const [activeFilter, setActiveFilter] = useState('all');

  const filteredBooks = useMemo(() => {
    if (activeFilter === 'all') return books;
    return books.filter((b) => b.status === activeFilter);
  }, [books, activeFilter]);

  return { activeFilter, setActiveFilter, filteredBooks };
}
