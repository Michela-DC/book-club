import { useEffect, useState } from 'react';
import { useBooksState, useBooksActions } from '../context/BooksContext';
import { useFilter } from '../hooks/useFilter';
import { useToast } from '../hooks/useToast';

import Header from '../components/ui/Header';
import Toast from '../components/ui/Toast';
import Modal from '../components/ui/Modal';
import BookGrid from '../components/books/BookGrid';
import BookForm from '../components/books/BookForm';
import FilterTabs from '../components/books/FilterTabs';

import styles from './BookShelf.module.css';

export default function BookShelf() {
  const { books, loading } = useBooksState();
  const { fetchBooks, addBook, editBook, removeBook } = useBooksActions();
  const { activeFilter, setActiveFilter, filteredBooks } = useFilter();
  const { toast, showToast } = useToast();

  const [modalOpen, setModalOpen] = useState(false);
  const [editingBook, setEditingBook] = useState(null); // null = adding
  const [isSaving, setIsSaving] = useState(false);

  useEffect(() => { fetchBooks(); }, [fetchBooks]);

  function openAdd() { setEditingBook(null); setModalOpen(true); }
  function openEdit(book) { setEditingBook(book); setModalOpen(true); }
  function closeModal() { setModalOpen(false); setEditingBook(null); }

  async function handleSubmit(formData) {
    setIsSaving(true);
    try {
      if (editingBook) {
        await editBook(editingBook.id, formData);
        showToast('Book updated.');
      } else {
        await addBook(formData);
        showToast('Book added to your shelf.');
      }
      closeModal();
    } catch {
      showToast('Something went wrong — try again.', 'error');
    } finally {
      setIsSaving(false);
    }
  }

  async function handleDelete(book) {
    if (!window.confirm(`Remove "${book.title}" from your shelf?`)) return;
    try {
      await removeBook(book.id);
      showToast('Book removed.');
    } catch {
      showToast('Could not remove this book.', 'error');
    }
  }

  return (
    <>
      <Header onAddBook={openAdd} />

      <main className={styles.main}>
        <FilterTabs
          active={activeFilter}
          onChange={setActiveFilter}
          books={books}
        />
        <BookGrid
          books={filteredBooks}
          loading={loading}
          onEdit={openEdit}
          onDelete={handleDelete}
        />
      </main>

      <Modal
        isOpen={modalOpen}
        onClose={closeModal}
        title={editingBook ? 'Edit Book' : 'Add a Book'}
      >
        <BookForm
          book={editingBook}
          onSubmit={handleSubmit}
          onCancel={closeModal}
          isSaving={isSaving}
        />
      </Modal>

      <Toast toast={toast} />
    </>
  );
}
