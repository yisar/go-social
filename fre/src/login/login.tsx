import { h, useState } from 'fre'
import { A, push } from '../use-route'
import { post } from '../util/post'
import './login.css'

export default function Login() {
    const [name, setName] = useState("")
    const [pwd, setPwd] = useState("")

    function changeName(v) {
        setName(v)
    }

    function changePwd(v) {
        setPwd(v)
    }

    function login() {
        alert("还没搞完呢")
    }
    return <div class="login">
        <li><h1>海棠。</h1><h2>文学城</h2></li>
        <li><input type="text" placeholder="笔名" onInput={(e) => changeName(e.target.value)} /></li>
        <li><input type="text" placeholder="密码" onInput={(e) => changePwd(e.target.value)} /></li>
        <li><button onClick={login}>登录</button></li>
        <li><A href="/register">注册</A></li>
    </div>
}
