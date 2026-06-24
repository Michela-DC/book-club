import { FILTER_TABS } from '../../api/constants';
import styles from './FilterTabs.module.css';

export default function FilterTabs({ active, onChange, books }) {
  function countFor(key) {
    if (key === 'all') return books.length;
    return books.filter((b) => b.status === key).length;
  }

  return (
    <div className={styles.tabs}>
      {FILTER_TABS.map((tab) => (
        <button
          key={tab.key}
          className={`${styles.tab} ${active === tab.key ? styles.active : ''}`}
          onClick={() => onChange(tab.key)}
        >
          {tab.label}
          <span className={styles.count}>{countFor(tab.key)}</span>
        </button>
      ))}
    </div>
  );
}
