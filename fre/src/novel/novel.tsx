import { h, useEffect, useState } from 'fre'
import { A } from '../use-route'
import { getNovel } from '../util/api'
import './novel.css'

export default function Novel(props) {
    const id = props.id.slice(2)

    useEffect(() => {
        getNovel(id).then(res => {
            setNovel(res.data)
        })
    }, [])

    const [novel, setNovel] = useState({} as any)

    // const tags = 

    return <div class='wrapper'>
        <div class='detail'>
            <h1>{novel.title}</h1>
            <p>{novel.content}</p>
            <ul class='info'>
                <li>{novel.status}</li>
                <li>{novel.size}</li>
                <li>{novel.aptitude}</li>
                <li>{novel.sort}</li>
            </ul>
            <ul class='tag'>
                {novel.tag && novel.tag.split(' ').filter(t => t.length > 0).map(tag => {
                    return <li>#{tag}</li>
                })}
            </ul>

        </div>
        <div class='reply'>
            <textarea name="" id="" rows="10"></textarea>
            <button>发布</button>

        </div>
    </div>
}