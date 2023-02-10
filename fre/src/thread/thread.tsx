import { h, useEffect, useState } from 'fre'
import { A, push } from '../use-route'
import { addPost, getPostDetail, getPosts, getThread, getUser } from '../util/api'
import './thread.css'

export default function Thread(props) {
    const id = props.id

    const [thread, setThread] = useState({} as any)
    const [data, setData] = useState({} as any)
    const [list, setList] = useState([])
    const [current, setIndex] = useState(0)

    useEffect(() => {
        getThread(id).then(res => {
            setThread(res.data)
            getPosts(res.data._id).then(res2 => {
                setList(res2.data)
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
            tid: thread._id
        }).then(res => {
            alert(res.msg)
        })
    }

    function open(id) {
        getPostDetail(id).then(res=>{
            console.log(res.data)
        })
    }

    const user = getUser()||{}

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
                <li onClick={()=>push(`/publish/${thread._id}`)}>编辑小说</li>
            </ul>
        </div>

        <div class='list'>
            {list.map((item, index) => {
                // const content = item.content.replace(/\s+/g,'\n')

                return <div class='post'>
                    <div onClick={() => {
                        open(item._id)
                    }}>
                        <h2>{item.title}</h2>
                    </div>
                </div>
            })}
        </div>

        <div class='reply'>
            {isUser && <input type="text" placeholder='请输入章节序号' onInput={e => changeData('oid', parseInt(e.target.value))} />}
            {isUser && <input type="text" placeholder='请输入标题' onInput={e => changeData('title', e.target.value)} />}
            <textarea name="" id="" rows="10" onInput={e => changeData('content', e.target.value.replace(/\s+/g,'\n'))}></textarea>
            <button onClick={publish}>发布</button>
        </div>
    </div>
}