import { h, useState, useEffect } from 'fre'
import { A, push } from '../use-route'
import { registerUser } from '../util/api'
import { post } from '../util/post'
import './login.css'

export default function Register({ id }) {

    const [name, setName] = useState(null)
    const [pwd, setPwd] = useState(null)
    const [email, setEmail] = useState(null)
    const [loading, setLoading] = useState(false)
    const [level, setLevel] = useState(0)


    useEffect(() => {
        if (id) {
            console.log('编辑用户')
        }

    }, [])


    function changeName(v) {
        setName(v)
    }

    function changePwd(v) {
        setPwd(v)
    }

    function changeEmail(v) {
        setEmail(v)
    }

    function changeLevel(v) {
        setLevel(v)
    }

    async function register() {
        if (id != null) {
            console.log('修改用户')
            return
        }

        setLoading(true)

        registerUser({
            name,
            pwd,
            email
        }).then(res=>{
          alert("注册成功啦~")
          setLoading(false)
        })
        
    }
    function logout() {
        localStorage.clear()
        location.reload()
    }
    return <div class="login">
        <li><h1>海棠。</h1><h2>文学城</h2></li>
        <li><input type="text" placeholder="邮箱" onInput={(e) => changeEmail(e.target.value)} value={email} /></li>
        <li><input type="text" placeholder="笔名" onInput={(e) => changeName(e.target.value)} value={name} /></li>
        <li><input type="text" placeholder={id ? "留空则不改" : "密码"} onInput={(e) => changePwd(e.target.value)} /></li>
        {id && <select value={level} onInput={e => changeLevel(e.target.value)}>
            <option value="1">游客</option>
            <option value="2">作者</option>
            <option value="3">审核</option>
            <option value="4">管理</option>
        </select>}
        <li><button onClick={register} disabled={loading}>{loading ? '少年注册中...' : id ? '修改' : '注册'}</button></li>
        {id && <li><button onClick={logout} style={{ background: '#ff2b79' }}>退出登陆</button></li>}
        {!id && <li><A href="/login">登录</A></li>}
    </div>
}