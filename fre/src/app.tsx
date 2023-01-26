import { render, Fragment, h, } from "fre"
import { useRoutes } from './use-route'
import Header from './header/header'
import Publish from './publish/publish'


import './app.css'

const routes = {
    '/': import('./home/home'),
    '/publish': import('./publish/publish'),
    '/login': import('./login/login'),
    '/register': import('./login/register'),
}

const App = () => {
    let route = useRoutes(routes)
    return <div>
        {window.location.pathname !== '/login' && window.location.pathname !== '/register' && <Header />}
        <div>{route}</div>
    </div>

}

render(<App />, document.getElementById("app"))
