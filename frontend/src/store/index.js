import { createStore } from 'vuex'
import api from '../services/api'
import router from '../router'

// const API_URL = 'http://localhost:8080/api'

export default createStore({
  state: {
    token: localStorage.getItem('token') || '',
    user: null,
    bookmarks: [],
    categories: [],
    userSettings: {}
  },
  mutations: {
    setToken(state, token) {
      state.token = token
    },
    setUser(state, user) {
      state.user = user
    },
    setBookmarks(state, bookmarks) {
      state.bookmarks = bookmarks
    },
    setCategories(state, categories) {
      state.categories = categories
    },
    setUserSettings(state, settings) {
      state.userSettings = settings
    },
    addBookmark(state, bookmark) {
        state.bookmarks.push(bookmark)
    },
    updateBookmark(state, updatedBookmark) {
        const index = state.bookmarks.findIndex(b => b.id === updatedBookmark.id)
        if (index !== -1) {
            state.bookmarks.splice(index, 1, updatedBookmark)
        }
    },
    removeBookmark(state, bookmarkId) {
      state.bookmarks = state.bookmarks.filter(bookmark => bookmark.ID !== bookmarkId)
    },
    addCategory(state, category) {
      state.categories.push(category)
    },
    updateCategory(state, updatedCategory) {
      const index = state.categories.findIndex(c => c.ID === updatedCategory.ID)
      if (index !== -1) {
        state.categories.splice(index, 1, updatedCategory)
      }
    },
    removeCategory(state, categoryId) {
      state.categories = state.categories.filter(c => c.ID !== categoryId)
    }
  },
  actions: {
    async login({ commit, dispatch }, credentials) {
        try {
            const response = await api.login(credentials)
            const token = response.data.token
            localStorage.setItem('token', token)
            commit('setToken', token)
            await dispatch('fetchUser')
            router.push('/home')  // 添加這行
            return true
        } catch (error) {
            console.error('Login error:', error)
            return false
        }
    },
    async register(_, userData) {
        await api.register(userData)
    },
    async fetchUser({ commit }) {
        try {
            const response = await api.getUser()
            commit('setUser', response.data)
        } catch (error) {
            console.error('Fetch user error:', error)
            // Maybe redirect to login page or clear token then show error message...
            // commit('setToken', '')
            // localStorage.removeItem('token')
        }
    },
    async fetchBookmarks({ commit }) {
        const response = await api.getBookmarks()
        commit('setBookmarks', response.data)
    },
    async fetchCategories({ commit }) {
        try {
            const response = await api.getCategories()
            commit('setCategories', response.data)
        } catch (error) {
            console.error('Error fetching categories:', error)
        }
    },
    async createBookmark({ commit }, bookmarkData) {
        try {
            const response = await api.createBookmark(bookmarkData)
            commit('addBookmark', response.data)
            return response.data
        } catch (error) {
            console.error('Error creating bookmark:', error)
            throw error
        }
    },
  
    async updateBookmark({ commit }, { id, bookmarkData }) {
      try {
        const response = await api.updateBookmark(id, bookmarkData)
        commit('updateBookmark', response.data)
      } catch (error) {
        console.error('Error updating bookmark:', error)
        throw error
      }
    },
  
    async deleteBookmark({ commit }, id) {
      try {
        await api.deleteBookmark(id)
        commit('removeBookmark', id)
      } catch (error) {
        console.error('Error deleting bookmark:', error)
        throw error
      }
    },
    async createCategory({ commit }, categoryData) {
      try {
        const response = await api.createCategory(categoryData)
        commit('addCategory', response.data)
        return response.data
      } catch (error) {
        console.error('Error creating category:', error)
        throw error
      }
    },
    async updateCategory({ commit }, { id, categoryData }) {
      try {
        const response = await api.updateCategory(id, categoryData)
        commit('updateCategory', response.data)
        return response.data
      } catch (error) {
        console.error('Error updating category:', error)
        throw error
      }
    },
    async deleteCategory({ commit }, id) {
      try {
        await api.deleteCategory(id)
        commit('removeCategory', id)
      } catch (error) {
        console.error('Error deleting category:', error)
        throw error
      }
    },
    async fetchSettings({ commit }) {
      try {
        const response = await api.getSettings()
        commit('setUserSettings', response.data)
        return response.data
      } catch (error) {
        console.error('Error fetching settings:', error)
        throw error
      }
    },
    async updateSettings({ commit }, settings) {
      try {
        const response = await api.updateSettings(settings)
        commit('setUserSettings', response.data)
        return response.data
      } catch (error) {
        console.error('Error updating settings:', error)
        throw error
      }
    },
    logout({ commit }) {
      localStorage.removeItem('token')
      commit('setToken', '')
      commit('setUser', null)
    }
  },
  getters: {
    isLoggedIn: state => !!state.token
  }
})