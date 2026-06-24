const API_BASE = import.meta.env.VITE_API_URL || 'http://localhost:8080';

async function request(method, path, body) {
  const opts = {
    method,
    headers: { 'Content-Type': 'application/json' },
  };
  if (body !== undefined) opts.body = JSON.stringify(body);

  const res = await fetch(`${API_BASE}${path}`, opts);

  if (!res.ok) {
    const text = await res.text().catch(() => res.statusText);
    throw new Error(text || `HTTP ${res.status}`);
  }

  if (res.status === 204 || res.headers.get('content-length') === '0') return null;
  return res.json();
}

// GET /v1/books
export function getBooks() {
  return request('GET', '/v1/books');
}

// PUT /v1/books
export function createBook(data) {
  return request('PUT', '/v1/books', data);
}

// PATCH /v1/books/:id
export function updateBook(id, data) {
  return request('PATCH', `/v1/books/${id}`, data);
}

// DELETE /v1/books/:id
export function deleteBook(id) {
  return request('DELETE', `/v1/books/${id}`);
}
