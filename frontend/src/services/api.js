import axios from 'axios'

const API_URL = 'http://localhost:8080/api'

const api = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json'
  }
})

api.interceptors.request.use(config => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

export default {
  login(credentials) {
    return api.post('/login', credentials)
  },
  register(userData) {
    return api.post('/register', userData)
  },
  getUser() {
    return api.get('/user')
  },
  getBookmarks() {
    return api.get('/bookmarks')
  },
  createBookmark(bookmarkData) {
    return api.post('/bookmarks', bookmarkData)
  },
  updateBookmark(id, bookmarkData) {
    return api.put(`/bookmarks/${id}`, bookmarkData)
  },
  deleteBookmark(id) {
    return api.delete(`/bookmarks/${id}`)
  },
  getCategories() {
    return api.get('/categories')
  },
  createCategory(categoryData) {
    return api.post('/categories', categoryData)
  },
  updateCategory(id, categoryData) {
    return api.put(`/categories/${id}`, categoryData)
  },
  deleteCategory(id) {
    return api.delete(`/categories/${id}`)
  },
  getSettings() {
    return api.get('/settings')
  },
  updateSettings(settings) {
    return api.put('/settings', settings)
  },
  getPublicBookmarks(publicId, accessCode) {
    return api.get(`/share/${publicId}`, {
      params: accessCode ? { access_code: accessCode } : {}
    })
  }
}