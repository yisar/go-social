import { h, useState } from 'fre'
import { A, push } from '../use-route'
import { loginPost } from '../util/api'
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
        loginPost({name, pwd}).then((res:any)=>{
            if(res.code > 0){
                alert("登录成功")
                window.localStorage.setItem('token', res.data.token)
                window.localStorage.setItem('author', JSON.stringify(res.data.author))
                push('/')
               
            }
        })
    }
    return <div class="login">
        <li><h1>海棠。</h1><h2>文学城</h2></li>
        <li><input type="text" placeholder="笔名" onInput={(e) => changeName(e.target.value)} /></li>
        <li><input type="text" placeholder="密码" onInput={(e) => changePwd(e.target.value)} /></li>
        <li><button onClick={login}>登录</button></li>
        <li><A href="/register">注册</A></li>
    </div>
}
