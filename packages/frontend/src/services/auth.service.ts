import axios from 'axios'
import { ACCESS_TOKEN_KEY, API_URL } from '../config'

let accessToken: string

export const setAccessToken = (token: string) => {
  accessToken = token
}

export const getAccessToken = () => accessToken

type AuthInput = {
  email: string
  password: string
}

// todo: update error message
// "pq: duplicate key value violates unique constraint \"users_email_key\""
const signUp = async (input: AuthInput) => {
  try {
    const { data } = await axios.post(`${API_URL}/auth/sign-up`, input, {
      headers: {
        'Content-Type': 'application/json'
      }
    })
    return data
  } catch (e) {
    console.error('signUp:', e)
  }
}

// todo: update error message
// "\"crypto/bcrypt: hashedPassword is not the hash of the given password\""
const signIn = async (input: AuthInput): Promise<string> => {
  try {
    const {
      data: { token }
    } = await axios.post(`${API_URL}/auth/sign-in`, input, {
      headers: {
        'Content-Type': 'application/json'
      }
    })
    return token as string
  } catch (e) {
    console.error('signIn:', e)
    return ''
  }
}

export const authenticate = async (input: AuthInput) => {
  await signUp(input) // signUp if needed
  const token = await signIn(input)
  if (!token) throw new Error('authenticate: no token provided')
  setAccessToken(token)
  localStorage.setItem(ACCESS_TOKEN_KEY, token)
}

export const checkIfAuthenticated = (): boolean => {
  const token = accessToken || localStorage.getItem(ACCESS_TOKEN_KEY)
  if (!token) return false
  setAccessToken(token)
  return true
}
