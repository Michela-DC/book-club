import Button from '../ui/Button';
import styles from './Header.module.css';

export default function Header({ onAddBook }) {
  return (
    <header className={styles.header}>
      <div className={styles.logo}>
        Book<span>Club</span>
      </div>
      <Button variant="ghost" onClick={onAddBook}>
        ＋ Add Book
      </Button>
    </header>
  );
}
