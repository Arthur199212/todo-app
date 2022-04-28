import { useEffect, useState } from 'react'
import { Routes, Route, useNavigate } from 'react-router-dom'
import { SignIn } from './components/SignIn'
import { Lists } from './components/Lists'
import { checkIfAuthenticated } from './services/auth.service'

export const App = () => {
  const [isUserSignedIn, setIsUserSignedIn] = useState(false)
  const navigate = useNavigate()

  useEffect(() => {
    setIsUserSignedIn(checkIfAuthenticated())
  }, [])

  if (!isUserSignedIn) {
    navigate('/sign-in')
  }

  return (
    <div className='h-full w-full flex justify-center items-center'>
      <div className='flex flex-col p-3 w-80 h-[38rem] rounded-md shadow'>
        <h1 className='font-bold text-lg text-center'>🎯 Todo App</h1>

        <Routes>
          <Route path='/' element={<Lists />} />
          <Route path='lists/:id' element={<Items />} />
          <Route
            path='sign-in'
            element={<SignIn setIsUserSignedIn={setIsUserSignedIn} />}
          />
        </Routes>
      </div>
    </div>
  )
}

function Items () {
  return <div>Items of a list</div>
}
