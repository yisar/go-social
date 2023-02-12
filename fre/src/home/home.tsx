import { h, useEffect, useState } from 'fre'
import { getThreads } from '../util/api'
import './home.css'
import List from '../list/list'

export default function Home() {
    const [list, setList] = useState([])
    const [list2, setList2] = useState([])
    useEffect(() => {
        getThreads('原创').then((res: any) => {
            setList(res.data)
        })
        getThreads('交流').then((res: any) => {
            setList2(res.data)
        })
    }, [])
    return <div class='wrapper'>
        <List name='原创区' list={list} />
        <List name='交流区' list={list2} />

    </div>
}