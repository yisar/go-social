import { h, useEffect, useState } from 'fre'
import { A } from '../use-route'
import { addChapter, getNovel, getUser } from '../util/api'
import './novel.css'

export default function Novel(props) {
    const id = props.id.slice(2)

    useEffect(() => {
        getNovel(id).then(res => {
            setNovel(res.data)
        })
    }, [])

    const [novel, setNovel] = useState({} as any)
    const [data, setData] = useState({} as any)

    function changeData(key, val) {
        console.log(key, val)
        setData({
            ...data,
            [key]: val,
        })
    }


    function publish() {
        addChapter({
            ...data,
            status: '发布',
            nid: novel._id
        }).then(res => {
            console.log(res)
        })
    }

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
            <input type="text" placeholder='请输入章节序号' onInput={e => changeData('oid', e.target.value)} />
            <input type="text" placeholder='请输入标题' onInput={e => changeData('title', e.target.value)} />
            <textarea name="" id="" rows="10" onInput={e => changeData('content', e.target.value)}></textarea>
            <button onClick={publish}>发布</button>

        </div>
    </div>
}