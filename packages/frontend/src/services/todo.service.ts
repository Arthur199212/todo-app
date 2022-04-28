import axios from 'axios'
import { getAccessToken } from './auth.service'

const instance = axios.create()

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
