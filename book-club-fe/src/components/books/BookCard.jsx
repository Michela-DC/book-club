import StatusBadge from '../ui/StatusBadge';
import Button from '../ui/Button';
import styles from './BookCard.module.css';

export default function BookCard({ book, onEdit, onDelete }) {
  return (
    <article className={styles.card} data-status={book.status}>
      {book.genre && <span className={styles.genre}>{book.genre}</span>}
      <h3 className={styles.title}>{book.title}</h3>
      <p className={styles.author}>{book.author}</p>
      <StatusBadge status={book.status} />
      {book.published_year && (
        <p className={styles.year}>{book.published_year}</p>
      )}
      <div className={styles.actions}>
        <Button size="sm" variant="primary" onClick={() => onEdit(book)}>
          Edit
        </Button>
        <Button size="sm" variant="danger" onClick={() => onDelete(book)}>
          Remove
        </Button>
      </div>
    </article>
  );
}
