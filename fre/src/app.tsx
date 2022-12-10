import { render, Fragment, h, } from "fre"
import { useRoutes } from './use-route'
import './app.css'

const routes = {
    '/': import('./home/home'),
    '/login': import('./login/login'),
    '/register': import('./login/register'),
}

const App = () => {
    let route = useRoutes(routes)
    return <div>{route}</div>

}

render(<App />, document.getElementById("app"))
