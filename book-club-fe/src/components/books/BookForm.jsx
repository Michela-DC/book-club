import { useState, useEffect } from 'react';
import Button from '../ui/Button';
import { STATUS_OPTIONS_ALL, STATUS_OPTIONS_CREATE } from '../../api/constants';
import styles from './BookForm.module.css';

const EMPTY = { title: '', author: '', genre: '', year: '', status: 'to_read' };

function toForm(book) {
  if (!book) return EMPTY;
  return {
    title: book.title || '',
    author: book.author || '',
    genre: book.genre || '',
    year: book.published_year ? String(book.published_year) : '',
    status: book.status || 'to_read',
  };
}

export default function BookForm({ book, onSubmit, onCancel, isSaving }) {
  const isEdit = Boolean(book);
  const [fields, setFields] = useState(() => toForm(book));
  const [errors, setErrors] = useState({});

  useEffect(() => {
    setFields(toForm(book));
    setErrors({});
  }, [book]);

  function set(key, value) {
    setFields((f) => ({ ...f, [key]: value }));
    setErrors((e) => ({ ...e, [key]: undefined }));
  }

  function validate() {
    const errs = {};
    if (!fields.title.trim()) errs.title = 'Title is required.';
    if (!fields.author.trim()) errs.author = 'Author is required.';
    return errs;
  }

  function handleSubmit() {
    const errs = validate();
    if (Object.keys(errs).length) { setErrors(errs); return; }
    onSubmit({
      title: fields.title.trim(),
      author: fields.author.trim(),
      genre: fields.genre.trim() || undefined,
      year: fields.year ? parseInt(fields.year, 10) : undefined,
      status: fields.status,
    });
  }

  const statusOptions = isEdit ? STATUS_OPTIONS_ALL : STATUS_OPTIONS_CREATE;

  return (
    <div className={styles.form}>
      <div className={styles.row}>
        <label className={styles.label}>Title *</label>
        <input
          className={`${styles.input} ${errors.title ? styles.inputError : ''}`}
          value={fields.title}
          onChange={(e) => set('title', e.target.value)}
          placeholder="The Name of the Rose"
          autoFocus
        />
        {errors.title && <p className={styles.error}>{errors.title}</p>}
      </div>

      <div className={styles.row}>
        <label className={styles.label}>Author *</label>
        <input
          className={`${styles.input} ${errors.author ? styles.inputError : ''}`}
          value={fields.author}
          onChange={(e) => set('author', e.target.value)}
          placeholder="Umberto Eco"
        />
        {errors.author && <p className={styles.error}>{errors.author}</p>}
      </div>

      <div className={styles.grid2}>
        <div className={styles.row}>
          <label className={styles.label}>Genre</label>
          <input
            className={styles.input}
            value={fields.genre}
            onChange={(e) => set('genre', e.target.value)}
            placeholder="Historical Fiction"
          />
        </div>
        <div className={styles.row}>
          <label className={styles.label}>Year</label>
          <input
            className={styles.input}
            type="number"
            value={fields.year}
            onChange={(e) => set('year', e.target.value)}
            placeholder="1980"
            min="0"
            max="2099"
          />
        </div>
      </div>

      <div className={styles.row}>
        <label className={styles.label}>Status *</label>
        <select
          className={styles.input}
          value={fields.status}
          onChange={(e) => set('status', e.target.value)}
        >
          {statusOptions.map((o) => (
            <option key={o.value} value={o.value}>{o.label}</option>
          ))}
        </select>
      </div>

      <div className={styles.actions}>
        <Button variant="secondary" onClick={onCancel} disabled={isSaving}>
          Cancel
        </Button>
        <Button variant="primary" onClick={handleSubmit} disabled={isSaving}>
          {isSaving ? 'Saving…' : isEdit ? 'Save Changes' : 'Add Book'}
        </Button>
      </div>
    </div>
  );
}
