import { render, Fragment, h, } from "fre"
import { useRoutes } from './use-route'
import './app.css'

const routes = {
    '/login': import('./login/login'),
    '/register': import('./login/register'),
}

function A(){
    return <div>222</div>
}

const App = () => {
    let route = useRoutes(routes)
    return <div>{route}</div>

}

render(<App />, document.getElementById("app"))


// // 以下都是时间戳对比
// if (Date.now() < 1670256000000 && window.location.pathname === '/') {
//     document.getElementById('app').style = `filter: grayscale(100%);position:absolute;top:0;bottom:0;left:0;right:0;`
// }


