import axios from 'axios'
import { API_URL } from '../config'
import { getAccessToken } from './auth.service'

const instance = axios.create({
  baseURL: `${API_URL}/api`
})

instance.interceptors.request.use(
  function (config) {
    config.headers.Authorization = 'Bearer ' + getAccessToken()
    return config
  },
  function (error) {
    // todo: refresh token
    return Promise.reject(error)
  }
)

export const createList = async (title: string): Promise<string> => {
  const {
    data: { id }
  } = await instance.post('/lists/', { title })
  if (!id) throw new Error('createList: list was not created')
  return id
}

export const getAllLists = async (): Promise<any> => {
  const { data } = await instance.get('/lists/')
  if (!data) throw new Error('getAllLists: could not get todo lists')
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
  if (!id) throw new Error('createItem: list was not created')
  return id
}

export const getAllItemsByListId = async (listId: string): Promise<any> => {
  const { data } = await instance.get(`/lists/${listId}/items/`)
  if (!data) throw new Error('getAllItemsByListId: could not get todo lists')
  return data
}
