import { SignIn } from './components/SignIn/SignIn'

export const App = () => {
  const isUserSignedIn = false

  if (!isUserSignedIn) return <SignIn />

  return (
    <div>
      Todo App
    </div>
  )
}
