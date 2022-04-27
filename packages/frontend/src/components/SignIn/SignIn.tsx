import { useState } from 'react'

export const SignIn = () => {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')

  return (
    <div className='h-full w-full flex justify-center items-center'>
      <div className='w-80'>
        <h1 className='text-lg font-bold flex-1 text-center'>Todo App</h1>
        <h2 className='mt-1 text-lg font-bold flex-1 text-center'>
          👋 Welcome
        </h2>
        <form className='mt-4 flex flex-wrap justify-center' action='submit'>
          <input
            className='p-2 border border-gray-800 rounded-md w-full'
            type='text'
            name='email'
            id='email'
            placeholder='your.email@company.com'
            value={email}
            onChange={e => {
              setEmail(e.target.value)
            }}
          />
          <input
            className='mt-2 p-2 border border-gray-800 rounded-md w-full'
            type='password'
            name='password'
            id='password'
            placeholder='password'
            value={password}
            onChange={e => {
              setPassword(e.target.value)
            }}
          />
          <button
            className='mt-3 w-20 p-2 rounded-md bg-gray-900 text-neutral-50 font-bold'
            type='button'
          >
            Sign In
          </button>
        </form>
      </div>
    </div>
  )
}
