import { useEffect, useState } from 'react'
import { Routes, Route, useNavigate } from 'react-router-dom'
import { Items, Lists, SignIn } from './components'
import { checkIfAuthenticated } from './services/auth.service'

export const App = () => {
  const navigate = useNavigate()

  useEffect(() => {
    if (!checkIfAuthenticated()) navigate('/sign-in')
  }, [])

  return (
    <div className='h-full w-full flex justify-center items-center'>
      <div className='flex flex-col p-3 w-80 h-[38rem] rounded-md shadow'>
        <div className='flex justify-center items-center'>
          <h1
            className='w-32 font-bold text-lg text-center cursor-pointer'
            onClick={() => navigate('/')}
          >
            🎯 Todo App
          </h1>
        </div>

        <Routes>
          <Route path='/' element={<Lists />} />
          <Route path='lists/:id' element={<Items />} />
          <Route path='sign-in' element={<SignIn />} />
        </Routes>
      </div>
    </div>
  )
}
