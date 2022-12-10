import { h, render } from 'fre'
import { getUser } from '../util/api'
import { getAvatar } from '../util/avatar'
import './header.css'

export default function Header() {
    const user = getUser()
    return <header>
        <div className="wrapper">
            <div className="logo">HenTai</div>
            <nav>
                <li>Library</li>
                <li>Forum</li>
            </nav>
            <div className="avatar"><img src={getAvatar(user.email)} alt="" /></div>
        </div>
    </header>
}