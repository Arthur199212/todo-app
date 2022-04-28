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
  if (!id) throw new Error('createTodoList: list was not created')
  return id
}

export const getAllLists = async (): Promise<any> => {
  const { data } = await instance.get('/lists/')
  if (!data) throw new Error('getAllLists: could not get todo lists')
  return data
}
