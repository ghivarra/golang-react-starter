import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter, Routes, Route } from "react-router"
import './index.css'
import App from './App.tsx'
import AuthView from './views/auth-view.tsx'
import LoginPage from './views/pages/auth/login-page.tsx'
import RegisterPage from './views/pages/auth/register-page.tsx'
import routeCollection from './lib/route-collection.ts'

createRoot(document.getElementById('root')!).render(
    <StrictMode>
        <BrowserRouter>
            <Routes>

                {/* HOME ROUTE  */}
                <Route path={routeCollection.home} element={ <App /> } />

                {/* USER ROUTE  */}
                <Route element={ <AuthView /> }>
                    <Route path={routeCollection.user_login} element={ <LoginPage /> } />
                    <Route path={routeCollection.user_register} element={ <RegisterPage /> } />
                </Route>

            </Routes>
        </BrowserRouter>
    </StrictMode>,
)
