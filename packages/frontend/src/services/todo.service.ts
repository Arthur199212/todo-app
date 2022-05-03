import axios from 'axios'
import { ACCESS_TOKEN_KEY, API_URL } from '../config'
import { getAccessToken, setAccessToken } from './auth.service'

const instance = axios.create({
  baseURL: `${API_URL}/api`
})

instance.interceptors.request.use(function (config) {
  config.headers.Authorization = 'Bearer ' + getAccessToken()
  return config
})

instance.interceptors.response.use(undefined, function (err) {
  if (
    err.response.status === 400 &&
    err.response.data.includes('token is expired')
  ) {
    // todo: add refresh token logic
    console.warn('Token is expired')
    localStorage.removeItem(ACCESS_TOKEN_KEY)
    setAccessToken(undefined)
    window.location.reload()
  }
  return Promise.reject(err)
})

export const createList = async (title: string): Promise<string> => {
  const {
    data: { id }
  } = await instance.post('/lists/', { title })
  if (!id) throw new Error('createList: list was not created')
  return id
}

export const getAllLists = async () => {
  const { data } = await instance.get('/lists/')
  if (!data) throw new Error('getAllLists: could not get todo lists')
  return data
}

export const getListById = async (listId: string) => {
  const { data } = await instance.get(`/lists/${listId}`)
  if (!data) throw new Error('getListById: could not get todo lists')
  return data
}

export const createItem = async ({
  title,
  listId
}: {
  title: string
  listId: string
}): Promise<string> => {
  const {
    data: { id }
  } = await instance.post(`/lists/${listId}/items/`, { title })
  if (!id) throw new Error('createItem: item was not created')
  return id
}

export const getAllItemsByListId = async (listId: string) => {
  const { data } = await instance.get(`/lists/${listId}/items/`)
  if (!data) throw new Error('getAllItemsByListId: could not get todo item')
  return data
}

export const toggleTodo = async ({
  id,
  done
}: {
  id: string
  done: boolean
}) => {
  const { data } = await instance.put(`/items/${id}`, { done: !done })
  if (!data) throw new Error('toggleTodo: could not toggle todo')
  return data
}
