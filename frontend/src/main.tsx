import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter, Routes, Route } from "react-router"
import './index.css'
import App from './App.tsx'
import AuthView from './views/auth-view.tsx'
import LoginPage from './views/pages/auth/login-page.tsx'
import RegisterPage from './views/pages/auth/register-page.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>

        <Route path="/" element={ <App /> } />

        <Route path="user">
          <Route element={ <AuthView /> }>
            <Route path="login" element={ <LoginPage /> } />
            <Route path="register" element={ <RegisterPage /> } />
          </Route>
        </Route>

      </Routes>
    </BrowserRouter>
  </StrictMode>,
)
