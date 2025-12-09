import {
  SignedIn,
  SignedOut,
  SignInButton,
  SignUpButton,
  UserButton,
} from '@clerk/clerk-react'
import WebSocketTestClient from './components/WebSocketTestClient.tsx'
import InterviewList from './components/InterviewList.tsx'

export default function App() {
  return (
    <>
      <header>
        <SignedOut>
          <SignInButton />
          <SignUpButton />
        </SignedOut>
        <SignedIn>
          <UserButton />
          <InterviewList />
        </SignedIn>
      </header>
      <main>
        <WebSocketTestClient />
      </main>
    </>
  )
}
