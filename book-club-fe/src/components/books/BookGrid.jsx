import BookCard from './BookCard';
import styles from './BookGrid.module.css';

export default function BookGrid({ books, loading, onEdit, onDelete }) {
  if (loading) {
    return (
      <div className={styles.state}>
        <div className={styles.spinner} />
        <p>Loading your shelf…</p>
      </div>
    );
  }

  if (!books.length) {
    return (
      <div className={styles.state}>
        <p className={styles.empty}>Your shelf is empty here — add a book to begin.</p>
      </div>
    );
  }

  return (
    <div className={styles.grid}>
      {books.map((book) => (
        <BookCard
          key={book.id}
          book={book}
          onEdit={onEdit}
          onDelete={onDelete}
        />
      ))}
    </div>
  );
}
