import { useEffect, useState } from 'react'
import { SignIn } from './components/SignIn/SignIn'
import { checkIfAuthenticated } from './services/auth.service'

export const App = () => {
  const [isUserSignedIn, setIsUserSignedIn] = useState(false)

  useEffect(() => {
    setIsUserSignedIn(checkIfAuthenticated())
  }, [])

  if (!isUserSignedIn) return <SignIn setIsUserSignedIn={setIsUserSignedIn} />

  return <div>Todo App</div>
}
