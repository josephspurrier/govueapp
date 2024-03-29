import shortid from 'shortid'
import { HTTP2 } from '~/modules/http-common'

const STORAGE_KEY = 'app-note'

export const state = () => ({
  list: []
})

export const actions = {
  addItem: function(state, text) {
    const form = {
      message: text
    }

    HTTP2.send(
      'post',
      `v1/note`,
      form,
      state.rootState.auth,
      'Created.',
      function(data) {
        state.commit('addListItem', {
          id: data.record_id || shortid.generate(),
          message: text
        })
        state.commit('save')
      }
    )
  },
  updateItem: function(state, { index, key, text }) {
    const form = {
      message: text
    }

    HTTP2.send(
      'put',
      `v1/note/` + key,
      form,
      state.rootState.auth,
      'Saved.',
      function(data) {
        state.commit('updateItem', { index, text })
        state.commit('save')
      }
    )
  },
  deleteItem: function(state, { index, key }) {
    HTTP2.send(
      'delete',
      `v1/note/` + key,
      null,
      state.rootState.auth,
      'Deleted.',
      function(data) {
        state.commit('deleteItem', index)
        state.commit('save')
      }
    )
  },
  loadItems: function(state) {
    HTTP2.send(
      'get',
      `v1/note`,
      null,
      state.rootState.auth,
      'Loaded.',
      function(data) {
        state.commit('setItems', data.notes)
        state.commit('save')
      }
    )
  }
}

export const mutations = {
  save: function(state) {
    localStorage.setItem(STORAGE_KEY, JSON.stringify(state.list))
  },
  setItems: function(state, arr) {
    state.list = arr
  },
  loadItems: function(state) {
    state.list = JSON.parse(localStorage.getItem(STORAGE_KEY) || '[]')
  },
  addListItem: function(state, item) {
    state.list.push(item)
  },
  updateItem: function(state, { index, text }) {
    state.list[index].message = text
  },
  deleteItem: function(state, index) {
    state.list.splice(index, 1)
  }
}

export default {
  state,
  actions,
  mutations
}
