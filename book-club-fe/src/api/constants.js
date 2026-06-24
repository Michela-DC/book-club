export const DEMO_BOOKS = [
  {
    id: '1',
    title: 'The Name of the Rose',
    author: 'Umberto Eco',
    genre: 'Historical Fiction',
    published_year: 1980,
    status: 'completed',
  },
  {
    id: '2',
    title: 'Invisible Cities',
    author: 'Italo Calvino',
    genre: 'Literary Fiction',
    published_year: 1972,
    status: 'completed',
  },
  {
    id: '3',
    title: 'The Leopard',
    author: 'Giuseppe Tomasi di Lampedusa',
    genre: 'Historical Novel',
    published_year: 1958,
    status: 'reading',
  },
  {
    id: '4',
    title: 'My Brilliant Friend',
    author: 'Elena Ferrante',
    genre: 'Literary Fiction',
    published_year: 2011,
    status: 'reading',
  },
  {
    id: '5',
    title: "If on a winter's night a traveler",
    author: 'Italo Calvino',
    genre: 'Postmodern',
    published_year: 1979,
    status: 'to_read',
  },
  {
    id: '6',
    title: 'Piranesi',
    author: 'Susanna Clarke',
    genre: 'Fantasy',
    published_year: 2020,
    status: 'to_read',
  },
];

export const STATUS_LABELS = {
  to_read: 'To Read',
  reading: 'Reading',
  completed: 'Completed',
  discarded: 'Discarded',
};

export const STATUS_OPTIONS_CREATE = [
  { value: 'to_read', label: 'To Read' },
  { value: 'reading', label: 'Reading' },
];

export const STATUS_OPTIONS_ALL = [
  { value: 'to_read', label: 'To Read' },
  { value: 'reading', label: 'Reading' },
  { value: 'completed', label: 'Completed' },
  { value: 'discarded', label: 'Discarded' },
];

export const FILTER_TABS = [
  { key: 'all', label: 'All' },
  { key: 'to_read', label: 'To Read' },
  { key: 'reading', label: 'Reading' },
  { key: 'completed', label: 'Completed' },
  { key: 'discarded', label: 'Discarded' },
];
