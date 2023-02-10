import { render, Fragment, h, } from "fre"
import { useRoutes } from './use-route'
import Header from './header/header'


import './app.css'

const routes = {
    '/': import('./home/home'),
    '/publish/:id': import('./publish/publish'),
    '/login': import('./login/login'),
    '/register': import('./login/register'),
    '/novel/:id': import('./novel/novel'),
}

const App = () => {
    let route = useRoutes(routes)
    return <div>
        {window.location.pathname !== '/login' && window.location.pathname !== '/register' && <Header />}
        <div>{route}</div>
    </div>

}

render(<App />, document.getElementById("app"))
