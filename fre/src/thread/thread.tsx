import { h, useEffect, useState } from 'fre'
import { A, push } from '../use-route'
import { addPost, getPostDetail, getPosts, getThread, getUser, getUserDetail } from '../util/api'
import './thread.css'

export default function Thread(props) {
    const id = props.id

    const [thread, setThread] = useState({} as any)
    const [data, setData] = useState({} as any)
    const [list, setList] = useState([])
    const [content, setConent] = useState("")
    const [current, setCurrent] = useState(0)

    useEffect(() => {
        getThread(id).then(res => {
            setThread(res.data)
            getPosts(res.data._id).then(res2 => {
                const arr = res2.data.map(async item => {
                    const user = await getUserDetail(item.uid)
                    item.user = user
                    if (item.title === "") {
                        const c = await getPostDetail(item._id).then(res => res.data.content)
                        item.content = c
                        return item
                    } else {
                        return item
                    }
                })

                Promise.all(arr).then(res => {
                    console.log(res)
                    setList(res2.data)
                })
            })
        })
    }, [])

    function changeData(key, val) {
        console.log(key, val)
        setData({
            ...data,
            [key]: val,
        })
    }


    function publish() {
        addPost({
            ...data,
            status: '发布',
            tid: thread._id,
            length: data.content.length,
            uname: getUser().name
        }).then(res => {
            alert(res.msg)
        })
    }

    function open(id, index) {
        getPostDetail(id).then(res => {
            setConent("    " + res.data.content.replace(/\n/g, '\n    '))
            setCurrent(index)
        })
    }

    const user = getUser() || {}

    const isUser = thread.uid === user._id

    return <div class='wrapper'>
        <div class='detail'>
            <h1>《{thread.title}》</h1>
            <p>{thread.content}</p>
            <ul class='info'>
                <li>{thread.status}</li>
                <li>{thread.size}</li>
                <li>{thread.aptitude}</li>
                <li>{thread.sort}</li>
            </ul>
            <ul class='tag'>
                {thread.tag && thread.tag.split(' ').filter(t => t.length > 0).map(tag => {
                    return <li>#{tag}</li>
                })}
                <li onClick={() => push(`/publish/${thread._id}`)}>#编辑小说</li>
            </ul>
        </div>

        <div class='list'>
            {list.map((item, index) => {

                return <div class='post'>
                    <div >
                        <h2 onClick={() => {
                            open(item._id, index)
                        }}>{item.title}</h2>
                        {current === index && <pre>{content}</pre>}
                        {item.title === "" && <pre>{item.content}</pre>}
                    </div>
                </div>
            })}
        </div>

        <div class='reply'>
            {isUser && <input type="text" placeholder='请输入标题' onInput={e => changeData('title', e.target.value)} />}
            {isUser && <input type="text" placeholder='请输入概述' onInput={e => changeData('summary', e.target.value)} />}
            <textarea name="" id="" rows="10" onInput={e => changeData('content', !data.title ? e.target.value : e.target.value.replace(/\s+/g, '\n'))}></textarea>
            <button onClick={publish}>发布</button>
        </div>
    </div>
}