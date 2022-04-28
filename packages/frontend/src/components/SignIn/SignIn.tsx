import isEmpty from 'lodash/isEmpty'
import { FormEvent, useEffect, useState } from 'react'
import { useMutation } from 'react-query'
import { authenticate } from '../../services/auth.service'
import { ErrorScreen } from '../ErrorScreen'
import { Spinner } from '../Spinner'
import { validateSignInInput } from './validation'

type SignInProps = {
  setIsUserSignedIn: (isSignedIn: boolean) => void
}

export const SignIn = ({ setIsUserSignedIn }: SignInProps) => {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [errors, setErrors] = useState<{ email?: string; password?: string }>(
    {}
  )
  const [triedToSubmit, setTriedToSubmit] = useState(false)
  const [isError, setIsError] = useState(false)
  const { mutate, isLoading } = useMutation(authenticate)

  useEffect(() => {
    if (!triedToSubmit) return
    validateSignInInput({ email, password })
      .then(errs => setErrors(errs))
      .catch(e => console.error('useEffect: validateSignInInput:', e))
  }, [email, password, triedToSubmit])

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    setTriedToSubmit(true)
    try {
      const input = { email, password }
      const validationErrors = await validateSignInInput(input)
      setErrors(validationErrors)
      if (!isEmpty(validationErrors)) return
      mutate(input, {
        onSuccess: () => {
          setIsUserSignedIn(true)
        },
        onError: () => {
          setIsUserSignedIn(false)
          setIsError(true)
        }
      })
    } catch (e) {
      console.error('handleSubmit:', e)
    }
  }

  if (isLoading) {
    return (
      <div className='h-full w-full flex justify-center items-center'>
        <Spinner />
      </div>
    )
  }

  if (isError) {
    return <ErrorScreen />
  }

  return (
    <div className='h-full w-full flex justify-center items-center'>
      <div className='w-80'>
        <h1 className='text-lg font-bold flex-1 text-center'>Todo App</h1>
        <h2 className='mt-1 text-lg font-bold flex-1 text-center'>
          👋 Welcome
        </h2>
        <form
          className='mt-4 flex flex-wrap justify-center'
          onSubmit={handleSubmit}
        >
          <input
            className={`p-2 border ${
              errors.email ? 'border-red-500' : 'border-gray-800'
            } rounded-md w-full`}
            type='text'
            name='email'
            id='email'
            placeholder='your.email@company.com'
            value={email}
            onChange={e => {
              setEmail(e.target.value)
            }}
          />
          {errors.email && (
            <p className='mt-1 w-full text-sm'>{errors.email}</p>
          )}
          <input
            className={`mt-2 p-2 border ${
              errors.password ? 'border-red-500' : 'border-gray-800'
            } rounded-md w-full`}
            type='password'
            name='password'
            id='password'
            placeholder='password'
            value={password}
            onChange={e => {
              setPassword(e.target.value)
            }}
          />
          {errors.password && (
            <p className='mt-1 w-full text-sm'>{errors.password}</p>
          )}
          <button
            className='mt-3 w-20 p-2 rounded-md bg-gray-900 text-neutral-50 font-bold'
            type='submit'
          >
            Sign In
          </button>
        </form>
      </div>
    </div>
  )
}
