import { h, useEffect, useState } from 'fre'
import { getNovels } from '../util/api'
import './home.css'
import List from '../list/list'

export default function Home() {
    const [list, setList] = useState([])
    useEffect(() => {
        getNovels('原创').then((res: any) => {
            setList(res.data)
        })
    }, [])
    console.log(list)
    return <div class='wrapper'>
        <List name='原创区' list={list} />
        <List name='同人区' list={list} />

    </div>
}