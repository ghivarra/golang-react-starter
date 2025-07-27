import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter, Routes, Route } from "react-router"
import './index.css'
import App from './App.tsx'
import AuthLayoutView from './views/auth-layout-view.tsx'
import LoginPage from './views/pages/auth/login/login-page.tsx'
import RegisterPage from './views/pages/auth/register/register-page.tsx'
import routeCollection from './lib/route-collection.ts'
import Loader from './components/loader.tsx'
import PanelLayoutView from './views/panel-layout-view.tsx'
import DashboardPage from './views/pages/panel/dashboard/dashboard-page.tsx'

createRoot(document.getElementById('root')!).render(
    <StrictMode>
        <Loader />
        <BrowserRouter>
            <Routes>

                {/* HOME ROUTE  */}
                <Route path={routeCollection.home} element={ <App /> } />

                {/* USER ROUTE  */}
                <Route element={ <AuthLayoutView /> }>
                    <Route path={routeCollection.user_login} element={ <LoginPage /> } />
                    <Route path={routeCollection.user_register} element={ <RegisterPage /> } />
                </Route>

                {/* PANEL ROUTE  */}
                <Route element={ <PanelLayoutView /> }>
                    <Route path={routeCollection.panel_dashboard} element={ <DashboardPage /> } />
                </Route>

            </Routes>
        </BrowserRouter>
    </StrictMode>,
)
